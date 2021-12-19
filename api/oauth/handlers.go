package oauth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/terrariumcloud/terrarium/pkg/registry/data/vcs"
	"github.com/terrariumcloud/terrarium/pkg/registry/responses"
	"github.com/terrariumcloud/terrarium/pkg/registry/stores"
)

// OAuthAPI is a struct implementing the handlers for the OAuthAPIInterface from the endpoints package in Terrarium
type OAuthAPI struct {
	Router          *mux.Router
	ErrorHandler    responses.APIErrorWriter
	ResponseHandler responses.APIResponseWriter
	VCSStore        stores.VCSSConnectionStore
}

// TODO: Not yet implemented
func (o *OAuthAPI) LoginHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

// GithubCallbackHandler implements an API route to handle the exchange of an OAuth code for a Github access token
// The resulting access token is then stored in the VCS store for later use by Terrarium on the users behalf. Typically
// this will be to sync source code stored in Git to the registry in response to webhook events
// TODO: Remove the mongo specfic code check for database agnostic approach
func (o *OAuthAPI) GithubCallbackHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		params := mux.Vars(r)
		vcsID := params["id"]
		if code == "" {
			o.ErrorHandler.Write(rw, errors.New("invalid code"), http.StatusBadRequest)
			return
		}
		vcsItem, err := o.VCSStore.ReadOne(vcsID, true)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				o.ErrorHandler.Write(rw, errors.New("vcs connection does not exist"), http.StatusNotFound)
				return
			}
			o.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		req, err := http.NewRequest(http.MethodPost, "https://github.com/login/oauth/access_token", nil)
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		req.Header.Add("Accept", "application/json")
		q := req.URL.Query()
		q.Add("client_id", vcsItem.OAuth.ClientID)
		q.Add("client_secret", vcsItem.OAuth.ClientSecret)
		q.Add("code", code)
		req.URL.RawQuery = q.Encode()
		client := http.DefaultClient
		resp, err := client.Do(req)
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		ghToken := &vcs.VCSToken{}
		err = json.Unmarshal(data, ghToken)
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		err = o.VCSStore.UpdateVCSToken(vcsItem.OAuth.ClientID, ghToken)
		if err != nil {
			o.ErrorHandler.Write(rw, err, http.StatusInternalServerError)
			return
		}
		o.ResponseHandler.Redirect(rw, r, "http://localhost:3000")
	})
}

// SetupRoutes Sets up the various callback handlers for supported OAuth providers in the Terraform login flow
// as defined here https://www.terraform.io/internals/login-protocol. Ultimately these routes will allow
// terraform to login with providers such as Github, Google etc
func (o *OAuthAPI) SetupRoutes() {
	o.Router.StrictSlash(true)
	o.Router.Handle("/github/{id}/callback", o.GithubCallbackHandler()).Methods(http.MethodGet)
}
