package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts a server",
	Run: func(cmd *cobra.Command, args []string) {
		addr := fmt.Sprintf("%s:%d", Address, Port)
		log.Println("server running on", addr)
		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			_, _ = fmt.Fprintln(c.Writer, hostname())
		})
		err := r.Run(addr)
		if err != nil {
			log.Println("ERROR:", err)
			return
		}
	},
}

func hostname() string {
	name, ok := os.LookupEnv("HOSTNAME")
	if ok {
		return name
	} else {
		return "not set"
	}
}

var Port int
var Address string

func init() {
	rootCmd.AddCommand(serveCmd)

	envCmd.PersistentFlags().IntVarP(&Port, "port", "p", 3000, "port to listen")
	envCmd.PersistentFlags().StringVarP(&Address, "address", "b", "0.0.0.0", "bind address")

}
