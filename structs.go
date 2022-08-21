package gourl

// PathElement breaks down path elements based on protocol rules
type PathElement map[string][]string

// QueryValue holds information regarding query arguments of key values,
// including repeated versions of them (based on RFCs).
type QueryValue map[string][]string

// CredentialsElement holds information regarding the username and password
// provided by the URL.
type CredentialsElement struct {
	// The username part of the credentials
	Username string

	// Password is a clear text version of what was provided, please use with
	// care.
	Password string

	// user:pass or user only as appears at the string in raw form
	RawCredentails string
}

// ParserRegister holds callback regarding each type of parser.
// Only non empty callback will override existed parser.
type ParserRegister struct {

	// RawParser is the main parser after finding out what is the Scheme.
	RawParser func(rawURL string) (URL, error)

	// AuthorityParser parses rawURL and place authority components to
	// Credentials, host and port.
	AuthorityParser func(rawURL string) (CredentialsElement, string, int, error)

	// PathParser get a raw path and breaks it out for it's component
	PathParser func(rawPath string) (PathElement, error)

	// QueryParser get a raw query and place it for it's components.
	QueryParser func(rawQuery string) (QueryValue, error)

	// ExtraInfoParser parses from rawURL the protocol extra components.
	ExtraInfoParser func(rawURL string) (interface{}, error)
}

// GeneratorRegister holds callbacks
type GeneratorRegister struct {
}

// URL holds information regarding the full components of address.
//
// Important: Order of struct is based on memory alignment for better
// optimization.
type URL struct {

	// Path elements of the URL.
	Path PathElement

	// Query elements based on.
	Query QueryValue

	// Credentials Holds information about user name and password.
	Credentials CredentialsElement

	// Scheme is the type of protocol component.
	Scheme string

	// Host name or IPv4/IPv6 element.
	Host string

	// Fragment holds the name of the fragment part
	Fragment string

	// ExtraInfo provides extra information based on specific protocol and it's
	// parsing mechanism
	ExtraInfo interface{}

	// RawPath hold the original path provided to be parsed.
	RawPath string

	// RawRequest holds the raw URL prior of parsing.
	RawRequest string

	// RawQuery holds the raw query prior of parsing.
	RawQuery string

	// Port holds information regarding the port component.
	Port int
}
