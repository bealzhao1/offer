package graph

import (
	"container/heap"
	"sort"
)

// Prim 最小生成树（加点法，堆优化）。
// 要求图是无向连通图，返回 MST 的总权重。
func (g *Graph) Prim() int {
	if g.n == 0 {
		return 0
	}
	visited := make([]bool, g.n)
	total := 0
	pq := &priorityQueue{{node: 0, dist: 0}} // 从节点 0 开始
	heap.Init(pq)
	for pq.Len() > 0 {
		it := heap.Pop(pq).(pqItem)
		if visited[it.node] {
			continue
		}
		visited[it.node] = true
		total += it.dist
		for _, e := range g.edges[it.node] {
			if !visited[e.To] {
				heap.Push(pq, pqItem{node: e.To, dist: e.Weight})
			}
		}
	}
	return total
}

// Kruskal 最小生成树（加边法，并查集 + 按权重排序）。
// edges 为无向边 [u, v, w]，n 为节点数，返回 MST 总权重。
// 若图不连通，返回的只是各连通块内部的 MST 权重之和。
func Kruskal(n int, edges [][3]int) int {
	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })
	uf := NewUnionFind(n)
	total := 0
	for _, e := range edges {
		if uf.Union(e[0], e[1]) {
			total += e[2]
		}
	}
	return total
}
