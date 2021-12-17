// Package api is a meta package implementing the complete Terrarium product.
// Terrarium aims to implement the protocols provided by Terraform for a terraform cli compliant registry experience.
// The meta package is made up of sub packages to subdivide and organize code by function. For detailed explanations
// on various API's please review these
package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dylanrhysscott/terrarium/api/discovery"
	"github.com/dylanrhysscott/terrarium/api/modules"
	"github.com/dylanrhysscott/terrarium/api/oauth"
	"github.com/dylanrhysscott/terrarium/api/organizations"
	"github.com/dylanrhysscott/terrarium/api/sources"
	"github.com/dylanrhysscott/terrarium/api/vcs"

	"github.com/dylanrhysscott/terrarium/internal/endpoints"
	"github.com/dylanrhysscott/terrarium/pkg/registry/drivers"
	"github.com/dylanrhysscott/terrarium/pkg/registry/responses"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Terrarium is a struct which contains methods for initialising the private Terraform Registry
// The Terrarium struct is a complete implementation of the product fully instantiated. An instance
// of this struct is created by the CLI when `terrarium serve modules` is called from the command line
type Terrarium struct {
	Port             int
	DataStore        drivers.TerrariumDatabaseDriver
	FileStore        drivers.TerrariumStorageDriver
	Source           drivers.TerrariumSourceDriver
	OrganizationAPI  endpoints.OrganizationAPIInterface
	ModuleAPI        endpoints.ModuleAPIInterface
	VCSConnectionAPI endpoints.VCSConnAPIInterface
	OAuthAPI         endpoints.OAuthAPIInterface
	SourceAPI        endpoints.SourceAPIInterface
	DiscoveryAPI     endpoints.DiscoveryAPIInterface
	Router           *mux.Router
	Responder        responses.APIResponseWriter
	Errorer          responses.APIErrorWriter
}

// Serve starts the Terrarium Registry listening on the specified port. A web server will be listening ready to
// serve API requests
func (t *Terrarium) Serve() error {
	bindAddress := fmt.Sprintf(":%d", t.Port)
	t.Init()
	log.Println(fmt.Sprintf("Listening on %s", bindAddress))
	return http.ListenAndServe(bindAddress, handlers.CombinedLoggingHandler(os.Stdout, t.Router))
}

// Init calls the various API sub packages to setup routers for endpoints. This is a central function that wires all API routers together
// providing OAuth, VCS, Discovery and many more features of Terrarium
func (t *Terrarium) Init() {
	t.VCSConnectionAPI = vcs.NewVCSAPI(t.Router, "/v1/oauth-clients", t.DataStore.VCSConnections(), t.DataStore.Organizations(), t.Responder, t.Errorer)
	t.OAuthAPI = oauth.NewOAuthAPI(t.Router, "/oauth", t.DataStore.VCSConnections(), t.Responder, t.Errorer)
	t.SourceAPI = sources.NewSourceAPI(t.Router, "/v1/sources", t.DataStore.VCSConnections(), t.Source, t.Responder, t.Errorer)
	t.OrganizationAPI = organizations.NewOrganizationAPI(t.Router, "/v1/organizations", t.DataStore.Organizations(), t.VCSConnectionAPI, t.Responder, t.Errorer)
	t.ModuleAPI = modules.NewModuleAPI(t.Router, "/v1/modules", t.DataStore.Modules(), t.FileStore, t.Responder, t.Errorer)
	// TODO: Should this be it's own binary / sub command?
	t.DiscoveryAPI = discovery.NewDiscoveryAPI(nil, "/v1/modules", t.Responder, t.Errorer)
	t.Router.Handle("/.well-known/terraform.json", t.DiscoveryAPI.DiscoveryHandler())
}

// NewTerrarium creates a new Terrarium instance setting up the required API routes
func NewTerrarium(port int, driver drivers.TerrariumDatabaseDriver, storageDriver drivers.TerrariumStorageDriver, sourceDriver drivers.TerrariumSourceDriver, responder responses.APIResponseWriter, errorer responses.APIErrorWriter) *Terrarium {
	return &Terrarium{
		Port:      port,
		DataStore: driver,
		FileStore: storageDriver,
		Source:    sourceDriver,
		Router:    mux.NewRouter(),
		Responder: responder,
		Errorer:   errorer,
	}
}
