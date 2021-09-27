package organization

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dylanrhysscott/terrarium/pkg/types"
)

func extractLimitAndOffset(qs url.Values) (int, int, error) {
	var limit int = 0
	var offset int = 0
	if qs.Has("limit") {
		// If we have a limit value set in QS attempt to convert to int
		i, err := strconv.Atoi(qs.Get("limit"))
		if err != nil {
			return 0, 0, errors.New("limit is not an integer")
		}
		limit = i
	}
	if qs.Has("offset") {
		i, err := strconv.Atoi(qs.Get("offset"))
		if err != nil {
			return 0, 0, errors.New("offset is not an integer")
		}
		offset = i
	}
	return limit, offset, nil
}

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
	ErrorHandler      types.APIErrorWriter
	ResponseHandler   types.APIResponseWriter
}

// CreateOrganizationHandler is a handler for creating an organization (POST)
func (o *OrganizationAPI) CreateOrganizationHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	jsonData, err := types.NewTerrariumBadRequest(err.Error())
		// 	if err != nil {
		// 		log.Println(err.Error())
		// 		rw.WriteHeader(http.StatusInternalServerError)
		// 		return
		// 	}
		// 	rw.WriteHeader(http.StatusBadRequest)
		// 	rw.Write(jsonData)
		// 	return
		// }
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

// DeleteOrganizationHandler is a handler for deleting an organization (DELETE)
func (o *OrganizationAPI) DeleteOrganizationHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

// NewOrganizationAPI creates an instance of the organization API with the reqired database
// driver support
func NewOrganizationAPI(store types.OrganizationStore, responseHandler types.APIResponseWriter, errorHandler types.APIErrorWriter) *OrganizationAPI {
	return &OrganizationAPI{
		OrganziationStore: store,
		ResponseHandler:   responseHandler,
		ErrorHandler:      errorHandler,
	}
}
