package types

// SourceStore is a generic data interface for implementaing database operations relating to modules
type SourceProvider interface {
	FetchVCSSources(token string)
	FetchVCSSource(token string, vcsRepoName string, vcsRepoOwner string) (SourceData, error)
}

type SourceData interface {
	ToModuleDocument() *Module
	GetVersionList() *ModuleVersion
}
