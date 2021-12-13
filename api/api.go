// Package api implements the Terrarium API
package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dylanrhysscott/terrarium/internal/terrariumsource"
	"github.com/dylanrhysscott/terrarium/pkg/types"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Terrarium is a struct which contains methods for initialising the private Terraform Registry
type Terrarium struct {
	Port             int
	Store            types.TerrariumDatabaseDriver
	Source           types.TerrariumSourceDriver
	OrganizationAPI  OrganizationAPIInterface
	ModuleAPI        ModuleAPIInterface
	VCSConnectionAPI VCSConnAPIInterface
	OAuthAPI         OAuthAPIInterface
	SourceAPI        SourceAPIInterface
	DiscoveryAPI     DiscoveryAPIInterface
	Router           *mux.Router
	Responder        types.APIResponseWriter
	Errorer          types.APIErrorWriter
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
	t.VCSConnectionAPI = NewVCSAPI(t.Router, "/v1/oauth-clients", t.Store.VCSConnections(), t.Store.Organizations(), t.Responder, t.Errorer)
	t.OAuthAPI = NewOAuthAPI(t.Router, "/oauth", t.Store.VCSConnections(), t.Responder, t.Errorer)
	t.SourceAPI = NewSourceAPI(t.Router, "/v1/sources", t.Store.VCSConnections(), t.Source, t.Responder, t.Errorer)
	t.OrganizationAPI = NewOrganizationAPI(t.Router, "/v1/organizations", t.Store.Organizations(), t.VCSConnectionAPI, t.Responder, t.Errorer)
	t.ModuleAPI = NewModuleAPI(t.Router, "/v1/modules", nil, t.Responder, t.Errorer)
	// TODO: Should this be it's own binary / sub command?
	t.DiscoveryAPI = NewDiscoveryAPI(nil, "/v1/modules", t.Responder, t.Errorer)
	// Register discovery route for Terraform native discovery
	t.Router.Handle("/.well-known/terraform.json", t.DiscoveryAPI.DiscoveryHandler())
}

// NewTerrarium creates a new Terrarium instance setting up the required API routes
func NewTerrarium(port int, driver types.TerrariumDatabaseDriver, responder types.APIResponseWriter, errorer types.APIErrorWriter) *Terrarium {
	return &Terrarium{
		Port:      port,
		Store:     driver,
		Source:    terrariumsource.NewTerrariumSourceDriver(),
		Router:    mux.NewRouter(),
		Responder: responder,
		Errorer:   errorer,
	}
}
