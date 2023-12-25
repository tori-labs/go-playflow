package go_playflow

import "net/http"

const PlayFlowProductionAPI = "https://api.cloud.playflow.app"

type Client struct {
	HTTP  *http.Client
	url   string
	token string
}

func NewClient(token string, url ...string) (*Client, error) {
	if len(url) == 0 {
		url = append(url, PlayFlowProductionAPI)
	}

	return &Client{
		url:   url[0],
		token: token,
		HTTP:  http.DefaultClient,
	}, nil
}
