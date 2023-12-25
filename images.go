package go_playflow

func (c *Client) ServerTags() ([]string, error) {
	return nil, nil
}

func (c *Client) DeleteServerTags() error {
	return nil
}

func (c *Client) UploadServerFiles(fileToUpload *[]byte, serverTag ...string) error {
	return nil
}

func (c *Client) UploadVersion() (string, error) {
	return "", nil
}
