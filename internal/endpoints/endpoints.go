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
	CreateModuleHandler() http.Handler
	GetModuleHandler() http.Handler
	GetModuleVersionHandler() http.Handler
	DownloadModuleHandler() http.Handler
	ArchiveHandler() http.Handler
	UpdateModuleHandler() http.Handler
	ListModulesHandler() http.Handler
	ListOrganizationModulesHandler() http.Handler
	DeleteModuleHandler() http.Handler
}

// ModuleAPIInterface specifies the required HTTP handlers for a Terrarium OAuth API implementation
type OAuthAPIInterface interface {
	SetupRoutes()
	LoginHandler() http.Handler
	GithubCallbackHandler() http.Handler
}

// OrganizationAPIInterface specifies the required HTTP handlers for a Terrarium Organizations API
type OrganizationAPIInterface interface {
	CreateOrganizationHandler() http.Handler
	GetOrganizationHandler() http.Handler
	UpdateOrganizationHandler() http.Handler
	ListOrganizationsHandler() http.Handler
	DeleteOrganizationHandler() http.Handler
}

// OrganizationAPIInterface specifies the required HTTP handlers for a Terrarium Sources API
type SourceAPIInterface interface {
	SetupRoutes()
	CreateVCSModule() http.Handler
}

// OrganizationAPIInterface specifies the required HTTP handlers for a Terrarium VCS Connections API
type VCSConnAPIInterface interface {
	CreateVCSHandler() http.Handler
	GetVCSHandler() http.Handler
	UpdateVCSHandler() http.Handler
	ListVCSHandler() http.Handler
	DeleteVCSHandler() http.Handler
}
