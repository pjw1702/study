#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include "..\List_Base_Stack\header\ListBaseStack.h"
#include "..\..\BinaryTree2\header\BinaryTree2.h"

// 피연산자는 스택으로 옮김
// 연산자를 만나면 스택에서 두 개의 피 연산자를 꺼내어 자식 노드로 연결
// 자식 노드를 연결해서 만들어진 트리는 다시 스택으로 옮김
BTreeNode * MakeExpTree(char exp[]) {
	Stack stack;
	BTreeNode * pnode;
	int expLen = strlen(exp);
	int i;
	for(i = 0; i < expLen; i++){
		pnode = MakeBTreeNode();
		if(isdigit(exp[i])) {				// 피 연산자라면
			SetData(pnode, exp[i]-'0');		// 문자를 정수로 바꿔서 저장(트리에 저장)
		} else {					// 연산자라면
			MakeRightSubTree(pnode, SPop(&stack));
			MakeLeftSubTree(pnode, SPop(&stack));
			SetData(pnode, exp[i]);		// 연산자 정보 또한 저장
		}
		SPush(&stack, pnode);	// 스택에 저장 (서브 트리 또한 저장해야 하므로 연산자, 피연산자 경우의 수 상관없이 무조건 진행)
	}
	return SPop(&stack);	// 수식 트리의 루트 노드의 주소 값을 반환
}

int EvaluateExpTree(BTreeNode * bt) {
	int op1, op2;

	if(GetLeftSubTree(bt) == NULL && GetRightSubTree(bt) == NULL)		// 단만 노드일 경우
		return GetData(bt);

  // 피연산자가 서브트리일 경우엔 재귀적으로 구현하여 연산을 진행한다
	op1 = EvaluateExpTree(GetLeftSubTree(bt));		// 첫 번째 피연산자
	op2 = EvaluateExpTree(GetRightSubTree(bt));		// 두 번째 피연산자
  
	switch(GetData(bt)) {			// 연산자를 확인하여 연산을 진행
	case '+':
		return op1 + op2;
	case '-':
		return op1 - op2;
	case '*':
		return op1 * op2;
	case '/':
		return op1 / op2;
	}
	return 0;
}

void ShowPrefixTypeExp(BTreeNode *bt) {
	PreorderTraverse(bt, ShowNodeData);		// 전위 표기법 수식 출력
}
void InPrefixTypeExp(BTreeNode *bt) {
	InorderTraverse(bt, ShowNodeData);		// 중위 표기법 수식 출력
}
void ShowPostfixTypeExp(BTreeNode *bt) {
	PostorderTraverse(bt, ShowNodeData);		// 후위 표기법 수식 출력
}

void ShowNodeData(int data) {
	if(0<=data && data<=9)
		printf("%d", data);		// 피연산자 출력
	else
		printf("%c", data);		// 연산자 출력
}
