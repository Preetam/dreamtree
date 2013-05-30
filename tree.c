#include <stdlib.h>
#include <stdio.h>
#include <string.h>

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
	//	printf("left insert\n");
		tree->left = insert_(tree->left, n);
		return tree;
	}

	tree->right = insert_(tree->right, n);
	//printf("right insert\n");
	return tree;
}

Node*
insert(Node* tree, void* data, int length) {
/*
	if(tree == NULL) {
		printf("	- tree is null\n");
		tree = new_tree(data, length);
		return tree;
	}

	tree->size++;
	if(tree->left == NULL && tree->right == NULL) tree->height++;

	if(memcmp(tree->data, data, (length > tree->data_length ? length : tree->data_length)) > 0) {
		printf("	- left insert\n");
		tree->left = insert(tree->left, data, length);
		return tree;
	}
	
	printf("	- right insert\n");
	tree->right = insert(tree->right, data, length);
	return tree;
*/
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
//	return left_height+right_height+1;
	if(left_height >= right_height)
		return left_height+1;
	else
		return right_height+1;
}

Node*
balance(Node* tree) {
	if(tree == NULL) return NULL;

	tree->left = balance(tree->left);
	tree->right = balance(tree->right);

	int left_height = height(tree->left);
	int right_height = height(tree->right);
	if( abs(left_height - right_height) <= 1) {
	//	printf("No change to: %s\n", tree->data);
		return tree;
	}

	if( left_height < right_height ) {
	//	printf("%s is right-heavy\n", tree->data);
		Node* root = tree->right;
		tree->right = NULL;
		root = insert_(root, tree);
		return root;
	}
	else {
	//	printf("%s is left-heavy\n", tree->data);
		Node* root = tree->left;
		tree->left = NULL;
		root = insert_(root, tree);
		return root;
	}
}

void
insert_many(Node* tree) {
	for(int i = 55; i < 72; i++) {
		int* dat = malloc(sizeof(int));
		*dat = i;
		insert(tree, dat, sizeof(int));
	}
}

int main() {
	Node* t = new_tree("a",1);
	insert(t, "b", 1);
	insert(t, "c", 1);
	insert(t, "d", 1);
	insert(t, "e", 1);
	//insert_many(t);
	traverse(t,0);
	printf("Height: %d\n", height(t));
	t = balance(t);
	t = balance(t);
	traverse(t, 0);
	printf("Height: %d\n", height(t));
	printf("\n%s\n", t->data);
	return 0;
}
