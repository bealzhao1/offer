package graph

// IsBipartite 判断图是否为二分图（BFS 染色法）。
// 用 0/1 两种颜色交替染色，若相邻节点同色则不是二分图。
// 注意：通常针对无向图使用（每条边双向添加）。
func (g *Graph) IsBipartite() bool {
	color := make([]int, g.n)
	for i := range color {
		color[i] = -1 // 未染色
	}
	for s := 0; s < g.n; s++ {
		if color[s] != -1 {
			continue
		}
		color[s] = 0
		q := []int{s}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, e := range g.edges[v] {
				switch color[e.To] {
				case -1:
					color[e.To] = 1 - color[v] // 交替染色
					q = append(q, e.To)
				case color[v]:
					return false // 相邻同色，冲突
				}
			}
		}
	}
	return true
}
