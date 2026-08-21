package main

import (
	"fmt"

	"offer/algorithm/graph"
)

func main() {
	// 1. 遍历 DFS/BFS
	fmt.Println("========== 1. 遍历 DFS/BFS ==========")
	g := graph.New(6)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(4, 5, 1)
	fmt.Println("DFS 递归:", g.DFS(0))
	fmt.Println("DFS 栈  :", g.DFSIterative(0))
	fmt.Println("BFS     :", g.BFS(0))

	// 2. 拓扑排序
	fmt.Println("\n========== 2. 拓扑排序 ==========")
	if order, ok := g.TopologicalSort(); ok {
		fmt.Println("Kahn 拓扑序:", order)
	} else {
		fmt.Println("Kahn: 存在环")
	}
	if order, ok := g.TopologicalSortDFS(); ok {
		fmt.Println("DFS 拓扑序 :", order)
	}

	// 3. 环检测
	fmt.Println("\n========== 3. 环检测 ==========")
	fmt.Println("有向 DAG 是否有环:", g.HasCycleDirected())
	gCycle := graph.New(3)
	gCycle.AddEdge(0, 1, 1)
	gCycle.AddEdge(1, 2, 1)
	gCycle.AddEdge(2, 0, 1)
	if _, ok := gCycle.TopologicalSort(); ok {
		fmt.Println("环图拓扑排序: 成功")
	} else {
		fmt.Println("环图拓扑排序: 失败（有环）")
	}

	// 4. 最短路径
	fmt.Println("\n========== 4. 最短路径 ==========")
	sp := graph.New(5)
	sp.AddEdge(0, 1, 10)
	sp.AddEdge(0, 2, 3)
	sp.AddEdge(1, 2, 1)
	sp.AddEdge(1, 3, 2)
	sp.AddEdge(2, 1, 4)
	sp.AddEdge(2, 3, 8)
	sp.AddEdge(2, 4, 2)
	sp.AddEdge(3, 4, 7)
	sp.AddEdge(4, 3, 9)
	fmt.Println("Dijkstra 从 0:", sp.Dijkstra(0))
	if d, ok := sp.BellmanFord(0); ok {
		fmt.Println("BellmanFord 从 0:", d)
	}
	fmt.Println("Floyd 第 0 行:", graph.FloydWarshall(sp.ToMatrix())[0])

	// 5. 最小生成树
	fmt.Println("\n========== 5. 最小生成树 ==========")
	ug := graph.New(5)
	ug.AddUndirectedEdge(0, 1, 2)
	ug.AddUndirectedEdge(0, 3, 6)
	ug.AddUndirectedEdge(1, 2, 3)
	ug.AddUndirectedEdge(1, 3, 8)
	ug.AddUndirectedEdge(1, 4, 5)
	ug.AddUndirectedEdge(2, 4, 7)
	ug.AddUndirectedEdge(3, 4, 9)
	fmt.Println("Prim 总权重:", ug.Prim())
	mstEdges := [][3]int{{0, 1, 2}, {0, 3, 6}, {1, 2, 3}, {1, 3, 8}, {1, 4, 5}, {2, 4, 7}, {3, 4, 9}}
	fmt.Println("Kruskal 总权重:", graph.Kruskal(5, mstEdges))

	// 6. 并查集
	fmt.Println("\n========== 6. 并查集 ==========")
	uf := graph.NewUnionFind(5)
	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(3, 4)
	fmt.Println("0 与 2 连通:", uf.Same(0, 2))
	fmt.Println("0 与 3 连通:", uf.Same(0, 3))
	fmt.Println("连通分量个数:", uf.Count())

	// 7. 二分图
	fmt.Println("\n========== 7. 二分图 ==========")
	even := graph.New(4)
	even.AddUndirectedEdge(0, 1, 1)
	even.AddUndirectedEdge(1, 2, 1)
	even.AddUndirectedEdge(2, 3, 1)
	even.AddUndirectedEdge(3, 0, 1)
	fmt.Println("偶环(应为 true):", even.IsBipartite())
	tri := graph.New(3)
	tri.AddUndirectedEdge(0, 1, 1)
	tri.AddUndirectedEdge(1, 2, 1)
	tri.AddUndirectedEdge(2, 0, 1)
	fmt.Println("三角形(应为 false):", tri.IsBipartite())

	// 8. 连通分量 / 强连通分量
	fmt.Println("\n========== 8. 连通分量 / 强连通分量 ==========")
	cc := graph.New(5)
	cc.AddUndirectedEdge(0, 1, 1)
	cc.AddUndirectedEdge(1, 2, 1)
	cc.AddUndirectedEdge(3, 4, 1)
	fmt.Println("无向连通分量:", cc.ConnectedComponents())
	scg := graph.New(5)
	scg.AddEdge(0, 1, 1)
	scg.AddEdge(1, 2, 1)
	scg.AddEdge(2, 0, 1)
	scg.AddEdge(2, 3, 1)
	scg.AddEdge(3, 4, 1)
	scg.AddEdge(4, 3, 1)
	fmt.Println("强连通分量:", scg.StronglyConnectedComponents())

	// 9. 最大流
	fmt.Println("\n========== 9. 最大流 ==========")
	flow := [][3]int{
		{0, 1, 16}, {0, 2, 13}, {1, 2, 10}, {1, 3, 12},
		{2, 1, 4}, {2, 4, 14}, {3, 2, 9}, {3, 5, 20},
		{4, 3, 7}, {4, 5, 4},
	}
	fmt.Println("最大流(应为 23):", graph.MaxFlow(6, flow, 0, 5))
}
