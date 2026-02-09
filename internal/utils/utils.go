// Copyright © 2026 Ping Identity Corporation

package utils

func Pointer[T any](t T) *T {
	return &t
}
