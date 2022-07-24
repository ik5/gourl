package gourl

import "testing"

func TestParseScheme(t *testing.T) {
	schemeList := []struct {
		url      string
		expected string
		err      error
	}{
		// Valid scheme
		{"http://foo.bar/", "http", nil},
		{"https://foo.bar", "https", nil},
		{"ftp://foo.bar/", "ftp", nil},
		{"mailto:me@example.com", "mailto", nil},
		{`sip:"John Doe" <500@127.0.0.1>`, "sip", nil},
		{"file:///path/to/file", "file", nil},
		{"scheme1+scheme2://foo.bar/", "scheme1+scheme2", nil},
		{"z39.50r://127.0.0.1/foo?12", "z39.50r", nil},

		// EmptyURL Provided
		{"", "", ErrNoURLProvided},

		// Invalid prefix
		{"1one:foo.bar", "", ErrInvalidScheme},
		{".dot://[:::1]/", "", ErrInvalidScheme},
		{"-minus://127.0.0.1/", "", ErrInvalidScheme},
		{"+plus:xyz", "", ErrInvalidScheme},
	}

	for _, scheme := range schemeList {
		result, err := ParseScheme(scheme.url)
		if err != scheme.err {
			t.Errorf(
				"Run over '%s', expected err '%#v' got '%#v'",
				scheme.url, scheme.err, err,
			)
		}

		if result != scheme.expected {
			t.Errorf(
				"Run over '%s', expected '%s', got '%s'",
				scheme.url, scheme.expected, result,
			)
		}
	}
}
