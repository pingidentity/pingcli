// Copyright © 2025 Ping Identity Corporation

package customtypes_test

import (
	"net/http"
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/require"
)

func Test_HeaderSlice_Set(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		cType         *customtypes.HeaderSlice
		headerStr     string
		expectedError error
	}{
		{
			name:      "Happy path - single header",
			cType:     new(customtypes.HeaderSlice),
			headerStr: "key:value",
		},
		{
			name:      "Happy path - multiple headers",
			cType:     new(customtypes.HeaderSlice),
			headerStr: "key1:value1,key2:value2",
		},
		{
			name:      "Happy path - empty",
			cType:     new(customtypes.HeaderSlice),
			headerStr: "",
		},
		{
			name:          "Invalid value",
			cType:         new(customtypes.HeaderSlice),
			headerStr:     "invalid-value",
			expectedError: customtypes.ErrInvalidHeaderFormat,
		},
		{
			name:          "Disallowed auth header",
			cType:         new(customtypes.HeaderSlice),
			headerStr:     "Authorization:some-token",
			expectedError: customtypes.ErrDisallowedAuthHeader,
		},
		{
			name:          "Nil custom header type",
			cType:         nil,
			headerStr:     "key:value",
			expectedError: customtypes.ErrCustomTypeNil,
		},
		{
			name:      "valid header with space after colon",
			cType:     new(customtypes.HeaderSlice),
			headerStr: "key: value",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := tc.cType.Set(tc.headerStr)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_HeaderSlice_Type(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.HeaderSlice
		expectedType string
	}{
		{
			name:         "Happy path",
			cType:        new(customtypes.HeaderSlice),
			expectedType: "[]string",
		},
		{
			name:         "Nil custom header type",
			cType:        nil,
			expectedType: "[]string",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			actualType := tc.cType.Type()

			require.Equal(t, tc.expectedType, actualType)
		})
	}
}

func Test_HeaderSlice_String(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.HeaderSlice
		expectedStr string
	}{
		{
			name:        "Happy path - single header",
			cType:       utils.Pointer(customtypes.HeaderSlice{{Key: "key", Value: "value"}}),
			expectedStr: "key:value",
		},
		{
			name:        "Happy path - multiple headers",
			cType:       utils.Pointer(customtypes.HeaderSlice{{Key: "key1", Value: "value1"}, {Key: "key2", Value: "value2"}}),
			expectedStr: "key1:value1,key2:value2",
		},
		{
			name:        "Happy path - empty",
			cType:       new(customtypes.HeaderSlice),
			expectedStr: "",
		},
		{
			name:        "Nil custom header type",
			cType:       nil,
			expectedStr: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			actualStr := tc.cType.String()

			require.Equal(t, tc.expectedStr, actualStr)
		})
	}
}

func Test_HeaderSlice_StringSlice(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name             string
		cType            *customtypes.HeaderSlice
		expectedStrSlice []string
	}{
		{
			name:             "Happy path - single header",
			cType:            utils.Pointer(customtypes.HeaderSlice{{Key: "key", Value: "value"}}),
			expectedStrSlice: []string{"key:value"},
		},
		{
			name:             "Happy path - multiple headers",
			cType:            utils.Pointer(customtypes.HeaderSlice{{Key: "key2", Value: "value2"}, {Key: "key1", Value: "value1"}}),
			expectedStrSlice: []string{"key1:value1", "key2:value2"},
		},
		{
			name:             "Happy path - empty",
			cType:            new(customtypes.HeaderSlice),
			expectedStrSlice: []string{},
		},
		{
			name:             "Nil custom header type",
			cType:            nil,
			expectedStrSlice: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			actualStrSlice := tc.cType.StringSlice()

			require.Equal(t, tc.expectedStrSlice, actualStrSlice)
		})
	}
}

func Test_HeaderSlice_SetHttpRequestHeaders(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name           string
		cType          *customtypes.HeaderSlice
		expectedHeader http.Header
	}{
		{
			name:           "Happy path - single header",
			cType:          utils.Pointer(customtypes.HeaderSlice{{Key: "key", Value: "value"}}),
			expectedHeader: http.Header{"Key": []string{"value"}},
		},
		{
			name:           "Happy path - multiple headers",
			cType:          utils.Pointer(customtypes.HeaderSlice{{Key: "key1", Value: "value1"}, {Key: "key2", Value: "value2"}}),
			expectedHeader: http.Header{"Key1": []string{"value1"}, "Key2": []string{"value2"}},
		},
		{
			name:           "Happy path - empty",
			cType:          new(customtypes.HeaderSlice),
			expectedHeader: http.Header{},
		},
		{
			name:           "Nil custom header type",
			cType:          nil,
			expectedHeader: http.Header{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			req, err := http.NewRequest(http.MethodGet, "http://localhost", nil)
			require.NoError(t, err)

			tc.cType.SetHttpRequestHeaders(req)

			require.Equal(t, tc.expectedHeader, req.Header)
		})
	}
}
