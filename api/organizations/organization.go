// Package organization implements the Terrarium Organization API.
// This package provides one such implementation. Other implementations are encouraged
// via the endpoints package interfaces
package organizations

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/dylanrhysscott/terrarium/pkg/registry/endpoints"
	"github.com/dylanrhysscott/terrarium/pkg/registry/responses"
	"github.com/dylanrhysscott/terrarium/pkg/registry/stores"
	"github.com/gorilla/mux"
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

// NewOrganizationAPI Creates a new instance of the organization API setting up routes as well as any backend storage responses
// and VCS integrations This function will call SetupRoutes() on your behalf to configure the router and associated handlers
func NewOrganizationAPI(router *mux.Router, path string, store stores.OrganizationStore, vcsAPI endpoints.VCSConnAPIInterface, responseHandler responses.APIResponseWriter, errorHandler responses.APIErrorWriter) *OrganizationAPI {
	o := &OrganizationAPI{
		Router:            router.PathPrefix(path).Subrouter(),
		OrganziationStore: store,
		ResponseHandler:   responseHandler,
		ErrorHandler:      errorHandler,
	}
	o.SetupRoutes(vcsAPI)
	return o
}
