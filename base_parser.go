package gourl

// Parse provides generic url parser when nothing else is known, it will
// try to provide a filled version of URL.
func Parse(url string) (URL, error) {
	return URL{}, nil
}

// ParseScheme parses the base request scheme as first action.
// The parser uses the definition based on RFC 3986 specification.
func ParseScheme(url string) (string, error) {
	result := ""

	firstRune := url[0]

	switch {
	case 'a' <= firstRune && firstRune <= 'z',
		'A' <= firstRune && firstRune <= 'Z':
	default:
		return "", ErrInvalidScheme
	}

	for _, ch := range url {
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

	return result, nil
}
