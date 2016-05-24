package tumblgo

// BlogLikesRequest .
type BlogLikesRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Before int `json:"before"`
	After  int `json:"after"`
}

// BlogLikesResponse .
type BlogLikesResponse struct {
	Meta     Meta      `json:"meta"`
	Response BlogLikes `json:"response"`
}

// BlogLikes .
type BlogLikes struct {
	LikedPosts int `json:"liked_posts"`
	LikedCount int `json:"liked_count"`
}

// BlogLikes .
func (c *Client) BlogLikes(bi string, param BlogLikesRequest) (*BlogLikesResponse, error) {
	req, err := c.Request("GET", c.BlogEndpoint(bi, "likes"), nil)
	if err != nil {
		return nil, err
	}

	resp := BlogLikesResponse{}
	if err := c.ReadResponse(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
