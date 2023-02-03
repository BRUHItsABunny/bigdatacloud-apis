package bigdatacloud_apis

import (
	bdcc "github.com/BRUHItsABunny/bigdatacloud-apis/client"
	andutils "github.com/BRUHItsABunny/go-android-utils"
	"net/http"
)

func NewBigDataCloudClient(hClient *http.Client, device *andutils.Device, apiKey string) *bdcc.BigDataCloudClient {
	return bdcc.NewBigDataCloudClient(hClient, device, apiKey)
}
