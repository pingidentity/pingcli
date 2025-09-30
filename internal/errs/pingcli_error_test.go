// Copyright © 2025 Ping Identity Corporation

package errs_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/stretchr/testify/require"
)

func Test_PingCLIError_Error(t *testing.T) {
	testErr := errors.New("test error")
	prefix1 := "prefix 1"
	prefix2 := "prefix 2"

	testCases := []struct {
		name         string
		err          *errs.PingCLIError
		expectedStr  string
		expectedAs   error
		expectedIs   error
		assertUnwrap require.ErrorAssertionFunc
	}{
		{
			name: "Happy path",
			err: &errs.PingCLIError{
				Prefix: prefix1,
				Err:    testErr,
			},
			expectedStr:  fmt.Sprintf("%s: %s", prefix1, testErr.Error()),
			expectedAs:   &errs.PingCLIError{},
			expectedIs:   testErr,
			assertUnwrap: require.Error,
		},
		{
			name: "Nested PingCLIError with same prefix",
			err: &errs.PingCLIError{
				Prefix: prefix1,
				Err: &errs.PingCLIError{
					Prefix: prefix1,
					Err:    testErr,
				},
			},
			expectedStr:  fmt.Sprintf("%s: %s", prefix1, testErr.Error()),
			expectedAs:   &errs.PingCLIError{},
			expectedIs:   testErr,
			assertUnwrap: require.Error,
		},
		{
			name: "Nested PingCLIError with different prefix",
			err: &errs.PingCLIError{
				Prefix: prefix2,
				Err: &errs.PingCLIError{
					Prefix: prefix1,
					Err:    testErr,
				},
			},
			expectedStr:  fmt.Sprintf("%s: %s: %s", prefix2, prefix1, testErr.Error()),
			expectedAs:   &errs.PingCLIError{},
			expectedIs:   testErr,
			assertUnwrap: require.Error,
		},
		{
			name: "Nil inner error",
			err: &errs.PingCLIError{
				Prefix: prefix1,
				Err:    nil,
			},
			expectedStr:  "",
			expectedAs:   nil,
			expectedIs:   nil,
			assertUnwrap: require.NoError,
		},
		{
			name:         "Nil PingCLIError",
			err:          nil,
			expectedStr:  "",
			expectedAs:   nil,
			expectedIs:   nil,
			assertUnwrap: require.NoError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err != nil {
				require.Equal(t, tc.expectedStr, tc.err.Error())
			} else {
				require.Equal(t, tc.expectedStr, "")
			}

			if tc.expectedAs != nil {
				var target *errs.PingCLIError
				require.ErrorAs(t, tc.err, &target)
			}

			if tc.expectedIs != nil {
				require.ErrorIs(t, tc.err, tc.expectedIs)
			}

			unwrappedErr := errors.Unwrap(tc.err)
			tc.assertUnwrap(t, unwrappedErr)
		})
	}
}
