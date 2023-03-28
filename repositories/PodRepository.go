package repositories

import (
	"context"
	"fmt"
	"sre-cli/entities"
	"sre-cli/usecases/repositories"

	"k8s.io/apimachinery/pkg/labels"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type podRepository struct {
	client      kubernetes.Interface
	environment string
}

// NewPodRepository get the struct with the kubernetes client
func NewPodRepository(client kubernetes.Interface) repositories.PodRepository {
	return &podRepository{client: client}
}

func (op *podRepository) GetAll(namespace string) ([]entities.Pod, error) {
	podList, err := op.client.CoreV1().Pods(namespace).
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	podResources := []entities.Pod{}
	for _, pod := range podList.Items {
		podResource := addPodtoEntity(pod)
		podResources = append(podResources, podResource)
	}
	return podResources, nil
}

func (op *podRepository) GetPodByLabels(namespace string, labelSelector map[string]string) ([]entities.Pod, error) {
	options := metav1.ListOptions{
		LabelSelector: labels.SelectorFromSet(labelSelector).String(),
	}
	podList, err := op.client.CoreV1().Pods(namespace).
		List(context.Background(), options)
	if err != nil {
		fmt.Println(err)
	}
	podResources := []entities.Pod{}
	for _, pod := range podList.Items {
		podResource := addPodtoEntity(pod)
		podResources = append(podResources, podResource)

	}
	return podResources, nil
}

func (op *podRepository) getPod(name string) (*entities.Pod, error) {
	pod, err := op.client.CoreV1().Pods("").Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	podResource := addPodtoEntity(*pod)
	return &podResource, nil
}

func addPodtoEntity(pod v1.Pod) entities.Pod {
	statuses := []entities.ContainerStatuses{}
	for _, status := range pod.Status.ContainerStatuses {
		statusResource := entities.ContainerStatuses{
			Name:         status.Name,
			State:        status.State.String(),
			RestartCount: status.RestartCount,
			Image:        status.Image,
		}
		statuses = append(statuses, statusResource)
	}
	podResource := entities.Pod{
		Name:              pod.Name,
		Namespace:         pod.Namespace,
		State:             string(pod.Status.Phase),
		ContainerStatuses: statuses,
	}
	return podResource
}
