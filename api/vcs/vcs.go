package vcs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/dylanrhysscott/terrarium/pkg/registry/data/vcs"
	"github.com/dylanrhysscott/terrarium/pkg/registry/responses"
	"github.com/dylanrhysscott/terrarium/pkg/registry/stores"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VCSAPI struct {
	Router            *mux.Router
	OrganziationStore stores.OrganizationStore
	ErrorHandler      responses.APIErrorWriter
	ResponseHandler   responses.APIResponseWriter
	VCSStore          stores.VCSSConnectionStore
}

// CreateVCSHandler is a handler for creating an organization VCS connection (POST)
func (v *VCSAPI) CreateVCSHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		orgName := params["organization_name"]
		org, err := v.OrganziationStore.ReadOne(orgName)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				v.ErrorHandler.Write(rw, errors.New("organization does not exist"), http.StatusNotFound)
				return
			}
			v.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		link := &vcs.VCSOAuthClientLink{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			v.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(body, link)
		if err != nil {
			v.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		err = link.Validate()
		if err != nil {
			v.ErrorHandler.Write(rw, err, http.StatusUnprocessableEntity)
			return
		}
		var id string
		if v.OrganziationStore.GetBackendType() == "mongo" {
			id = org.ID.(primitive.ObjectID).Hex()
		}
		vcsConnection, err := v.VCSStore.Create(id, orgName, link)
		if err != nil {
			v.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		v.ResponseHandler.Write(rw, vcsConnection, http.StatusCreated)
	})
}

// UpdateVCSHandler is a handler for updating an organization VCS connection (PUT)
func (v *VCSAPI) UpdateVCSHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

// GetVCSHandler is a handler for getting a single organization VCS connection (GET)
func (v *VCSAPI) GetVCSHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		vcsconnID := params["id"]
		vcsConnection, err := v.VCSStore.ReadOne(vcsconnID, false)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				v.ErrorHandler.Write(rw, errors.New("vcs connection does not exist"), http.StatusNotFound)
				return
			}
			v.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		v.ResponseHandler.Write(rw, vcsConnection, http.StatusOK)
	})
}

// ListVCSHandler is a handler for listing all organization VCS connections (GET)
func (v *VCSAPI) ListVCSHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

// DeleteVCSHandler is a handler for deleting an organization VCS connection (DELETE)
func (v *VCSAPI) DeleteVCSHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

func (v *VCSAPI) SetupRoutes() {
	v.Router.StrictSlash(true)
	v.Router.Handle("/{id}", v.GetVCSHandler()).Methods(http.MethodGet)
	v.Router.Handle("/{id}", v.UpdateVCSHandler()).Methods(http.MethodPatch)
	v.Router.Handle("/{id}", v.DeleteVCSHandler()).Methods(http.MethodDelete)
}

// NewVCSAPI creates an instance of the VCS API with the reqired database
// driver support
func NewVCSAPI(router *mux.Router, path string, vcsconnstore stores.VCSSConnectionStore, orgstore stores.OrganizationStore, responseHandler responses.APIResponseWriter, errorHandler responses.APIErrorWriter) *VCSAPI {
	v := &VCSAPI{
		Router:            router.PathPrefix(path).Subrouter(),
		OrganziationStore: orgstore,
		VCSStore:          vcsconnstore,
		ResponseHandler:   responseHandler,
		ErrorHandler:      errorHandler,
	}
	v.SetupRoutes()
	return v
}
