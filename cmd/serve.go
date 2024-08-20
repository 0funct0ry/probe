package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts a server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	},
}

var Port int
var Address string

func init() {
	rootCmd.AddCommand(serveCmd)

	envCmd.PersistentFlags().IntVarP(&Port, "port", "p", 3000, "port to listen")
	envCmd.PersistentFlags().StringVarP(&Address, "address", "b", "0.0.0.0", "bind address")

}
