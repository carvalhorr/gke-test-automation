package main

import (
	"fmt"
	"github.com/carvalhorr/k8s-test-automation/gke"
)

func main() {
	var clusterName = "cluster1"
	gcpService, err := gke.NewGKEService("./config/testing-k8s-poc-deem-email.json", "testing-k8s-poc", "us-central1")
	if err != nil {
		panic("application could not start")
	}
	gcpService.CreateCluster(clusterName)
	fmt.Println(gcpService.GetClusterInfo(clusterName))
	gcpService.WaitClusterReady(clusterName)
	gcpService.DeleteCluster(clusterName)
}
