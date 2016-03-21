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
