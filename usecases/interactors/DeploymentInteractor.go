package interactors

import (
	"fmt"
	"sre-cli/entities"
	"sre-cli/usecases/repositories"
	"sync"
)

type deploymentInteractor struct {
	DeploymentRepo map[string]repositories.DeploymentRepository
	PodRepo        map[string]repositories.PodRepository
}

// DeploymentInteractor is an interface for connect to deployment interactor
type DeploymentInteractor interface {
	GetAll(string) (map[string][]entities.Deployment, error)
}

// NewDeploymentInteractor return a new struct with deploymentInteractor
func NewDeploymentInteractor(
	deploymentRepo map[string]repositories.DeploymentRepository,
	podRepo map[string]repositories.PodRepository,
) DeploymentInteractor {
	return &deploymentInteractor{
		DeploymentRepo: deploymentRepo,
		PodRepo:        podRepo,
	}
}

func (dI *deploymentInteractor) GetAll(namespace string) (map[string][]entities.Deployment, error) {
	deploymentLists := map[string][]entities.Deployment{}
	wg := sync.WaitGroup{}
	var err error
	for key, repo := range dI.DeploymentRepo {
		wg.Add(1)
		go func(namespace string, key string, repo repositories.DeploymentRepository) {
			deploymentLists[key], err = repo.GetAll(namespace)
			if err != nil {
				fmt.Println(err)
			}
			defer wg.Done()
		}(namespace, key, repo)
	}
	wg.Wait()
	return deploymentLists, nil
}
