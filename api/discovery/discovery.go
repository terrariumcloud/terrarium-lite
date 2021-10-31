package discovery

import (
	"net/http"

	"github.com/dylanrhysscott/terrarium/pkg/types"
)

type DiscoveryAPI struct {
	ErrorHandler    types.APIErrorWriter
	ResponseHandler types.APIResponseWriter
	LoginConfig     *LoginConfig
	ModuleEndpoint  string
}

func (d *DiscoveryAPI) DiscoveryHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		resp := &ServiceDiscoveryResponse{
			LoginV1:  d.LoginConfig,
			ModuleV1: d.ModuleEndpoint,
		}
		d.ResponseHandler.WriteRaw(rw, resp, http.StatusOK)
	})
}

func NewDiscoveryAPI(loginConfig *LoginConfig, moduleEndpoint string, responseHandler types.APIResponseWriter, errorHandler types.APIErrorWriter) *DiscoveryAPI {
	return &DiscoveryAPI{
		LoginConfig:     loginConfig,
		ModuleEndpoint:  moduleEndpoint,
		ResponseHandler: responseHandler,
		ErrorHandler:    errorHandler,
	}
}
