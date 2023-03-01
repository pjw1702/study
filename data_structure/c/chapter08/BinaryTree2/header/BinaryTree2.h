#ifndef __BINARY_TREE2_H__
#define __BINARY_TREE2_H__

typedef int BTData;

typedef struct _bTreeNode
{
  BTData data;
  struct _bTreeNode * left;
  struct _bTreeNode * right;
  
} BTreeNode;

BTreeNode * MakeBTreeNode(void);              // 노드의 생성

BTData GetData(BTreeNode * bt);               // 노드에 저장된 데이터를 반환
void SetData(BTreeNode * bt, BTData data);    // 노드에 데이터를 저장

BTreeNode * GetLeftSubTree(BTreeNode * bt);   // 왼쪽 서브 트리의 주소 값(루트 노드 주소 값) 반환
BTreeNode * GetRightSubTree(BTreeNode * bt);  // 오른쪽 서브 트리의 주소 값(루트 노드 주소 값) 반환

// 루트 노드, 또는 마지막 노드(terminal node)의 주소 값과 자식 노드의 주소 값을 전달하여 각각 왼쪽, 오른쪽에 연결함으로써 새 이진 트리 형태의 (서브)트리를 생성한다
// 단일 노드 뿐만 아니라 트리와 트리를 연결할 수도 있다
void MakeLeftSubTree(BTreeNode * main, BTreeNode * sub);    // main의 서브 왼쪽 서브 트리로 sub를 연결
void MakeRightSubTree(BTreeNode * main, BTreeNode * sub);   // main의 서브 오른쪽 서브 트리로 sub를 연결

// 함수 포인터 형  VisitFuncPtr의 정의
// 함수 정의 앞에 typedef로 선언하면 해당 함수의 이름이 해당 함수의 형식에 해당하는 타입으로서 선언된다
typedef void VisitFuncPtr(BTData data);		// (*VisitFncPtr)으로 대신해도 된다

void PreorderTraverse(BTreeNode * bt, VisitFuncPtr action);
void InorderTraverse(BTreeNode * bt, VisitFuncPtr action);
void PostorderTraverse(BTreeNode * bt, VisitFuncPtr action);

#endif