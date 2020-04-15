package k8s

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Create a new k8s service
// Takes a ./kube/config file path and a namespace. The namespace is optional.
func NewK8sService(k8sConfigFile string, namespace string) K8sFacade {
	return &k8sService{
		ConfigFile: k8sConfigFile,
		Namespace:  namespace,
	}
}

// Expose functions to interact with a kubernetes cluster
type K8sFacade interface {
	GetPortForService(serviceName string, portName string) (int32, error)
}

type k8sService struct {
	ConfigFile string
	Namespace  string
}

// Return the external port that a service is running.
// Takes the service name and the port name as parameters. This function expects the port is named
func (k8s *k8sService) GetPortForService(serviceName string, portName string) (int32, error) {
	config, err := clientcmd.BuildConfigFromFlags("", k8s.ConfigFile)
	if err != nil {
		return 0, fmt.Errorf("error getting configurations: %s", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return 0, fmt.Errorf("error creating client: %s", err)
	}

	services, err := clientset.CoreV1().Services(k8s.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return 0, fmt.Errorf("error getting services in the cluster: %s", err)
	}
	for _, service := range services.Items {
		if service.Name == serviceName {
			for _, port := range service.Spec.Ports {
				if port.Name == portName {
					return port.NodePort, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("could not find service/port")
}
