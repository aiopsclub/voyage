package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"voyage/internal/config"
)

var exampleNodeAdd = `
	# k8s init by password with three master nodes
	voyage node add --role node  \
	--username root  --passwd passwd \
	192.168.1.6
`

var nodeAddConfig = config.NewNodeConfig()

var nodeAddCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add nodes into  exist kubernets cluster.",
	Long:    `Add nodes into  exist kubernets cluster. The node's role can be master or node.`,
	Args:    cobra.ExactArgs(1),
	Example: exampleNode,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v", nodeAddConfig)
	},
}

func init() {
	nodeCmd.AddCommand(nodeAddCmd)

	// Here you will define your flags and configuration settings.
	nodeAddCmd.Flags().StringVar(&nodeConfig.UserName, "username", "root", "servers user name for ssh")
	nodeAddCmd.Flags().StringVar(&nodeConfig.Passwd, "passwd", "", "password for ssh")
	nodeAddCmd.Flags().StringVar(&nodeConfig.Role, "role", "node", "node role name")
}
