package modules

type Module struct {
	Name         string
	Organization string
	Provider     string
	Version      string
	Source       string
}

type ModuleVersionItem struct {
	Version string `json:"version"`
}

type ModuleVersions struct {
	Versions []*ModuleVersionItem `json:"versions"`
}

type ModuleVersionResponse struct {
	Modules []*ModuleVersions `json:"modules"`
}
