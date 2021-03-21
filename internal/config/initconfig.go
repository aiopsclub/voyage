package config

type initConfig struct {
	UserName  string
	Passwd    string
	MasterIPs []string
	NodeIPs   []string
	Version   string
}

func NewInitConfig() *initConfig {
	return &initConfig{}

}
