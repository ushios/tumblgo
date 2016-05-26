package tumblgo

import "testing"

func TestBlogInfo(t *testing.T) {
	test := func(bi string, ak string, name string) {
		c := NewClient(testAPIKey1)

		res, err := c.BlogInfo(testBlogIdentifier)
		if err != nil {
			t.Errorf("BlogInfo got error: %s", err)
		}

		if res.Meta.Status != 200 {
			t.Errorf("blog-info status expected(%d) but (%d) ", 200, res.Meta.Status)
		}

		if res.Response.Blog.Name != name {
			t.Errorf("blog-info name expected (%s) but (%s)", name, res.Response.Blog.Name)
		}
	}

	test(testBlogIdentifier, testAPIKey1, "scipsy")
}
