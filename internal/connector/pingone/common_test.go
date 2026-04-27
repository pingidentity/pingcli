// Copyright © 2026 Ping Identity Corporation

package pingone

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/stretchr/testify/require"
)

func TestGetManagementAPIObjectsFromIterator_StopsOnRepeatingPageURL(t *testing.T) {
	cursors := []management.PagedCursor{
		newManagementCursor(t, "https://example.test/page-a", "https://example.test/page-b"),
		newManagementCursor(t, "https://example.test/page-b", "https://example.test/page-a"),
		newManagementCursor(t, "https://example.test/page-a", "https://example.test/page-b"),
		newManagementCursor(t, "https://example.test/page-c", "https://example.test/page-d"),
	}
	yieldCount := 0
	iter := management.EntityArrayPagedIterator(func(yield func(management.PagedCursor, error) bool) {
		for _, cursor := range cursors {
			yieldCount++
			if !yield(cursor, nil) {
				return
			}
		}
	})

	apiObjects, err := GetManagementAPIObjectsFromIterator[management.TemplateContent](
		iter,
		"ReadAllTemplateContents",
		"GetContents",
		"pingone_notification_template_content",
	)

	require.NoError(t, err)
	require.Len(t, apiObjects, 2)
	require.Equal(t, 3, yieldCount)
}

func newManagementCursor(t *testing.T, pageURL, nextLink string) management.PagedCursor {
	t.Helper()
	return management.PagedCursor{
		EntityArray:  newManagementEntityArray(nextLink),
		HTTPResponse: newOKResponse(t, pageURL),
	}
}

func newManagementEntityArray(nextLink string) *management.EntityArray {
	embedded := management.EntityArrayEmbedded{
		Contents: []management.TemplateContent{{}},
	}
	entityArray := management.EntityArray{Embedded: &embedded}
	if nextLink != "" {
		links := map[string]management.LinksHATEOASValue{
			management.PAGINATION_HAL_LINK_INDEX_NEXT: {
				Href: nextLink,
			},
		}
		entityArray.Links = &links
	}

	return &entityArray
}

func newOKResponse(t *testing.T, pageURL string) *http.Response {
	t.Helper()
	parsedURL, err := url.Parse(pageURL)
	require.NoError(t, err)

	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     "200 OK",
		Body:       io.NopCloser(strings.NewReader("{}")),
		Request:    &http.Request{URL: parsedURL},
	}
}
