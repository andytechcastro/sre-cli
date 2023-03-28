package registries

import (
	"sre-cli/controllers"

	"k8s.io/client-go/kubernetes"
)

type registry struct {
	clients map[string]kubernetes.Interface
}

// Registry registry for all the layers
type Registry interface {
	NewAppController() controllers.AppController
}

// NewRegistry return a new registry
func NewRegistry(clients map[string]kubernetes.Interface) Registry {
	return &registry{clients}
}

func (r *registry) NewAppController() controllers.AppController {
	return controllers.AppController{
		Pod:        r.NewPodController(),
		Deployment: r.NewDeploymentController(),
	}
}
