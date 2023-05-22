package cmd

import "github.com/spf13/cobra"

var podCmd = &cobra.Command{
	Use:   "pods",
	Short: "This command is for use the pod Options",
	Long:  "With this command you can call the pods option",
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("namespace")
		boolAllNamespace, _ := cmd.Flags().GetBool("allnamespaces")
		if boolAllNamespace {
			namespace = ""
		}
		controller.Pod.GetAllPodsTerminal(namespace)
	},
}

func init() {
	kubeCmd.AddCommand(podCmd)
}
