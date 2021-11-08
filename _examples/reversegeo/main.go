package main

import (
	"context"
	"fmt"
	"github.com/BRUHItsABunny/bigdatacloud-apis/client"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	userAgent := ""
	key := ""

	c := client.NewBigDataCloudClient(nil, key, userAgent)
	res, err := c.ReverseGeocodeWithTimezoneRequest(context.Background(), "", 52.370216, 4.895168)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}
