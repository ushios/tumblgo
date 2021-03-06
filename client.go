package tumblgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// Scheme http or https.
	Scheme string = "https"
	// Host tumblr api host.
	Host string = "api.tumblr.com"
	// Version api version.
	Version string = "v2"
	// TumblgoVersion this version.
	TumblgoVersion string = " 0.1"
)

// A Client for Tumblr
type Client struct {
	Client *http.Client
	APIKey string
}

// NewClient create tumblgo client
// using tumblgo.Client{..} if you want to options
func NewClient(apiKey string) *Client {
	return &Client{
		Client: &http.Client{Timeout: time.Duration(10) * time.Second},
		APIKey: apiKey,
	}
}

// BlogEndpoint attache scheme,host and etc.
func (c *Client) BlogEndpoint(bi string, path string) string {
	return fmt.Sprintf("%s://%s/%s/blog/%s.tumblr.com/%s",
		Scheme,
		Host,
		Version,
		bi,
		path,
	)
}

// Request create request.
func (c *Client) Request(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}

	// Set User-Agent.
	req.Header.Set("User-Agent", "golang tumblr client tumblgo/"+TumblgoVersion)

	// Set api token.
	val := req.URL.Query()
	val.Add("api_key", c.APIKey)
	req.URL.RawQuery = val.Encode()

	return req, nil
}

// ReadResponse .
func (c *Client) ReadResponse(req *http.Request, v interface{}) error {
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(v); err != nil {
		return err
	}

	return err
}
