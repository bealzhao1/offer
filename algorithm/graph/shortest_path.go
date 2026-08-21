package graph

import "container/heap"

// inf 用作"无穷大"，表示不可达或无边
const inf = int(1e9)

// ---------- 优先队列（小顶堆）实现 ----------

type pqItem struct {
	node int
	dist int
}

type priorityQueue []pqItem

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq priorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(x any)        { *pq = append(*pq, x.(pqItem)) }
func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	it := old[n-1]
	*pq = old[:n-1]
	return it
}

// Dijkstra 单源最短路径（堆优化），要求所有边权非负。
// 返回从 start 到各节点的最短距离，不可达节点为 inf。
func (g *Graph) Dijkstra(start int) []int {
	dist := make([]int, g.n)
	for i := range dist {
		dist[i] = inf
	}
	dist[start] = 0
	pq := &priorityQueue{{start, 0}}
	heap.Init(pq)
	for pq.Len() > 0 {
		it := heap.Pop(pq).(pqItem)
		if it.dist > dist[it.node] {
			continue // 过期的旧记录，跳过
		}
		for _, e := range g.edges[it.node] {
			if nd := it.dist + e.Weight; nd < dist[e.To] {
				dist[e.To] = nd
				heap.Push(pq, pqItem{e.To, nd})
			}
		}
	}
	return dist
}

// BellmanFord 单源最短路径，支持负权边，并能检测负环。
// 返回距离数组和一个布尔值：false 表示存在从 start 可达的负环。
func (g *Graph) BellmanFord(start int) ([]int, bool) {
	dist := make([]int, g.n)
	for i := range dist {
		dist[i] = inf
	}
	dist[start] = 0
	// 最多松弛 n-1 轮
	for i := 0; i < g.n-1; i++ {
		updated := false
		for v := 0; v < g.n; v++ {
			if dist[v] == inf {
				continue
			}
			for _, e := range g.edges[v] {
				if dist[v]+e.Weight < dist[e.To] {
					dist[e.To] = dist[v] + e.Weight
					updated = true
				}
			}
		}
		if !updated {
			break
		}
	}
	// 第 n 轮还能继续松弛 => 存在负环
	for v := 0; v < g.n; v++ {
		if dist[v] == inf {
			continue
		}
		for _, e := range g.edges[v] {
			if dist[v]+e.Weight < dist[e.To] {
				return dist, false
			}
		}
	}
	return dist, true
}

// ToMatrix 把图转为邻接矩阵：无边用 inf 表示，对角线为 0。
// 多重边取最小权。
func (g *Graph) ToMatrix() [][]int {
	m := make([][]int, g.n)
	for i := range m {
		m[i] = make([]int, g.n)
		for j := range m[i] {
			m[i][j] = inf
		}
		m[i][i] = 0
	}
	for v := 0; v < g.n; v++ {
		for _, e := range g.edges[v] {
			if e.Weight < m[v][e.To] {
				m[v][e.To] = e.Weight
			}
		}
	}
	return m
}

// FloydWarshall 多源最短路径（Floyd-Warshall）。
// 输入为邻接矩阵 dist：dist[i][j] 表示 i->j 边权，无边为 inf，对角线为 0。
// 返回任意两点间的最短距离矩阵。
func FloydWarshall(dist [][]int) [][]int {
	n := len(dist)
	d := make([][]int, n)
	for i := range d {
		d[i] = append([]int(nil), dist[i]...)
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if d[i][k] < inf && d[k][j] < inf && d[i][k]+d[k][j] < d[i][j] {
					d[i][j] = d[i][k] + d[k][j]
				}
			}
		}
	}
	return d
}
