package graph

// ConnectedComponents 求无向图的连通分量，返回每个连通分量的节点列表。
// 注意：针对无向图使用（每条边双向添加）。
func (g *Graph) ConnectedComponents() [][]int {
	visited := make([]bool, g.n)
	var comps [][]int
	for s := 0; s < g.n; s++ {
		if visited[s] {
			continue
		}
		var comp []int
		q := []int{s}
		visited[s] = true
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			comp = append(comp, v)
			for _, e := range g.edges[v] {
				if !visited[e.To] {
					visited[e.To] = true
					q = append(q, e.To)
				}
			}
		}
		comps = append(comps, comp)
	}
	return comps
}

// StronglyConnectedComponents 求有向图的强连通分量（Tarjan 算法）。
// 返回每个强连通分量（SCC）的节点列表。
func (g *Graph) StronglyConnectedComponents() [][]int {
	index := 0
	dfn := make([]int, g.n) // 深度优先编号（发现时间）
	low := make([]int, g.n) // 能回溯到的最小编号
	for i := range dfn {
		dfn[i] = -1
	}
	onStack := make([]bool, g.n)
	var stack []int
	var sccs [][]int

	var dfs func(v int)
	dfs = func(v int) {
		dfn[v] = index
		low[v] = index
		index++
		stack = append(stack, v)
		onStack[v] = true

		for _, e := range g.edges[v] {
			if dfn[e.To] == -1 { // 树边
				dfs(e.To)
				if low[e.To] < low[v] {
					low[v] = low[e.To]
				}
			} else if onStack[e.To] { // 回边（指向还在栈中的节点）
				if dfn[e.To] < low[v] {
					low[v] = dfn[e.To]
				}
			}
		}

		// 自己是 SCC 的根
		if dfn[v] == low[v] {
			var scc []int
			for {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				onStack[top] = false
				scc = append(scc, top)
				if top == v {
					break
				}
			}
			sccs = append(sccs, scc)
		}
	}

	for v := 0; v < g.n; v++ {
		if dfn[v] == -1 {
			dfs(v)
		}
	}
	return sccs
}
