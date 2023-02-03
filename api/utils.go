package api

import (
	"github.com/BRUHItsABunny/bigdatacloud-apis/constants"
	andutils "github.com/BRUHItsABunny/go-android-utils"
	"net/http"
	"net/url"
)

func DefaultUserAgent(in string) string {
	if len(in) < 1 {
		return constants.DefaultUserAgent
	}
	return in
}

func DefaultHeaders(device *andutils.Device) http.Header {
	return http.Header{
		"user-agent": {"Dalvik/2.1.0 " + device.GetUserAgent()},
	}
}

func DefaultLanguage(in string) string {
	if len(in) != 2 {
		return "en"
	}
	return in
}

func DefaultParameters(apiKey, language string) url.Values {
	return url.Values{
		constants.ParameterKeyAPIKey:   {apiKey},
		constants.ParameterKeyLanguage: {DefaultLanguage(language)},
	}
}
