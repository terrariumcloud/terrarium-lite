// Package discovery implements native Terraform Discovery as part of the Registry Protocol defined here
// https://www.terraform.io/internals/module-registry-protocol#service-discovery. The package provides
// a single .well-known route allowing a Terraform client to probe the registry and discover endpoints
// of supported protocols such as modules or login
package discovery

import (
	"github.com/terrariumcloud/terrarium/pkg/registry/data/discovery"
	"github.com/terrariumcloud/terrarium/pkg/registry/responses"
)

// NewDiscoveryAPI Creates a new instance of the discovery API that defines a static well known route pointing
// to endpoints for modules. In future once implemented and enabled this will extend the discovery document for futher
// protocols
func NewDiscoveryAPI(loginConfig *discovery.LoginConfig, moduleEndpoint string, responseHandler responses.APIResponseWriter, errorHandler responses.APIErrorWriter) *DiscoveryAPI {
	return &DiscoveryAPI{
		LoginConfig:     loginConfig,
		ModuleEndpoint:  moduleEndpoint,
		ResponseHandler: responseHandler,
		ErrorHandler:    errorHandler,
	}
}
