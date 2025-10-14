// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"net/http"
	"regexp"
	"slices"
	"strings"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/spf13/pflag"
)

var (
	headerErrorPrefix = "custom type header error"
	headerRegex       = regexp.MustCompile(`(^[^\s:]+):(.*)`)
)

type Header struct {
	Key   string
	Value string
}

type HeaderSlice []Header

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*HeaderSlice)(nil)

func newHeader(header string) (Header, error) {
	matches := headerRegex.FindStringSubmatch(header)
	if len(matches) != 3 {
		return Header{}, fmt.Errorf("%w: %s", ErrInvalidHeaderFormat, header)
	}

	key := matches[1]
	if strings.EqualFold(key, "Authorization") {
		return Header{}, fmt.Errorf("%w: %s", ErrDisallowedAuthHeader, key)
	}

	return Header{
		Key:   key,
		Value: strings.TrimSpace(matches[2]), // Trim space as tabs and spaces are allowed after the colon in Header format
	}, nil
}

func (h *HeaderSlice) Set(val string) error {
	if h == nil {
		return &errs.PingCLIError{Prefix: headerErrorPrefix, Err: ErrCustomTypeNil}
	}

	if val == "" || val == "[]" {
		return nil
	}

	for header := range strings.SplitSeq(val, ",") {
		headerVal, err := newHeader(header)
		if err != nil {
			return &errs.PingCLIError{Prefix: headerErrorPrefix, Err: err}
		}
		*h = append(*h, headerVal)
	}

	return nil
}

func (h *HeaderSlice) SetHttpRequestHeaders(request *http.Request) {
	if h == nil {
		return
	}

	for _, header := range *h {
		request.Header.Add(header.Key, header.Value)
	}
}

func (h *HeaderSlice) Type() string {
	return "[]string"
}

func (h *HeaderSlice) String() string {
	if h == nil {
		return ""
	}

	return strings.Join(h.StringSlice(), ",")
}

func (h *HeaderSlice) StringSlice() []string {
	if h == nil {
		return []string{}
	}

	headers := []string{}
	for _, header := range *h {
		headers = append(headers, fmt.Sprintf("%s:%s", header.Key, header.Value))
	}

	slices.Sort(headers)

	return headers
}
