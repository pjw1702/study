#include <stdio.h>
#include <stdlib.h>
#include "ArrayList.h"
#include "Point.h"

int main(void) {
	List list;
	Point compPos;
	Point * ppos;

	ListInit(&list);

	/*** 4개의 데이터 저장 ***/
	ppos = (Point*)malloc(sizeof(Point));
	SetPointPos(ppos, 2,1);
	LInsert(&list, ppos);

	ppos = (Point*)malloc(sizeof(Point));
	SetPointPos(ppos, 2,1);
	LInsert(&list, ppos);

	ppos = (Point*)malloc(sizeof(Point));
	SetPointPos(ppos, 3,1);
	LInsert(&list, ppos);

	ppos = (Point*)malloc(sizeof(Point));
	SetPointPos(ppos, 3,2);
	LInsert(&list, ppos);

	/*** 저장된 데이터의 출력 ***/
	printf("현재 데이터의 수: %d \n", LCount(&list));

	if(LFirst(&list, &ppos)) {
		ShowPointPos(ppos);

		while(LNext(&list, &ppos))
			showPointPos(ppos);
	}
	printf("\n");

	/*** xpos가 2인 모든 데이터 삭제 ***/
	compPos.xpos = 2;
	compPos.ypos = 0;		// 실제로 ypos 값이 0으로 초기화 한 건 없다. 그냥 xpos가 2인 것에만 초점을 두고 ypos는 따로 신경쓰지 않는다는 뜻에서 0을 저장한다

	if(LFirst(&list, &ppos)) {
		if(PointComp(ppos, &compPos)==1) {
			ppos=LRemove(&list);		// 동적할당된 메모리(공간)까지 소멸되는 것은 아니다 (리스트 자료구조의 책임 범위는 주소 값 참조와 값(또는 주소 값)의 저장과 삭제 까지일 뿐이다)
			free(ppos);			// 리스트 자료구조의 책임범위가 아니므로 직접 메모리 공간을 소멸시킨다
		}
	}
	while(LNext(&list, &ppos)) {
		if(PointComp(ppos, &compPos)==1) {
			ppos=LRemove(&list);
			free(ppos);
		}
	}

	/*** 삭제 후 남은 데이터 전체 출력 ***/
	printf("현재 데이터의 수: %d \n", LCount(&list));

	if(LFirst(&list, &ppos)) {
		showPointPos(ppos);
		
		while(LNext(&list, &ppos))
			ShowPointPos(ppos);
	}
	printf("\n");

	return 0;
}

