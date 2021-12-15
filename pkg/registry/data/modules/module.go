package modules

type Module struct {
	ID             string `json:"_id" dynamodbav:"_id"`
	OrganizationID string `json:"_organization_id" dynamodbav:"_organization_id"`
	Name           string `json:"name" dynamodbav:"name"`
	Organization   string `json:"organization" dynamodbav:"organization"`
	Provider       string `json:"provider" dynamodbav:"provider"`
	Version        string `json:"version" dynamodbav:"version"`
	Description    string `json:"description" dynamodbav:"description"`
	Source         string `json:"source" dynamodbav:"source"`
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
