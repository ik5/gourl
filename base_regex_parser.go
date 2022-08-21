package gourl

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	// REGEX based on https://www.rfc-editor.org/rfc/rfc3986#appendix-B
	baseRegex = regexp.MustCompile(
		`^(?P<raw_scheme>(?P<scheme>[^:/?#]+):)?(?P<raw_authority>//(?P<authority>[^/?#]*))?(?P<path>[^?#]*)(?P<raw_query>\?(?P<query>[^#]*))?(?P<raw_fragment>#(?P<fragment>.*))?`,
	)

	baseGroupNames = baseRegex.SubexpNames()
)

func ParseRegex(rawURL string) (URL, error) {
	var url URL
	var err error
	rawData := map[string]string{
		"": "",
	}

	subMatched := baseRegex.FindAllStringSubmatch(rawURL, -1)
	for _, match := range subMatched {
		for groupIdx, group := range match {
			name := baseGroupNames[groupIdx]
			rawData[name] = group
		}
	}

	url.RawRequest = rawData[""]
	url.Scheme = rawData["scheme"]
	url.RawQuery = rawData["raw_query"]
	url.RawPath = rawData["path"]
	url.Fragment = rawData["fragment"]

	if rawData["authority"] != "" {
		url.Credentials, url.Host, url.Port, err = ParseAuthorityRegex(
			rawData["authority"],
		)

		if err != nil {
			return URL{}, err
		}
	}

	//fmt.Printf("url: %+v | err: %+v\n", url, err)

	return url, err
}

// ParseAuthorityRegex takes "raw" authority of user:password@path:port and breaks it down
func ParseAuthorityRegex(authority string) (CredentialsElement, string, int, error) {
	var (
		credentials, username, password, host string
		port                                  int
		err                                   error
	)

	elements := strings.Split(authority, "@")
	var tmpCredentials []string
	tmpHost := elements[len(elements)-1]

	if len(elements) == 2 {
		credentials = elements[0]
		tmpCredentials = strings.Split(credentials, ":")
	}

	switch len(tmpCredentials) {
	case 0:
		tmpCredentials = []string{"", ""}
	case 1:
		tmpCredentials = append(tmpCredentials, "")
	}

	username = tmpCredentials[0]
	password = tmpCredentials[1]

	var hostPort []string
	// is it IPv6?
	if strings.HasPrefix(tmpHost, "[") {
		hostPort = strings.SplitAfter(tmpHost, "]")
	} else { // anything else ...
		hostPort = strings.Split(tmpHost, ":")
	}

	switch len(hostPort) {
	case 1:
		hostPort = append(hostPort, "0")
	}

	host = hostPort[0]
	port, _ = strconv.Atoi(hostPort[1])

	return CredentialsElement{username, password, credentials}, host, port, err
}
