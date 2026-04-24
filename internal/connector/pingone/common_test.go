// Copyright © 2026 Ping Identity Corporation

package pingone

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/stretchr/testify/require"
)

func TestGetManagementAPIObjectsFromIterator_StopsOnRepeatingNextLink(t *testing.T) {
	cursors := []management.PagedCursor{
		newManagementCursor("link-a"),
		newManagementCursor("link-b"),
		newManagementCursor("link-a"),
		newManagementCursor("link-c"),
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
	require.Len(t, apiObjects, 3)
	require.Equal(t, 3, yieldCount)
}

func newManagementCursor(nextLink string) management.PagedCursor {
	return management.PagedCursor{
		EntityArray:  newManagementEntityArray(nextLink),
		HTTPResponse: newOKResponse(),
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

func newOKResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     "200 OK",
		Body:       io.NopCloser(strings.NewReader("{}")),
	}
}
