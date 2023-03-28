package registries

import (
	"sre-cli/controllers"
	"sre-cli/repositories"
	"sre-cli/usecases/interactors"
	interRepo "sre-cli/usecases/repositories"
)

func (r *registry) NewDeploymentController() controllers.DeploymentController {
	deployRepos := map[string]interRepo.DeploymentRepository{}
	podRepos := map[string]interRepo.PodRepository{}
	for key, client := range r.clients {
		deployRepo := repositories.NewDeploymentRepository(client)
		podRepo := repositories.NewPodRepository(client)
		deployRepos[key] = deployRepo
		podRepos[key] = podRepo
	}
	deploymentInteractor := interactors.NewDeploymentInteractor(deployRepos, podRepos)
	return controllers.NewDeploymentController(deploymentInteractor)
}
