package interactors

import (
	"fmt"
	"sre-cli/entities"
	"sre-cli/usecases/repositories"
	"sync"
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
	wg := sync.WaitGroup{}
	var err error
	for key, repo := range oi.PodRepo {
		wg.Add(1)
		go func(namespace string, key string, repo repositories.PodRepository) {
			podLists[key], err = repo.GetAll(namespace)
			if err != nil {
				fmt.Println(err)
			}
			defer wg.Done()
		}(namespace, key, repo)
	}
	wg.Wait()
	return podLists, nil
}
