package main;

/*
#cgo CFLAGS: -std=c99
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <time.h>

struct Node;

typedef struct {
	void* left;
	void* right;
	void* data;
	int data_length;
} Node;

Node* balance(Node* t);

Node*
new_tree(void* data, int length) {
	Node* n = malloc(sizeof(Node));
	n->data = data;
	n->data_length = length;
	return n;
}

Node*
insert_(Node* tree, Node* n) {
	if(tree == NULL) return n;

	if(memcmp(tree->data, n->data, (tree->data_length < n->data_length ? n->data_length : tree->data_length)) >= 0) {
		tree->left = insert_(tree->left, n);
		return tree;
	}

	tree->right = insert_(tree->right, n);
	return tree;
}

Node*
insert(Node* tree, void* data, int length) {
	return insert_(tree, new_tree(data, length));
}

Node*
delete(Node* tree, void* data) {
	return tree;
}

void
traverse(Node* tree, int space) {
	char* spacer = malloc(space);
	for(int i = 0; i < space; i++)
		spacer[i] = 0;
	for(int i = 0; i < space; i++)
		spacer[i] = ' ';
	if(tree == NULL) return;
	traverse(tree->left, space+2);
	printf("%s%s\n", spacer, (char*)tree->data);
	traverse(tree->right, space+2);
}

Node*
get_least(Node* tree) {
	if(tree->left == NULL)
		return tree;
	else
		return get_least(tree->left);
}

int
height(Node* tree) {
	if(tree == NULL) return 0;

	int left_height = height(tree->left);
	int right_height = height(tree->right);
	if(left_height >= right_height)
		return left_height+1;
	else
		return right_height+1;
}


Node*
balance(Node* tree) {
	if(tree == NULL) return NULL;

	int left_height = height(tree->left);
	int right_height = height(tree->right);

	while(abs(left_height - right_height) >= 2) {
		
		if( left_height <= right_height ) {
		//	printf("%s is right-heavy\n", tree->data);
			Node* root = tree->right;
			tree->right = NULL;
			root = insert_(root, tree);
			tree = root;
		}
		else {
		//	printf("%s is left-heavy\n", tree->data);
			Node* root = tree->left;
			tree->left = NULL;
			root = insert_(root, tree);
			tree = root;
		}
	
		left_height = height(tree->left);
		right_height = height(tree->right);
	}

	tree->left = balance(tree->left);
	tree->right = balance(tree->right);
	return tree;
}

void
insert_many(Node* tree) {
	for(int i = 0; i < 100000; i++) {
		int* dat = malloc(sizeof(int));
		*dat = i;
		insert(tree, dat, sizeof(int));
	}
}
*/
import "C"

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
	Value string
}

func Create() *Tree {
	t := new(Tree)
	t.Size = 0
	return t
}

func Insert(root *Node, n *Node) (*Node) {
	if(root == nil) {
		return n
	}

	if(n.Value > root.Value) {
		root.Right = Insert(root.Right, n)
	} else {
		root.Left = Insert(root.Left, n)
	}

	return root

}

func Traverse(root *Node, padding int) {
	if(root == nil) {
		return
	}

	Traverse(root.Left, padding+1)
	tmp := padding
	for (tmp > 0) {
		fmt.Print("  ")
		tmp--
	}
	fmt.Println(root.Value)
	Traverse(root.Right, padding+1)
}

func Height(root *Node, ch chan int) {
	if(root == nil) {
		ch<-0
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

	if(h1 >= h2) {
		ch<-h1+1
	} else {
		ch<-h2+1
	}
}

func Balance(root *Node) *Node {
	if(root == nil) {
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

	for math.Abs( (float64)(h1 - h2)) >= 2 {
		if(h1 <= h2) {
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

func main() {
	t := Create()
	t.Root = Insert(t.Root, &Node{nil, nil, "b"})
	t.Root = Insert(t.Root, &Node{nil, nil, "d"})
	t.Root = Insert(t.Root, &Node{nil, nil, "a"})
	t.Root = Insert(t.Root, &Node{nil, nil, "c"})
	t.Root = Insert(t.Root, &Node{nil, nil, "e"})
	t.Root = Insert(t.Root, &Node{nil, nil, "f"})
	t.Root = Insert(t.Root, &Node{nil, nil, "g"})
	t.Root = Insert(t.Root, &Node{nil, nil, "h"})
	t.Root = Insert(t.Root, &Node{nil, nil, "i"})
	t.Root = Insert(t.Root, &Node{nil, nil, "j"})
	t.Root = Insert(t.Root, &Node{nil, nil, "k"})
	t.Root = Insert(t.Root, &Node{nil, nil, "l"})
	t.Root = Insert(t.Root, &Node{nil, nil, "m"})
	t.Root = Insert(t.Root, &Node{nil, nil, "n"})
	Traverse(t.Root, 0)
	ch := make(chan int)
	go func() {
		Height(t.Root, ch)
	}()
	h := <-ch
	fmt.Println("Height:", h)
	t.Root = Balance(t.Root)
	go func() {
		Height(t.Root, ch)
	}()
	h = <-ch
	fmt.Println("Height:", h)
	Traverse(t.Root, 0)
}
