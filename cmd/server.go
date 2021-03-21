package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"voyage/internal/config"
)

var exampleServer = `
	# start restful server 
	voyage server --host 0.0.0.0  \
	--port 80
`

var serverConfig = config.NewServerConfig()

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "start restful server",
	Long:    "start restful server for k8s maintain.",
	Example: exampleServer,
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})

		})
		r.Run(serverConfig.Host + ":" + serverConfig.Port)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.
	initCmd.Flags().StringVar(&serverConfig.Host, "host", "0.0.0.0", "http sever host")
	initCmd.Flags().StringVar(&serverConfig.Port, "port", "80", "http sever port")
}
