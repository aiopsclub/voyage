package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"voyage/internal/config"
)

var exampleNode = `
	# k8s init by password with three master nodes
	voyage node --action add --role node  \
	--username root  --passwd passwd \
	192.168.1.6
`

var nodeConfig = config.NewNodeConfig()

var nodeCmd = &cobra.Command{
	Use:     "node",
	Short:   "Add nodes into  exist kubernets cluster.",
	Long:    `Add nodes into  exist kubernets cluster. The node's role can be master or node.`,
	Args:    cobra.ExactArgs(1),
	Example: exampleNode,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v", nodeConfig)
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)

	// Here you will define your flags and configuration settings.
	nodeCmd.Flags().StringVar(&nodeConfig.UserName, "username", "root", "servers user name for ssh")
	nodeCmd.Flags().StringVar(&nodeConfig.Passwd, "passwd", "", "password for ssh")
	nodeCmd.Flags().StringVar(&nodeConfig.Role, "role", "node", "node role name")
	nodeCmd.Flags().StringVar(&nodeConfig.Action, "action", "add", "action for k8s node")
}
