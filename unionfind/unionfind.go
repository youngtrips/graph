package unionfind

type UnionFind struct {
	ranks []int32
	nodes []int32
	count int
}

func New(n int) *UnionFind {
	uf := &UnionFind{
		ranks: make([]int32, n),
		nodes: make([]int32, n),
		count: n,
	}
	for i := 0; i < n; i++ {
		uf.nodes[i] = int32(i)
		uf.ranks[i] = 1
	}
	return uf
}

func (uf *UnionFind) Union(p int32, q int32) {
	pr := uf.Find(p)
	qr := uf.Find(q)
	if pr == qr {
		return
	}
	if uf.ranks[pr] < uf.ranks[qr] {
		uf.ranks[qr] += uf.ranks[pr]
		uf.nodes[pr] = qr
	} else {
		uf.ranks[pr] += uf.ranks[qr]
		uf.nodes[qr] = pr
	}
	uf.count--
}

func (uf *UnionFind) Find(p int32) int32 {
	for p != uf.nodes[p] {
		p = uf.nodes[p]
	}
	return p
}

func (uf *UnionFind) Connected(p int32, q int32) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UnionFind) Count() int {
	return uf.count
}
