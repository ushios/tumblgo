package tumblgo

// BlogAvatarResponse .
type BlogAvatarResponse struct {
	Meta     Meta       `json:"meta"`
	Response BlogAvatar `json:"response"`
}

// BlogAvatar .
type BlogAvatar struct {
	AvatarURL string `json:"avatar_url"`
}

// AvatarSize .
type AvatarSize string

const (
	AvatarSize16  AvatarSize = "16"
	AvatarSize24  AvatarSize = "24"
	AvatarSize30  AvatarSize = "30"
	AvatarSize40  AvatarSize = "40"
	AvatarSize48  AvatarSize = "48"
	AvatarSize64  AvatarSize = "64"
	AvatarSize96  AvatarSize = "96"
	AvatarSize128 AvatarSize = "128"
	AvatarSize512 AvatarSize = "512"
)

// BlogAvatar .
func (c *Client) BlogAvatar(bi string, size AvatarSize) (*BlogAvatarResponse, error) {
	// req, err := c.Request("GET", c.BlogEndpoint(bi, "avatar/"+string(size)), nil)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// resp := BlogAvatarResponse{}
	// var v interface{}
	// if err := c.ReadResponse(req, &v); err != nil {
	// 	return nil, err
	// }
	//
	// return &resp, nil
	return nil, nil
}
