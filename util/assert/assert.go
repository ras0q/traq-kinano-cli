package assert

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Error(t *testing.T, wantErr bool, err error) {
	t.Helper()

	if !wantErr && err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func Equal(t *testing.T, want interface{}, got interface{}, opts ...cmp.Option) {
	t.Helper()

	if diff := cmp.Diff(want, got, opts...); len(diff) > 0 {
		t.Error(diff)
	}
}
