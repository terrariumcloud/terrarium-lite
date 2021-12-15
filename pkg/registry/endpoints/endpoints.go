package endpoints

import "net/http"

type DiscoveryAPIInterface interface {
	DiscoveryHandler() http.Handler
}

// ModuleAPIInterface specifies the required HTTP handlers for a Terrarium Modules API
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

type SourceAPIInterface interface {
	SetupRoutes()
	CreateVCSModule() http.Handler
}

type VCSConnAPIInterface interface {
	CreateVCSHandler() http.Handler
	GetVCSHandler() http.Handler
	UpdateVCSHandler() http.Handler
	ListVCSHandler() http.Handler
	DeleteVCSHandler() http.Handler
}
