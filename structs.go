package gourl

// PathElement breaks down path elements based on protocol rules
type PathElement map[string][]string

// QueryValue holds information regarding query arguments of key values, including repeated versions of them (based on RFCs).
type QueryValue map[string][]string

// CredentialsElement holds information regarding the username and password
// provided by the URL.
type CredentialsElement struct {
	Username string

	// Password is a clear text version of what was provided, please use with
	// care.
	Password string
}

// ParserRegister holds callback regarding each type of parser
type ParserRegister struct {
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

	// WithAuthority is set to true if the protocol supports authority
	// mechanism (username and password) such as HTTP, FTP and more.
	WithAuthority bool

	// Port holds information regarding the port component.
	Port int
}
