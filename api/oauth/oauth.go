package oauth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dylanrhysscott/terrarium/internal/terrariummongo/vcs"
	"github.com/dylanrhysscott/terrarium/pkg/types"

	"github.com/gorilla/mux"
)

type OAuthAPIInterface interface {
	SetupRoutes()
	LoginHandler() http.Handler
	GithubCallbackHandler() http.Handler
}

type OAuthAPI struct {
	Router          *mux.Router
	ErrorHandler    types.APIErrorWriter
	ResponseHandler types.APIResponseWriter
	VCSStore        types.VCSStore
}

func (o *OAuthAPI) LoginHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

func (o *OAuthAPI) GithubCallbackHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		params := mux.Vars(r)
		vcsID := params["id"]
		if code == "" {
			o.ErrorHandler.Write(rw, errors.New("invalid code"), http.StatusBadRequest)
			return
		}
		vcsItem, err := o.VCSStore.ReadOne(vcsID)
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
		log.Println(string(data))
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

func (o *OAuthAPI) SetupRoutes() {
	o.Router.StrictSlash(true)
	o.Router.Handle("/github/{id}/callback", o.GithubCallbackHandler()).Methods(http.MethodGet)
}

func NewOAuthAPI(router *mux.Router, path string, vcsstore types.VCSStore, responseHandler types.APIResponseWriter, errorHandler types.APIErrorWriter) *OAuthAPI {
	a := &OAuthAPI{
		Router:          router.PathPrefix(path).Subrouter(),
		VCSStore:        vcsstore,
		ResponseHandler: responseHandler,
		ErrorHandler:    errorHandler,
	}
	a.SetupRoutes()
	return a
}
