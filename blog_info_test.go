package tumblgo

import "testing"

func TestBlogInfo(t *testing.T) {
	test := func(bi string, ak string) {
		c := NewClient(TestAPIKey1)

		bir, err := c.BlogInfo(TestBlogIdentifier)
		if err != nil {
			t.Errorf("BlogInfo got error: %s", err)
		}

		if bir.Meta.Status != 200 {
			t.Errorf("blog-info status expected(%d) but (%d) ", 200, bir.Meta.Status)
		}
	}

	test(TestBlogIdentifier, TestAPIKey1)
}
