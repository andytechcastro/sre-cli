package registries

import (
	"sre-cli/controllers"
	"sre-cli/repositories"
	"sre-cli/usecases/interactors"
	interRepo "sre-cli/usecases/repositories"
)

func (r *registry) NewPodController() controllers.PodController {
	repos := map[string]interRepo.PodRepository{}
	for key, client := range r.clients {
		repo := repositories.NewPodRepository(client)
		repos[key] = repo
	}
	podInteractor := interactors.NewPodInteractor(repos)
	return controllers.NewPodController(podInteractor)
}
