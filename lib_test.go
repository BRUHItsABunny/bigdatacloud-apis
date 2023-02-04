package bigdatacloud_apis

import (
	"context"
	"fmt"
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
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

func TestValidatePhoneNumber(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	apiKey := os.Getenv("VAL_API_KEY")
	resultLang := os.Getenv("VAL_LANG")
	phoneCountry := os.Getenv("VAL_COUNTRY_PHONEVALIDATE")
	phoneNumber := os.Getenv("VAL_NUMBER_PHONEVALIDATE")
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewBigDataCloudClient(hClient, nil, apiKey)
	res, err := c.ValidatePhoneNumber(context.Background(), resultLang, phoneCountry, phoneNumber)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestVerifyEmailAddress(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	apiKey := os.Getenv("VAL_API_KEY")
	resultLang := os.Getenv("VAL_LANG")
	emailAddress := os.Getenv("VAL_EMAIL_EMAILVERIFY")
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewBigDataCloudClient(hClient, nil, apiKey)
	res, err := c.VerifyEmailAddress(context.Background(), resultLang, emailAddress)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestNetworkByIP(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	apiKey := os.Getenv("VAL_API_KEY")
	resultLang := os.Getenv("VAL_LANG")
	ipAddress := os.Getenv("VAL_IP_GEOLOCATION")
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewBigDataCloudClient(hClient, nil, apiKey)
	res, err := c.NetworkByIP(context.Background(), resultLang, ipAddress)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestUserRisk(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	apiKey := os.Getenv("VAL_API_KEY")
	resultLang := os.Getenv("VAL_LANG")
	ipAddress := os.Getenv("VAL_IP_GEOLOCATION")
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewBigDataCloudClient(hClient, nil, apiKey)
	res, err := c.UserRisk(context.Background(), resultLang, ipAddress)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestUserAgentInfo(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	apiKey := os.Getenv("VAL_API_KEY")
	resultLang := os.Getenv("VAL_LANG")
	userAgent := os.Getenv("VAL_USERAGENT_UAPARSER")
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewBigDataCloudClient(hClient, nil, apiKey)
	if len(userAgent) < 1 {
		userAgent = "Dalvik/2.1.0 " + c.Device.GetUserAgent()
	}

	res, err := c.UserAgentInfo(context.Background(), resultLang, userAgent)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestASNInfoFull(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	apiKey := os.Getenv("VAL_API_KEY")
	resultLang := os.Getenv("VAL_LANG")
	asn := os.Getenv("VAL_ASN_ASNFULL")
	peersCap, err := strconv.ParseInt(os.Getenv("VAL_PEERSCAP_ASNFULL"), 10, 64)
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
	res, err := c.ASNInfoFull(context.Background(), resultLang, asn, int(peersCap))
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestIPStatsPerCountry(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	apiKey := os.Getenv("VAL_API_KEY")
	resultLang := os.Getenv("VAL_LANG")
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewBigDataCloudClient(hClient, nil, apiKey)
	res, err := c.IPStatsPerCountry(context.Background(), resultLang)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestSubmitLocation(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	apiKey := os.Getenv("VAL_API_KEY_SUBMITLOCATION")
	resultLang := os.Getenv("VAL_LANG")
	latitude, err := strconv.ParseFloat(os.Getenv("VAL_LAT_SUBMITLOCATION"), 64)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	longitude, err := strconv.ParseFloat(os.Getenv("VAL_LONG_SUBMITLOCATION"), 64)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	accuracy, err := strconv.ParseFloat(os.Getenv("VAL_ACCURACY_SUBMITLOCATION"), 64)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	isVPN := os.Getenv("VAL_ISVPN_SUBMITLOCATION") == "true"
	isCell := os.Getenv("VAL_ISCELL_SUBMITLOCATION") == "true"
	isRoaming := os.Getenv("VAL_ISROAMING_SUBMITLOCATION") == "true"
	isProxy := os.Getenv("VAL_ISPROXY_SUBMITLOCATION") == "true"
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewBigDataCloudClient(hClient, nil, apiKey)
	err = c.SubmitLocation(context.Background(), resultLang, strings.ReplaceAll(uuid.New().String(), "-", ""), latitude, longitude, accuracy, isVPN, isCell, isRoaming, isProxy)
	if err != nil {
		fmt.Println("err: ", err)
	}
}
