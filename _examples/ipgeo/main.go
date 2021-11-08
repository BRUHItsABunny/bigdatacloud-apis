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
	res, err := c.IPGeolocationFull(context.Background(), "", "178.18.25.2")
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}
