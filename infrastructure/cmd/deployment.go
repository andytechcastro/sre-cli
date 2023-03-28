package cmd

import "github.com/spf13/cobra"

var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "This command if for use the deployment Options",
	Long:  "With this command you can call the deployments option",
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("namespace")
		boolAllNamespace, _ := cmd.Flags().GetBool("allnamespaces")
		pods, _ := cmd.Flags().GetBool("pods")
		if boolAllNamespace {
			namespace = ""
		}
		controller.Deployment.GetAllDeploymentsTerminal(namespace, pods)
	},
}

func init() {
	kubeCmd.AddCommand(deploymentCmd)
	deploymentCmd.Flags().BoolP("pods", "p", false, "Show their pods")
}
