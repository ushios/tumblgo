package tumblgo

// BlogInfoResponse .
type BlogInfoResponse struct {
	Meta     Meta     `json:"meta"`
	Response BlogInfo `json:"response"`
}

// BlogInfo .
type BlogInfo struct {
	Blog Blog `json:"blog"`
}

// Blog .
type Blog struct {
	Title                string `json:"title"`
	Posts                int    `json:"posts"`
	Name                 string `json:"name"`
	Updated              int    `json:"updated"`
	Description          string `json:"description"`
	Ask                  bool   `json:"ask"`
	AskAnon              bool   `json:"ask_anon"`
	Likes                int    `json:"likes"`
	IsBlockedFromPrimary bool   `json:"is_blocked_from_primary"`
}

// BlogInfo get blog info.
func (c *Client) BlogInfo(bi string) (*BlogInfoResponse, error) {
	req, err := c.Request("GET", c.BlogEndpoint(bi, "info"), nil)
	if err != nil {
		return nil, err
	}

	resp := BlogInfoResponse{}
	if err := c.ReadResponse(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
