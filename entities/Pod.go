package entities

// Pod is the struct for save the resource info
type Pod struct {
	Name              string
	Namespace         string
	State             string
	ContainerStatuses []ContainerStatuses
}

// ContainerStatuses the struct for save container status
type ContainerStatuses struct {
	Name         string
	State        string
	RestartCount int32
	Image        string
}
