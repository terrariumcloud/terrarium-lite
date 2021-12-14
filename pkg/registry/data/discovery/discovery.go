package discovery

type ServiceDiscoveryResponse struct {
	LoginV1  *LoginConfig `json:"login.v1,omitempty"`
	ModuleV1 string       `json:"modules.v1,omitempty"`
}

type LoginConfig struct {
	Client     string   `json:"client"`
	GrantTypes []string `json:"grant_types"`
	Authz      string   `json:"authz"`
	Token      string   `json:"token"`
	Ports      []int    `json:"ports"`
}
