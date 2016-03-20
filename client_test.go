package tumblgo

import "testing"

const (
	TestBlogIdentifier string = "scipsy"
	TestAPIKey1        string = "fuiKNFp9vQFvjLNvx4sUwti4Yb5yGutBN4Xh10LXZhhRKjWlV4"
)

func TestClient(t *testing.T) {
	test := func(bi string, ak string) {
		c := NewClient(ak)

		if c == nil {
			t.Errorf("Client is null.")
		}
	}

	test(TestBlogIdentifier, TestAPIKey1)
}

func TestBlogInfo(t *testing.T) {
	test := func(bi string, ak string) {
		c := NewClient(TestAPIKey1)

		bir, err := c.BlogInfo(TestBlogIdentifier)
		if err != nil {
			t.Errorf("Blog Info got error: %s", err)
		}

		if bir.Meta.Status != 200 {
			t.Errorf("blog-info status expected(%d) but (%d) ", 200, bir.Meta.Status)
		}
	}

	test(TestBlogIdentifier, TestAPIKey1)
}
