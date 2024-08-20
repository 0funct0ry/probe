package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"time"
)

var interval int

var clockCmd = &cobra.Command{
	Use:   "clock",
	Short: "log current time at regular intervals",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			log.Println(time.Now())
			time.Sleep(time.Duration(interval) * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(clockCmd)
	clockCmd.Flags().IntVarP(&interval, "interval", "i", 5, "Interval in seconds")
}
