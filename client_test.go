package tumblgo

import (
	"fmt"
	"testing"
)

const (
	TestBlogIdentifier string = "scipsy"
	TestAPIKey1        string = "fuiKNFp9vQFvjLNvx4sUwti4Yb5yGutBN4Xh10LXZhhRKjWlV4"
)

func TestClient(t *testing.T) {
	test := func(bi string, ak string) {
		c := NewClient(TestBlogIdentifier, ak)

		if c == nil {
			t.Errorf("Client is null.")
		}
	}

	test(TestBlogIdentifier, TestAPIKey1)
}

func TestBlogInfo(t *testing.T) {
	test := func(bi string, ak string) {
		c := NewClient(TestBlogIdentifier, TestAPIKey1)

		bir := BlogInfoResponse{}
		if err := c.BlogInfo(&bir); err != nil {
			t.Errorf("Blog Info got error: %s", err)
		}

		fmt.Println(bir.Response.Blog.Title)
	}

	test(TestBlogIdentifier, TestAPIKey1)
}
