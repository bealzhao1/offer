# 图算法实现文档

> 代码目录：`algorithm/graph/`（Go 库，`package graph`）
> 演示入口：`algorithm/graph/example/main.go`

## 1. 功能概述

实现图的常用算法，覆盖「表示 → 遍历 → 拓扑 → 最短路 → 生成树 → 连通性 → 网络流」全链路。

## 2. 文件组织

| 文件 | 内容 |
| --- | --- |
| `graph.go` | 图的邻接表表示 + DFS/BFS 遍历 |
| `topological.go` | 拓扑排序（Kahn / DFS）+ 环检测 |
| `shortest_path.go` | Dijkstra / Bellman-Ford / Floyd-Warshall |
| `mst.go` | Prim / Kruskal 最小生成树 |
| `unionfind.go` | 并查集 |
| `bipartite.go` | 二分图判定（染色法） |
| `connectivity.go` | 连通分量 / 强连通分量（Tarjan） |
| `maxflow.go` | 最大流（Edmonds-Karp） |
| `example/main.go` | 所有算法的演示 |

## 3. 图的表示

```go
type Graph struct {
    n     int      // 节点数，编号 0..n-1
    edges [][]Edge // 邻接表
}

type Edge struct {
    To     int // 目标节点
    Weight int // 边权
}
```

- `New(n)` 创建空图。
- `AddEdge(u, v, w)` 加有向边；`AddUndirectedEdge(u, v, w)` 加无向边（双向）。
- 邻接表空间 O(V+E)，适合稀疏图，是最通用的表示。

## 4. 遍历

| 函数 | 方式 | 用途 |
| --- | --- | --- |
| `DFS(start)` | 递归 | 深搜、找路径、连通块 |
| `DFSIterative(start)` | 栈 | 避免递归爆栈 |
| `BFS(start)` | 队列 | 无权最短路、层序遍历 |

- 三者时间 O(V+E)。
- DFS 迭代版逆序压栈，保证与递归版访问顺序一致。

## 5. 拓扑排序与环检测

| 函数 | 说明 | 环检测 |
| --- | --- | --- |
| `TopologicalSort()` | Kahn：入度 + BFS | 结果长度 < n 则有环 |
| `TopologicalSortDFS()` | DFS 逆后序 + 三色标记 | 遇到「访问中」节点则成环 |
| `HasCycleDirected()` | 有向图环检测 | 复用 DFS 三色 |
| `HasCycleUndirected()` | 无向图环检测 | 访问过且非父节点 |

- 时间 O(V+E)。
- 应用：课程表、任务依赖、编译器链接顺序。

## 6. 最短路径

| 函数 | 适用 | 复杂度 |
| --- | --- | --- |
| `Dijkstra(start)` | 非负权单源 | O((V+E)logV) |
| `BellmanFord(start)` | 含负权单源，检测负环 | O(VE) |
| `FloydWarshall(matrix)` | 多源全对 | O(V³) |

- `inf = 1e9` 表示不可达。
- `ToMatrix()` 把邻接表转邻接矩阵，供 Floyd 使用。

## 7. 最小生成树

| 函数 | 思路 | 复杂度 |
| --- | --- | --- |
| `Prim()` | 加点（贪心 + 堆） | O(E logV)，适合稠密图 |
| `Kruskal(n, edges)` | 加边（并查集 + 排序） | O(E logE)，适合稀疏图 |

- 输入均要求无向连通图。

## 8. 并查集

```go
uf := NewUnionFind(n)
uf.Union(x, y)  // 合并，返回是否真正合并
uf.Find(x)      // 找根，路径压缩
uf.Same(x, y)   // 判断连通
uf.Count()      // 连通分量个数
```

- 路径压缩 + 按秩合并，单次操作近乎 O(α(n)) ≈ O(1)。

## 9. 二分图

- `IsBipartite()`：BFS 染色法，0/1 交替染色，相邻同色则不是二分图。O(V+E)。
- 应用：任务分配、相亲/匹配问题。

## 10. 连通性

| 函数 | 说明 |
| --- | --- |
| `ConnectedComponents()` | 无向图连通分量 |
| `StronglyConnectedComponents()` | 有向图强连通分量（Tarjan） |

- Tarjan 用 `dfn`（发现时间）+ `low`（能回溯的最小编号），`dfn[v] == low[v]` 时 v 是 SCC 根。

## 11. 网络流

- `MaxFlow(n, edges, s, t)`：Edmonds-Karp（BFS 找增广路），`edges` 为 `[from, to, capacity]`。
- 复杂度 O(VE²)。
- 应用：物流、交通容量、二分图匹配（最大流 = 最小割）。

## 12. 复杂度总览

| 算法 | 时间 | 空间 |
| --- | --- | --- |
| DFS/BFS | O(V+E) | O(V) |
| 拓扑排序 | O(V+E) | O(V) |
| Dijkstra | O((V+E)logV) | O(V) |
| Bellman-Ford | O(VE) | O(V) |
| Floyd | O(V³) | O(V²) |
| Prim | O(E logV) | O(V) |
| Kruskal | O(E logE) | O(V) |
| 并查集 | ≈O(1)/次 | O(V) |
| 二分图判定 | O(V+E) | O(V) |
| Tarjan SCC | O(V+E) | O(V) |
| 最大流 (EK) | O(VE²) | O(V+E) |

## 13. 运行

```bash
# 运行演示
go run ./algorithm/graph/example/

# 编译检查
go build ./...
go vet ./algorithm/graph/...
```

作为库使用：

```go
import "offer/algorithm/graph"

g := graph.New(5)
g.AddEdge(0, 1, 10)
g.AddEdge(0, 2, 3)
dist := g.Dijkstra(0) // [0 7 3 9 5]
```
