package tumblgo

import "testing"

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
	test := func(bi string, ak string, ps []PostType) {
		c := NewClient(ak)

		param := BlogPostsRequest{}
		res, err := c.BlogPosts(bi, param)
		if err != nil {
			t.Fatalf("BlogPosts got error: %s", err)
		}

		if res.Meta.Status != 200 {
			t.Fatalf("blog-posts status expected(%d) but (%d)", 200, res.Meta.Status)
		}

		// Check Type.
		for i, postType := range ps {
			if res.Response.Posts[i].Type() != postType {
				t.Errorf("Type Expected (%s) but (%s)", postType, res.Response.Posts[i].Type())
			}
		}

		// Check post objects.
		for _, post := range res.Response.Posts {
			switch post.Type() {
			case PostTypeText:
				tp, err := post.TextPost()
				if err != nil {
					t.Errorf("TextPost failed.")
				}

				if tp.Body == "" {
					t.Errorf("The text post have no body.")
				}
			}
		}
	}

	test("scipsy", testAPIKey1, []PostType{
		PostTypeAnswer,
	})
}
