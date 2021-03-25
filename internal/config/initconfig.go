package config

type InitConfig struct {
	UserName  string
	Passwd    string
	MasterIPs []string
	NodeIPs   []string
	Version   string
	SshPort   string
}

func NewInitConfig() *InitConfig {
	return &InitConfig{}

}
