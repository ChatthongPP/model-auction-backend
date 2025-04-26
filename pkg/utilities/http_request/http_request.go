package http_request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Result struct
type Result struct {
	Error error
	Data  interface{}
}

type ErrorResponse struct {
	ErrorCode   interface{} `json:"code"`
	MessageCode string      `json:"messageCode"`
	Message     string      `json:"message"`
}

type other_http_header map[string]string

// SuccessFn type
type SuccessFn func(bodyByte []byte) interface{}

// FailureFn type
type FailureFn func(response *http.Response, bodyByte []byte) error

// HandleResponse func
func HandleResponse(request *http.Request, successFn SuccessFn, failureFn FailureFn) *Result {
	result := &Result{}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		result.Error = err

		return result
	}
	defer response.Body.Close()

	bodyByte, err := io.ReadAll(response.Body)
	if err != nil {
		result.Error = err
		return result
	}

	if response.StatusCode >= http.StatusOK && response.StatusCode < http.StatusMultipleChoices {
		result.Data = successFn(bodyByte)
	} else {
		result.Error = failureFn(response, bodyByte)
	}

	return result
}

func BuildRequest(method, url string, ApiKey string, rawBody interface{}, header ...other_http_header) (*http.Request, error) {
	var request *http.Request
	var err error

	if rawBody != nil {
		jsonStr, err := json.Marshal(rawBody)
		if err != nil {
			return nil, err
		}

		body := bytes.NewBuffer(jsonStr)

		request, err = http.NewRequest(method, url, body)
		if err != nil {
			return nil, err
		}

		request.Header.Set("Content-Type", "application/json")
	} else {
		request, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
	}

	request.Header.Add("x-api-key", ApiKey)
	if len(header) > 0 {
		for k, v := range header[0] {
			request.Header.Set(k, v)
		}
	}

	return request, nil
}

func GetErrorResponse(errorResponse ErrorResponse) string {
	errorMessage := errorResponse.MessageCode
	if errorMessage == "" {
		errorMessage = errorResponse.Message
	}

	return errorMessage
}
