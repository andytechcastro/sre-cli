package datastore

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// NewClients return kubernetes clients
func NewClients() (map[string]kubernetes.Interface, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")
	if err != nil {
		return nil, err
	}

	name := getAllContext(kubeConfigPath)

	clients := map[string]kubernetes.Interface{}

	for key := range name {
		kubeConfig, _ := buildConfigFromFlags(key, kubeConfigPath)
		//kubeConfig, err := clientcmd.BuildConfigFromFlags(cluster, kubeConfigPath)
		if err != nil {
			return nil, err
		}
		client, err := kubernetes.NewForConfig(kubeConfig)
		if err != nil {
			return nil, err
		}
		clients[key] = client
	}
	return clients, nil
}

func buildConfigFromFlags(context, kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}

func getAllContext(pathToKubeConfig string) map[string]*api.Context {
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: pathToKubeConfig},
		&clientcmd.ConfigOverrides{
			CurrentContext: "",
		}).RawConfig()

	if err != nil {
		fmt.Println(err)
	}

	return config.Contexts
}
