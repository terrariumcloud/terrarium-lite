package api

import (
	"net/http"

	"github.com/dylanrhysscott/terrarium/pkg/types"
)

type DiscoveryAPI struct {
	ErrorHandler    types.APIErrorWriter
	ResponseHandler types.APIResponseWriter
	LoginConfig     *types.LoginConfig
	ModuleEndpoint  string
}

type DiscoveryAPIInterface interface {
	DiscoveryHandler() http.Handler
}

func (d *DiscoveryAPI) DiscoveryHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		resp := &types.ServiceDiscoveryResponse{
			LoginV1:  d.LoginConfig,
			ModuleV1: d.ModuleEndpoint,
		}
		d.ResponseHandler.WriteRaw(rw, resp, http.StatusOK)
	})
}

func NewDiscoveryAPI(loginConfig *types.LoginConfig, moduleEndpoint string, responseHandler types.APIResponseWriter, errorHandler types.APIErrorWriter) *DiscoveryAPI {
	return &DiscoveryAPI{
		LoginConfig:     loginConfig,
		ModuleEndpoint:  moduleEndpoint,
		ResponseHandler: responseHandler,
		ErrorHandler:    errorHandler,
	}
}
