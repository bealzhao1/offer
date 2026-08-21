package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Val         int
	Left, Right *Node
}

// buildPreorder 按 根->左->右 的先序展开，-2 表示空节点
func buildPreorder(arr []int, i *int) *Node {
	if *i >= len(arr) {
		return nil
	}
	v := arr[*i]
	*i++
	if v == -2 {
		return nil
	}
	n := &Node{Val: v}
	n.Left = buildPreorder(arr, i)
	n.Right = buildPreorder(arr, i)
	return n
}

// printTree 自顶向下打印二叉树，带分支连线
// 思路：先用中序遍历给每个节点分配横坐标 x（左子树在左、右子树在右、父节点居中于两子树之间），
//
//	再按层输出节点，层与层之间画 ┌─┴─┐ 连线。
func printTree(root *Node) {
	xs := map[*Node]int{}
	var nextX int
	var inorder func(n *Node)
	inorder = func(n *Node) {
		if n == nil {
			return
		}
		inorder(n.Left)
		xs[n] = nextX
		nextX++
		inorder(n.Right)
	}
	inorder(root)

	maxD := 0
	var setDepth func(n *Node, d int)
	setDepth = func(n *Node, d int) {
		if n == nil {
			return
		}
		if d > maxD {
			maxD = d
		}
		setDepth(n.Left, d+1)
		setDepth(n.Right, d+1)
	}
	setDepth(root, 0)

	totalCols := nextX
	colW := 4
	width := totalCols * colW
	rows := (maxD+1)*2 - 1
	grid := make([][]rune, rows)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	levelNodes := make([][]*Node, maxD+1)
	var collect func(n *Node, d int)
	collect = func(n *Node, d int) {
		if n == nil {
			return
		}
		levelNodes[d] = append(levelNodes[d], n)
		collect(n.Left, d+1)
		collect(n.Right, d+1)
	}
	collect(root, 0)

	// 节点行
	for d := 0; d <= maxD; d++ {
		for _, n := range levelNodes[d] {
			base := xs[n]*colW + 1
			copy(grid[d*2][base:], []rune(strconv.Itoa(n.Val)))
		}
	}
	// 连线行
	for d := 0; d < maxD; d++ {
		r := d*2 + 1
		for _, n := range levelNodes[d] {
			px := xs[n]
			lx, rx := -1, -1
			if n.Left != nil {
				lx = xs[n.Left]
			}
			if n.Right != nil {
				rx = xs[n.Right]
			}
			if lx >= 0 {
				grid[r][lx*colW+1] = '┌'
				for c := lx*colW + 2; c < px*colW+1; c++ {
					grid[r][c] = '─'
				}
			}
			if rx >= 0 {
				for c := px*colW + 2; c < rx*colW+1; c++ {
					grid[r][c] = '─'
				}
				grid[r][rx*colW+1] = '┐'
			}
			if lx >= 0 || rx >= 0 {
				grid[r][px*colW+1] = '┴'
			}
		}
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func main() {
	arr := []int{10, 0, 1, -2, -2, 11, -2, 1, 15, -2, -2, 11, -2, -2, 12, -2, 0, 3, -2, 4, -2, -2, 1, -2, -2}
	i := 0
	root := buildPreorder(arr, &i)
	fmt.Printf("先序消费 %d/%d 个元素\n\n", i, len(arr))
	printTree(root)

	fmt.Println("\n===== 三种遍历：递归 vs 栈 =====")
	fmt.Println("先序(递归):", preorderRecursive(root))
	fmt.Println("先序(栈)  :", preorderIterative(root))
	fmt.Println("中序(递归):", inorderRecursive(root))
	fmt.Println("中序(栈)  :", inorderIterative(root))
	fmt.Println("后序(递归):", postorderRecursive(root))
	fmt.Println("后序(栈)  :", postorderIterative(root))
}
