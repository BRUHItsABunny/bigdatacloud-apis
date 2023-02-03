package bigdatacloud_apis

import (
	"context"
	"fmt"
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"testing"
)

// Loads the .env file in our repo - if you have one
func loadTestEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func TestIPGeo(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	apiKey := os.Getenv("VAL_API_KEY")
	resultLang := os.Getenv("VAL_LANG")
	ip := os.Getenv("VAL_IP_GEOLOCATION")
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewBigDataCloudClient(hClient, nil, apiKey)
	res, err := c.IPGeolocationFull(context.Background(), resultLang, ip)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestReverseGeo(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	apiKey := os.Getenv("VAL_API_KEY")
	resultLang := os.Getenv("VAL_LANG")
	latitude, err := strconv.ParseFloat(os.Getenv("VAL_LAT_REVERSEGEO"), 64)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	longitude, err := strconv.ParseFloat(os.Getenv("VAL_LONG_REVERSEGEO"), 64)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewBigDataCloudClient(hClient, nil, apiKey)
	res, err := c.ReverseGeocodeWithTimezone(context.Background(), resultLang, latitude, longitude)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}
