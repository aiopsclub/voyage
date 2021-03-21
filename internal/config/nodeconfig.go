package config

type nodeConfig struct {
	UserName string
	Passwd   string
	Action   string
	Role     string
}

func NewNodeConfig() *nodeConfig {
	return &nodeConfig{}

}
