package gourl

import "errors"

// ErrNoURLProvided returns when the URL string or structs are empty
var ErrNoURLProvided = errors.New("url is empty")

// ErrInvalidScheme returned when given scheme parser detects an invalid char
// at a given place that is not expected by RFC 3986 definition of
//
//     scheme = ALPHA *( ALPHA / DIGIT / "+" / "-" / "." )
var ErrInvalidScheme = errors.New("invalid scheme structure")

// ErrNoScheme returns when scheme is empty, after the parse.
var ErrNoScheme = errors.New("no scheme provided")
