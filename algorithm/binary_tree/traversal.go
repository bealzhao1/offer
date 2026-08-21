package main

// 二叉树的三种遍历：先序、中序、后序。
// 每种遍历均提供「递归」与「栈（迭代）」两种实现。
//
// 先序：根 -> 左 -> 右
// 中序：左 -> 根 -> 右
// 后序：左 -> 右 -> 根

// ---------- 先序遍历 ----------

// preorderRecursive 先序遍历（递归）
func preorderRecursive(root *Node) []int {
	var res []int
	var dfs func(n *Node)
	dfs = func(n *Node) {
		if n == nil {
			return
		}
		res = append(res, n.Val)
		dfs(n.Left)
		dfs(n.Right)
	}
	dfs(root)
	return res
}

// preorderIterative 先序遍历（栈）
// 栈里始终存"待访问的根"。弹出即访问，然后先压右、再压左，
// 这样左孩子会先出栈，符合 根-左-右 的顺序。
func preorderIterative(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := []*Node{root}
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, n.Val)
		if n.Right != nil {
			stack = append(stack, n.Right)
		}
		if n.Left != nil {
			stack = append(stack, n.Left)
		}
	}
	return res
}

// ---------- 中序遍历 ----------

// inorderRecursive 中序遍历（递归）
func inorderRecursive(root *Node) []int {
	var res []int
	var dfs func(n *Node)
	dfs = func(n *Node) {
		if n == nil {
			return
		}
		dfs(n.Left)
		res = append(res, n.Val)
		dfs(n.Right)
	}
	dfs(root)
	return res
}

// inorderIterative 中序遍历（栈）
// 一路向左把整条左链压栈，弹出栈顶访问它，再转向它的右子树。
func inorderIterative(root *Node) []int {
	var res []int
	var stack []*Node
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

// ---------- 后序遍历 ----------

// postorderRecursive 后序遍历（递归）
func postorderRecursive(root *Node) []int {
	var res []int
	var dfs func(n *Node)
	dfs = func(n *Node) {
		if n == nil {
			return
		}
		dfs(n.Left)
		dfs(n.Right)
		res = append(res, n.Val)
	}
	dfs(root)
	return res
}

// postorderIterative 后序遍历（栈）
// 后序最麻烦：根必须等左右子树都访问完才能访问。
// 用 lastVisited 记录上一个被访问的节点，判断栈顶的右子树是否已处理：
//   - 右子树为空，或右子树就是上次访问的节点 → 可以访问栈顶；
//   - 否则 → 先转向右子树继续压栈。
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
		n := stack[len(stack)-1]
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
