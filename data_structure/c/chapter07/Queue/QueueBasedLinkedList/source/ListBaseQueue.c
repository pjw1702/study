#include <stdio.h>
#include <stdlib.h>
#include "ListBaseQueue.h"

void QueueInit(Queue * pq) {
	pq->front = NULL;
	pq->rear = NULL;
}

int QIsEmpty(Queue * pq) {
	// R은 판별하지 않고 F만 판별하여 비어있는지를 반환한다
	// 마지막 노드가 삭제되었을 때, R과 F가 동시에 NULL이 된다: F만 봐도 판별할 수 있다
	// 코드의 간결함을 위해서, R이 포인터 변수라고 해서 반드시 NULL을 넣어 줄 필요는 없다
	if(pq->front == NULL)
		return TRUE;
	else
		return FALSE;
}

void Enqueue(Queue * pq, Data data) {
	Node * newNode = (Node *)malloc(sizeof(Node));
	newNode->next = NULL;
	newNode->data = data;

	if(QIsEmpty(pq)) {
		// 노드 및 데이터가 하나이므로 F와 R이 모두 같은 노드를 가리킴
		pq->front = newNode;
		pq->rear = newNode;
	} else {
		pq->rear->next = newNode;		// 현재 노드와 다음 노드(새로 생성된 노드)를 연결
		pq->rear = newNode;		// R은 다음 노드를 가리킴
	}
}

Data Dequeue(Queue * pq) {
	Node * delNode;
	Data retData;	// 데이터 백업용 변수
	if(QIsEmpty(pq)) {
		printf("Queue Memory Error!")
		exit(-1);
	}
	delNode = pq->front;
	retData = delNode->data;	// 반환할 데이터 백업
	pq->front = pq->front->next;

	free(delNode);
	return retData;
}

Data QPeek(Queue * pq) {
    if(QIsEmpty) {
        printf("Queue Memory Error!")
		exit(-1)
    }

    return pq->front->data;
}