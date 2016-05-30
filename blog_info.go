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
	TotalPosts           int    `json:"total_posts"`
	Posts                int    `json:"posts"`
	Name                 string `json:"name"`
	Updated              int    `json:"updated"`
	Description          string `json:"description"`
	IsNsfw               bool   `json:"is_nsfw"`
	Ask                  bool   `json:"ask"`
	AskPageTitle         string `json:"ask_page_title"`
	AskAnon              bool   `json:"ask_anon"`
	ShareLikes           bool   `json:"share_likes"`
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
