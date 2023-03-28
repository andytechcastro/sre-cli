package controllers

import (
	"fmt"
	"os"
	"sre-cli/usecases/interactors"
	"text/tabwriter"

	"log"
)

type podController struct {
	Interactor interactors.PodInteractor
}

// PodController in an interface for connect with controller
type PodController interface {
	GetAllPodsTerminal(string)
}

// NewPodController return an operatorController struct
func NewPodController(interactor interactors.PodInteractor) PodController {
	return &podController{
		Interactor: interactor,
	}
}

func (oc *podController) GetAllPodsTerminal(namespace string) {
	podLists, err := oc.Interactor.GetAll(namespace)
	if err != nil {
		fmt.Println(err)
		log.Panic()
	}
	tab := tabwriter.NewWriter(os.Stdout, 0, 0, 10, ' ', tabwriter.TabIndent)
	fmt.Fprintf(tab, "%s\t%s\t%s\t%s\n", "NAME", "STATUS", "CLUSTER", "NAMESPACE")
	for key, podList := range podLists {
		for _, pod := range podList {
			fmt.Fprintf(tab, "%s\t%s\t%s\t%s\n", pod.Name, pod.State, key, pod.Namespace)
		}
	}
	tab.Flush()
}
