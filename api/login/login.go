package login

import (
	"encoding/json"
	"net/http"
)

type LoginAPIInterface interface {
	DiscoveryHandler() http.Handler
}

type ServiceDiscoveryResponse struct {
	LoginV1 *ServiceDiscoveryConfig `json:"login.v1"`
}

type ServiceDiscoveryError struct {
	Code    int
	Message string
}

type ServiceDiscoveryConfig struct {
	Client     string   `json:"client"`
	GrantTypes []string `json:"grant_types"`
	Authz      string   `json:"authz"`
	Token      string   `json:"token"`
	Ports      []int    `json:"ports"`
}

type LoginAPI struct {
	ClientID      string
	AuthEndpoint  string
	TokenEndpoint string
	Ports         []int
}

func (l *LoginAPI) DiscoveryHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		conf, err := l.constructLoginObject()
		if err != nil {
			resp, err := constructErrorObject("Failed creating login config", http.StatusInternalServerError)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			rw.Header().Add("Content-Type", "application/json")
			http.Error(rw, string(resp), http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		rw.Write(conf)
	}
}

func (l *LoginAPI) constructLoginObject() ([]byte, error) {
	conf := &ServiceDiscoveryResponse{
		LoginV1: &ServiceDiscoveryConfig{
			Client:     l.ClientID,
			GrantTypes: []string{"authz_code"},
			Authz:      l.AuthEndpoint,
			Token:      l.TokenEndpoint,
			Ports:      l.Ports,
		},
	}

	return json.Marshal(conf)
}

func constructErrorObject(message string, code int) ([]byte, error) {
	errResp := &ServiceDiscoveryError{
		Message: message,
		Code:    code,
	}
	return json.Marshal(errResp)
}

func NewLoginAPI(client string, authEndpoint string, tokenEndpoint string, ports []int) *LoginAPI {
	return &LoginAPI{
		ClientID:      client,
		AuthEndpoint:  authEndpoint,
		TokenEndpoint: tokenEndpoint,
		Ports:         ports,
	}
}
