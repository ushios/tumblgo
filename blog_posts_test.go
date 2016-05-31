package tumblgo

import (
	"fmt"
	"testing"
)

func TestBlogPosts(t *testing.T) {
	test := func(bi string, ak string) {
		c := NewClient(ak)

		param := BlogPostsRequest{}
		res, err := c.BlogPosts(bi, param)
		if err != nil {
			t.Fatalf("BlogPosts got error: %s", err)
		}

		if res.Meta.Status != 200 {
			t.Fatalf("blog-posts status expected(%d) but (%d)", 200, res.Meta.Status)
		}

		// fmt.Println(res.Response.Posts[0])
	}

	test("scipsy", testAPIKey1)
}

func TestPostType(t *testing.T) {
	test := func(bi string, ak string) {
		c := NewClient(ak)

		param := BlogPostsRequest{}
		res, err := c.BlogPosts(bi, param)
		if err != nil {
			t.Fatalf("BlogPosts got error: %s", err)
		}

		if res.Meta.Status != 200 {
			t.Fatalf("blog-posts status expected(%d) but (%d)", 200, res.Meta.Status)
		}

		fmt.Println(res.Response.Posts[0].PostType())
	}

	test("scipsy", testAPIKey1)
}
