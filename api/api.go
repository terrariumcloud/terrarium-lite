// Package api implements the Terrarium API
package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dylanrhysscott/terrarium/api/organization"
	"github.com/dylanrhysscott/terrarium/pkg/types"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Terrarium is a struct which contains methods for initialising the private Terraform Registry
type Terrarium struct {
	Port      int
	Store     types.TerrariumDriver
	Router    *mux.Router
	Responder types.APIResponseWriter
	Errorer   types.APIErrorWriter
}

// Serve starts the Terrarium Registry
func (t *Terrarium) Serve() error {
	bindAddress := fmt.Sprintf(":%d", t.Port)
	log.Println(fmt.Sprintf("Listening on %s", bindAddress))
	return http.ListenAndServe(bindAddress, handlers.CombinedLoggingHandler(os.Stdout, t.Router))
}

// setupOrganizationRoutes configures the organization API subrouter
func (t *Terrarium) setupOrganizationRoutes(path string) {
	apiHandlers := organization.NewOrganizationAPI(t.Store.Organizations(), t.Responder, t.Errorer)
	s := t.Router.PathPrefix(path).Subrouter()
	s.StrictSlash(true)
	s.Handle("/", apiHandlers.ListOrganizationsHandler()).Methods(http.MethodGet)
	s.Handle("/", apiHandlers.CreateOrganizationHandler()).Methods(http.MethodPost)
	s.Handle("/{id}", apiHandlers.GetOrganizationHandler()).Methods(http.MethodGet)
	s.Handle("/{id}", apiHandlers.UpdateOrganizationHandler()).Methods(http.MethodPatch)
	s.Handle("/{id}", apiHandlers.DeleteOrganizationHandler()).Methods(http.MethodDelete)
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
	t.setupOrganizationRoutes("/v1/organizations")
	return t
}