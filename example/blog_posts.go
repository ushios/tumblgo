package example

import (
	"fmt"

	"github.com/ushios/tumblgo"
)

// SeeBlogPosts example
func SeeBlogPosts() {
	client := tumblgo.NewClient("fuiKNFp9vQFvjLNvx4sUwti4Yb5yGutBN4Xh10LXZhhRKjWlV4")

	res, err := client.BlogPosts("ushio", tumblgo.BlogPostsRequest{})
	if err != nil {
		// error
	}

	if res.Meta.Status != 200 {
		// api error
	}

	for _, post := range res.Response.Posts {
		switch post.Type() {
		case tumblgo.PostTypeText:
			text, err := post.TextPost()
			if err != nil {
				// error
			}
			fmt.Println(text)
		case tumblgo.PostTypePhoto:
			photo, _ := post.PhotoPost()
			fmt.Println(photo)
		}
	}
}
