# URI

There is a group of three terms:
  - URI - Uniform Resource Identifier
  - URL - Uniform Resource Locator
  - URN - Uniform Resource Name

The group of the three terms holds information about the network resources
how to get the information, and how to identify the type of information.

The following package holds implementation for RFC implementation for URI
using the Go (Golang) language.

The reason for it to live, is that the [url](https://golang.org/pkg/net/url/)
package is domain specific implementation rather then using an RFC implementation.

It means that there are large set of URI (and URL) that cannot be parsed by the
package and requires their own implementation.

This package aims to solve it.

# TODO:
 - [ ] [rfc3986](https://www.rfc-editor.org/rfc/rfc3986)
 - [ ] [rfc6874](https://www.rfc-editor.org/rfc/rfc6874)
 - [ ] [rfc7320](https://www.rfc-editor.org/rfc/rfc7320) / [rfc8820](https://www.rfc-editor.org/rfc/rfc8820)
 - [ ] [rfc8141](https://www.rfc-editor.org/rfc/rfc8141.html)

# License

The following package is under the Apache 2.0 license.
