package apitools

import (
	"context"
	"time"

	"github.com/lab5e/go-spanapi/v5"
)

// ContextWithAuth returns a context with the API token set
func ContextWithAuth(token string) context.Context {
	ctx := context.Background()

	keys := make(map[string]spanapi.APIKey)
	keys["APIToken"] = spanapi.APIKey{Key: token, Prefix: ""}
	return context.WithValue(ctx, spanapi.ContextAPIKeys, keys)
}

// ContextWithAuthAndTimeout returns a context with the API token set and timeout
func ContextWithAuthAndTimeout(token string, timeout time.Duration) (context.Context, func()) {
	ctx, done := context.WithTimeout(context.Background(), timeout)

	keys := make(map[string]spanapi.APIKey)
	keys["APIToken"] = spanapi.APIKey{Key: token, Prefix: ""}
	return context.WithValue(ctx, spanapi.ContextAPIKeys, keys), done
}
