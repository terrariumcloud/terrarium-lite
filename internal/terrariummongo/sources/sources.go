package sources

type SourceVCSRepo struct {
	VCSUser string `json:"vcs_user"`
	Branch  string `json:"branch"`
	Repo    string `json:"repo"`
}

type Module struct {
	ID       string           `json:"id" bson:"_id"`
	Owner    string           `json:"owner" bson:"owner"`
	Name     string           `json:"name" bson:"name"`
	Provider string           `json:"provider" bson:"provider"`
	Versions []*ModuleVersion `json:"versions" bson:"versions"`
}

type ModuleVersion struct {
	Version string `json:"version" bson:"version"`
}
