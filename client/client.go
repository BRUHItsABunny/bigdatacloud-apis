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

	result, err := api.IPGeolocationParser(resp)
	if err != nil {
		return nil, fmt.Errorf("api.IPGeolocationParser: %w", err)
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

	result, err := api.IPGeolocationParser(resp)
	if err != nil {
		return nil, fmt.Errorf("api.IPGeolocationParser: %w", err)
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

	result, err := api.IPGeolocationParser(resp)
	if err != nil {
		return nil, fmt.Errorf("api.IPGeolocationParser: %w", err)
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

	result, err := api.ReverseGeolocationParser(resp)
	if err != nil {
		return nil, fmt.Errorf("api.ReverseGeolocationParser: %w", err)
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

	result, err := api.ReverseGeolocationParser(resp)
	if err != nil {
		return nil, fmt.Errorf("api.ReverseGeolocationParser: %w", err)
	}
	return result, nil
}
