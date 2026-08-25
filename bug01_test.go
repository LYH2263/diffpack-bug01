package diffpack_test

import (
	"context"
	"testing"
)

func TestBug01_BuildDeltaInsertAlias(t *testing.T) {
	p := newPacker(t)
	base := []byte("aaaaaa")
	target := []byte("aaabbb")
	b := mustBuild(t, p, "j1", base, target)
	target[3] = 'z'
	out, err := p.ApplyDelta(context.Background(), b, base)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte("aaabbb")
	if string(out) != string(want) {
		t.Fatalf("target mutation leaked: want %q got %q", want, out)
	}
}
