// Copyright © 2025 Ping Identity Corporation

package resources

import "errors"

var (
	ErrFlowPolicyNameNotFound   = errors.New("flow policy name not found for flow policy ID")
	ErrResourceNameNotFound     = errors.New("resource name not found for grant resource ID")
	ErrRoleNameNotFound         = errors.New("role name not found for role ID")
	ErrUnexpectedResponse       = errors.New("unexpected response - worker apps cannot read their own secret")
	ErrSignOnPolicyNameNotFound = errors.New("sign-on policy name not found for sign-on policy ID")
)
