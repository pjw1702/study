#ifndef __SIMPLE_HEAP_H__
#define __SIMPLE_HEAP_H__

#define TRUE        1
#define FALSE       0

#define HEAP_LEN    100

typedef char HData;
typedef int Priority;

// 데이터와 우선순위 정보를 각각 구분
typedef struct _heapElem {
	Priority pr;	// 값이 작을수록 높은 우선순위
	HData data;	
} HeapElem;

typedef struct _heap {
	int numOfData;
	HeapElem heapArr[HEAP_LEN];
} Heap;

void HeapInit(Heap * ph);
int HIsEmpty(Heap * ph);

void HInsert(Heap * ph, HData data, Priority pr);
HData HDelete(Heap * ph);		// 우선순위가 가장 높은 데이터가 삭제되도록 정의

#endif