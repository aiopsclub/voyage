package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"voyage/internal/config"
	"voyage/pkg/install"
)

var contactMe = `
__   _____  _   _  __ _  __ _  ___ 
\ \ / / _ \| | | |/ _  |/ _  |/ _ \
 \ V / (_) | |_| | (_| | (_| |  __/
  \_/ \___/ \__, |\__,_|\__, |\___|
            |___/       |___/      
官方文档：voyage.aiopsclub.com
项目地址：github.com/aiopsclub/voyage
`
var exampleInit = `
	# k8s init by password with three master nodes
	voyage init --passwd server_password  \
	--master 192.168.1.2,192.168.1.3,192.168.1.4 \
	--node 192.168.1.5 --username root \
	--version v1.18.0
`

var initConfig = config.NewInitConfig()

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "One key to init kubernets HA cluster",
	Long: `voyage init --master 192.168.1.2,192.168.1.3,192.168.1.4 \
	--node 192.168.1.5 --username root --passwd server_password \
	--version v1.18.0`,
	Example: exampleInit,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(contactMe)
		err := install.Run(initConfig)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.
	initCmd.Flags().StringVar(&initConfig.UserName, "username", "root", "servers user name for ssh")
	initCmd.Flags().StringVar(&initConfig.Passwd, "passwd", "", "password for ssh")
	initCmd.Flags().StringSliceVar(&initConfig.MasterIPs, "master", []string{}, "kubernetes multi-masters.")
	initCmd.Flags().StringSliceVar(&initConfig.NodeIPs, "node", []string{}, "kubernetes multi-nodes .")
	initCmd.Flags().StringVar(&initConfig.Version, "version", "1.18.0", "kubeadm log level")
	initCmd.Flags().StringVar(&initConfig.SshPort, "sshport", "22", "ssh port")
}
