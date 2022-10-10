#include <stdio.h>
#include <stdlib.h>
#include "DLinkedList.h"

// 초기화 함수의 정의
void ListInit(List * plist) {
	plist->head = (Node*)malloc(sizeof(Node));		// 더미 노드의 생성
	plist->head->next = NULL;
	plist->comp = NULL;
	plist->numOfData = 0;		// 초기화
}

// 더미 노드 연결 리스트 구현: 삽입1
void FInsert(List * plist, LData data) {
	Node * newNode = (Node*)malloc(sizeof(Node));
	newNode->data = data;
	
	newNode->next = plist->head->next;		// plist의 헤드이다
	plist->head->next = newNode;		// 헤드의 next값에 newNode의 주소 값 저장

	(plist->numOfData)++;
}

/*
// 더미 노드 연결 리스트 구현: 삽입2
// 정렬 삽입을 구현하지 않을 경우
void SInsert(List * plist, LData data) {
	Node * newNode = (Node*)malloc(sizeof(Node));	// 새 노드 생성
	newNode->data = data;				// 새 노드에 데이터 저장

	newNode->next = plist->head->next;			// 새 노드가 다른 노드를 가리키게 함
	plist->head->next = newNode;			// 더미 노드가 새 노드를 가리키게 함

	(plist->numOfData)++;				// 저장된 노드의 수를 하나 증가시킴
}
*/

// 더미 노드 연결 리스트 구현: 삽입2
// 정렬 삽입을 구현할 경우
void SInsert(List * plist, LData data) {
	Node * newNode = (Node*)malloc(sizeof(Node));
	Node * pred = plist->head;			// 더미노드의 값으로 초기화(직전 노드의 주소 값을 알기 위함)
	newNode->data = data;               // while 문을 통해서 처음부터 순차적으로 접근하기 위함

	// 새 노드가 들어갈 위치를 찾기 위한 반복문
    // comp가 0을 반환한다는 것은 첫 번째 인자인 data가 정렬 순서상 앞서므로 head에 더 가까워야 한다는 의미이다
	// 0 반환: data > pred->next->data
	// 1 반환: data < pred->next->data
	while(pred->next != NULL && plist->comp(data, pred->next->data) != 0)
		pred = pred->next;	// 다음 노드로 이동
        
	newNode->next = pred->next;
	pred->next = newNode;
	(plist->numOfData)++;
}

void LInsert(List * plist, LData data) {
	if(plist->comp == NULL)				// 정렬기준이 마련되지 않았다면,
		FInsert(plist, data);				// 머리에 노드를 추가
	else						// 정렬기준이 마련되었다면,
		SInsert(plist, data);				// 정렬기준에 근거하여 노드를 추가
}

// 더미 노드 연결 리스트 구현: 참조1
int LFirtst(List * plist, LData * pdata) {
	if(plist->head->next == NULL)		// 더미 노드가 NULL을 가리킨다면,
		return FALSE;			// 변환할 데이터가 없다!
	plist->before = plist->head;			// before는 더미 노드를 가리키게 함 (before의 목적: 노드의 삭제에 필요)
	plist->cur = plist->head->next;		// cur은 첫 번째 노드를 가리키게 함

	*pdata = plist->cur->data;			// 첫 번째 노드의 데이터를 전달
	return TRUE;				// 데이터 반환 성공!
}

// 더미 노드 연결 리스트 구현: 참조2
int LNext(List * plist, LData * pdata) {
	if(plist->cur->next == NULL)			// 더미 노드가 NULL을 가리킨다면,
		return FALSE;				// 반환할 데이터가 없다!

	// 실질적 차이점
	plist->before = plist->cur;				// cur이 가리키던 것을 before가 가리킴
	plist->cur = plist->cur->next;			// cur은 그 다음 노드를 가리킴

	*pdata = plist->cur->data;				// cur이 가리키는 노드의 데이터 전달
	return TRUE;					// 데이터 반환 성공!
}

// 더미 노드 연결 리스트 구현: 삭제1
LData LRemove(List * plist) {
	Node * rpos = plist->cur;		// rpos: cur의 주소 값(삭제하려는 노드)을 백업(주소 값을 이용해 free를 해야하므로)
	LData rdata = rpos->data;		// rdata: cur의 데이터 값(삭제하려는 노드)을 백업

	plist->before->next = plist->cur->next;		// 삭제하려는 노드의 다음 노드의 주소 값을 직전 노드의 next 값으로 저장하면 삭제하려는 노드(직전 노드의 다음 노드)가 아니라 삭제하려는 노드의 다음 노드를 가리키게 된다
	plist->cur = plist->before;				// before는 직전 노드로 옮길 필요 없고(LF또는 LN을 호출하면 자동으로 before는 왼쪽으로 옮겨지게 된다), cur은 삭제한 다음 노드로 옮기면 안된다 (cur은 현재 참조 위치이므로 아직 참조하지 않은 cur의 다음 위치를 가리키면 안되기 때문이다)

	free(rpos);
	(plist->numOfData)--;
	return rdata;
}

// 저장된 데이터의 갯수 참조
int LCount(List * plist) {
	return plist->numOfData;
}

// 정렬기준이 되는 함수를 등록
void SetSortRule(List * plist, int (*comp)(LData d1, LData d2)) {
	plist->comp = comp;			// SetSortRule 함수를 통해 전달된 함수정보 저장을 위한 LinkedList의 멤버 comp
}