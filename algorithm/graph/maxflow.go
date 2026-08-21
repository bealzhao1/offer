package graph

// MaxFlow 用 Edmonds-Karp（BFS 找增广路）计算从 s 到 t 的最大流。
// edges 为有向边 [from, to, capacity]，n 为节点数。
func MaxFlow(n int, edges [][3]int, s, t int) int {
	// 残量网络：邻接表存 (to, rev, cap)，rev 指向反向边在 to 的邻接表中的下标
	type fEdge struct {
		to, rev, cap int
	}
	g := make([][]fEdge, n)
	addEdge := func(u, v, c int) {
		g[u] = append(g[u], fEdge{v, len(g[v]), c})
		g[v] = append(g[v], fEdge{u, len(g[u]) - 1, 0})
	}
	for _, e := range edges {
		addEdge(e[0], e[1], e[2])
	}

	total := 0
	for {
		// BFS 找一条增广路，同时记录父节点与父边
		parent := make([]int, n)
		parentEdge := make([]int, n)
		for i := range parent {
			parent[i] = -1
		}
		parent[s] = s
		q := []int{s}
		for len(q) > 0 && parent[t] == -1 {
			v := q[0]
			q = q[1:]
			for i, e := range g[v] {
				if e.cap > 0 && parent[e.to] == -1 {
					parent[e.to] = v
					parentEdge[e.to] = i
					q = append(q, e.to)
				}
			}
		}
		if parent[t] == -1 {
			break // 没有增广路了
		}
		// 找瓶颈流量
		f := inf
		for v := t; v != s; v = parent[v] {
			if c := g[parent[v]][parentEdge[v]].cap; c < f {
				f = c
			}
		}
		// 更新残量网络：正向减、反向加
		for v := t; v != s; v = parent[v] {
			p := parent[v]
			ei := parentEdge[v]
			g[p][ei].cap -= f
			g[v][g[p][ei].rev].cap += f
		}
		total += f
	}
	return total
}
