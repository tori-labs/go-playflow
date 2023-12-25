package go_playflow

import (
	"net/http"
	"slices"
)

var validMethods = []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodOptions}

func (c *Client) request(endpoint string, method string, headers ...map[string]string) (*http.Response, error) {
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

	//TODO!
	return nil, nil

}
