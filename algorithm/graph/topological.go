package graph

// TopologicalSort 拓扑排序（Kahn 算法，基于入度 + BFS）。
// 返回拓扑序和一个布尔值：false 表示图中有环，无法完成拓扑排序。
func (g *Graph) TopologicalSort() ([]int, bool) {
	indeg := make([]int, g.n)
	for v := 0; v < g.n; v++ {
		for _, e := range g.edges[v] {
			indeg[e.To]++
		}
	}
	q := make([]int, 0, g.n)
	for v := 0; v < g.n; v++ {
		if indeg[v] == 0 {
			q = append(q, v)
		}
	}
	order := make([]int, 0, g.n)
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		order = append(order, v)
		for _, e := range g.edges[v] {
			indeg[e.To]--
			if indeg[e.To] == 0 {
				q = append(q, e.To)
			}
		}
	}
	if len(order) != g.n {
		return order, false // 有环
	}
	return order, true
}

// TopologicalSortDFS 拓扑排序（DFS 逆后序）。
// 后序遍历顺序反过来就是拓扑序；同时用三色标记检测环。
func (g *Graph) TopologicalSortDFS() ([]int, bool) {
	const (
		unvisited = iota // 0
		visiting         // 1
		visited          // 2
	)
	color := make([]int, g.n)
	order := make([]int, 0, g.n)
	var dfs func(v int) bool
	dfs = func(v int) bool {
		color[v] = visiting
		for _, e := range g.edges[v] {
			switch color[e.To] {
			case visiting:
				return false // 遇到仍在递归栈中的节点，说明成环
			case unvisited:
				if !dfs(e.To) {
					return false
				}
			}
		}
		color[v] = visited
		order = append(order, v) // 后序
		return true
	}
	for v := 0; v < g.n; v++ {
		if color[v] == unvisited {
			if !dfs(v) {
				return nil, false
			}
		}
	}
	// 逆后序即拓扑序
	for i, j := 0, len(order)-1; i < j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}
	return order, true
}

// HasCycleDirected 判断有向图是否存在环（DFS 三色标记）
func (g *Graph) HasCycleDirected() bool {
	_, ok := g.TopologicalSortDFS()
	return !ok
}

// HasCycleUndirected 判断无向图是否存在环（DFS，忽略来路 parent）。
// 注意：调用方需保证图是无向图（每条边都双向添加）。
func (g *Graph) HasCycleUndirected() bool {
	visited := make([]bool, g.n)
	var dfs func(v, parent int) bool
	dfs = func(v, parent int) bool {
		visited[v] = true
		for _, e := range g.edges[v] {
			if !visited[e.To] {
				if dfs(e.To, v) {
					return true
				}
			} else if e.To != parent {
				return true // 访问过且不是父节点，说明有环
			}
		}
		return false
	}
	for v := 0; v < g.n; v++ {
		if !visited[v] {
			if dfs(v, -1) {
				return true
			}
		}
	}
	return false
}
