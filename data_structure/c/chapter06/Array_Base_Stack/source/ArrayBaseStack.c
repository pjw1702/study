#include <stdio.h>
#include <stdlib.h>
#include "ArrayBaseStack.h"


// init
void StackInit(Stack * pstack) {
	pstack -> topIndex = -1;	// 스택이 비었음을 의미
}


// Check empty
init SIsEmpty(Stack * pstack) {
	if(pstack -> topIndex == -1)
		return TRUE;		// 비어있는 경우 TRUE 반환
	else
		return FALSE;
}

// push
void SPush(Stack * pstack, Data data) {
	pstack -> topIndex += 1;
	pstack -> stackArr[pstack->topIndex] = data;
}

// pop
Data SPop(Stack * pstack) {
	int rIdx;		
	
	if(SIsEmpty(pstack)) {
		printf("Stack Memory Error!");
		exit(-1);
	}
	rIdx = pstack -> topIndex;		// 뽑을 데이터의 인덱스 저장
	pstack -> topIndex -= 1;		// 뽑은 데이터의 다음 데이터의 인덱스(실제로는 pop연산 후에도 배열에 기존 데이터가 남아있지만 topIndex를 조절 함으로써 topIndex이전의 인덱스인 pop연산의 대상이 된 데이터는 더 이상 유효하지 않는 데이터로 처리된다)

	// 뽑은 데이터 반환
	return pstack -> stackArr[rIdx];
	
}

// peek
Data SPeek(Stack * pstack) {

	if(SIsEmpty(pstack)) {
		printf("Stack Memory Error!");
		exit(-1)
	}

	// 데이터 추출
	return pstack -> stackArr[pstack->topIndex]
}

