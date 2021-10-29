package sources

import (
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
}

func (s *SourceAPI) CreateVCSModule() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

func (s *SourceAPI) SetupRoutes() {
	s.Router.StrictSlash(true)
	s.Router.Handle("/{id}", s.CreateVCSModule()).Methods(http.MethodGet)
}

func NewSourceAPI(router *mux.Router, path string, vcsstore types.VCSStore, responseHandler types.APIResponseWriter, errorHandler types.APIErrorWriter) *SourceAPI {
	s := &SourceAPI{
		Router:          router.PathPrefix(path).Subrouter(),
		VCSStore:        vcsstore,
		ResponseHandler: responseHandler,
		ErrorHandler:    errorHandler,
	}
	s.SetupRoutes()
	return s
}
