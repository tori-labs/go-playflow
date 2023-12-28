package go_playflow

import (
	"io"
	"net/http"
	"slices"
)

var validMethods = []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodOptions}

func (c *Client) request(endpoint string, method string, headers ...map[string]string) ([]byte, error) {
	if !slices.Contains(validMethods, method) {
		return nil, ErrInvalidMethod
	}

	request, _ := http.NewRequest(method, c.url+endpoint, nil)
	request.Header.Add("token", c.token)

	for _, header := range headers {
		for k, v := range header {
			request.Header.Add(k, v)
		}
	}

	response, err := c.HTTP.Do(request)
	if err != nil {
		return nil, err
	}

	var body []byte
	body, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
