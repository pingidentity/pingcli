// Copyright © 2026 Ping Identity Corporation

package connector

// A connector that allows authentication
type Authenticatable interface {
	Login() error
	Logout() error
}
