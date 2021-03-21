package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"voyage/internal/config"
)

var exampleRemoveNode = `
	# k8s init by password with three master nodes
	voyage node remove --role node  \
	--username root  --passwd passwd \
	192.168.1.6
`

var nodeRemoveConfig = config.NewNodeConfig()

var nodeRemoveCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove nodes into  exist kubernets cluster.",
	Long:    `Remove nodes into  exist kubernets cluster. The node's role can be master or node.`,
	Example: exampleRemoveNode,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v", nodeRemoveConfig)
	},
}

func init() {
	nodeCmd.AddCommand(nodeRemoveCmd)

	// Here you will define your flags and configuration settings.
	nodeRemoveCmd.Flags().StringVar(&nodeConfig.UserName, "username", "root", "servers user name for ssh")
	nodeRemoveCmd.Flags().StringVar(&nodeConfig.Passwd, "passwd", "", "password for ssh")
	nodeRemoveCmd.Flags().StringVar(&nodeConfig.Role, "role", "node", "node role name")
}
