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

// BasePost .
type BasePost struct {
	BlogName    string   `json:"blog_name"`
	ID          int64    `json:"id"`
	PostURL     string   `json:"post_url"`
	Type        PostType `json:"type"`
	Timestump   int      `json:"timestump"`
	Date        string   `json:"date"`
	Format      string   `json:"format"`
	ReblogKey   string   `json:"reblog_key"`
	Tags        []string `json:"tags"`
	Bookmarklet bool     `json:"bookmarklet"`
	Mobile      bool     `json:"mobile"`
	SourceURL   string   `json:"source_url"`
	SourceTitle string   `json:"source_title"`
	Liked       bool     `json:"liked"`
	State       string   `json:"state"`
	TotalPosts  int      `json:"total_posts"`
}

// TextPost text post.
type TextPost struct {
	BasePost
	Title string `json:"title"`
	Body  string `json:"body"`
}

// PhotoSize .
type PhotoSize struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

// Photo .
type Photo struct {
	Caption      string      `json:"caption"`
	AltSizes     []PhotoSize `json:"alt_size"`
	OriginalSize PhotoSize   `json:"original_size"`
}

// PhotoPost .
type PhotoPost struct {
	BasePost
	Photos []Photo `json:"photos"`
}

// QuotePost .
type QuotePost struct {
	BasePost
	Text   string `json:"text"`
	Source string `json:"source"`
}

// LinkPost .
type LinkPost struct {
	BasePost
	Title       string  `json:"title"`
	URL         string  `json:"url"`
	Author      string  `json:"author"`
	Excerpt     string  `json:"excerpt"`
	Publisher   string  `json:"publisher"`
	Photos      []Photo `json:"photos"`
	Description string  `json:"description"`
}

// Dialogue .
type Dialogue struct {
	Name   string `json:"name"`
	Label  string `json:"label"`
	Phrase string `json:"phrase"`
}

// ChatPost .
type ChatPost struct {
	BasePost
	Title     string     `json:"title"`
	Post      string     `json:"post"`
	Dialogues []Dialogue `json:"dialogue"`
}

// AudioPost .
type AudioPost struct {
	BasePost
	Caption     string `json:"caption"`
	Player      string `json:"player"`
	Plays       int64  `json:"plays"`
	AlbumArt    string `json:"almub_art"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	TrackName   string `json:"track_name"`
	TrackNumber int    `json:"track_number"`
	Year        int    `json:"year"`
}

// Player .
type Player struct {
	Width     int    `json:"width"`
	EmbedCode string `json:"embed_code"`
}

// VideoPost .
type VideoPost struct {
	BasePost
	Caption string   `json:"caption"`
	Players []Player `json:"player"`
}

// AnswerPost .
type AnswerPost struct {
	BasePost
	AskingName string `json:"asking_name"`
	AskingURL  string `json:"asking_url"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
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
