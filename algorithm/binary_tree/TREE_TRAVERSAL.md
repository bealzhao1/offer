# 二叉树三种遍历（递归 + 栈）核心流程

> 代码文件：`traversal.go`

## 1. 功能概述

实现二叉树的三种深度优先遍历，每种遍历均提供两种实现：

1. **递归**：利用函数调用栈，代码简洁直观。
2. **栈（迭代）**：用显式 `stack` 模拟调用栈，避免递归深度过大导致栈溢出。

三种遍历的访问顺序如下（约定「根」为当前节点）：

| 遍历 | 顺序 | 记忆方式 |
| --- | --- | --- |
| 先序 | 根 → 左 → 右 | 「根」在最前 |
| 中序 | 左 → 根 → 右 | 「根」在中间 |
| 后序 | 左 → 右 → 根 | 「根」在最后 |

## 2. 核心流程总览

```
            10
          /    \
         0      12
        / \
       1   11

先序：10 0 1 11 12
中序：1 0 11 10 12
后序：1 11 0 12 10
```

- **先序**：路过一个节点就先记下它的值。
- **中序**：先钻到最左，再回头记值。
- **后序**：把左右子树都看完，最后才记根的值。

## 3. 先序遍历（根-左-右）

### 3.1 递归

```go
func preorderRecursive(root *Node) []int {
    var res []int
    var dfs func(n *Node)
    dfs = func(n *Node) {
        if n == nil { return }
        res = append(res, n.Val) // 先访问根
        dfs(n.Left)              // 再左子树
        dfs(n.Right)             // 再右子树
    }
    dfs(root)
    return res
}
```

### 3.2 栈

- **思路**：栈里始终存「待访问的根」。弹出即访问，然后**先压右、再压左**，这样左孩子会先出栈，符合 根-左-右 顺序。
- **关键**：先压右、再压左（栈是 LIFO，后压的左孩子先弹出）。

```go
func preorderIterative(root *Node) []int {
    if root == nil { return nil }
    var res []int
    stack := []*Node{root}
    for len(stack) > 0 {
        n := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, n.Val)
        if n.Right != nil { stack = append(stack, n.Right) } // 先压右
        if n.Left != nil  { stack = append(stack, n.Left) }  // 再压左
    }
    return res
}
```

## 4. 中序遍历（左-根-右）

### 4.1 递归

```go
func inorderRecursive(root *Node) []int {
    var res []int
    var dfs func(n *Node)
    dfs = func(n *Node) {
        if n == nil { return }
        dfs(n.Left)              // 先左子树
        res = append(res, n.Val) // 再访问根
        dfs(n.Right)             // 再右子树
    }
    dfs(root)
    return res
}
```

### 4.2 栈

- **思路**：一路向左把整条左链压栈，弹出栈顶访问它，再转向它的右子树。
- **关键**：中序最「左」，所以先 `for cur != nil` 把最左路径全部压栈。

```go
func inorderIterative(root *Node) []int {
    var res []int
    var stack []*Node
    cur := root
    for cur != nil || len(stack) > 0 {
        for cur != nil { // 把整条左链压栈
            stack = append(stack, cur)
            cur = cur.Left
        }
        cur = stack[len(stack)-1] // 弹出并访问
        stack = stack[:len(stack)-1]
        res = append(res, cur.Val)
        cur = cur.Right            // 转向右子树
    }
    return res
}
```

## 5. 后序遍历（左-右-根）

### 5.1 递归

```go
func postorderRecursive(root *Node) []int {
    var res []int
    var dfs func(n *Node)
    dfs = func(n *Node) {
        if n == nil { return }
        dfs(n.Left)              // 先左子树
        dfs(n.Right)             // 再右子树
        res = append(res, n.Val) // 最后访问根
    }
    dfs(root)
    return res
}
```

### 5.2 栈（重点）

- **难点**：后序的根必须等左右子树**都访问完**才能访问，栈里弹出的时机最难把握。
- **思路**：用 `lastVisited` 记录上一个被访问的节点，判断栈顶节点的右子树是否已处理：
  - 右子树为空，或右子树就是上次访问的节点 → 可以访问栈顶；
  - 否则 → 先转向右子树继续压栈。

```go
func postorderIterative(root *Node) []int {
    var res []int
    var stack []*Node
    var lastVisited *Node
    cur := root
    for cur != nil || len(stack) > 0 {
        for cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        }
        n := stack[len(stack)-1] // 只看栈顶，先不弹
        if n.Right == nil || n.Right == lastVisited {
            res = append(res, n.Val)
            stack = stack[:len(stack)-1]
            lastVisited = n
        } else {
            cur = n.Right
        }
    }
    return res
}
```

## 6. 三种栈实现记忆口诀

| 遍历 | 栈的核心技巧 | 一句话 |
| --- | --- | --- |
| 先序 | 先压右、再压左 | 弹出就访问，右先左后压 |
| 中序 | 一路向左压左链 | 左链全压栈，弹出转右 |
| 后序 | `lastVisited` 标记 | 右子树没看先别看根 |

## 7. 复杂度

| 指标 | 复杂度 |
| --- | --- |
| 时间 | O(N)，每个节点恰好访问一次 |
| 空间（递归） | O(H)，H 为树高（最坏退化成链表时为 O(N)） |
| 空间（栈） | O(H)，同上 |

> 递归与栈的空间复杂度理论上相同，但栈实现避免了「递归调用栈」在极深树下的爆栈风险（Go 栈会动态增长，但极端情况下迭代更稳）。

## 8. 主要用法

```go
arr := []int{10, 0, 1, -2, -2, 11, -2, -2, 12, -2, -2}
i := 0
root := buildPreorder(arr, &i)

fmt.Println(preorderRecursive(root))  // [10 0 1 11 12]
fmt.Println(preorderIterative(root))  // [10 0 1 11 12]
fmt.Println(inorderRecursive(root))   // [1 0 11 10 12]
fmt.Println(inorderIterative(root))   // [1 0 11 10 12]
fmt.Println(postorderRecursive(root)) // [1 11 0 12 10]
fmt.Println(postorderIterative(root)) // [1 11 0 12 10]
```
