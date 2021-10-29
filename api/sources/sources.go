package sources

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dylanrhysscott/terrarium/pkg/types"
	"github.com/gorilla/mux"
)

type SourceAPIInterface interface {
	SetupRoutes()
	CreateVCSModule() http.Handler
}

type SourceAPI struct {
	Router          *mux.Router
	ErrorHandler    types.APIErrorWriter
	ResponseHandler types.APIResponseWriter
	VCSStore        types.VCSSConnectionStore
	SourceStores    *SourcesMap
}

type SourcesMap struct {
	Github types.SourceStore
}

func (s *SourceAPI) detectStoreType(t string) (types.SourceStore, error) {
	var genericStore types.SourceStore = nil
	switch t {
	case "github":
		genericStore = s.SourceStores.Github
	default:
		return nil, fmt.Errorf("vcs provider: %s is not supported", t)
	}
	return genericStore, nil
}

func (s *SourceAPI) CreateVCSModule() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		vcsID := params["id"]
		provider := params["provider"]
		genericStore, err := s.detectStoreType(provider)
		if err != nil {
			s.ErrorHandler.Write(rw, err, http.StatusNotImplemented)
			return
		}
		vcs, err := s.VCSStore.ReadOne(vcsID)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				s.ErrorHandler.Write(rw, errors.New("vcs provider does not exist"), http.StatusNotFound)
				return
			}
			s.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		if vcs.OAuth.ServiceProvider != provider {
			s.ErrorHandler.Write(rw, errors.New("vcs provider mismatch"), http.StatusBadRequest)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.ErrorHandler.Write(rw, err, http.StatusUnprocessableEntity)
			return
		}
		reqBody := &SourceVCSRepoBody{}
		err = json.Unmarshal(body, reqBody)
		if err != nil {
			s.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		err = reqBody.Validate()
		if err != nil {
			s.ErrorHandler.Write(rw, err, http.StatusUnprocessableEntity)
			return
		}
		sourceRepo, err := genericStore.FetchVCSSource(vcs.OAuth.Token.AccessToken, reqBody.Repo, reqBody.Owner)
		if err != nil {
			s.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		s.ResponseHandler.Write(rw, sourceRepo, http.StatusOK)
	})
}

func (s *SourceAPI) SetupRoutes() {
	s.Router.StrictSlash(true)
	s.Router.Handle("/{provider}/{id}", s.CreateVCSModule()).Methods(http.MethodPost)
}

func NewSourceAPI(router *mux.Router, path string, vcsconnstore types.VCSSConnectionStore, vcsProviders *SourcesMap, responseHandler types.APIResponseWriter, errorHandler types.APIErrorWriter) *SourceAPI {
	s := &SourceAPI{
		Router:          router.PathPrefix(path).Subrouter(),
		VCSStore:        vcsconnstore,
		SourceStores:    vcsProviders,
		ResponseHandler: responseHandler,
		ErrorHandler:    errorHandler,
	}
	s.SetupRoutes()
	return s
}
