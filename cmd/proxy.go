package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var remoteUrl string

// proxyCmd represents the proxy command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "reverse proxy to a remote url",
	Run: func(cmd *cobra.Command, args []string) {
		remote, err := url.Parse(remoteUrl)
		if err != nil {
			log.Println("ERROR:", err)
			return
		}

		r := gin.Default()

		r.Any("/*proxyPath", func(c *gin.Context) {

			proxy := httputil.NewSingleHostReverseProxy(remote)
			proxy.Director = func(req *http.Request) {
				req.Header = c.Request.Header
				req.Host = remote.Host
				req.URL.Scheme = remote.Scheme
				req.URL.Host = remote.Host
				req.URL.Path = c.Param("proxyPath")
			}

			proxy.ServeHTTP(c.Writer, c.Request)
		})

		err = r.Run(fmt.Sprintf("%s:%d", Address, Port))
		if err != nil {
			log.Println("ERROR:", err)
			return
		}
	},
}

func init() {
	serveCmd.AddCommand(proxyCmd)
	proxyCmd.Flags().StringVarP(&remoteUrl, "proxy", "p", "", "remote url")
}
