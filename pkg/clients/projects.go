package clients

import (
	"github.com/gophercloud/gophercloud"
	masterdata "github.com/sapcc/gophercloud-sapcc/billing/masterdata/projects"
)

type ProjectClient struct {
	billingClient *gophercloud.ServiceClient
}

func NewProjectClient(billingClient *gophercloud.ServiceClient) *ProjectClient {
	return &ProjectClient{
		billingClient: billingClient,
	}
}

func (pc *ProjectClient) GetProject(projectId string) (*masterdata.Project, error) {
	projectMeta := masterdata.Get(pc.billingClient, projectId)
	project, err := projectMeta.Extract()

	if err != nil {
		logger.Printf("failed to fetch project with id: %w", projectId)
		return nil, err
	}
	logger.Printf("project with id: %s, has been fetched successfully", projectId)
	return project, nil
}
