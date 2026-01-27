// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"testing"
)

func TestStorageType_Set_ValidValues(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"file_system", ENUM_STORAGE_TYPE_FILE_SYSTEM},
		{"secure_local", ENUM_STORAGE_TYPE_SECURE_LOCAL},
		{"secure_remote", ENUM_STORAGE_TYPE_SECURE_REMOTE},
		{"none", ENUM_STORAGE_TYPE_NONE},
		{"FILE_SYSTEM", ENUM_STORAGE_TYPE_FILE_SYSTEM}, // case-insensitive
		{"SECURE_LOCAL", ENUM_STORAGE_TYPE_SECURE_LOCAL},
		{"SECURE_REMOTE", ENUM_STORAGE_TYPE_SECURE_REMOTE},
		{"NONE", ENUM_STORAGE_TYPE_NONE},
	}

	for _, tc := range cases {
		var st StorageType
		if err := (&st).Set(tc.in); err != nil {
			t.Fatalf("Set(%q) unexpected error: %v", tc.in, err)
		}
		if got := st.String(); got != tc.want {
			t.Fatalf("Set(%q) => %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestStorageType_Set_EmptyDefaultsToSecureLocal(t *testing.T) {
	var st StorageType
	if err := (&st).Set(""); err != nil {
		t.Fatalf("Set(\"\") error: %v", err)
	}
	if got, want := st.String(), ENUM_STORAGE_TYPE_SECURE_LOCAL; got != want {
		t.Fatalf("Set(\"\") => %q, want %q", got, want)
	}
}

func TestStorageType_Set_Invalid(t *testing.T) {
	var st StorageType
	if err := (&st).Set("invalid_value"); err == nil {
		t.Fatalf("Set(invalid_value) expected error, got nil with value %q", st.String())
	}
}

func TestStorageType_String_NilReceiver(t *testing.T) {
	var st *StorageType
	if got := st.String(); got != "" {
		t.Fatalf("nil.String() => %q, want empty string", got)
	}
}

func TestStorageType_Type(t *testing.T) {
	var st StorageType
	if got, want := (&st).Type(), "string"; got != want {
		t.Fatalf("Type() => %q, want %q", got, want)
	}
}
