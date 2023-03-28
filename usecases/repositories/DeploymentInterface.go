package repositories

import (
	"sre-cli/entities"
)

// DeploymentRepository Interface for comunicate with deployment repo
type DeploymentRepository interface {
	GetAll(namespace string) ([]entities.Deployment, error)
	GetDeployment(string) (*entities.Deployment, error)
}
