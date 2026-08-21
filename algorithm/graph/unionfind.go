package graph

// UnionFind 并查集（路径压缩 + 按秩合并），
// 用于动态维护「连通性」：判断两个元素是否在同一个集合、合并集合、统计连通块个数。
type UnionFind struct {
	parent []int
	rank   []int
	count  int // 连通分量个数
}

// NewUnionFind 初始化 n 个相互独立的元素（0..n-1）
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		count:  n,
	}
	for i := range uf.parent {
		uf.parent[i] = i
	}
	return uf
}

// Find 返回 x 所在集合的根，沿途做路径压缩
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union 合并 x、y 所在集合，返回是否真正发生了合并（原本不连通）
func (uf *UnionFind) Union(x, y int) bool {
	rx, ry := uf.Find(x), uf.Find(y)
	if rx == ry {
		return false
	}
	// 按秩合并：把矮树挂到高树上
	if uf.rank[rx] < uf.rank[ry] {
		rx, ry = ry, rx
	}
	uf.parent[ry] = rx
	if uf.rank[rx] == uf.rank[ry] {
		uf.rank[rx]++
	}
	uf.count--
	return true
}

// Same 判断 x、y 是否连通
func (uf *UnionFind) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// Count 返回当前连通分量个数
func (uf *UnionFind) Count() int { return uf.count }
