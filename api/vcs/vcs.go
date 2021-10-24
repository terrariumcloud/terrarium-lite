package vcs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dylanrhysscott/terrarium/pkg/types"
	"github.com/gorilla/mux"
)

type VCSAPIInterface interface {
	CreateVCSHandler() http.Handler
	GetVCSHandler() http.Handler
	UpdateVCSHandler() http.Handler
	ListVCSHandler() http.Handler
	DeleteVCSHandler() http.Handler
	GithubCallbackHander() http.Handler
}

type VCSAPI struct {
	OrganziationStore types.OrganizationStore
	ErrorHandler      types.APIErrorWriter
	ResponseHandler   types.APIResponseWriter
	VCSStore          types.VCSStore
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
		link := &types.VCSOAuthClientLink{}
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
		link.CallbackURI = fmt.Sprintf("/v1/oauth-clients/%s/github/callback", link.ClientID)
		err = link.Validate()
		if err != nil {
			v.ErrorHandler.Write(rw, err, http.StatusUnprocessableEntity)
			return
		}
		vcsConnection, err := v.VCSStore.Create(org.ID.Hex(), orgName, link)
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

// GithubCallbackHandler is a handler for handling callback redirection from Github
func (v *VCSAPI) GithubCallbackHander() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		params := mux.Vars(r)
		clientID := params["id"]
		if code == "" {
			v.ErrorHandler.Write(rw, errors.New("invalid code"), http.StatusBadRequest)
			return
		}
		vcs, err := v.VCSStore.ReadOneByClientID(clientID)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				v.ErrorHandler.Write(rw, errors.New("vcs connection does not exist"), http.StatusNotFound)
				return
			}
			v.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		req, err := http.NewRequest(http.MethodPost, "https://github.com/login/oauth/access_token", nil)
		if err != nil {
			v.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		req.Header.Add("Accept", "application/json")
		q := req.URL.Query()
		q.Add("client_id", vcs.OAuth.ClientID)
		q.Add("client_secret", vcs.OAuth.ClientSecret)
		q.Add("code", code)
		req.URL.RawQuery = q.Encode()
		client := http.DefaultClient
		resp, err := client.Do(req)
		if err != nil {
			v.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			v.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		ghToken := &types.VCSToken{}
		err = json.Unmarshal(data, ghToken)
		if err != nil {
			v.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		v.ResponseHandler.Write(rw, ghToken, http.StatusOK)
	})
}

// NewVCSAPI creates an instance of the VCS API with the reqired database
// driver support
func NewVCSAPI(vcsstore types.VCSStore, orgstore types.OrganizationStore, responseHandler types.APIResponseWriter, errorHandler types.APIErrorWriter) *VCSAPI {
	return &VCSAPI{
		OrganziationStore: orgstore,
		VCSStore:          vcsstore,
		ResponseHandler:   responseHandler,
		ErrorHandler:      errorHandler,
	}
}
