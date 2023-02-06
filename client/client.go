package client

import (
	"context"
	"fmt"
	"github.com/BRUHItsABunny/bigdatacloud-apis/api"
	"github.com/BRUHItsABunny/bigdatacloud-apis/constants"
	andutils "github.com/BRUHItsABunny/go-android-utils"
	"net/http"
)

type BigDataCloudClient struct {
	Client *http.Client
	Device *andutils.Device
	APIKey string
}

func NewBigDataCloudClient(client *http.Client, device *andutils.Device, apiKey string) *BigDataCloudClient {
	if client == nil {
		client = http.DefaultClient
	}

	if device == nil {
		device = andutils.GetRandomDevice()
	}

	if len(apiKey) == 0 {
		apiKey = constants.DefaultAPIKey
	}

	return &BigDataCloudClient{
		Client: client,
		Device: device,
		APIKey: apiKey,
	}
}

func (c *BigDataCloudClient) IPGeolocationFull(ctx context.Context, language, ip string) (*api.IPGeolocation, error) {
	req, err := api.IPGeolocationFullRequest(ctx, c.Device, c.APIKey, language, ip)
	if err != nil {
		return nil, fmt.Errorf("api.IPGeolocationFullRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.IPGeolocation{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *BigDataCloudClient) IPGeolocation(ctx context.Context, language, ip string) (*api.IPGeolocation, error) {
	req, err := api.IPGeolocationRequest(ctx, c.Device, c.APIKey, language, ip)
	if err != nil {
		return nil, fmt.Errorf("api.IPGeolocationRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.IPGeolocation{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *BigDataCloudClient) IPGeolocationWithConfidence(ctx context.Context, language, ip string) (*api.IPGeolocation, error) {
	req, err := api.IPGeolocationWithConfidenceRequest(ctx, c.Device, c.APIKey, language, ip)
	if err != nil {
		return nil, fmt.Errorf("api.IPGeolocationWithConfidenceRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.IPGeolocation{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *BigDataCloudClient) ReverseGeocode(ctx context.Context, language string, latitude, longitude float64) (*api.ReverseGeolocation, error) {
	req, err := api.ReverseGeocodeRequest(ctx, c.Device, c.APIKey, language, latitude, longitude)
	if err != nil {
		return nil, fmt.Errorf("api.ReverseGeocodeRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.ReverseGeolocation{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *BigDataCloudClient) ReverseGeocodeWithTimezone(ctx context.Context, language string, latitude, longitude float64) (*api.ReverseGeolocation, error) {
	req, err := api.ReverseGeocodeWithTimezoneRequest(ctx, c.Device, c.APIKey, language, latitude, longitude)
	if err != nil {
		return nil, fmt.Errorf("api.ReverseGeocodeWithTimezoneRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.ReverseGeolocation{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *BigDataCloudClient) ValidatePhoneNumber(ctx context.Context, language, countryCode, phoneNumber string) (*api.PhoneValidation, error) {
	req, err := api.ValidatePhoneNumberRequest(ctx, c.Device, c.APIKey, language, countryCode, phoneNumber)
	if err != nil {
		return nil, fmt.Errorf("api.ValidatePhoneNumberRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.PhoneValidation{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *BigDataCloudClient) ValidatePhoneNumberByIP(ctx context.Context, language, ipAddress, phoneNumber string) (*api.PhoneValidation, error) {
	req, err := api.ValidatePhoneNumberByIPRequest(ctx, c.Device, c.APIKey, language, ipAddress, phoneNumber)
	if err != nil {
		return nil, fmt.Errorf("api.ValidatePhoneNumberRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.PhoneValidation{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *BigDataCloudClient) VerifyEmailAddress(ctx context.Context, language, emailAddress string) (*api.EmailVerification, error) {
	req, err := api.VerifyEmailAddressRequest(ctx, c.Device, c.APIKey, language, emailAddress)
	if err != nil {
		return nil, fmt.Errorf("api.VerifyEmailAddressRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.EmailVerification{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *BigDataCloudClient) NetworkByIP(ctx context.Context, language, ipAddress string) (*api.Network, error) {
	req, err := api.NetworkByIPRequest(ctx, c.Device, c.APIKey, language, ipAddress)
	if err != nil {
		return nil, fmt.Errorf("api.NetworkByIPRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.Network{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *BigDataCloudClient) UserRisk(ctx context.Context, language, ipAddress string) (*api.UserRisk, error) {
	req, err := api.UserRiskRequest(ctx, c.Device, c.APIKey, language, ipAddress)
	if err != nil {
		return nil, fmt.Errorf("api.UserRiskRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.UserRisk{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *BigDataCloudClient) UserAgentInfo(ctx context.Context, language, userAgent string) (*api.UserAgentInfo, error) {
	req, err := api.UserAgentInfoRequest(ctx, c.Device, c.APIKey, language, userAgent)
	if err != nil {
		return nil, fmt.Errorf("api.UserAgentInfoRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.UserAgentInfo{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *BigDataCloudClient) ASNInfoFull(ctx context.Context, language, asn string, peersCap int) (*api.Carrier, error) {
	req, err := api.ASNInfoFullRequest(ctx, c.Device, c.APIKey, language, asn, peersCap)
	if err != nil {
		return nil, fmt.Errorf("api.ASNInfoFullRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.Carrier{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *BigDataCloudClient) IPStatsPerCountry(ctx context.Context, language string) (*api.IPStatsPerCountry, error) {
	req, err := api.IPStatsPerCountryRequest(ctx, c.Device, c.APIKey, language)
	if err != nil {
		return nil, fmt.Errorf("api.IPStatsPerCountryRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.IPStatsPerCountry{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *BigDataCloudClient) SubmitLocation(ctx context.Context, language, uuid string, latitude, longitude, accuracy float64, isVPN, isCell, isRoaming, isProxy bool) error {
	req, err := api.SubmitLocationRequest(ctx, c.Device, c.APIKey, language, uuid, latitude, longitude, accuracy, isVPN, isCell, isRoaming, isProxy)
	if err != nil {
		return fmt.Errorf("api.SubmitLocationRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return fmt.Errorf("c.Client.Do: %w", err)
	}

	err = api.SubmitResultParser(resp)
	if err != nil {
		return fmt.Errorf("api.SubmitResultParser: %w", err)
	}
	return nil
}
