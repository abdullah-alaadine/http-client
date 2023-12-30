package httpmethods

import (
	"fmt"
	"net/http"

	"github.com/knbr13/http-client/internal/utilities"
)

// Options sends an HTTP OPTIONS request.
func options(input Input) (*http.Response, error) {
	if input.HTTPMethod == "" || input.URL == "" {
		return nil, fmt.Errorf("missing some required arguments")
	}

	// Create an HTTP request
	httpRequest, err := http.NewRequest(http.MethodOptions, input.URL, nil)
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

	// Send the HTTP request
	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}
