package cmd

import (
	"os"
	"sre-cli/controllers"

	"github.com/spf13/cobra"
)

var controller controllers.AppController

var rootCmd = &cobra.Command{
	Use:   "sre-cli",
	Short: "Its an app for do easy the sre job",
	Long:  "The idea is manage the different resources fo sre with one application, starting for kubernetes",
}

// Execute add all child commands to the root command
func Execute(c controllers.AppController) {
	controller = c
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toogle", "t", false, "Help message for toggle")
}
