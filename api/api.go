package api

import (
	"context"
	"fmt"
	"github.com/BRUHItsABunny/bigdatacloud-apis/constants"
	"github.com/BRUHItsABunny/gOkHttp/requests"
	andutils "github.com/BRUHItsABunny/go-android-utils"
	"net/http"
	"strconv"
)

func IPGeolocationFullRequest(ctx context.Context, device *andutils.Device, apiKey, language, ip string) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyIP] = []string{ip}

	req, err := requests.MakeGETRequest(ctx, constants.EndpointIPGeolocationFull,
		requests.NewHeaderOption(DefaultHeaders(device)),
		requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func IPGeolocationRequest(ctx context.Context, device *andutils.Device, apiKey, language, ip string) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyIP] = []string{ip}

	req, err := requests.MakeGETRequest(ctx, constants.EndpointIPGeolocation,
		requests.NewHeaderOption(DefaultHeaders(device)),
		requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func IPGeolocationWithConfidenceRequest(ctx context.Context, device *andutils.Device, apiKey, language, ip string) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyIP] = []string{ip}

	req, err := requests.MakeGETRequest(ctx, constants.EndpointIPGeolocationWithConfidence,
		requests.NewHeaderOption(DefaultHeaders(device)),
		requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func ReverseGeocodeRequest(ctx context.Context, device *andutils.Device, apiKey, language string, latitude, longitude float64) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyLatitude] = []string{strconv.FormatFloat(latitude, 'f', -1, 64)}
	parameters[constants.ParameterKeyLongitude] = []string{strconv.FormatFloat(longitude, 'f', -1, 64)}

	req, err := requests.MakeGETRequest(ctx, constants.EndpointReverseGeocode,
		requests.NewHeaderOption(DefaultHeaders(device)),
		requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func ReverseGeocodeWithTimezoneRequest(ctx context.Context, device *andutils.Device, apiKey, language string, latitude, longitude float64) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyLatitude] = []string{strconv.FormatFloat(latitude, 'f', -1, 64)}
	parameters[constants.ParameterKeyLongitude] = []string{strconv.FormatFloat(longitude, 'f', -1, 64)}

	req, err := requests.MakeGETRequest(ctx, constants.EndpointReverseGeocodeWithTimezone,
		requests.NewHeaderOption(DefaultHeaders(device)),
		requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}
	return req, nil
}
