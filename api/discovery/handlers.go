package discovery

import (
	"net/http"

	"github.com/terrariumcloud/terrarium/pkg/registry/data/discovery"
	"github.com/terrariumcloud/terrarium/pkg/registry/responses"
)

// DiscoveryAPI is a struct implementing the handlers for the DiscoveryAPIInterface from the endpoints package in Terrarium
type DiscoveryAPI struct {
	ErrorHandler    responses.APIErrorWriter
	ResponseHandler responses.APIResponseWriter
	LoginConfig     *discovery.LoginConfig
	ModuleEndpoint  string
}

// DiscoveryHandler Handles an API request for service discovery from a Terraform client
// This is the first step a Terraform client will take in determining if the registry is a valid
// implementation and where to look for other endpoints
func (d *DiscoveryAPI) DiscoveryHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		resp := &discovery.ServiceDiscoveryResponse{
			LoginV1:  d.LoginConfig,
			ModuleV1: d.ModuleEndpoint,
		}
		d.ResponseHandler.WriteRaw(rw, resp, http.StatusOK)
	})
}
