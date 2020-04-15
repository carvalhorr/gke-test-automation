package main

import (
	"fmt"
	"github.com/carvalhorr/k8s-test-automation/gke"
	"github.com/carvalhorr/k8s-test-automation/k8s"
)

func main() {
	getK8sServicePort()
	// createClusterOnGKE()
}

func getK8sServicePort() {
	k8sService := k8s.NewK8sService("./config/kube_config_qa9", "rcqa9")
	port, _ := k8sService.GetPortForService("grpc-mock-rcqa9", "grpc-service")
	fmt.Println(port)
}

func createClusterOnGKE() {
	var clusterName = "cluster4"
	gcpService, err := gke.NewGKEService("./config/testing-k8s-poc-deem-email.json", "testing-k8s-poc", "us-central1")
	if err != nil {
		panic("application could not start")
	}
	gcpService.CreateCluster(clusterName)
	fmt.Println(gcpService.GetClusterInfo(clusterName))
	gcpService.WaitClusterReady(clusterName)
	fmt.Println("")
	fmt.Println("")
	fmt.Println(gcpService.CreateK8sConfig(clusterName))
	fmt.Println("")
	fmt.Println("")
	gcpService.DeleteCluster(clusterName)

}
