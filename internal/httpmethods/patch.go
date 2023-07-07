package httpmethods

import (
	"net/http"
	"strings"

	"github.com/abdullah-alaadine/http-client/internal/utilities"
)

func patch(input Input) (*http.Response, error) {
	// Create an HTTP request with the request body
	httpRequest, err := http.NewRequest(http.MethodPatch, input.URL, strings.NewReader(input.Body))
	if err != nil {
		return nil, err
	}

	// Set headers
	headers, err := utilities.ParseHeaders(input.Header)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		httpRequest.Header.Set(key, value)
	}

	httpRequest.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}
