// Copyright © 2025 Ping Identity Corporation

package plugins

import (
	"context"
	"errors"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

type mockTokenSource struct {
	token *oauth2.Token
	err   error
}

func (m *mockTokenSource) Token() (*oauth2.Token, error) {
	return m.token, m.err
}

func TestPluginAuthenticator_GetToken_Success(t *testing.T) {
	a := &PluginAuthenticator{getValidTokenSource: func(ctx context.Context) (oauth2.TokenSource, error) {
		return &mockTokenSource{token: &oauth2.Token{AccessToken: "abc", TokenType: "Bearer", Expiry: time.Now().Add(time.Hour)}}, nil
	}}
	tok, err := a.GetToken()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if tok != "abc" {
		t.Fatalf("expected token %q, got %q", "abc", tok)
	}
}

func TestPluginAuthenticator_GetToken_TokenSourceError(t *testing.T) {
	expectedErr := errors.New("no config")
	a := &PluginAuthenticator{getValidTokenSource: func(ctx context.Context) (oauth2.TokenSource, error) {
		return nil, expectedErr
	}}
	_, err := a.GetToken()
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected wrapped error %v, got %v", expectedErr, err)
	}
}

func TestPluginAuthenticator_GetToken_TokenError(t *testing.T) {
	expectedErr := errors.New("token failure")
	a := &PluginAuthenticator{getValidTokenSource: func(ctx context.Context) (oauth2.TokenSource, error) {
		return &mockTokenSource{err: expectedErr}, nil
	}}
	_, err := a.GetToken()
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected wrapped error %v, got %v", expectedErr, err)
	}
}

func TestPluginAuthenticator_GetToken_EmptyToken(t *testing.T) {
	a := &PluginAuthenticator{getValidTokenSource: func(ctx context.Context) (oauth2.TokenSource, error) {
		return &mockTokenSource{token: &oauth2.Token{AccessToken: ""}}, nil
	}}
	_, err := a.GetToken()
	if err == nil {
		t.Fatalf("expected error")
	}
}
