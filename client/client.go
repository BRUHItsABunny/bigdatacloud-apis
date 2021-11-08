package client

import (
	"context"
	"github.com/BRUHItsABunny/bigdatacloud-apis/api"
	"net/http"
)

type BigDataCloudClient struct {
	Client            *http.Client
	APIKey, UserAgent string
}

func NewBigDataCloudClient(client *http.Client, apiKey, userAgent string) *BigDataCloudClient {
	if client == nil {
		client = http.DefaultClient
	}

	return &BigDataCloudClient{
		Client:    client,
		APIKey:    apiKey,
		UserAgent: userAgent,
	}
}

func (c *BigDataCloudClient) IPGeolocationFull(ctx context.Context, language, ip string) (*api.IPGeolocation, error) {
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)

	req, err = api.IPGeolocationFullRequest(ctx, c.APIKey, c.UserAgent, language, ip)
	if err == nil {
		resp, err = c.Client.Do(req)
		if err == nil {
			return api.IPGeolocationParser(resp)
		}
	}
	return nil, err
}

func (c *BigDataCloudClient) IPGeolocation(ctx context.Context, language, ip string) (*api.IPGeolocation, error) {
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)

	req, err = api.IPGeolocationRequest(ctx, c.APIKey, c.UserAgent, language, ip)
	if err == nil {
		resp, err = c.Client.Do(req)
		if err == nil {
			return api.IPGeolocationParser(resp)
		}
	}
	return nil, err
}

func (c *BigDataCloudClient) IPGeolocationWithConfidence(ctx context.Context, language, ip string) (*api.IPGeolocation, error) {
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)

	req, err = api.IPGeolocationWithConfidenceRequest(ctx, c.APIKey, c.UserAgent, language, ip)
	if err == nil {
		resp, err = c.Client.Do(req)
		if err == nil {
			return api.IPGeolocationParser(resp)
		}
	}
	return nil, err
}

func (c *BigDataCloudClient) ReverseGeocodeRequest(ctx context.Context, language string, latitude, longitude float64) (*api.ReverseGeolocation, error) {
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)

	req, err = api.ReverseGeocodeRequest(ctx, c.APIKey, c.UserAgent, language, latitude, longitude)
	if err == nil {
		resp, err = c.Client.Do(req)
		if err == nil {
			return api.ReverseGeolocationParser(resp)
		}
	}
	return nil, err
}

func (c *BigDataCloudClient) ReverseGeocodeWithTimezoneRequest(ctx context.Context, language string, latitude, longitude float64) (*api.ReverseGeolocation, error) {
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)

	req, err = api.ReverseGeocodeWithTimezoneRequest(ctx, c.APIKey, c.UserAgent, language, latitude, longitude)
	if err == nil {
		resp, err = c.Client.Do(req)
		if err == nil {
			return api.ReverseGeolocationParser(resp)
		}
	}
	return nil, err
}
