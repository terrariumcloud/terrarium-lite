package types

// SourceProvider is a generic version control provider interface for implementaing source control operations
// FetchVCSSource should accept an OAuth token and a repo name returning some form of source data. If there
// is an error this should be returned and source data is expected to be nil
type SourceProvider interface {
	FetchVCSSource(token string, vcsRepoName string) (SourceData, error)
}

// SourceData is an interface to standardise the kind of data returned by the SourceProvider interface. This
// interface should provide the ability to return repo names, descriptions and owners information for the piece
// of source data. Typically the owner will be the username on the source provider
type SourceData interface {
	GetRepoName() string
	GetRepoDescription() string
	GetRepoOwner() string
}
