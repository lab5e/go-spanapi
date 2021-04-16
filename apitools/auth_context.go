package apitools

import (
	"context"
	"time"

	"github.com/lab5e/go-spanapi/v4"
)

// ContextWithAuth returns a context with the API token set
func ContextWithAuth(token string, timeout time.Duration) (ctx context.Context, done func()) {
	ctx, done = context.WithTimeout(context.Background(), timeout)

	keys := make(map[string]spanapi.APIKey)
	keys["APIToken"] = spanapi.APIKey{Key: token, Prefix: ""}
	ctx = context.WithValue(ctx, spanapi.ContextAPIKeys, keys)
	return
}
