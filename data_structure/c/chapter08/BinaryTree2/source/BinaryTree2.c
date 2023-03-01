#include <stdio.h>
#include <stdlib.h>
#include "..\header\BinaryTree2.h"

BTreeNode * MakeBTreeNode(void) {
  BTreeNode * nd = (BTreeNode *)malloc(sizeof(BTreeNode));
  nd->left = NULL;
  nd->right = NULL;
  return nd;
}

BTData GetData(BTreeNode * bt) {
  return bt->data;
}

void SetData(BTreeNode * bt, BTData data) {
  bt->data = data;
}

BTreeNode * GetLeftSubTree(BTreeNode * bt) {
  return bt->left;
}


BTreeNode * GetRightSubTree(BTreeNode * bt) {
  return bt->right;
}

void MakeLeftSubTree(BTreeNode * main, BTreeNode * sub) {
  // 기존에 연결된 노드는 삭제하고 새로 연결
  // 기존에 연결된 노드가 단일 노드가 아닌 서브 트리일 경우가 있으므로 고민해야할 부분이다: 순회(Traversal)을 통해 해결할 수 있다
  if(main->left != NULL)
    free(main->left);
  
  main->left = sub;
}

void MakeRightSubTree(BTreeNode * main, BTreeNode * sub) {
  // 기존에 연결된 노드는 삭제하고 새로 연결
  // 기존에 연결된 노드가 단일 노드가 아닌 서브 트리일 경우가 있으므로 고민해야할 부분이다: 순회(Traversal)을 통해 해결할 수 있다
  if(main->right != NULL)
    free(main->right);
  
  main->right = sub;
}

void PreorderTraverse(BTreeNode * bt, VisitFuncPtr action) {
	if(bt == NULL) {		// bt가 NULL이면 재귀 탈출
		return;
	}
	InorderTraverse(bt->left, action);
	action(bt->data);		// printf()함수는 데이터 출력을 목적으로 하는 함수인데, 노드의 방문을 나타내는 기능으로서 printf()함수를 사용하면 가독성 및 기능 표현 측면에서 깔끔하지 못하다.
	InorderTraverse(bt->right, action);
}

void InorderTraverse(BTreeNode * bt, VisitFuncPtr action) {
	if(bt == NULL) {		
		return;
	}
	InorderTraverse(bt->left, action);
	action(bt->data);		
	InorderTraverse(bt->right, action);
}

void PostorderTraverse(BTreeNode * bt, VisitFuncPtr action) {
	if(bt == NULL) {		
		return;
	}
	InorderTraverse(bt->left, action);
	action(bt->data);		
	InorderTraverse(bt->right, action);
}