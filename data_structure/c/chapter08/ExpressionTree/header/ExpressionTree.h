#ifdef __EXPRESSION_TREE_H__
#define __EXPRESSION_TREE_H__

#include "..\BinaryTree2\header\BinaryTree2.h"

BTreeNode * MakeExpTree(char exp[]);		// 수식 트리 구성 (후위 표기법의 수식을 인자로 받아서 수식 트리를 구성하고 루트 노드의 주소 값을 반환)

int EvaluateExpTree(BTreeNode * bt);		// 수식 트리 계산(MakeExpTree가 구성한 수식 트리의 수식을 계산하여 그 결과를 반환)

void ShowPrefixTypeExp(BTreeNode * bt);	// 전위 표기법 기반 출력
void ShowInfixTypeExp(BTreeNode * bt);	// 중위 표기법 기반 출력
void ShowPostfixTypeExp(BTreeNode * bt);	// 후위 표기법 기반 출력

// 전위, 중위, 후위 순회하여 출력 시 각각 전위, 중위 후위 표기법의 수식이 출력된다

#endif

