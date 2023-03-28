package interactors

import (
	"sre-cli/entities"
	"sre-cli/usecases/repositories"
)

type podInteractor struct {
	PodRepo map[string]repositories.PodRepository
}

// PodInteractor is the interface for connect with this struct
type PodInteractor interface {
	GetAll(string) (map[string][]entities.Pod, error)
}

// NewPodInteractor return an struct of tyoe operatorInteractor
func NewPodInteractor(podRepo map[string]repositories.PodRepository) PodInteractor {
	return &podInteractor{
		PodRepo: podRepo,
	}
}

func (oi *podInteractor) GetAll(namespace string) (map[string][]entities.Pod, error) {
	podLists := map[string][]entities.Pod{}
	for key, repo := range oi.PodRepo {
		podList, err := repo.GetAll(namespace)
		if err != nil {
			return podLists, nil
		}
		podLists[key] = podList
	}
	return podLists, nil
}
