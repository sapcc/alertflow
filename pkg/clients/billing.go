package clients

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

// TODO: add Basic Auth
func NewBillingClient(endpoint string, auth string) (*gophercloud.ServiceClient, error) {
	client, err := openstack.NewClient("")
	if err != nil {
		return nil, err
	}
	billingClient := &gophercloud.ServiceClient{
		ProviderClient: client,
		Endpoint:       endpoint,
		Type:           "sapcc-billing",
		MoreHeaders: map[string]string{
			"X-Auth-Token": auth,
		},
	}
	return billingClient, nil
}
