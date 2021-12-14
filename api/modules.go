package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/dylanrhysscott/terrarium/pkg/registry/data/modules"
	"github.com/dylanrhysscott/terrarium/pkg/registry/drivers"
	"github.com/dylanrhysscott/terrarium/pkg/registry/responses"
	"github.com/gorilla/mux"
)

// ModuleAPIInterface specifies the required HTTP handlers for a Terrarium Modules API
type ModuleAPIInterface interface {
	CreateModuleHandler() http.Handler
	GetModuleHandler() http.Handler
	GetModuleVersionHandler() http.Handler
	DownloadModuleHandler() http.Handler
	ArchiveHandler() http.Handler
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
	FileStore       drivers.TerrariumStorageDriver
	ErrorHandler    responses.APIErrorWriter
	ResponseHandler responses.APIResponseWriter
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
	// TODO: Make this dynamic and ensure header handling is compliant
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("X-Terraform-Get", "./archive?archive=zip")
		m.ResponseHandler.Write(rw, nil, http.StatusNoContent)
	})
}

func (m *ModuleAPI) GetModuleVersionHandler() http.Handler {
	// TODO: Make this dynamic
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		vr := &modules.ModuleVersionResponse{
			Modules: []*modules.ModuleVersions{
				{
					Versions: []*modules.ModuleVersionItem{
						{
							Version: "0.0.1",
						},
					},
				},
			},
		}
		data, _ := json.Marshal(vr)
		rw.Header().Add("Content-Type", "application/json")
		rw.Write(data)
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

func (m *ModuleAPI) ArchiveHandler() http.Handler {
	// TODO: Make this flexible to fetch archives dynamically
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		zipData, err := m.FileStore.FetchModuleSource(context.TODO(), "terrarium-dev", "test.zip")
		if err != nil {
			log.Printf("[FILE STORE] Error: %s", err.Error())
			m.ErrorHandler.Write(rw, errors.New("failed fetching module source from file store"), http.StatusInternalServerError)
			return
		}
		r.Header.Set("Content-Type", "application/zip")
		rw.Write(zipData)
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
	m.Router.Handle("/{organization_name}/{name}/{provider}/{version}/archive", m.ArchiveHandler()).Methods(http.MethodGet)

}

func NewModuleAPI(router *mux.Router, path string, store interface{}, fileStore drivers.TerrariumStorageDriver, responseHandler responses.APIResponseWriter, errorHandler responses.APIErrorWriter) *ModuleAPI {
	m := &ModuleAPI{
		Router:          router.PathPrefix(path).Subrouter(),
		ModuleStore:     store,
		FileStore:       fileStore,
		ErrorHandler:    errorHandler,
		ResponseHandler: responseHandler,
	}
	m.SetupRoutes()
	return m
}
