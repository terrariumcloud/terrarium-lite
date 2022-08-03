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

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/terrariumcloud/terrarium/api/discovery"
	"github.com/terrariumcloud/terrarium/api/modules"
	"github.com/terrariumcloud/terrarium/internal/endpoints"
	"github.com/terrariumcloud/terrarium/pkg/registry/drivers"
	"github.com/terrariumcloud/terrarium/pkg/registry/responses"
)

// Terrarium is a struct which contains methods for initialising the private Terraform Registry
// The Terrarium struct is a complete implementation of the product fully instantiated. An instance
// of this struct is created by the CLI when `terrarium serve modules` is called from the command line
type Terrarium struct {
	Port         int
	CertFile     string
	KeyFile      string
	DataStore    drivers.TerrariumDatabaseDriver
	FileStore    drivers.TerrariumStorageDriver
	ModuleAPI    endpoints.ModuleAPIInterface
	DiscoveryAPI endpoints.DiscoveryAPIInterface
	Router       *mux.Router
	Responder    responses.APIResponseWriter
	Errorer      responses.APIErrorWriter
}

// Serve starts the Terrarium Registry listening on the specified port. A web server will be listening ready to
// serve API requests
func (t *Terrarium) Serve() error {
	bindAddress := fmt.Sprintf(":%d", t.Port)
	t.Init()
	log.Println(fmt.Sprintf("Listening on %s", bindAddress))
	return http.ListenAndServeTLS(bindAddress, t.CertFile, t.KeyFile, handlers.CombinedLoggingHandler(os.Stdout, t.Router))
}

// Init calls the various API sub packages to setup routers for endpoints. This is a central function that wires all API routers together
func (t *Terrarium) Init() {
	t.ModuleAPI = modules.NewModuleAPI(t.Router, "/v1/modules", t.DataStore.Modules(), t.FileStore, t.Responder, t.Errorer)
	// TODO: Should this be it's own binary / sub command?
	t.DiscoveryAPI = discovery.NewDiscoveryAPI(nil, "/v1/modules", t.Responder, t.Errorer)
	t.Router.Handle("/.well-known/terraform.json", t.DiscoveryAPI.DiscoveryHandler())
}

// NewTerrarium creates a new Terrarium instance setting up the required API routes
func NewTerrarium(port int, certFile string, keyFile string, driver drivers.TerrariumDatabaseDriver, storageDriver drivers.TerrariumStorageDriver, responder responses.APIResponseWriter, errorer responses.APIErrorWriter) *Terrarium {
	return &Terrarium{
		Port:      port,
		CertFile:  certFile,
		KeyFile:   keyFile,
		DataStore: driver,
		FileStore: storageDriver,
		Router:    mux.NewRouter(),
		Responder: responder,
		Errorer:   errorer,
	}
}
