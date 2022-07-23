package gourl

import "testing"

func TestParseScheme(t *testing.T) {
	validSchemeList := []string{
		"http://foo.bar/",
		"https://foo.bar",
		"ftp://foo.bar/",
		"mailto:me@example.com",
		`sip:"John Doe" <500@127.0.0.1>`,
		"file:///path/to/file",
		"scheme1+scheme2://foo.bar/",
	}

}
