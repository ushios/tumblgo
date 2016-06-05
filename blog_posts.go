package tumblgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/guregu/null"
)

// PostType .
type PostType string

const (
	// PostTypeUnknown .
	PostTypeUnknown PostType = "unknown"
	// PostTypeText .
	PostTypeText = "text"
	// PostTypeQuote .
	PostTypeQuote = "quote"
	// PostTypeLink .
	PostTypeLink = "link"
	// PostTypeAnswer .
	PostTypeAnswer = "answer"
	// PostTypeVideo .
	PostTypeVideo = "video"
	// PostTypeAudio .
	PostTypeAudio = "audio"
	// PostTypePhoto .
	PostTypePhoto = "photo"
	// PostTypeChat .
	PostTypeChat = "chat"
)

const (
	// TypeFieldName TumblrAPI's spec.
	TypeFieldName = "type"
)

var (
	// ErrPostTypeNotMatched the post type is not mathced.
	ErrPostTypeNotMatched = errors.New("post type not matched")
)

// BlogPostsRequest .
type BlogPostsRequest struct {
	Type       *PostType   `json:"type"`
	ID         null.Int    `json:"id"`
	Tag        null.String `json:""`
	Limit      int         `json:"limit"`
	Offset     int         `json:"offset"`
	ReblogInfo bool        `json:"reblog_info"`
	NotesInfo  bool        `json:"notes_info"`
	Filter     null.String `json:"filter"`
}

// BlogPostsResponse .
type BlogPostsResponse struct {
	Meta     Meta      `json:"meta"`
	Response BlogPosts `json:"response"`
}

// BlogPosts .
type BlogPosts struct {
	Blog  Blog   `json:"blog"`
	Posts []Post `json:"posts"`
}

// Post .
type Post struct {
	jsonData []byte
	Raw      interface{}
}

// MarshalJSON encode to json.
func (p *Post) MarshalJSON() ([]byte, error) {
	return json.Marshal(p)
}

// UnmarshalJSON get data from json!
func (p *Post) UnmarshalJSON(data []byte) error {
	p.jsonData = data

	if err := json.Unmarshal(data, &(p.Raw)); err != nil {
		return err
	}
	return nil
}

// Type get post type.
func (p Post) Type() PostType {
	switch p.Raw.(type) {
	case map[string]interface{}:
		vo := reflect.ValueOf(p.Raw)
		if vo.IsValid() {
			keys := vo.MapKeys()
			for _, key := range keys {
				if key.String() == TypeFieldName {
					typeName := fmt.Sprintf("%s", vo.MapIndex(key))
					return PostType(typeName)
				}
			}
		}
	}

	return PostTypeUnknown
}

// postModel restruct Post to XxxxPost.
func (p Post) postModel(postType PostType, v interface{}) error {
	if p.Type() != postType {
		return ErrPostTypeNotMatched
	}

	if err := json.Unmarshal(p.jsonData, v); err != nil {
		return err
	}

	return nil
}

// TextPost restruct to TextPost.
func (p Post) TextPost() (*TextPost, error) {
	tp := TextPost{}
	if err := p.postModel(PostTypeText, &tp); err != nil {
		return &tp, err
	}
	return &tp, nil
}

// PhotoPost restruct to PhotoPost.
func (p Post) PhotoPost() (*PhotoPost, error) {
	pp := PhotoPost{}
	if err := p.postModel(PostTypePhoto, &pp); err != nil {
		return &pp, err
	}
	return &pp, nil
}

// QuotePost restruct to QuotePost.
func (p Post) QuotePost() (*QuotePost, error) {
	qp := QuotePost{}
	if err := p.postModel(PostTypeQuote, &qp); err != nil {
		return &qp, err
	}
	return &qp, nil
}

// LinkPost restruct to LinkPost.
func (p Post) LinkPost() (*LinkPost, error) {
	lp := LinkPost{}
	if err := p.postModel(PostTypeLink, &lp); err != nil {
		return &lp, err
	}
	return &lp, nil
}

// ChatPost restruct to ChatPost.
func (p Post) ChatPost() (*ChatPost, error) {
	cp := ChatPost{}
	if err := p.postModel(PostTypeChat, &cp); err != nil {
		return &cp, err
	}
	return &cp, nil
}

// AudioPost restruct to AudioPost.
func (p Post) AudioPost() (*AudioPost, error) {
	ap := AudioPost{}
	if err := p.postModel(PostTypeAudio, &ap); err != nil {
		return &ap, err
	}
	return &ap, nil
}

// VideoPost restruct to VideoPost.
func (p Post) VideoPost() (*VideoPost, error) {
	vp := VideoPost{}
	if err := p.postModel(PostTypeVideo, &vp); err != nil {
		return &vp, err
	}
	return &vp, nil
}

// AnswerPost restruct to AnswerPost.
func (p Post) AnswerPost() (*AnswerPost, error) {
	ap := AnswerPost{}
	if err := p.postModel(PostTypeAnswer, &ap); err != nil {
		return &ap, err
	}
	return &ap, nil
}

// BlogPosts Retrieve Published Posts
func (c *Client) BlogPosts(bi string, param BlogPostsRequest) (*BlogPostsResponse, error) {
	req, err := c.Request("GET", c.BlogEndpoint(bi, "posts"), nil)
	if err != nil {
		return nil, err
	}

	resp := BlogPostsResponse{}
	if err := c.ReadResponse(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
