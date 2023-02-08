package useragent

import (
	"context"
	"fmt"
)

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

// userAgentKey is the key for useragent values in Contexts. It is
// unexported; clients use useragent.NewContext and useragent.FromContext
// instead of using this key directly.
var userAgentKey key

func NewContext(ctx context.Context, product, version string) context.Context {
	ua := fmt.Sprintf("%s/%s", product, version)
	return context.WithValue(ctx, userAgentKey, ua)
}

func FromContext(ctx context.Context) string {
	got, ok := ctx.Value(userAgentKey).(string)
	if !ok {
		return ""
	}
	return got
}
