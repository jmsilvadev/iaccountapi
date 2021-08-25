package requests

import (
	"bytes"
	"net/http"
	"time"
)

// Client is the setup of a conection with the Form3API
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// SendRequest send a HTTP request using the configuration in the Client Object
func (cli *Client) SendRequest(method string, endpoint string, body []byte) (*http.Response, error) {
	request, err := http.NewRequest(method, cli.BaseURL+endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/vnd.api+json")
	cli.HTTPClient.Timeout = time.Duration(5 * time.Second)
	resp, err := cli.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
