package cmd

import (
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
	Short:   "Add or remove nodes into  exist kubernets cluster.",
	Long:    `Add or remove nodes into  exist kubernets cluster. The node's role can be master or node.`,
	Example: exampleNode,
}

func init() {
	rootCmd.AddCommand(nodeCmd)
}
