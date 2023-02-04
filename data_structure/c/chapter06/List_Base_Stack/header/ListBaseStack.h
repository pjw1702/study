#ifndef __LB_STACK_H__
#define __LB_STACK_H__

#define TRUE    1
#define FALSE   0

typedef int Data;

// 노드
typedef struct _node {
	Data data;
	struct _node * next;

} Node;

// 스택
// 삽입, 삭제, 검색 등에 필요한 연산은 헤드 하나로도 충분하다
typedef struct _liststack {
	Node * head;
} ListStack;

typedef ListStack Stack;

void StackInit(Stack * pstack);
int SIsEmpty(Stack * pstack);

void SPush(Stack * pstack, Data data);
Data SPop(Stack * pstack);
Data SPeek(Stack * pstack);

#endif