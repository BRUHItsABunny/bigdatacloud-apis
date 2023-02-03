package api

import (
	"fmt"
	"github.com/BRUHItsABunny/gOkHttp/responses"
	"net/http"
)

func IPGeolocationParser(resp *http.Response) (*IPGeolocation, error) {
	result := new(IPGeolocation)
	err := responses.ResponseJSON(resp, result)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return result, nil
}

func ReverseGeolocationParser(resp *http.Response) (*ReverseGeolocation, error) {
	result := new(ReverseGeolocation)
	err := responses.ResponseJSON(resp, result)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return result, nil
}
