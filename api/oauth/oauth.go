// Package oauth provides API routing for handling OAuth login flows via terraform login. Many OAuth
// provider handlers maybe added here in future. Currently this package only supports Github.
// Note: This package is currently a WIP and has no guarantee to work correctly when running
// the login API.
package oauth

import (
	"github.com/dylanrhysscott/terrarium/pkg/registry/responses"
	"github.com/dylanrhysscott/terrarium/pkg/registry/stores"

	"github.com/gorilla/mux"
)

// NewModuleAPI Creates a new instance of the OAuth API setting up routes for various support
// OAuth provider callbacks. It is backed by a VCS connection store which contains OAuth client ID's and secrets
// This function will call SetupRoutes() on your behalf to configure the router and associated handlers
// used in the authentication flow.
func NewOAuthAPI(router *mux.Router, path string, vcsconnstore stores.VCSSConnectionStore, responseHandler responses.APIResponseWriter, errorHandler responses.APIErrorWriter) *OAuthAPI {
	a := &OAuthAPI{
		Router:          router.PathPrefix(path).Subrouter(),
		VCSStore:        vcsconnstore,
		ResponseHandler: responseHandler,
		ErrorHandler:    errorHandler,
	}
	a.SetupRoutes()
	return a
}
