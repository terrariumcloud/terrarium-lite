package modules

import (
	"net/http"

	"github.com/dylanrhysscott/terrarium/pkg/types"
	"github.com/gorilla/mux"
)

// ModuleAPIInterface specifies the required HTTP handlers for a Terrarium Modules API
type ModuleAPIInterface interface {
	CreateOrganizationHandler() http.Handler
	GetOrganizationHandler() http.Handler
	UpdateOrganizationHandler() http.Handler
	ListOrganizationsHandler() http.Handler
	DeleteOrganizationHandler() http.Handler
}

// ModuleAPI is a struct implementing the handlers for the Module API in Terrarium
type ModuleAPI struct {
	Router          *mux.Router
	ModuleStore     types.OrganizationStore
	ErrorHandler    types.APIErrorWriter
	ResponseHandler types.APIResponseWriter
}

func NewModulesAPI(router *mux.Router, path string, store types.OrganizationStore, responseHandler types.APIResponseWriter, errorHandler types.APIErrorWriter) *ModuleAPI {
	return &ModuleAPI{
		Router:      router,
		ModuleStore: nil,
	}
}
