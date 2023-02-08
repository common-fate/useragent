package useragent

import (
	"context"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := context.Background()

	ctx = NewContext(ctx, "test", "v1.0.0")
	got := FromContext(ctx)
	want := "test/v1.0.0"

	if got != want {
		t.Fatalf("got = %s, want = %s", got, want)
	}
}
