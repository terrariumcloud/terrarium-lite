package api

import (
	"net/http"

	"github.com/dylanrhysscott/terrarium/pkg/registry/data/discovery"
	"github.com/dylanrhysscott/terrarium/pkg/registry/responses"
)

type DiscoveryAPI struct {
	ErrorHandler    responses.APIErrorWriter
	ResponseHandler responses.APIResponseWriter
	LoginConfig     *discovery.LoginConfig
	ModuleEndpoint  string
}

func (d *DiscoveryAPI) DiscoveryHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		resp := &discovery.ServiceDiscoveryResponse{
			LoginV1:  d.LoginConfig,
			ModuleV1: d.ModuleEndpoint,
		}
		d.ResponseHandler.WriteRaw(rw, resp, http.StatusOK)
	})
}

func NewDiscoveryAPI(loginConfig *discovery.LoginConfig, moduleEndpoint string, responseHandler responses.APIResponseWriter, errorHandler responses.APIErrorWriter) *DiscoveryAPI {
	return &DiscoveryAPI{
		LoginConfig:     loginConfig,
		ModuleEndpoint:  moduleEndpoint,
		ResponseHandler: responseHandler,
		ErrorHandler:    errorHandler,
	}
}
