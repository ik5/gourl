package gourl

import "github.com/ik5/gostrutils"

// Parse provides generic url parser when nothing else is known, it will
// try to provide a filled version of URL.
func Parse(url string) (URL, error) {
	return URL{}, nil
}

// ParseScheme parses the base request scheme as first action.
// The parser uses the definition based on RFC 3986 specification.
func ParseScheme(url string) (string, error) {
	if gostrutils.IsEmpty(url) {
		return "", ErrNoURLProvided
	}

	result := ""

	// extract string to runes, instead of going over bytes, so non ASCII
	// or extended ASCII chars can still be looked upon single char.
	// Note that UTF-8 (the go's string) is multi-byte, that is a char can have
	// one, two, three or four bytes.
	//
	// Once the runes exists, get the first rune to make sure it is valid...
	runes := []rune(url)
	firstRune := runes[0]

	switch {
	case 'a' <= firstRune && firstRune <= 'z',
		'A' <= firstRune && firstRune <= 'Z':
	default:
		return "", ErrInvalidScheme
	}

	for _, ch := range runes {
		if ch == ':' {
			break
		}

		switch { //
		case '0' <= ch && ch <= '9':
			fallthrough
		case 'a' <= ch && ch <= 'z':
			fallthrough
		case 'A' <= ch && ch <= 'Z':
			fallthrough
		case ch == '+', ch == '-', ch == '.':
			result += string(ch)
		default:
			return "", ErrInvalidScheme
		}
	}

	if gostrutils.IsEmpty(result) {
		return "", ErrNoScheme
	}
	return result, nil
}
