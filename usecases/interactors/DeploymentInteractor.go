package interactors

import (
	"fmt"
	"sre-cli/entities"
	"sre-cli/usecases/repositories"
)

type deploymentInteractor struct {
	DeploymentRepo map[string]repositories.DeploymentRepository
	PodRepo        map[string]repositories.PodRepository
}

// DeploymentInteractor is an interface for connect to deployment interactor
type DeploymentInteractor interface {
	GetAll(string, bool) (map[string][]entities.Deployment, error)
}

// NewDeploymentInteractor return a new struct with deploymentInteractor
func NewDeploymentInteractor(deploymentRepo map[string]repositories.DeploymentRepository, podRepo map[string]repositories.PodRepository) DeploymentInteractor {
	return &deploymentInteractor{
		DeploymentRepo: deploymentRepo,
		PodRepo:        podRepo,
	}
}

func (dI *deploymentInteractor) GetAll(namespace string, pods bool) (map[string][]entities.Deployment, error) {
	DeploymentLists := map[string][]entities.Deployment{}
	for key, repo := range dI.DeploymentRepo {
		DeploymentList, err := repo.GetAll(namespace)
		if err != nil {
			return DeploymentLists, nil
		}
		if pods {
			for keyDeployment, deployment := range DeploymentList {
				DeploymentList[keyDeployment].PodList, err = dI.PodRepo[key].GetPodByLabels(namespace, deployment.MatchLabels)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		DeploymentLists[key] = DeploymentList
	}
	return DeploymentLists, nil
}
