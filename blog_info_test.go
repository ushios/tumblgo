package tumblgo

import "testing"

func TestBlogInfo(t *testing.T) {
	test := func(bi string, ak string) {
		c := NewClient(testAPIKey1)

		bir, err := c.BlogInfo(testBlogIdentifier)
		if err != nil {
			t.Errorf("BlogInfo got error: %s", err)
		}

		if bir.Meta.Status != 200 {
			t.Errorf("blog-info status expected(%d) but (%d) ", 200, bir.Meta.Status)
		}
	}

	test(testBlogIdentifier, testAPIKey1)
}
