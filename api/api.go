// Package api implements the Terrarium API
package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dylanrhysscott/terrarium/api/oauth"
	"github.com/dylanrhysscott/terrarium/api/organization"
	"github.com/dylanrhysscott/terrarium/api/sources"
	"github.com/dylanrhysscott/terrarium/api/vcs"
	"github.com/dylanrhysscott/terrarium/internal/terrariumvcs/github"
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
	VCSProviders    *sources.SourcesMap
	OAuthAPI        oauth.OAuthAPIInterface
	SourceAPI       sources.SourceAPIInterface
	Router          *mux.Router
	Responder       types.APIResponseWriter
	Errorer         types.APIErrorWriter
}

// Serve starts the Terrarium Registry
func (t *Terrarium) Serve() error {
	bindAddress := fmt.Sprintf(":%d", t.Port)
	t.Init()
	log.Println(fmt.Sprintf("Listening on %s", bindAddress))
	return http.ListenAndServe(bindAddress, handlers.CombinedLoggingHandler(os.Stdout, t.Router))
}

// Init calls the various API sub packages to setup routers for endpoints
func (t *Terrarium) Init() {
	t.VCSAPI = vcs.NewVCSAPI(t.Router, "/v1/oauth-clients", t.Store.VCSConnections(), t.Store.Organizations(), t.Responder, t.Errorer)
	t.OAuthAPI = oauth.NewOAuthAPI(t.Router, "/oauth", t.Store.VCSConnections(), t.Responder, t.Errorer)
	t.SourceAPI = sources.NewSourceAPI(t.Router, "/v1/sources", t.Store.VCSConnections(), t.VCSProviders, t.Responder, t.Errorer)
	t.OrganizationAPI = organization.NewOrganizationAPI(t.Router, "/v1/organizations", t.Store.Organizations(), t.VCSAPI, t.Responder, t.Errorer)
}

// NewTerrarium creates a new Terrarium instance setting up the required API routes
func NewTerrarium(port int, driver types.TerrariumDriver, responder types.APIResponseWriter, errorer types.APIErrorWriter) *Terrarium {
	return &Terrarium{
		Port:      port,
		Store:     driver,
		Router:    mux.NewRouter(),
		Responder: responder,
		Errorer:   errorer,
		VCSProviders: &sources.SourcesMap{
			Github: github.NewGithubBackend(),
		},
	}
}
