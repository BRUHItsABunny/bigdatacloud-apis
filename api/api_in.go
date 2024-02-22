package api

import (
	"fmt"
	"github.com/BRUHItsABunny/gOkHttp/responses"
	"net/http"
)

func ErrorParser(resp *http.Response, fallbackErr error) error {
	errRsp := new(ErrorResponse)
	err := gokhttp_responses.ResponseJSON(resp, errRsp)
	if err != nil {
		return fmt.Errorf("gokhttp_responses.CheckHTTPCode: %w", fallbackErr)
	}
	return errRsp
}

func SubmitResultParser(resp *http.Response) error {
	err := gokhttp_responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return ErrorParser(resp, err)
	}

	_, err = gokhttp_responses.ResponseBytes(resp)
	if err != nil {
		return fmt.Errorf("gokhttp_responses.ResponseJSON: %w", err)
	}
	return nil
}

func ObjectParser[T any](obj T, resp *http.Response) (*T, error) {
	result := new(T)
	err := gokhttp_responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return nil, ErrorParser(resp, err)
	}

	err = gokhttp_responses.ResponseJSON(resp, result)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_responses.ResponseJSON: %w", err)
	}
	return result, nil
}
