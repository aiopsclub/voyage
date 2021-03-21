package config

type serverConfig struct {
	Host string
	Port string
}

func NewServerConfig() *serverConfig {
	return &serverConfig{}

}
