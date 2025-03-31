// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/spf13/pflag"
)

type Header struct {
	Key   string
	Value string
}

type HeaderSlice []Header

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*HeaderSlice)(nil)

func IsValidHeader(header string) (Header, bool) {
	headerNameRegex := regexp.MustCompile(`(^[^\s]+):[\t ]{0,1}(.*)$`)
	matches := headerNameRegex.FindStringSubmatch(header)
	if len(matches) != 3 {
		return Header{}, false
	}

	return Header{
		Key:   matches[1],
		Value: matches[2],
	}, true
}

func (h *HeaderSlice) Set(val string) error {
	if h == nil {
		return fmt.Errorf("failed to set Headers value: %s. Headers is nil", val)
	}

	if val == "" || val == "[]" {
		return nil
	} else {
		valH := strings.SplitSeq(val, ",")
		for header := range valH {
			headerVal, isValid := IsValidHeader(header)
			if !isValid {
				return fmt.Errorf("failed to set Headers: Invalid header: %s. Headers must be in the proper format", header)
			}
			*h = append(*h, headerVal)
		}
	}

	return nil
}

func (h HeaderSlice) SetHttpRequestHeaders(request *http.Request) {
	for _, header := range h {
		request.Header.Add(header.Key, header.Value)
	}
}

func (h HeaderSlice) Type() string {
	return "[]string"
}

func (h HeaderSlice) String() string {
	if h == nil {
		return ""
	}

	var headers []string
	for _, header := range h {
		headers = append(headers, fmt.Sprintf("%s:%s", header.Key, header.Value))
	}

	return strings.Join(headers, ",")
}

func (h HeaderSlice) StringSlice() []string {
	if h == nil {
		return []string{}
	}

	var headers []string
	for _, header := range h {
		headers = append(headers, fmt.Sprintf("%s:%s", header.Key, header.Value))
	}

	return headers
}
