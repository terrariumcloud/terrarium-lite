package organization

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dylanrhysscott/terrarium/pkg/types"
	"gopkg.in/errgo.v2/errors"
)

// OrganizationAPIInterface specifies the required HTTP handlers for a Terrarium Organizations API
type OrganizationAPIInterface interface {
	CreateOrganizationHandler() http.Handler
	GetOrganizationHandler() http.Handler
	UpdateOrganizationHandler() http.Handler
	ListOrganizationsHandler() http.Handler
	DeleteOrganizationHandler() http.Handler
}

// OrganizationAPI is a struct implementing the handlers for the Organization API in Terrarium
type OrganizationAPI struct {
	Path              string
	OrganziationStore types.OrganizationStore
}

// CreateOrganizationHandler is a handler for creating an organization (POST)
func (o *OrganizationAPI) CreateOrganizationHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

// UpdateOrganizationHandler is a handler for updating an organization (PUT)
func (o *OrganizationAPI) UpdateOrganizationHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

// GetOrganizationHandler is a handler for getting a single organization (GET)
func (o *OrganizationAPI) GetOrganizationHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

// ListOrganizationsHandler is a handler for listing all organizations (GET)
func (o *OrganizationAPI) ListOrganizationsHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		qs := r.URL.Query()
		limit := 0
		offset := 0
		if qs.Has("limit") {
			// If we have a limit value set in QS attempt to convert to int
			i, err := strconv.Atoi(qs.Get("limit"))
			if err != nil {
				// limit is not an int generate a bad quest error
				jsonData, err := types.NewTerrariumBadRequest("limit query parameter must be a whole number")
				if err != nil {
					// Something went wrong with generating response - return a server error and log stack trace
					log.Println(err.Error())
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
				// Responded with 400 bad request
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write(jsonData)
				return
			}
			limit = i
		}
		if qs.Has("offset") {
			i, err := strconv.Atoi(qs.Get("offset"))
			if err != nil {
				jsonData, err := types.NewTerrariumBadRequest("offset query parameter must be a whole number")
				if err != nil {
					log.Println(err.Error())
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write(jsonData)
				return
			}
			offset = i
		}
		// Query organization store for all orgs using limit and offset
		orgs, err := o.OrganziationStore.ReadAll(limit, offset)
		if err != nil {
			// Something went wrong return a 500 to the user
			data, err := types.NewTerrariumServerError(err.Error())
			if err != nil {
				// Something went wrong further return a 500 and stack trace
				log.Printf("%+v", errors.Wrap(err))
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(data)
			return
		}
		// We have the orgs now return a 200
		data, err := types.NewTerrariumOK(orgs)
		if err != nil {
			data, err := types.NewTerrariumServerError(err.Error())
			if err != nil {
				log.Printf("%+v", errors.Wrap(err))
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(data)
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write(data)
	})
}

// DeleteOrganizationHandler is a handler for deleting an organization (DELETE)
func (o *OrganizationAPI) DeleteOrganizationHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

// NewOrganizationAPI creates an instance of the organization API with the reqired database
// driver support
func NewOrganizationAPI(store types.OrganizationStore) *OrganizationAPI {
	return &OrganizationAPI{
		OrganziationStore: store,
	}
}
