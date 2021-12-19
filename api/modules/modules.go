// Package modules implements the Terrarium Modules API and Terraform Registry Protocol
// https://www.terraform.io/internals/module-registry-protocol. This package provides one such implementation
// Other implementations are encouraged via the endpoints package interfaces
package modules

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/terrariumcloud/terrarium/pkg/registry/drivers"
	"github.com/terrariumcloud/terrarium/pkg/registry/responses"
	"github.com/terrariumcloud/terrarium/pkg/registry/stores"
)

// extractLimitAndOffset is a convience method to extract pagination and limit values passed
// to a handler that supports pages. For internal module use only
// TODO: Remove duplicated function between API packages
func extractLimitAndOffset(qs url.Values) (int, int, error) {
	var limit int = 10
	var offset int = 0
	if qs.Has("limit") {
		// If we have a limit value set in QS attempt to convert to int
		i, err := strconv.Atoi(qs.Get("limit"))
		if err != nil {
			return 0, 0, errors.New("limit is not an integer")
		}
		limit = i
	}
	if qs.Has("offset") {
		i, err := strconv.Atoi(qs.Get("offset"))
		if err != nil {
			return 0, 0, errors.New("offset is not an integer")
		}
		offset = i
	}
	return limit, offset, nil
}

// NewModuleAPI Creates a new instance of the module API setting up routes as well as any backend storage and responses
// This function will call SetupRoutes() on your behalf to configure the router and associated handlers
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
