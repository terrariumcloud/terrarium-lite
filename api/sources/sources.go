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
	SourceStore     types.TerrariumSourceDriver
}

func sourceToModule(data types.SourceData, provider string, orgLink *types.ResourceLink, vcsConnectionID string) *types.VCSModule {
	return &types.VCSModule{
		Name:        data.GetRepoName(),
		Provider:    provider,
		Description: data.GetRepoDescription(),
		VCSConnection: &types.ResourceLink{
			ID:   vcsConnectionID,
			Link: fmt.Sprintf("/v1/oauth-clients/%s", vcsConnectionID),
		},
		Organization: orgLink,
		VCSRepo:      data,
	}
}

func (s *SourceAPI) detectStoreType(t string) (types.SourceProvider, error) {
	var genericStore types.SourceProvider = nil
	switch t {
	case "github":
		genericStore = s.SourceStore.GithubSources()
	default:
		return nil, fmt.Errorf("vcs provider: %s is not supported", t)
	}
	return genericStore, nil
}

func (s *SourceAPI) CreateVCSModule() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		vcsConnID := params["vcsconnid"]
		vcsProvider := params["vcsprovider"]
		repoName := params["reponame"]
		genericStore, err := s.detectStoreType(vcsProvider)
		if err != nil {
			s.ErrorHandler.Write(rw, err, http.StatusNotImplemented)
			return
		}
		vcs, err := s.VCSStore.ReadOne(vcsConnID, true)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				s.ErrorHandler.Write(rw, errors.New("vcs provider does not exist"), http.StatusNotFound)
				return
			}
			s.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		if vcs.OAuth.ServiceProvider != vcsProvider {
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
		sourceRepo, err := genericStore.FetchVCSSource(vcs.OAuth.Token.AccessToken, repoName)
		if err != nil {
			s.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		module := sourceToModule(sourceRepo, reqBody.Provider, vcs.Organization, vcsConnID)
		s.ResponseHandler.Write(rw, module, http.StatusOK)
	})
}

func (s *SourceAPI) SetupRoutes() {
	s.Router.StrictSlash(true)
	s.Router.Handle("/{vcsprovider}/{vcsconnid}/{reponame}", s.CreateVCSModule()).Methods(http.MethodPost)
}

func NewSourceAPI(router *mux.Router, path string, vcsconnstore types.VCSSConnectionStore, sourceProviders types.TerrariumSourceDriver, responseHandler types.APIResponseWriter, errorHandler types.APIErrorWriter) *SourceAPI {
	s := &SourceAPI{
		Router:          router.PathPrefix(path).Subrouter(),
		VCSStore:        vcsconnstore,
		SourceStore:     sourceProviders,
		ResponseHandler: responseHandler,
		ErrorHandler:    errorHandler,
	}
	s.SetupRoutes()
	return s
}
