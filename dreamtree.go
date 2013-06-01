package dreamtree

import (
	"fmt"
	"math"
)

type Tree struct {
	Root *Node
	Size int
}

type Node struct {
	Left, Right *Node
	Value       string
}

func Create() *Tree {
	t := new(Tree)
	t.Size = 0
	return t
}

func Insert(root *Node, n *Node) *Node {
	if root == nil {
		return n
	}

	if n.Value > root.Value {
		root.Right = Insert(root.Right, n)
	} else {
		root.Left = Insert(root.Left, n)
	}

	return root

}

func Remove(root *Node, n *Node) *Node {
	if root == nil {
		return nil
	}

	if n == nil {
		return root
	}

	if root.Value == n.Value {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil && root.Right != nil {
			return root.Right
		} else if root.Left != nil && root.Right == nil {
			return root.Left
		} else if root.Left != nil && root.Right != nil {
			max := Max(root.Left)
			root.Value = max.Value
			root.Left = Remove(root.Left, max)
			return root
		}
	}

	if root.Value > n.Value {
		root.Left = Remove(root.Left, n)
	}

	if root.Value < n.Value {
		root.Right = Remove(root.Right, n)
	}

	return root
}

func Traverse(root *Node, padding int) {
	if root == nil {
		return
	}

	Traverse(root.Left, padding+1)
	tmp := padding
	for tmp > 0 {
		fmt.Print("  ")
		tmp--
	}
	fmt.Println(root.Value)
	Traverse(root.Right, padding+1)
}

func Get(root *Node, val string) *Node {
	if root == nil {
		return nil
	}

	if root.Value > val {
		return Get(root.Left, val)
	}

	if root.Value < val {
		return Get(root.Right, val)
	}

	return root
}

func Max(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Right == nil {
		return root
	}

	return Max(root.Right)
}

func Min(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Left == nil {
		return root
	}

	return Max(root.Left)
}

func Height(root *Node, ch chan int) {
	if root == nil {
		ch <- 0
		return
	}

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		Height(root.Left, c1)
	}()
	go func() {
		Height(root.Right, c2)
	}()

	h1 := <-c1
	h2 := <-c2

	if h1 >= h2 {
		ch <- h1 + 1
	} else {
		ch <- h2 + 1
	}
}

func Balance(root *Node) *Node {
	if root == nil {
		return nil
	}

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		Height(root.Left, c1)
	}()
	go func() {
		Height(root.Right, c2)
	}()

	h1 := <-c1
	h2 := <-c2

	for math.Abs((float64)(h1-h2)) >= 2 {
		if h1 <= h2 {
			newRoot := root.Right
			root.Right = nil
			Insert(newRoot, root)
			root = newRoot

		} else {
			newRoot := root.Left
			root.Left = nil
			Insert(newRoot, root)
			root = newRoot

		}
		go func() {
			Height(root.Left, c1)
		}()
		go func() {
			Height(root.Right, c2)
		}()

		h1 = <-c1
		h2 = <-c2
	}
	root.Left = Balance(root.Left)
	root.Right = Balance(root.Right)
	return root
}

func (t *Tree) Insert(val string) {
	t.Root = Insert(t.Root, &Node{nil, nil, val})
	t.Size++
}

func (t *Tree) Remove(val string) {
	node := Get(t.Root, val)
	if node != nil {
		t.Size--
	}

	t.Root = Remove(t.Root, node)
}

func (t *Tree) Balance() {
	t.Root = Balance(t.Root)
}

func (t *Tree) Height(ch chan int) {
	Height(t.Root, ch)
}
