package api

import (
	"fmt"
	"github.com/BRUHItsABunny/gOkHttp/responses"
	"net/http"
)

func ErrorParser(resp *http.Response, fallbackErr error) error {
	errRsp := new(ErrorResponse)
	err := responses.ResponseJSON(resp, errRsp)
	if err != nil {
		return fmt.Errorf("responses.CheckHTTPCode: %w", fallbackErr)
	}
	return errRsp
}

func IPGeolocationParser(resp *http.Response) (*IPGeolocation, error) {
	result := new(IPGeolocation)
	err := responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return nil, ErrorParser(resp, err)
	}

	err = responses.ResponseJSON(resp, result)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return result, nil
}

func ReverseGeolocationParser(resp *http.Response) (*ReverseGeolocation, error) {
	result := new(ReverseGeolocation)
	err := responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return nil, ErrorParser(resp, err)
	}

	err = responses.ResponseJSON(resp, result)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return result, nil
}

func PhoneValidationParser(resp *http.Response) (*PhoneValidation, error) {
	result := new(PhoneValidation)
	err := responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return nil, ErrorParser(resp, err)
	}

	err = responses.ResponseJSON(resp, result)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return result, nil
}

func EmailVerificationParser(resp *http.Response) (*EmailVerification, error) {
	result := new(EmailVerification)
	err := responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return nil, ErrorParser(resp, err)
	}

	err = responses.ResponseJSON(resp, result)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return result, nil
}

func NetworkParser(resp *http.Response) (*Network, error) {
	result := new(Network)
	err := responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return nil, ErrorParser(resp, err)
	}

	err = responses.ResponseJSON(resp, result)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return result, nil
}

func UserRiskParser(resp *http.Response) (*UserRisk, error) {
	result := new(UserRisk)
	err := responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return nil, ErrorParser(resp, err)
	}

	err = responses.ResponseJSON(resp, result)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return result, nil
}

func UserAgentInfoParser(resp *http.Response) (*UserAgentInfo, error) {
	result := new(UserAgentInfo)
	err := responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return nil, ErrorParser(resp, err)
	}

	err = responses.ResponseJSON(resp, result)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return result, nil
}

func CarrierParser(resp *http.Response) (*Carrier, error) {
	result := new(Carrier)
	err := responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return nil, ErrorParser(resp, err)
	}

	err = responses.ResponseJSON(resp, result)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return result, nil
}

func IPStatsPerCountryParser(resp *http.Response) (*IPStatsPerCountry, error) {
	result := new(IPStatsPerCountry)
	err := responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return nil, ErrorParser(resp, err)
	}

	err = responses.ResponseJSON(resp, result)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return result, nil
}

func SubmitResultParser(resp *http.Response) error {
	result := new(interface{})
	err := responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return ErrorParser(resp, err)
	}

	err = responses.ResponseJSON(resp, result)
	if err != nil {
		return fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return nil
}
