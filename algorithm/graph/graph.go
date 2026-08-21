// Package graph 提供图数据结构的常用算法实现：
// 遍历、拓扑排序、环检测、最短路径、最小生成树、并查集、二分图、连通性与网络流。
package graph

// Graph 邻接表表示的图（带权边）。
// 通过 AddEdge 构造有向图，通过 AddUndirectedEdge 构造无向图。
type Graph struct {
	n     int
	edges [][]Edge
}

// Edge 一条带权边
type Edge struct {
	To     int
	Weight int
}

// New 创建包含 n 个节点（编号 0..n-1）的空图
func New(n int) *Graph {
	return &Graph{n: n, edges: make([][]Edge, n)}
}

// N 返回节点数
func (g *Graph) N() int { return g.n }

// Adj 返回节点 v 的所有出边
func (g *Graph) Adj(v int) []Edge { return g.edges[v] }

// AddEdge 添加有向边 u -> v，权重 w
func (g *Graph) AddEdge(u, v, w int) {
	g.edges[u] = append(g.edges[u], Edge{To: v, Weight: w})
}

// AddUndirectedEdge 添加无向边 u <-> v，权重 w
func (g *Graph) AddUndirectedEdge(u, v, w int) {
	g.AddEdge(u, v, w)
	g.AddEdge(v, u, w)
}

// DFS 深度优先遍历（递归），返回从 start 可达节点的访问顺序
func (g *Graph) DFS(start int) []int {
	visited := make([]bool, g.n)
	var order []int
	var dfs func(v int)
	dfs = func(v int) {
		visited[v] = true
		order = append(order, v)
		for _, e := range g.edges[v] {
			if !visited[e.To] {
				dfs(e.To)
			}
		}
	}
	dfs(start)
	return order
}

// DFSIterative 深度优先遍历（栈）。
// 逆序压栈是为了让编号靠前的邻点先出栈，从而和递归版的访问顺序一致。
func (g *Graph) DFSIterative(start int) []int {
	visited := make([]bool, g.n)
	var order []int
	stack := []int{start}
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[v] {
			continue
		}
		visited[v] = true
		order = append(order, v)
		for i := len(g.edges[v]) - 1; i >= 0; i-- {
			if !visited[g.edges[v][i].To] {
				stack = append(stack, g.edges[v][i].To)
			}
		}
	}
	return order
}

// BFS 广度优先遍历（队列），返回从 start 可达节点的访问顺序
func (g *Graph) BFS(start int) []int {
	visited := make([]bool, g.n)
	var order []int
	q := []int{start}
	visited[start] = true
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		order = append(order, v)
		for _, e := range g.edges[v] {
			if !visited[e.To] {
				visited[e.To] = true
				q = append(q, e.To)
			}
		}
	}
	return order
}
