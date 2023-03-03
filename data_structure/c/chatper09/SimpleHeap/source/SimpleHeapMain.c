#include <stdio.h>
#include "..\header\SimpleHeap.h"

// 아래와 같은 경우에는 우선순위에 대한 정보와 데이터를 함께 전달한다
// 데이터를 저장할 때 우선순위 정보를 별도로 전달하는 것은 적합하지 않은 경우가 많다. 일반적으로 데이터의 우선순위는 데이터를 근거로 판단이 이루어지기 때문이다
int main(void) {
	Heap heap;
	HeapInit(&heap);

	HInsert(&heap, 'A', 1);	// 문자 'A'를 최고의 우선순위로 저장
	HInsert(&heap, 'B', 2);	// 문자 'B'를 두 번째 우선순위로 저장
	HInsert(&heap, 'C', 3);	// 문자 'C'를 세 번째 우선순위로 저장
	printf("%c \n", HDelete(&heap));

	HInsert(&heap, 'A', 1);	// 문자 'A'를 한 번 더 저장
	HInsert(&heap, 'B', 2);	// 문자 'B'를 한 번 더 저장
	HInsert(&heap, 'C', 3);	// 문자 'C'를 한 번 더 저장
	printf("%c \n", HDelete(&heap));

	while(!HIsEmpty(&heap))
		printf("%c \n", HDelete(&heap));

	return 0;

  // 실행 결과: A A B B C C
}