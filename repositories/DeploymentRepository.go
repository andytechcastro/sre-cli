package repositories

import (
	"context"
	"fmt"
	"sre-cli/entities"
	"sre-cli/usecases/repositories"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type deploymentRepository struct {
	client      kubernetes.Interface
	environment string
}

// NewDeploymentRepository return a deploymentRepository struct
func NewDeploymentRepository(client kubernetes.Interface) repositories.DeploymentRepository {
	return &deploymentRepository{client: client}
}

func (repo *deploymentRepository) GetAll(namespace string) ([]entities.Deployment, error) {
	deploymentList, err := repo.client.AppsV1().Deployments(namespace).
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	deploymentResources := []entities.Deployment{}
	for _, deployment := range deploymentList.Items {
		deploymentResource := addDeploymentEntity(deployment)
		deploymentResources = append(deploymentResources, deploymentResource)
	}
	return deploymentResources, nil
}

func (repo *deploymentRepository) GetDeployment(name string) (*entities.Deployment, error) {
	return &entities.Deployment{}, nil
}

func addDeploymentEntity(deployment appsv1.Deployment) entities.Deployment {
	deploymentResource := entities.Deployment{
		Name:              deployment.Name,
		Namespace:         deployment.Namespace,
		Replicas:          deployment.Status.Replicas,
		AvailableReplicas: deployment.Status.AvailableReplicas,
		ReadyReplicas:     deployment.Status.ReadyReplicas,
		UpdatedReplicas:   deployment.Status.UpdatedReplicas,
		MatchLabels:       deployment.Spec.Selector.MatchLabels,
	}
	return deploymentResource
}
