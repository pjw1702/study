#include <stdio.h>
#include "..\header\BinaryTree2.h"

void ShowIntData(int data);

int main(void) {
  BTreeNode * bt1 = MakeBTreeNode();    // 노드 bt1 생성
  BTreeNode * bt2 = MakeBTreeNode();    // 노드 bt2 생성
  BTreeNode * bt3 = MakeBTreeNode();    // 노드 bt3 생성
  BTreeNode * bt4 = MakeBTreeNode();    // 노드 bt4 생성
  BTreeNode * bt5 = MakeBTreeNode();    // 노드 bt5 생성
  BTreeNode * bt6 = MakeBTreeNode();    // 노드 bt6 생성

  SetData(bt1, 1);    // bt1에 1 저장
  SetData(bt2, 2);    // bt2에 2 저장
  SetData(bt3, 3);    // bt3에 3 저장
  SetData(bt4, 4);    // bt4에 4 저장
  SetData(bt5, 5);    // bt5에 5 저장
  SetData(bt6, 6);    // bt6에 6 저장

  MakeLeftSubTree(bt1, bt2);    // bt2를 bt1의 왼쪽 자식 노드로 연결
  MakeRightSubTree(bt1, bt3);   // bt3을 bt1의 오른쪽 자식 노드로 연결
  MakeLeftSubTree(bt2, bt4);    // bt4를 bt2의 왼쪽 자식 노드로 연결
  MakeRightSubTree(bt2, bt5);   // bt5을 bt2의 오른쪽 자식 노드로 연결
  MakeRightSubTree(bt3, bt6);   // bt6을 bt3의 오른쪽 자식 노드로 연결

  // 순회
  PreorderTraverse(bt1, ShowIntData);
  printf("\n");
  InorderTraverse(bt1, ShowIntData);
  printf("\n");
  PostorderTraverse(bt1, ShowIntData);
  printf("\n");

  return 0;
}

void ShowIntData(int data) {
	printf("%d", data);
}