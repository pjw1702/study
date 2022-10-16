#include <stdio.h>
#include <stdlib.h>
#include "CLinkedList.h"

// 리스트 초기화
void ListInit(List * plist) {
    plist->tail = NULL;
    plist->cur = NULL;
    plist->before = NULL;
    plist->numOfData = 0;
}

// 데이터 및 노드 삽입

// 머리 추가 부분
// 꼬리가 새 노드를 가리키면 머리가 된다
// 새 노드는 꼬리를 가리키면 원형 연결 형태가 된다
void LInsertFront(List * plist, Data data) {
    Node * newNode = (Node*)malloc(sizeof(Node));
    newNode->data = data;

    if(plist->tail == NULL) {
		plist->tail = newNode;
		newNode->next = newNode;	// 자기 자신을 가리킴
	} else {
		newNode->next = plist->tail->next;		// 새 노드의 오른쪽(next) 연결: 원형 구조를 가지므로 tail의 next 부분이 연결할 노드의 주소 값을 가지고 있음
		plist->tail->next = newNode;		// 새 노드의 왼쪽(tail) 연결
	}
	(plist->numOfData)++;
}

// 꼬리 추가 부분
// tail의 위치가 유일한 차이점이다
void LInsert(List * plist, Data data) {
	Node * newNode  = (Node*)malloc(sizeof(Node));
	newNode->data = data;

	if(plist->tail == NULL) {
		plist->tail = newNode;
		newNode->next = newNode;	// 자기 자신을 가리킴
	} else {
		newNode->next = plist->tail->next;		// 새 노드의 오른쪽(next) 연결: 원형 구조를 가지므로 tail의 next 부분이 연결할 노드의 주소 값을 가지고 있음
		plist->tail->next = newNode;		// 새 노드의 왼쪽(tail) 연결
		plist->tail = newNode;			// LInsertFront 함수와의 유일한 차이점: tail이 새 노드를 가리키도록 하는 코드
	}
	(plist->numOfData)++;
}

// 데이터 조회
int LFirst(List * plist, Data * pdata) {
	if(plist->tail == NULL)
		return FALSE;

	// before는 tail이 가리키는 노드를, cur은 머리(tail이 가리키는 다음 노드) 노드로 초기화 한다
	plist->before = plist->tail;			// before는 cur의 직전 노드 이므로 원형 연결 리스트에서는 머리와 연결된 꼬리이다
	plist->cur = plist->tail->next;		// cur이 가리키는 노드: 머리

	*pdata = plist->cur->data;
	return TRUE;
}

// 원형 연결 리스트이므로 리스트의 끝을 검사하는 코드가 없다
int LNext(List * plist, Data * pdata) {
	if(plist->tail == NULL)
		return FALSE;

	// before과 cur은 한 칸씩 앞으로 전진한다
	plist->before = plist->cur;
	plist->cur = plist->cur->next;

	// 한 칸씩 앞으로 전진하면서 cur이 가리키는 노드의 data를 조회한다
	*pdata = plist->cur->data;
	return TRUE;
}

// 데이터 및 노드 삭제
Data LRemove(List * plist) {
	Node * rpos = plist->cur;
	Data rdata = rpos->data;

	// 삭제할 노드가 tail이 가리키는 노드라면
	if(rpos == plist->tail) {
		// tail이 가리키는 노드인데 마지막 노드이기 까지 한다면
		if(plist->tail == plist->tail->next) 
			plist->tail = NULL;
		else
			plist->tail = plist->before;
	}

	// 삭제할 노드의 다음 노드와 연결
	plist->before->next = plist->cur->next;
	plist->cur = plist->before;

	free(rpos);
	(plist->numOfData)--;
	return rdata;
}

int LCount(List * plist) {
    return plist->numOfData;
}