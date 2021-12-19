package modules

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/terrariumcloud/terrarium/pkg/registry/data/modules"
	"github.com/terrariumcloud/terrarium/pkg/registry/drivers"
	"github.com/terrariumcloud/terrarium/pkg/registry/responses"
	"github.com/terrariumcloud/terrarium/pkg/registry/stores"
)

// ModuleAPI is a struct implementing the handlers for the ModuleAPIInterface from the endpoints package in Terrarium
type ModuleAPI struct {
	Router          *mux.Router
	ModuleStore     stores.ModuleStore
	FileStore       drivers.TerrariumStorageDriver
	ErrorHandler    responses.APIErrorWriter
	ResponseHandler responses.APIResponseWriter
}

// TODO: Not yet implemented
func (m *ModuleAPI) CreateModuleHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

// GetModuleHandler Handles an API request to fetch a single module via organization name, module name and provider.
// This will return the module from the backend if found or 404 if the module or organization doesn't exist
func (m *ModuleAPI) GetModuleHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		orgName := params["organization_name"]
		moduleName := params["name"]
		providerName := params["provider"]
		module, err := m.ModuleStore.ReadOne(orgName, moduleName, providerName)
		if err != nil {
			m.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		if module == nil {
			m.ErrorHandler.Write(rw, errors.New("module not found"), http.StatusNotFound)
			return
		}
		m.ResponseHandler.Write(rw, module, http.StatusOK)
	})
}

// DownloadModuleHandler will return a header indicating where the requesting CLI can download module content from
// This handler complies with the following implementation from the module protocol
// https://www.terraform.io/internals/module-registry-protocol#download-source-code-for-a-specific-module-version
func (m *ModuleAPI) DownloadModuleHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("X-Terraform-Get", "./archive?archive=zip")
		m.ResponseHandler.Write(rw, nil, http.StatusNoContent)
	})
}

// GetModuleVersionHandler will return a list of available versions for a given module.
// This signifies to the requesting CLI if that module is available to consume from the registry.
// Will return a 404 if a non existent organization and/or module is requested.
// This handler complies with the following implementation from the module protocol
// https://www.terraform.io/internals/module-registry-protocol#download-source-code-for-a-specific-module-version
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
		if moduleItems == nil {
			m.ErrorHandler.Write(rw, errors.New("organization does not exist"), http.StatusNotFound)
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

// TODO: Not yet implemented
func (m *ModuleAPI) UpdateModuleHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

// ListModulesHandler Lists all modules currently available in the registry.
// This is a convience handler to allow for module browsing and discovery
// Will return an error if communicating with the backend failed
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

// ListModulesHandler Lists all modules for a given organization available in the registry.
// This is a convience handler to allow for module browsing and discovery
// Will return an error if communicating with the backend failed or a 404 not found if the organization does not exist
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
		if modules == nil {
			m.ErrorHandler.Write(rw, errors.New("organization does not exist"), http.StatusNotFound)
			return
		}
		m.ResponseHandler.Write(rw, modules, http.StatusOK)
	})
}

// TODO: Not yet implemented
func (m *ModuleAPI) DeleteModuleHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

// ArchiveHandler performs a fetch of the requested module source code from the chosen backing store and presents it to the client
// As part of the module flow clients are redirected here from the DownloadModuleHandler x-terraform-get header. This handler
// makes the stored registry code available to the client
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
		zipData, err := m.FileStore.FetchModuleSource(context.TODO(), key)
		if err != nil {
			log.Printf("[FILE STORE] Error: %s", err.Error())
			m.ErrorHandler.Write(rw, errors.New("failed fetching module source from file store"), http.StatusInternalServerError)
			return
		}
		r.Header.Set("Content-Type", "application/zip")
		rw.Write(zipData)
	})
}

// SetupRoutes Sets up the various endpoints for the modules API by registering handlers from this struct to it's
// corresponding routes. This will register the routes required by the module registry protocol as defined here
// https://www.terraform.io/internals/module-registry-protocol Additional routes not part of the specification will also be registered.
// These are used to provide additional functionality for a more complete registry experience
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
