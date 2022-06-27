package fake

import (
	"net/url"
	"testing"
)

func TestURL(t *testing.T) {
	t.Parallel()

	t.Run("should work with configuration", func(t *testing.T) {
		resultURL := URL([]string{"www.example.com"}, []string{"/"})
		want := "https://www.example.com/"
		got := resultURL.String()
		if got != want {
			t.Fatalf("unexpected URL string, want: %#v, got: %#v", want, got)
		}
	})

	t.Run("should work without configuration", func(t *testing.T) {
		resultURL := URL([]string{}, nil) // check nil and empty slice
		got := resultURL.String()
		_, err := url.Parse(got)
		if err != nil {
			t.Fatalf("invalid URL string: %#v", got)
		}
	})
}
