package tumblgo

import "testing"

func TestBlogInfo(t *testing.T) {
	test := func(bi string, ak string, name string) {
		c := NewClient(ak)

		res, err := c.BlogInfo(bi)
		if err != nil {
			t.Fatalf("BlogInfo got error: %s", err)
		}

		if res.Meta.Status != 200 {
			t.Fatalf("blog-info status expected(%d) but (%d) ", 200, res.Meta.Status)
		}

		if res.Response.Blog.Name != name {
			t.Errorf("blog-info name expected (%s) but (%s)", name, res.Response.Blog.Name)
		}
	}

	test("scipsy", testAPIKey1, "scipsy")
}
