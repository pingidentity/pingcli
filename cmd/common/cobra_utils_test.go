// Copyright © 2026 Ping Identity Corporation

package common_test

import (
	"testing"

	"github.com/pingidentity/pingcli/cmd/common"
	"github.com/stretchr/testify/require"
)

func Test_ExactArgs(t *testing.T) {
	testCases := []struct {
		name        string
		numArgs     int
		args        []string
		expectedErr error
	}{
		{
			name:        "No arguments, expecting 0",
			numArgs:     0,
			args:        []string{},
			expectedErr: nil,
		},
		{
			name:        "One argument, expecting 1",
			numArgs:     1,
			args:        []string{"arg1"},
			expectedErr: nil,
		},
		{
			name:        "Two arguments, expecting 2",
			numArgs:     2,
			args:        []string{"arg1", "arg2"},
			expectedErr: nil,
		},
		{
			name:        "Three arguments, expecting 2",
			numArgs:     2,
			args:        []string{"arg1", "arg2", "arg3"},
			expectedErr: common.ErrExactArgs,
		},
		{
			name:        "No arguments, expecting 1",
			numArgs:     1,
			args:        []string{},
			expectedErr: common.ErrExactArgs,
		},
		{
			name:        "One argument, expecting 0",
			numArgs:     0,
			args:        []string{"arg1"},
			expectedErr: common.ErrExactArgs,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			posArgsFunc := common.ExactArgs(tc.numArgs)
			err := posArgsFunc(nil, tc.args)

			if tc.expectedErr != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_RangeArgs(t *testing.T) {
	testCases := []struct {
		name        string
		minArgs     int
		maxArgs     int
		args        []string
		expectedErr error
	}{
		{
			name:        "No arguments, expecting 0 to 2",
			minArgs:     0,
			maxArgs:     2,
			args:        []string{},
			expectedErr: nil,
		},
		{
			name:        "One argument, expecting 1 to 2",
			minArgs:     1,
			maxArgs:     2,
			args:        []string{"arg1"},
			expectedErr: nil,
		},
		{
			name:        "Two arguments, expecting 1 to 2",
			minArgs:     1,
			maxArgs:     2,
			args:        []string{"arg1", "arg2"},
			expectedErr: nil,
		},
		{
			name:        "Three arguments, expecting 1 to 2",
			minArgs:     1,
			maxArgs:     2,
			args:        []string{"arg1", "arg2", "arg3"},
			expectedErr: common.ErrRangeArgs,
		},
		{
			name:        "No arguments, expecting 1 to 2",
			minArgs:     1,
			maxArgs:     2,
			args:        []string{},
			expectedErr: common.ErrRangeArgs,
		},
		{
			name:        "One argument, expecting 0 to 0",
			minArgs:     0,
			maxArgs:     0,
			args:        []string{"arg1"},
			expectedErr: common.ErrRangeArgs,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			posArgsFunc := common.RangeArgs(tc.minArgs, tc.maxArgs)
			err := posArgsFunc(nil, tc.args)

			if tc.expectedErr != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
