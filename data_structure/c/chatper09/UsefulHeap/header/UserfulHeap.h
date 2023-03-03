#ifndef __USEFUL_HEAP_H__
#define __USEFUL_HEAP_H__

#define TRUE        1
#define FALSE       0

#define HEAP_LEN    100

typedef char HData;
typedef int PriorityComp(HData d1, HData d2);

// 함수 포인터를 선언할 경우 함수 명 앞에 *를 붙여야 하는 것이 원칙 적이지만, 구조체의 멤버로 사용되는 경우 멤버에 *를 붙이면 typedef 선언에서 *을 안 붙일 수 있다
// 반면에 멤버에 *을 안 붙일 경우엔 typedef 선언에서 *을 붙여야 한다

typedef struct _heap {
  PriorityComp * comp;    // typedef int (* PriorityComp)(HData d1, HData d2); 로 선언하고 PriorityComp comp; 로 멤버로 선언한 것과 같다
	int numOfData;
	HData heapArr[HEAP_LEN];
} Heap;

void HeapInit(Heap * ph, PriorityComp pc);
int HIsEmpty(Heap * ph);

void HInsert(Heap * ph, HData data);
HData HDelete(Heap * ph);		

#endif