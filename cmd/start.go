package cmd

import (
	"fmt"

	"untils/server"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动服务",
	Long:  `启动收敛服务`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		startServer()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func startServer() {
	server.Server.Start()
	// utils.Server.Start()
}
