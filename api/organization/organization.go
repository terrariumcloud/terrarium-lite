package organization

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dylanrhysscott/terrarium/pkg/types"
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
			i, err := strconv.Atoi(qs.Get("limit"))
			if err != nil {
				jsonData, err := types.NewTerrariumBadRequest("limit query parameter must be a whole number")
				if err != nil {
					log.Println(err.Error())
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
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
		orgs, err := o.OrganziationStore.ReadAll(limit, offset)
		if err != nil {
			data, err := types.NewTerrariumServerError(err.Error())
			if err != nil {
				log.Printf("Error generating a server response - %s", err.Error())
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(data)
			return
		}
		data, err := types.NewTerrariumOK(orgs)
		if err != nil {
			data, err := types.NewTerrariumServerError(err.Error())
			if err != nil {
				log.Printf("Error generating a server response - %s", err.Error())
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
