package controllers

import (
	"fmt"
	"log"
	"os"
	"sre-cli/usecases/interactors"
	"strconv"
	"text/tabwriter"
)

type deploymentController struct {
	DeploymentInteractor interactors.DeploymentInteractor
}

// DeploymentController is the interface for connect to this controller
type DeploymentController interface {
	GetAllDeploymentsTerminal(string)
}

// NewDeploymentController return a controller
func NewDeploymentController(interactor interactors.DeploymentInteractor) DeploymentController {
	return &deploymentController{
		DeploymentInteractor: interactor,
	}
}

func (dC *deploymentController) GetAllDeploymentsTerminal(namespace string) {
	deploymentLists, err := dC.DeploymentInteractor.GetAll(namespace)
	if err != nil {
		fmt.Println(err)
		log.Panic()
	}

	tab := tabwriter.NewWriter(os.Stdout, 0, 0, 10, ' ', tabwriter.TabIndent)
	fmt.Fprintf(tab, "%s\t%s\t%s\t%s\t%s\t%s\n", "NAME", "READY", "UP-TO-DATE", "AVAILABLE", "CLUSTER", "NAMESPACE")
	for key, deploymentList := range deploymentLists {
		for _, deployment := range deploymentList {
			fmt.Fprintf(
				tab,
				"%s\t%s\t%v\t%v\t%s\t%s\n",
				deployment.Name,
				strconv.Itoa(int(deployment.ReadyReplicas))+"/"+strconv.Itoa(int(deployment.Replicas)),
				deployment.UpdatedReplicas,
				deployment.AvailableReplicas,
				key,
				deployment.Namespace,
			)
			for _, pod := range deployment.PodList {
				fmt.Fprintf(
					tab,
					"-   %s\t%s\n",
					pod.Name,
					pod.State,
				)
			}
		}
	}
	tab.Flush()
}
