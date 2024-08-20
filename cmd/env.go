package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Runs HTTP server providing OS environment variables in JSON or text",
	Run: func(cmd *cobra.Command, args []string) {
		addr := fmt.Sprintf("%s:%d", Address, Port)
		log.Println("Starting env server", addr)
		envVars := os.Environ()
		r := gin.Default()
		r.GET("/env.json", func(c *gin.Context) {
			c.JSON(200, toMap(envVars))
		})
		r.GET("/env", func(c *gin.Context) {
			for _, envVar := range envVars {
				_, _ = fmt.Fprintln(c.Writer, envVar)
			}
		})
		err := r.Run(addr)
		if err != nil {
			log.Println("ERROR:", err)
			return
		}
	},
}

func toMap(vars []string) map[string]string {
	result := make(map[string]string)
	for _, v := range vars {
		parts := strings.SplitN(v, "=", 2)
		result[parts[0]] = parts[1]
	}
	return result
}

func init() {
	serveCmd.AddCommand(envCmd)
}
