package go_playflow

import (
	"encoding/json"
	"net/http"
)

func (c *Client) ServerTags() (interface{}, error) {
	response, err := c.request("/server_tags", http.MethodGet)
	if err != nil {
		return nil, err
	}

	var responseData interface{} //Todo change this

	err = json.Unmarshal(response, responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func (c *Client) DeleteServerTag(serverTag string) (interface{}, error) {
	var params = map[string]string{
		"server-tag": serverTag,
	}

	response, err := c.request("/server_tags", http.MethodDelete, params)
	if err != nil {
		return nil, err
	}

	var responseData interface{} //Todo change this

	err = json.Unmarshal(response, responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func (c *Client) UploadServerFiles(fileToUpload *[]byte, serverTag ...string) error {
	return nil
}

func (c *Client) UploadVersion() (string, error) {
	return "", nil
}
