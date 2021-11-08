package bigdatacloud_apis

import (
	bdcc "github.com/BRUHItsABunny/bigdatacloud-apis/client"
	"net/http"
)

func NewBigDataCloudClient(client *http.Client, apiKey, userAgent string) *bdcc.BigDataCloudClient {
	if client == nil {
		client = http.DefaultClient
	}

	return &bdcc.BigDataCloudClient{
		Client:    client,
		APIKey:    apiKey,
		UserAgent: userAgent,
	}
}
