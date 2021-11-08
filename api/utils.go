package api

import (
	"github.com/BRUHItsABunny/bigdatacloud-apis/constants"
	"net/url"
)

func DefaultUserAgent(in string) string {
	if len(in) < 1 {
		return constants.DefaultUserAgent
	}
	return in
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
