package config

type nodeConfig struct {
	UserName string
	Passwd   string
	Role     string
}

func NewNodeConfig() *nodeConfig {
	return &nodeConfig{}

}
