// Helper functions for creating, getting information about and deleting kubernetes clusters in Google Cloud Platform
// It provide default cluster configurations.
package gke

import (
	container "cloud.google.com/go/container/apiv1"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
	"time"
)

type GKEFacade interface {
	CreateCluster(name string) error
	GetClusterInfo(name string) (*containerpb.Cluster, error)
	DeleteCluster(name string) error
	WaitClusterReady(name string) error
}

func NewGKEService(gcpConfigFilePath string, project string, region string) (GKEFacade, error) {
	log.Infof("Creating new cluster manager client using config file: %s", gcpConfigFilePath)
	ctx := context.Background()
	options := option.WithCredentialsFile(gcpConfigFilePath)
	clusterManager, err := container.NewClusterManagerClient(ctx, options)
	if err != nil {
		log.Errorf("Failed to create cluster manager client with error: %s", err.Error())
		return nil, err
	}
	return &gkeService{clusterManager: clusterManager, project: project, region: region}, nil
}

type gkeService struct {
	clusterManager *container.ClusterManagerClient
	project        string
	region         string
}

func (gcp *gkeService) CreateCluster(name string) error {
	log.Infof("Creating cluster: %s", name)
	createReq := GetCreateClusterRequest(name, gcp.project, gcp.region)
	_, createErr := gcp.clusterManager.CreateCluster(context.Background(), createReq)
	if createErr != nil {
		log.Infof("Failed to create cluster %s: %s", name, createErr.Error())
		return createErr
	}
	return nil
}

func (gcp *gkeService) GetClusterInfo(name string) (*containerpb.Cluster, error) {
	log.Infof("Retrieving cluster info for cluster: %s", name)
	clusterInfo, err := gcp.clusterManager.GetCluster(context.Background(), GetGetClusterRequest(name, gcp.project, gcp.region))
	if err != nil {
		log.Errorf("Failed to get cluster info for cluster %s with error: %s", name, err.Error())
		return nil, err
	}
	return clusterInfo, nil
}

func (gcp *gkeService) DeleteCluster(name string) error {
	_, deleteErr := gcp.clusterManager.DeleteCluster(context.Background(), GetDeleteClusterRequest(name, gcp.project, gcp.region))
	if deleteErr != nil {
		log.Errorf("Failed to delete cluster %s with error: %s", name, deleteErr.Error())
		return deleteErr
	}
	return nil
}

func (gcp *gkeService) WaitClusterReady(name string) error {
	for {
		clusterInfo, err := gcp.GetClusterInfo(name)
		if err != nil {
			log.Errorf("Failed to get cluster info for cluster %s with error: %s", name, err.Error())
			return err
		}
		switch clusterInfo.Status {
		case containerpb.Cluster_RUNNING:
			return nil
		case containerpb.Cluster_ERROR:
			return fmt.Errorf("creation of cluster %s failed", name)

		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
