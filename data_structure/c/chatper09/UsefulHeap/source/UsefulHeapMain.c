#include <stdio.h>
#include "..\header\UserfulHeap.h"

int DataPriorityComp(char ch1, char ch2) {
  return ch2-ch1;
 // return ch1-ch2;   // 아스키 코드 값이 작은 문자의 우선순위가 높으므로('B'보다 'A'의 값이 더 큼) ch2-ch1의 연산이 이루어져야 양수가 반환된다
}

int main(void) {
	Heap heap;
	HeapInit(&heap, DataPriorityComp);

	HInsert(&heap, 'A');	
	HInsert(&heap, 'B');	
	HInsert(&heap, 'C');	
	printf("%c \n", HDelete(&heap));

	HInsert(&heap, 'A');	
	HInsert(&heap, 'B');	
	HInsert(&heap, 'C');	
	printf("%c \n", HDelete(&heap));

	while(!HIsEmpty(&heap))
		printf("%c \n", HDelete(&heap));

	return 0;

  // 실행 결과: A A B B C C
}