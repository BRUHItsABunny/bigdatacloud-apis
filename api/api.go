package api

import (
	"context"
	"github.com/BRUHItsABunny/bigdatacloud-apis/constants"
	"net/http"
	"strconv"
)

func IPGeolocationFullRequest(ctx context.Context, apiKey, userAgent, language, ip string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", constants.EndpointIPGeolocationFull, nil)
	if err == nil {
		req.Header["user-agent"] = []string{DefaultUserAgent(userAgent)}
		parameters := DefaultParameters(apiKey, language)
		parameters[constants.ParameterKeyIP] = []string{ip}
		req.URL.RawQuery = parameters.Encode()
	}
	return req, err
}

func IPGeolocationRequest(ctx context.Context, apiKey, userAgent, language, ip string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", constants.EndpointIPGeolocation, nil)
	if err == nil {
		req.Header["user-agent"] = []string{DefaultUserAgent(userAgent)}
		parameters := DefaultParameters(apiKey, language)
		parameters[constants.ParameterKeyIP] = []string{ip}
		req.URL.RawQuery = parameters.Encode()
	}
	return req, err
}

func IPGeolocationWithConfidenceRequest(ctx context.Context, apiKey, userAgent, language, ip string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", constants.EndpointIPGeolocationWithConfidence, nil)
	if err == nil {
		req.Header["user-agent"] = []string{DefaultUserAgent(userAgent)}
		parameters := DefaultParameters(apiKey, language)
		parameters[constants.ParameterKeyIP] = []string{ip}
		req.URL.RawQuery = parameters.Encode()
	}
	return req, err
}

func ReverseGeocodeRequest(ctx context.Context, apiKey, userAgent, language string, latitude, longitude float64) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", constants.EndpointReverseGeocode, nil)
	if err == nil {
		req.Header["user-agent"] = []string{DefaultUserAgent(userAgent)}
		parameters := DefaultParameters(apiKey, language)
		parameters[constants.ParameterKeyLatitude] = []string{strconv.FormatFloat(latitude, 'f', -1, 64)}
		parameters[constants.ParameterKeyLongitude] = []string{strconv.FormatFloat(longitude, 'f', -1, 64)}
		req.URL.RawQuery = parameters.Encode()
	}
	return req, err
}

func ReverseGeocodeWithTimezoneRequest(ctx context.Context, apiKey, userAgent, language string, latitude, longitude float64) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", constants.EndpointReverseGeocodeWithTimezone, nil)
	if err == nil {
		req.Header["user-agent"] = []string{DefaultUserAgent(userAgent)}
		parameters := DefaultParameters(apiKey, language)
		parameters[constants.ParameterKeyLatitude] = []string{strconv.FormatFloat(latitude, 'f', -1, 64)}
		parameters[constants.ParameterKeyLongitude] = []string{strconv.FormatFloat(longitude, 'f', -1, 64)}
		req.URL.RawQuery = parameters.Encode()
	}
	return req, err
}
