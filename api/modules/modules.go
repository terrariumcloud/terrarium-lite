// Package modules implements the Terrarium Modules API and Terraform Registry Protocol
// https://www.terraform.io/internals/module-registry-protocol.
package modules

import (
	"github.com/gorilla/mux"
	"github.com/terrariumcloud/terrarium-lite/pkg/registry/drivers"
	"github.com/terrariumcloud/terrarium-lite/pkg/registry/responses"
	"github.com/terrariumcloud/terrarium-lite/pkg/registry/stores"
)

// NewModuleAPI Creates a new instance of the module API setting up routes as well as any backend storage and responses.
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
