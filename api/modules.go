package api

import (
	"net/http"

	"github.com/dylanrhysscott/terrarium/pkg/types"
	"github.com/gorilla/mux"
)

// ModuleAPIInterface specifies the required HTTP handlers for a Terrarium Modules API
type ModuleAPIInterface interface {
	CreateModuleHandler() http.Handler
	GetModuleHandler() http.Handler
	GetModuleVersionHandler() http.Handler
	DownloadModuleHandler() http.Handler
	UpdateModuleHandler() http.Handler
	ListModulesHandler() http.Handler
	ListOrganizationModulesHandler() http.Handler
	DeleteModuleHandler() http.Handler
}

// ModuleAPI is a struct implementing the handlers for the Module API in Terrarium
// TODO: Implement Module Store
type ModuleAPI struct {
	Router          *mux.Router
	ModuleStore     interface{}
	ErrorHandler    types.APIErrorWriter
	ResponseHandler types.APIResponseWriter
}

func (m *ModuleAPI) CreateModuleHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

func (m *ModuleAPI) GetModuleHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

func (m *ModuleAPI) DownloadModuleHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

func (m *ModuleAPI) GetModuleVersionHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

func (m *ModuleAPI) UpdateModuleHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

func (m *ModuleAPI) ListModulesHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

func (m *ModuleAPI) ListOrganizationModulesHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

func (m *ModuleAPI) DeleteModuleHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

func (m *ModuleAPI) SetupRoutes() {
	m.Router.StrictSlash(true)
	m.Router.Handle("/", m.ListModulesHandler()).Methods(http.MethodGet)
	m.Router.Handle("/{organization_name}", m.ListOrganizationModulesHandler()).Methods(http.MethodGet)
	m.Router.Handle("/", m.CreateModuleHandler()).Methods(http.MethodPost)
	m.Router.Handle("/{organization_name}/{name}/{provider}", m.GetModuleHandler()).Methods(http.MethodGet)
	m.Router.Handle("/{organization_name}/{name}/{provider}/versions", m.GetModuleVersionHandler()).Methods(http.MethodGet)
	m.Router.Handle("/{organization_name}/{name}/{provider}/{version}/download", m.DownloadModuleHandler()).Methods(http.MethodGet)
	m.Router.Handle("/{organization_name}/{name}/{provider}", m.UpdateModuleHandler()).Methods(http.MethodPatch)
	m.Router.Handle("/{organization_name}/{name}/{provider}", m.DeleteModuleHandler()).Methods(http.MethodDelete)
}

func NewModuleAPI(router *mux.Router, path string, store interface{}, responseHandler types.APIResponseWriter, errorHandler types.APIErrorWriter) *ModuleAPI {
	m := &ModuleAPI{
		Router:          router,
		ModuleStore:     store,
		ErrorHandler:    errorHandler,
		ResponseHandler: responseHandler,
	}
	m.SetupRoutes()
	return m
}
