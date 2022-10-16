#include <stdio.h>
#include <stdlib.h>
#include "DBLinkedList.h"

// 리스트 초기화
void ListInit(List * plist) {
	plist->head = NULL;
	plist->numOfData = 0;
}

// 데이터 및 노드 삽입
void LInsert(List * plist, Data data) {
	Node * newNode = (Node*)malloc(sizeof(Node));
	newNode->data = data;

	/* 양 방향 연결 */
	// 첫 번째 노드 삽입
	newNode->next = plist->head;		// 새로운 노드의 next가 head가 가리키는 노드를 가리킴
	// 두 번째 이후 노드 삽입
	if(plist->head != NULL)
		plist->head->prev = newNode;	// head가 가리키는 노드의 prev가 새로운 노드를 가리킴

	newNode->prev = NULL;
	plist->head = newNode;

	(plist->numOfData)++;
}

// 데이터 조회
int LFirst(List * plist, Data * pdata) {
	if(plist->head == NULL)
		return FALSE;

	plist->cur = plist->head;

	*pdata = plist->cur->data;
	return TRUE;
}

// LFirst 함수와 LNext함수는 사실상 단방향 연결 리스트의 경우가 차이가 없다
// LPrevious 함수는 LNext 함수와 이동 방향에서만 차이가 난다

// 정방향 조회
int LNext(List * plist, Data * pdata) {
	if(plist->cur->next == NULL)
		return FALSE;

	plist->cur = plist->cur->next;

	*pdata = plist->cur->data;

	return TRUE;
}

// 역방향 조회
intLPrevious(List * plist, Data * pdata) {
	if(plist->cur->prev = NULL)
		return FALSE;

	plist->cur = plist->cur->prev;

	* pdata = plist->cur->data;

	return TRUE;
}

// 현재 저장된 데이터 갯수
int LCount(List * plist) {
    return plist->numOfData;
}