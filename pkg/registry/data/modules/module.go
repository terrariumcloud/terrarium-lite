package modules

type Module struct {
	ID             string `json:"_id"`
	OrganizationID string `json:"_organization_id"`
	Name           string `json:"name"`
	Organization   string `json:"organization"`
	Provider       string `json:"provider"`
	Version        string `json:"version"`
	Description    string `json:"description"`
	Source         string `json:"source"`
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
