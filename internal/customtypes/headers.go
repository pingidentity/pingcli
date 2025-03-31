// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"net/http"
	"regexp"
	"slices"
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

func NewHeader(header string) (Header, error) {
	regexPattern := `(^[^\s]+):[\t ]{0,1}(.*)$`
	headerNameRegex := regexp.MustCompile(regexPattern)
	matches := headerNameRegex.FindStringSubmatch(header)
	if len(matches) != 3 {
		return Header{}, fmt.Errorf("failed to set Headers: Invalid header: %s. Headers must be in the proper format. Expected regex pattern: %s", header, regexPattern)
	}

	if matches[1] == "Authorization" {
		return Header{}, fmt.Errorf("failed to set Headers: Invalid header: %s. Authorization header is not allowed", matches[1])
	}

	return Header{
		Key:   matches[1],
		Value: matches[2],
	}, nil
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
			headerVal, err := NewHeader(header)
			if err != nil {
				return err
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
	return strings.Join(h.StringSlice(), ",")
}

func (h HeaderSlice) StringSlice() []string {
	if h == nil {
		return []string{}
	}

	headers := []string{}
	for _, header := range h {
		headers = append(headers, fmt.Sprintf("%s:%s", header.Key, header.Value))
	}

	slices.Sort(headers)

	return headers
}
