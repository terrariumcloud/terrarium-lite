package organizations

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/terrariumcloud/terrarium/internal/endpoints"
	"github.com/terrariumcloud/terrarium/pkg/registry/data/organizations"
	"github.com/terrariumcloud/terrarium/pkg/registry/responses"
	"github.com/terrariumcloud/terrarium/pkg/registry/stores"
)

// OrganizationAPI is a struct implementing the handlers for the OrganizationAPIInterface from the endpoints package in Terrarium
type OrganizationAPI struct {
	Router            *mux.Router
	OrganziationStore stores.OrganizationStore
	ErrorHandler      responses.APIErrorWriter
	ResponseHandler   responses.APIResponseWriter
}

// CreateOrganizationHandler Handles an API request to create a new organization in the registry. It accepts a JSON document
// in the format of an Organization struct in the Terrarium data package. If it is valid the organization will be created or rejected
// as unprocessable. Depending on the backing store of your choice organizations may be unique.
func (o *OrganizationAPI) CreateOrganizationHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		org := &organizations.Organization{}
		err = json.Unmarshal(body, org)
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		err = org.Validate()
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusUnprocessableEntity)
			return
		}
		org, err = o.OrganziationStore.Create(org.Name, org.Email)
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		o.ResponseHandler.Write(rw, org, http.StatusCreated)
	})
}

// UpdateOrganizationHandler Handles an API request to update a existing organization in the registry. It accepts a JSON document
// in the format of an Organization struct from the Terrarium data package. If it is valid the organization will be updated or rejected
// as unprocessable. The organization must exist before attempting to update it. It will not be upserted if it doesn't exist
func (o *OrganizationAPI) UpdateOrganizationHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		orgName := params["organization_name"]
		org, err := o.OrganziationStore.ReadOne(orgName)
		if org == nil {
			o.ErrorHandler.Write(rw, errors.New("organization does not exist"), http.StatusNotFound)
			return
		}
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusUnprocessableEntity)
			return
		}
		org = &organizations.Organization{}
		err = json.Unmarshal(body, org)
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		updatedOrg, err := o.OrganziationStore.Update(orgName, org.Email)
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		o.ResponseHandler.Write(rw, updatedOrg, http.StatusOK)
	})
}

// GetOrganizationHandler Handles an API request to fetch a single organization via organization name
// This will return the organization if found or 404 if the organization doesn't exist
func (o *OrganizationAPI) GetOrganizationHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		orgName := params["organization_name"]
		org, err := o.OrganziationStore.ReadOne(orgName)
		if org == nil {
			o.ErrorHandler.Write(rw, errors.New("organization does not exist"), http.StatusNotFound)
			return
		}
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		o.ResponseHandler.Write(rw, org, http.StatusOK)
	})
}

// ListOrganizationsHandler Handles an API request to list all organizations available in the registry
// This will return an error if the API cannot communicate with the backend
func (o *OrganizationAPI) ListOrganizationsHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		limit, offset, err := extractLimitAndOffset(r.URL.Query())
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusBadRequest)
			return
		}
		// Query organization store for all orgs using limit and offset
		orgs, err := o.OrganziationStore.ReadAll(limit, offset)
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		// We have the orgs now return a 200
		o.ResponseHandler.Write(rw, orgs, http.StatusOK)
	})
}

// DeleteOrganization Handles an API request to update a delete organization in the registry. It accepts an organization name via
// the REST URL and removes it from the registry. Note: This operation is idempotent and will return a success
// if a non existent organization was requested for deletion. Depending on your backing store modules underneath
// this organization may or may not be cascading deleted
func (o *OrganizationAPI) DeleteOrganizationHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		orgName := params["organization_name"]
		err := o.OrganziationStore.Delete(orgName)
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		o.ResponseHandler.Write(rw, nil, http.StatusNoContent)
	})
}

// SetupRoutes Sets up the various endpoints for the organizations API by registering handlers from this struct to it's
// corresponding routes. These are used to provide additional functionality for a more complete registry experience
// as well as logical groupings for modules
func (o *OrganizationAPI) SetupRoutes(vcsAPI endpoints.VCSConnAPIInterface) {
	o.Router.StrictSlash(true)
	o.Router.Handle("/", o.ListOrganizationsHandler()).Methods(http.MethodGet)
	o.Router.Handle("/", o.CreateOrganizationHandler()).Methods(http.MethodPost)
	o.Router.Handle("/{organization_name}", o.GetOrganizationHandler()).Methods(http.MethodGet)
	o.Router.Handle("/{organization_name}", o.UpdateOrganizationHandler()).Methods(http.MethodPatch)
	o.Router.Handle("/{organization_name}", o.DeleteOrganizationHandler()).Methods(http.MethodDelete)
	o.Router.Handle("/{organization_name}/oauth-clients", vcsAPI.ListVCSHandler()).Methods(http.MethodGet)
	o.Router.Handle("/{organization_name}/oauth-clients", vcsAPI.CreateVCSHandler()).Methods(http.MethodPost)
}
