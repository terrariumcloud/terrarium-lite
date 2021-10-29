// Package api implements the Terrarium API
package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dylanrhysscott/terrarium/api/oauth"
	"github.com/dylanrhysscott/terrarium/api/organization"
	"github.com/dylanrhysscott/terrarium/api/vcs"
	"github.com/dylanrhysscott/terrarium/pkg/types"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Terrarium is a struct which contains methods for initialising the private Terraform Registry
type Terrarium struct {
	Port            int
	Store           types.TerrariumDriver
	OrganizationAPI organization.OrganizationAPIInterface
	VCSAPI          vcs.VCSAPIInterface
	OAuthAPI        oauth.OAuthAPIInterface
	Router          *mux.Router
	Responder       types.APIResponseWriter
	Errorer         types.APIErrorWriter
}

// Serve starts the Terrarium Registry
func (t *Terrarium) Serve() error {
	bindAddress := fmt.Sprintf(":%d", t.Port)
	log.Println(fmt.Sprintf("Listening on %s", bindAddress))
	return http.ListenAndServe(bindAddress, handlers.CombinedLoggingHandler(os.Stdout, t.Router))
}

// setupOrganizationRoutes configures the organization API subrouter
func (t *Terrarium) setupOrganizationRoutes(path string) {
	s := t.Router.PathPrefix(path).Subrouter()
	s.StrictSlash(true)
	s.Handle("/", t.OrganizationAPI.ListOrganizationsHandler()).Methods(http.MethodGet)
	s.Handle("/", t.OrganizationAPI.CreateOrganizationHandler()).Methods(http.MethodPost)
	s.Handle("/{organization_name}", t.OrganizationAPI.GetOrganizationHandler()).Methods(http.MethodGet)
	s.Handle("/{organization_name}", t.OrganizationAPI.UpdateOrganizationHandler()).Methods(http.MethodPatch)
	s.Handle("/{organization_name}", t.OrganizationAPI.DeleteOrganizationHandler()).Methods(http.MethodDelete)
	s.Handle("/{organization_name}/oauth-clients", t.VCSAPI.ListVCSHandler()).Methods(http.MethodGet)
	s.Handle("/{organization_name}/oauth-clients", t.VCSAPI.CreateVCSHandler()).Methods(http.MethodPost)
}

// setupOAuthClientRoutes configures the VCS API subrouter
func (t *Terrarium) setupOAuthClientRoutes(path string) {
	s := t.Router.PathPrefix(path).Subrouter()
	s.StrictSlash(true)
	s.Handle("/{id}", t.VCSAPI.GetVCSHandler()).Methods(http.MethodGet)
	s.Handle("/{id}", t.VCSAPI.UpdateVCSHandler()).Methods(http.MethodPatch)
	s.Handle("/{id}", t.VCSAPI.DeleteVCSHandler()).Methods(http.MethodDelete)
}

func (t *Terrarium) setupAuthorizeRoutes(path string) {
	oauthHandlers := oauth.NewOAuthAPI(t.Store.VCS(), t.Responder, t.Errorer)
	s := t.Router.PathPrefix(path).Subrouter()
	s.StrictSlash(true)
	s.Handle("/github/{id}/callback", oauthHandlers.GithubCallbackHandler()).Methods(http.MethodGet)
}

// NewTerrarium creates a new Terrarium instance setting up the required API routes
func NewTerrarium(port int, driver types.TerrariumDriver, responder types.APIResponseWriter, errorer types.APIErrorWriter) *Terrarium {
	t := &Terrarium{
		Port:      port,
		Store:     driver,
		Router:    mux.NewRouter(),
		Responder: responder,
		Errorer:   errorer,
	}
	t.OrganizationAPI = organization.NewOrganizationAPI(t.Store.Organizations(), t.Responder, t.Errorer)
	t.VCSAPI = vcs.NewVCSAPI(t.Store.VCS(), t.Store.Organizations(), t.Responder, t.Errorer)
	t.OAuthAPI = oauth.NewOAuthAPI(t.Store.VCS(), t.Responder, t.Errorer)
	t.setupOrganizationRoutes("/v1/organizations")
	t.setupOAuthClientRoutes("/v1/oauth-clients")
	t.setupAuthorizeRoutes("/oauth")
	return t
}
