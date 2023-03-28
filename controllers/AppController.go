package controllers

// AppController Init for controller
type AppController struct {
	Pod        interface{ PodController }
	Deployment interface{ DeploymentController }
}
