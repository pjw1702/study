#include <stdio.h>
#include "DLinkedList.h"

// 정렬을 구현하는 함수
int WhoIsPrecede(int d1, int d2) {
	if(d1 < d2)
		return 0;		// 정렬 순서상 앞선다
	else
		return 1;		// d2가 정렬 순서상 앞서거나 같다
}

int main(void) {

List list;
int data;

// 리스트 초기화
ListInit(&list);

// 리스트의 멤버 comp를 초기화
// comp에 등록된 정렬기준(정렬 함수 Callback)을 근거로 데이터를 저장
SetSortRule(&list, WhoIsPrecede);

// 데이터 입력
LInsert(&list, 11); LInsert(&list, 11);
LInsert(&list, 22); LInsert(&list, 22);
LInsert(&list, 33);

printf("현재 데이터 수: %d \n", LCount(&list));

// 데이터 참조
if(LFirst(&list, &data)) {
    printf("%d", data);

    while (LNext(&list, &data))
        printf("%d", data);
}
printf("\n\n");

// 저장된 데이터 중 22를 모두 삭제
if(LFirst(&list, &data)) {
    if(data == 22)
        LRemove(&list);
    while (LNext(&list, &data))
    {
        if(data == 22)
            LRemove(&list);
    }

    
}
// 22를 모두 삭제 후 남은 데이터 수
printf("현재 데이터의 수: %d \n", LCount(&list));

// 22를 모두 삭제 후 남은 데이터 참조
if(LFirst(&list, &data)) {
    printf("%d", data);

    while (LNext(&list, &data))
        printf("%d", data);
}
printf("\n\n");
return 0;
}