package tumblgo

import (
	"fmt"
	"testing"
)

func TestOAuth(t *testing.T) {
	test := func(key, secret string) {
		o, err := NewOAuth(key, secret)
		if err != nil {
			t.Fatalf("NewOAuth got error: %s", err)
		}

		fmt.Println(o)
	}

	test(
		"",
		"",
	)
}
