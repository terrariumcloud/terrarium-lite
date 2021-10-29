package sources

import (
	"errors"
	"fmt"
	"log"
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
	VCSStore        types.VCSStore
	SourceStores    *SourcesMap
}

type SourcesMap struct {
	Github types.SourceStore
}

func (s *SourceAPI) CreateVCSModule() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		vcsID := params["id"]
		provider := params["provider"]
		var genericStore types.SourceStore = nil
		switch provider {
		case "github":
			genericStore = s.SourceStores.Github
		default:
			s.ErrorHandler.Write(rw, fmt.Errorf("vcs provider: %s is not supported", provider), http.StatusNotImplemented)
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
		log.Printf("%v", vcs.OAuth.Token)
		genericStore.FetchVCSSources(vcs.OAuth.Token.AccessToken)
	})
}

func (s *SourceAPI) SetupRoutes() {
	s.Router.StrictSlash(true)
	s.Router.Handle("/{provider}/{id}", s.CreateVCSModule()).Methods(http.MethodGet)
}

func NewSourceAPI(router *mux.Router, path string, vcsstore types.VCSStore, vcsProviders *SourcesMap, responseHandler types.APIResponseWriter, errorHandler types.APIErrorWriter) *SourceAPI {
	s := &SourceAPI{
		Router:          router.PathPrefix(path).Subrouter(),
		VCSStore:        vcsstore,
		SourceStores:    vcsProviders,
		ResponseHandler: responseHandler,
		ErrorHandler:    errorHandler,
	}
	s.SetupRoutes()
	return s
}
