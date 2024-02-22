package api

import (
	"context"
	"fmt"
	"github.com/BRUHItsABunny/bigdatacloud-apis/constants"
	"github.com/BRUHItsABunny/gOkHttp/requests"
	andutils "github.com/BRUHItsABunny/go-android-utils"
	"net/http"
	"net/url"
	"strconv"
)

func IPGeolocationFullRequest(ctx context.Context, device *andutils.Device, apiKey, language, ip string) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyIP] = []string{ip}

	req, err := gokhttp_requests.MakeGETRequest(ctx, constants.EndpointIPGeolocationFull,
		gokhttp_requests.NewHeaderOption(DefaultHeaders(device)),
		gokhttp_requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func IPGeolocationRequest(ctx context.Context, device *andutils.Device, apiKey, language, ip string) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyIP] = []string{ip}

	req, err := gokhttp_requests.MakeGETRequest(ctx, constants.EndpointIPGeolocation,
		gokhttp_requests.NewHeaderOption(DefaultHeaders(device)),
		gokhttp_requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func IPGeolocationWithConfidenceRequest(ctx context.Context, device *andutils.Device, apiKey, language, ip string) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyIP] = []string{ip}

	req, err := gokhttp_requests.MakeGETRequest(ctx, constants.EndpointIPGeolocationWithConfidence,
		gokhttp_requests.NewHeaderOption(DefaultHeaders(device)),
		gokhttp_requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func ReverseGeocodeRequest(ctx context.Context, device *andutils.Device, apiKey, language string, latitude, longitude float64) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyLatitude] = []string{strconv.FormatFloat(latitude, 'f', -1, 64)}
	parameters[constants.ParameterKeyLongitude] = []string{strconv.FormatFloat(longitude, 'f', -1, 64)}

	req, err := gokhttp_requests.MakeGETRequest(ctx, constants.EndpointReverseGeocode,
		gokhttp_requests.NewHeaderOption(DefaultHeaders(device)),
		gokhttp_requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func ReverseGeocodeWithTimezoneRequest(ctx context.Context, device *andutils.Device, apiKey, language string, latitude, longitude float64) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyLatitude] = []string{strconv.FormatFloat(latitude, 'f', -1, 64)}
	parameters[constants.ParameterKeyLongitude] = []string{strconv.FormatFloat(longitude, 'f', -1, 64)}

	req, err := gokhttp_requests.MakeGETRequest(ctx, constants.EndpointReverseGeocodeWithTimezone,
		gokhttp_requests.NewHeaderOption(DefaultHeaders(device)),
		gokhttp_requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func ValidatePhoneNumberRequest(ctx context.Context, device *andutils.Device, apiKey, language, countryCode, phoneNumber string) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyCountryCode] = []string{countryCode}
	parameters[constants.ParameterKeyNumber] = []string{phoneNumber}

	req, err := gokhttp_requests.MakeGETRequest(ctx, constants.EndpointPhoneNumberValidate,
		gokhttp_requests.NewHeaderOption(DefaultHeaders(device)),
		gokhttp_requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func ValidatePhoneNumberByIPRequest(ctx context.Context, device *andutils.Device, apiKey, language, ipAddress, phoneNumber string) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyIP] = []string{ipAddress}
	parameters[constants.ParameterKeyNumber] = []string{phoneNumber}

	req, err := gokhttp_requests.MakeGETRequest(ctx, constants.EndpointPhoneNumberValidateByIP,
		gokhttp_requests.NewHeaderOption(DefaultHeaders(device)),
		gokhttp_requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func VerifyEmailAddressRequest(ctx context.Context, device *andutils.Device, apiKey, language, emailAddress string) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyEmailAddress] = []string{emailAddress}

	req, err := gokhttp_requests.MakeGETRequest(ctx, constants.EndpointEmailVerify,
		gokhttp_requests.NewHeaderOption(DefaultHeaders(device)),
		gokhttp_requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func NetworkByIPRequest(ctx context.Context, device *andutils.Device, apiKey, language, ipAddress string) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyIP] = []string{ipAddress}

	req, err := gokhttp_requests.MakeGETRequest(ctx, constants.EndpointNetworkByIP,
		gokhttp_requests.NewHeaderOption(DefaultHeaders(device)),
		gokhttp_requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func UserRiskRequest(ctx context.Context, device *andutils.Device, apiKey, language, ipAddress string) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyIP] = []string{ipAddress}

	req, err := gokhttp_requests.MakeGETRequest(ctx, constants.EndpointUserRisk,
		gokhttp_requests.NewHeaderOption(DefaultHeaders(device)),
		gokhttp_requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func UserAgentInfoRequest(ctx context.Context, device *andutils.Device, apiKey, language, userAgent string) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyUserAgent] = []string{userAgent}

	req, err := gokhttp_requests.MakeGETRequest(ctx, constants.EndpointUserAgentInfo,
		gokhttp_requests.NewHeaderOption(DefaultHeaders(device)),
		gokhttp_requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func ASNInfoFullRequest(ctx context.Context, device *andutils.Device, apiKey, language, asn string, peersCap int) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters[constants.ParameterKeyASN] = []string{asn}
	parameters[constants.ParameterKeyPeersCap] = []string{strconv.FormatInt(int64(peersCap), 10)}

	req, err := gokhttp_requests.MakeGETRequest(ctx, constants.EndpointASNInfoFull,
		gokhttp_requests.NewHeaderOption(DefaultHeaders(device)),
		gokhttp_requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func IPStatsPerCountryRequest(ctx context.Context, device *andutils.Device, apiKey, language string) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)

	req, err := gokhttp_requests.MakeGETRequest(ctx, constants.EndpointIPGeoPerCountry,
		gokhttp_requests.NewHeaderOption(DefaultHeaders(device)),
		gokhttp_requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func SubmitLocationRequest(ctx context.Context, device *andutils.Device, apiKey, language, uuid string, latitude, longitude, accuracy float64, isVPN, isCell, isRoaming, isProxy bool) (*http.Request, error) {
	parameters := DefaultParameters(apiKey, language)
	parameters["uuid"] = []string{uuid}
	postForm := url.Values{
		"Lat":       {strconv.FormatFloat(latitude, 'f', 3, 64)},
		"Lon":       {strconv.FormatFloat(longitude, 'f', 3, 64)},
		"Acc":       {strconv.FormatFloat(accuracy, 'f', 0, 64) + ".0"},
		"isVPN":     {strconv.FormatBool(isVPN)},
		"isCell":    {strconv.FormatBool(isCell)},
		"isRoaming": {strconv.FormatBool(isRoaming)},
		"isProxy":   {strconv.FormatBool(isProxy)},
	}

	req, err := gokhttp_requests.MakePOSTRequest(ctx, constants.EndpointSubmitLocation,
		gokhttp_requests.NewPOSTFormOption(postForm),
		gokhttp_requests.NewHeaderOption(DefaultHeaders(device)),
		gokhttp_requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_requests.MakePOSTRequest: %w", err)
	}
	return req, nil
}
