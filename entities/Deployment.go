package entities

// Deployment the struct for the deployment information
type Deployment struct {
	Name              string
	Namespace         string
	Replicas          int32
	MatchLabels       map[string]string
	AvailableReplicas int32
	ReadyReplicas     int32
	UpdatedReplicas   int32
	PodList           []Pod
}
