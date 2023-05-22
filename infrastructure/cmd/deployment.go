package cmd

import "github.com/spf13/cobra"

var deploymentCmd = &cobra.Command{
	Use:   "deployments",
	Short: "This command if for use the deployment Options",
	Long:  "With this command you can call the deployments option",
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("namespace")
		boolAllNamespace, _ := cmd.Flags().GetBool("allnamespaces")
		if boolAllNamespace {
			namespace = ""
		}
		controller.Deployment.GetAllDeploymentsTerminal(namespace)
	},
}

func init() {
	kubeCmd.AddCommand(deploymentCmd)
}
