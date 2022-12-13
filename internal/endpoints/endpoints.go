// Package drivers provides interfaces to implement a Terrarium API. Generally these interfaces are not intended for 3rd
// party implementation and serve to allow for internal implementation testing of handlers for the various APIs
package endpoints

import "net/http"

// OrganizationAPIInterface specifies the required HTTP handlers for a Terrarium Discovery API
type DiscoveryAPIInterface interface {
	DiscoveryHandler() http.Handler
}

// ModuleAPIInterface specifies the required HTTP handlers for a Terrarium Modules API implementation
type ModuleAPIInterface interface {
	GetModuleVersionHandler() http.Handler
	DownloadModuleHandler() http.Handler
	ArchiveHandler() http.Handler
}
