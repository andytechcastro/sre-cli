package repositories

import "sre-cli/entities"

// PodRepository isn an interface for comunicate the operator repository
type PodRepository interface {
	GetAll(string) ([]entities.Pod, error)
	GetPodByLabels(string, map[string]string) ([]entities.Pod, error)
}
