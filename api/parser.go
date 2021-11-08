package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func IPGeolocationParser(resp *http.Response) (*IPGeolocation, error) {
	result := new(IPGeolocation)
	bytes, err := io.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(bytes, result)
	}
	return result, err
}

func ReverseGeolocationParser(resp *http.Response) (*ReverseGeolocation, error) {
	result := new(ReverseGeolocation)
	bytes, err := io.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(bytes, result)
	}
	return result, err
}
