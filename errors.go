package gourl

import "errors"

// ErrInvalidScheme returned when given scheme parser detects an invalid char
// at a given place that is not expected by RFC 3986 definition of
//
//     scheme = ALPHA *( ALPHA / DIGIT / "+" / "-" / "." )
var ErrInvalidScheme = errors.New("invalid scheme structure")
