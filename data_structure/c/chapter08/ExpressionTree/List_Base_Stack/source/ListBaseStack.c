#include <stdio.h>
#include <stdlib.h>
#include "ListBaseStack.h"

void StackInit(Stack * pstack) {
	pstack->head = NULL;
}

int SIsEmpty(Stack * pstack) {
	if(pstack->head == NULL) 
		return TRUE;
	else
		return FALSE;
	
}

void SPush(Stack * pstack, Data data) {
	Node * newNode = (Node*)malloc(sizeof(Node));

	newNode->data = data;
    // 새 노드는 기존의 head가 가리키던 노드를 가리키게 됨
	newNode->next = pstack->head;

    // head는 이제 push한 새 노드를 가리키게 됨
	pstack->head = newNode;
}

Data SPop(Stack * pstack) {
	Data rdata;
	Node * rnode;

	if(SIsEmpty(pstack)) {
		printf("Stack Memory Error!");
		exit(-1);
	}
	// 삭제할 노드의 주소 값, 데이터 백업
	rdata = pstack->head->data;
	rnode = pstack->head;

    // head는 삭제 할 노드가 가리키는 다음 노드(직전 데이터를 담은 노드)를 가리킴
	pstack->head = pstack->head->next;
	free(rnode);
	return rdata;
}

Data SPeek(Stack * pstack) {
	if(SIsEmpty(pstack)) {
		printf("Stack Memory Error!");
		exit(-1);
	}

	return pstack->head->data;
}