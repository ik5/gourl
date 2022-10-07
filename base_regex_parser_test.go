package gourl

import (
	"testing"

	"github.com/go-test/deep"
)

func TestParseRegex(t *testing.T) {
	urlToCheck := []struct {
		rawURL string
		err    error
		url    URL
	}{
		{
			rawURL: "https://www.example.com/",
			err:    nil,
			url: URL{
				RawRequest: "https://www.example.com/",
				Scheme:     "https",
				RawPath:    "/",
				Host:       "www.example.com",
			},
		},
		{
			rawURL: "http://user@ww1.domain.com:8080/",
			err:    nil,
			url: URL{
				Scheme:     "http",
				RawPath:    "/",
				Host:       "ww1.domain.com",
				Port:       8080,
				RawRequest: "http://user@ww1.domain.com:8080/",
				Credentials: CredentialsElement{
					Username:       "user",
					RawCredentails: "user",
				},
			},
		},
		{
			rawURL: "http://user:pass@ww1.domain.com:8080/",
			err:    nil,
			url: URL{
				Scheme:     "http",
				RawPath:    "/",
				Host:       "ww1.domain.com",
				Port:       8080,
				RawRequest: "http://user:pass@ww1.domain.com:8080/",
				Credentials: CredentialsElement{
					Username:       "user",
					Password:       "pass",
					RawCredentails: "user:pass",
				},
			},
		},
		{
			rawURL: "sftp://user:pass@/",
			err:    nil,
			url: URL{
				RawPath:    "/",
				Scheme:     "sftp",
				Host:       "",
				Port:       0,
				RawRequest: "sftp://user:pass@/",
				Credentials: CredentialsElement{
					Username:       "user",
					Password:       "pass",
					RawCredentails: "user:pass",
				},
			},
		},
		{
			rawURL: "ftp://[::1]/",
			err:    nil,
			url: URL{
				RawRequest: "ftp://[::1]/",
				RawPath:    "/",
				Host:       "[::1]",
				Scheme:     "ftp",
				Port:       0,
			},
		},
		{
			rawURL: "http://www.עברית.com/",
			err:    nil,
			url: URL{
				RawRequest: "http://www.עברית.com/",
				Scheme:     "http",
				RawPath:    "/",
				Host:       "www.עברית.com",
			},
		},
		{ // https://daniel.haxx.se/blog/2022/09/08/http-http-http-http-http-http-http/
			rawURL: "http://http://http://@http://http://?http://#http://",
			err:    nil,
			url: URL{
				RawRequest:  "http://http://http://@http://http://?http://#http://",
				RawPath:     "//http://@http://http://",
				Fragment:    "http://",
				Scheme:      "http",
				Host:        "http",
				RawQuery:    "?http://",
				Credentials: CredentialsElement{},
			},
		},
	}

	for idx, toCheck := range urlToCheck {
		url, err := ParseRegex(toCheck.rawURL)
		if err != toCheck.err {
			t.Errorf("%d. expected err: %s but got %s", idx, toCheck.err, err)
		}

		diff := deep.Equal(url, toCheck.url)
		if diff != nil {
			t.Errorf(
				"%d. expected url to be the same but %+v diff.\n%+v\n!=\n%+v\n",
				idx, diff, url, toCheck.url,
			)
		}
	}
}
