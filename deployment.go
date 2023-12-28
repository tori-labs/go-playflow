package go_playflow

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func (c *Client) StartGameServer(serverArguments, serverRegion, serverType, serverTag *string) (interface{}, error) {
	var params = map[string]string{}

	if serverArguments != nil {
		params["arguments"] = *serverArguments
	}

	if serverRegion != nil {
		params["region"] = *serverRegion
	}

	if serverType != nil {
		params["type"] = *serverType
	} else {
		params["type"] = "small"
	}

	if serverTag != nil {
		params["server-tag"] = *serverTag
	}

	response, err := c.request("/start_game_server", http.MethodPost, params)
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

func (c *Client) StopGameServer(matchId string) (interface{}, error) {
	var params = map[string]string{
		"match-id": matchId,
	}

	response, err := c.request("/stop_game_server", http.MethodDelete, params)
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

func (c *Client) RestartGameServer(serverArguments, serverType, serverTag *string) (interface{}, error) {
	var params = map[string]string{}

	if serverArguments != nil {
		params["arguments"] = *serverArguments
	}

	if serverType != nil {
		params["type"] = *serverType
	}

	if serverTag != nil {
		params["server-tag"] = *serverTag
	}

	response, err := c.request("/restart_game_server", http.MethodPost, params)
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

func (c *Client) ServerStatus(matchId string) (interface{}, error) {
	var params = map[string]string{}

	if matchId == "" {
		return nil, errors.New("match-id is required") // Todo move to the error list
	}

	params["match-id"] = matchId

	response, err := c.request("/get_server_status", http.MethodGet, params)
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

func (c *Client) Servers(includeLaunching ...bool) (interface{}, error) {
	var params = map[string]string{}

	if len(includeLaunching) > 0 {
		params["include-launching"] = strconv.FormatBool(includeLaunching[0])
	} else {
		params["include-launching"] = "false"
	}

	response, err := c.request("/list_servers", http.MethodGet, params)
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
