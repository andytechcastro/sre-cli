package cmd

import (
	"github.com/spf13/cobra"
)

var kubeCmd = &cobra.Command{
	Use:   "kube",
	Short: "the command for connect with kubernetes",
	Long:  "You can use it for connect to kubernetes and lunch different commands",
}

func init() {
	rootCmd.AddCommand(kubeCmd)
	kubeCmd.PersistentFlags().StringP("namespace", "n", "default", "Namespace for resource")
	kubeCmd.PersistentFlags().BoolP("allnamespaces", "A", false, "All namespaces")
	kubeCmd.PersistentFlags().StringP("environment", "e", "all", "Environment for the cluster")
}
