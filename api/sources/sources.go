package sources

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dylanrhysscott/terrarium/pkg/registry/data/relationships"
	"github.com/dylanrhysscott/terrarium/pkg/registry/data/vcs"
	"github.com/dylanrhysscott/terrarium/pkg/registry/drivers"
	"github.com/dylanrhysscott/terrarium/pkg/registry/responses"
	"github.com/dylanrhysscott/terrarium/pkg/registry/sources"
	"github.com/dylanrhysscott/terrarium/pkg/registry/stores"
	"github.com/gorilla/mux"
)

type SourceAPI struct {
	Router          *mux.Router
	ErrorHandler    responses.APIErrorWriter
	ResponseHandler responses.APIResponseWriter
	VCSStore        stores.VCSSConnectionStore
	SourceStore     drivers.TerrariumSourceDriver
}

func sourceToModule(data sources.SourceData, provider string, orgLink *relationships.ResourceLink, vcsConnectionID string) *vcs.VCSModule {
	return &vcs.VCSModule{
		Name:        data.GetRepoName(),
		Provider:    provider,
		Description: data.GetRepoDescription(),
		VCSConnection: &relationships.ResourceLink{
			ID:   vcsConnectionID,
			Link: fmt.Sprintf("/v1/oauth-clients/%s", vcsConnectionID),
		},
		Organization: orgLink,
		VCSRepo:      data,
	}
}

func (s *SourceAPI) detectStoreType(t string) (sources.SourceProvider, error) {
	var genericStore sources.SourceProvider = nil
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
		vcsItem, err := s.VCSStore.ReadOne(vcsConnID, true)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				s.ErrorHandler.Write(rw, errors.New("vcs provider does not exist"), http.StatusNotFound)
				return
			}
			s.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		if vcsItem.OAuth.ServiceProvider != vcsProvider {
			s.ErrorHandler.Write(rw, errors.New("vcs provider mismatch"), http.StatusBadRequest)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.ErrorHandler.Write(rw, err, http.StatusUnprocessableEntity)
			return
		}
		reqBody := &vcs.SourceVCSRepoBody{}
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
		sourceRepo, err := genericStore.FetchVCSSource(vcsItem.OAuth.Token.AccessToken, repoName)
		if err != nil {
			s.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		module := sourceToModule(sourceRepo, reqBody.Provider, vcsItem.Organization, vcsConnID)
		s.ResponseHandler.Write(rw, module, http.StatusOK)
	})
}

func (s *SourceAPI) SetupRoutes() {
	s.Router.StrictSlash(true)
	s.Router.Handle("/{vcsprovider}/{vcsconnid}/{reponame}", s.CreateVCSModule()).Methods(http.MethodPost)
}

func NewSourceAPI(router *mux.Router, path string, vcsconnstore stores.VCSSConnectionStore, sourceProviders drivers.TerrariumSourceDriver, responseHandler responses.APIResponseWriter, errorHandler responses.APIErrorWriter) *SourceAPI {
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
