package unionfind

import (
	"testing"
)

func TestUnionFind(t *testing.T) {
	uf := New(10)
	uf.Union(1, 9)
	uf.Union(2, 9)
	uf.Union(2, 6)
	uf.Union(5, 6)
	uf.Union(3, 7)
	uf.Union(5, 7)
	uf.Union(4, 7)
	uf.Union(4, 8)
	t.Log(uf.Connected(1, 8))
	t.Log(uf.Connected(2, 5))
	t.Log(uf.Connected(3, 5))
	t.Log(uf.Count())
}
