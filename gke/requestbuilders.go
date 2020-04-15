package gke

import (
	"fmt"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

func GetDeleteClusterRequest(name string, project string, region string) *containerpb.DeleteClusterRequest {
	return &containerpb.DeleteClusterRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s-c/clusters/%s", project, region, name),
	}
}

func GetGetClusterRequest(name string, project string, region string) *containerpb.GetClusterRequest {
	return &containerpb.GetClusterRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s-c/clusters/%s", project, region, name),
	}
}

func GetCreateClusterRequest(name string, project string, region string) *containerpb.CreateClusterRequest {
	return &containerpb.CreateClusterRequest{
		Cluster: &containerpb.Cluster{
			Name:            name,
			Description:     "A test cluster",
			Network:         fmt.Sprintf("projects/%s/global/networks/default", project),
			ClusterIpv4Cidr: "",
			AddonsConfig: &containerpb.AddonsConfig{
				KubernetesDashboard: &containerpb.KubernetesDashboard{
					Disabled: true,
				},
			},
			Subnetwork: fmt.Sprintf("projects/%s/regions/%s/subnetworks/default", project, region),
			NodePools: []*containerpb.NodePool{
				&containerpb.NodePool{
					Name: "default",
					Config: &containerpb.NodeConfig{
						MachineType: "f1-micro",
						DiskSizeGb:  100,
						OauthScopes: []string{
							"https://www.googleapis.com/auth/devstorage.read_only",
							"https://www.googleapis.com/auth/logging.write",
							"https://www.googleapis.com/auth/monitoring",
							"https://www.googleapis.com/auth/servicecontrol",
							"https://www.googleapis.com/auth/service.management",
							"https://www.googleapis.com/auth/trace.append",
							"https://www.googleapis.com/auth/cloud-platform",
							"https://www.googleapis.com/auth/compute",
						},
						Metadata:  map[string]string{"disable-legacy-endpoints": "true"},
						ImageType: "COS",
						DiskType:  "pd-standard",
					},
					InitialNodeCount: 3,
					Version:          "1.14.10-gke.27",
					Management: &containerpb.NodeManagement{
						AutoUpgrade: false,
						AutoRepair:  false,
					},
				},
			},
			IpAllocationPolicy: &containerpb.IPAllocationPolicy{
				UseIpAliases: true,
			},
			DefaultMaxPodsConstraint: &containerpb.MaxPodsConstraint{
				MaxPodsPerNode: 110,
			},
			DatabaseEncryption: &containerpb.DatabaseEncryption{
				State: containerpb.DatabaseEncryption_DECRYPTED,
			},
		},
		Parent: fmt.Sprintf("projects/%s/locations/us-central1-c", project),
	}
}
