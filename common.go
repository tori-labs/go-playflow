package go_playflow

type HTTPValidationError struct {
	ValidationError []struct {
		Loc  []any  `json:"loc"`
		Msg  string `json:"msg"`
		Type string `json:"type"`
	} `json:"detail"`
}

type ServerList struct {
	TotalServers int `json:"total_servers"`
	Servers      []struct {
	}
}

type ServerTags struct {
}
