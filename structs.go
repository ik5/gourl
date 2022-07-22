package gourl

// PathElement breaks down path elements based on protocol rules
type PathElement map[string][]string

// QueryValue holds information regarding query arguments of key values, including repeated versions of them (based on RFCs).
type QueryValue map[string][]string

// URL holds information regarding the full components of address.
type URL struct {
	// Order of struct is based on memory alignment for better optimization.

	// Path elements of the URL
	Path PathElement

	// Query elements based on
	Query QueryValue

	// Scheme is the type of protocol component.
	Scheme string

	Host string

	RawPath string

	// WithAuthority is set to true if the protocol supports authority mechanism (username and password) such as HTTP, FTP and more.
	WithAuthority bool

	// Port holds information regarding the port component.
	Port int
}
