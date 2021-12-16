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
	"github.com/dylanrhysscott/terrarium/pkg/registry/stores"
	"github.com/gorilla/mux"
)

// ModuleAPI is a struct implementing the handlers for the Module API in Terrarium
type ModuleAPI struct {
	Router          *mux.Router
	ModuleStore     stores.ModuleStore
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
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("X-Terraform-Get", "./archive?archive=zip")
		m.ResponseHandler.Write(rw, nil, http.StatusNoContent)
	})
}

func (m *ModuleAPI) GetModuleVersionHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		orgName := params["organization_name"]
		moduleName := params["name"]
		providerName := params["provider"]
		moduleItems, err := m.ModuleStore.ReadModuleVersions(orgName, moduleName, providerName)
		if err != nil {
			m.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		if len(moduleItems) == 0 {
			m.ErrorHandler.Write(rw, errors.New("module not found"), http.StatusNotFound)
			return
		}
		var versions []*modules.ModuleVersionItem
		for _, moduleItem := range moduleItems {
			versions = append(versions, &modules.ModuleVersionItem{
				Version: moduleItem.Version,
			})
		}
		vr := &modules.ModuleVersionResponse{
			Modules: []*modules.ModuleVersions{
				{
					Versions: versions,
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
		limit, offset, err := extractLimitAndOffset(r.URL.Query())
		if err != nil {
			m.ErrorHandler.Write(rw, err, http.StatusBadRequest)
			return
		}
		modules, err := m.ModuleStore.ReadAll(limit, offset)
		if err != nil {
			m.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		m.ResponseHandler.Write(rw, modules, http.StatusOK)
	})
}

func (m *ModuleAPI) ListOrganizationModulesHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		orgName := params["organization_name"]
		limit, offset, err := extractLimitAndOffset(r.URL.Query())
		if err != nil {
			m.ErrorHandler.Write(rw, err, http.StatusBadRequest)
			return
		}
		modules, err := m.ModuleStore.ReadOrganizationModules(orgName, limit, offset)
		if err != nil {
			m.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		m.ResponseHandler.Write(rw, modules, http.StatusOK)
	})
}

func (m *ModuleAPI) DeleteModuleHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

func (m *ModuleAPI) ArchiveHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		orgName := params["organization_name"]
		moduleName := params["name"]
		providerName := params["provider"]
		version := params["version"]
		key, err := m.ModuleStore.ReadModuleVersionSource(orgName, moduleName, providerName, version)
		if err != nil {
			log.Printf("[FILE STORE] Error: %s", err.Error())
			m.ErrorHandler.Write(rw, errors.New("failed finding module source"), http.StatusInternalServerError)
			return
		}
		zipData, err := m.FileStore.FetchModuleSource(context.TODO(), m.FileStore.GetBackingStoreName(), key)
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

func NewModuleAPI(router *mux.Router, path string, store stores.ModuleStore, fileStore drivers.TerrariumStorageDriver, responseHandler responses.APIResponseWriter, errorHandler responses.APIErrorWriter) *ModuleAPI {
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
