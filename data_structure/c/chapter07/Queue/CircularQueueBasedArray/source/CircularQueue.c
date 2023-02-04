#include <stdio.h>
#include <stdlib.h>
#include "CircularQueue.h"

void QueueInit(Queue * pq) {
	pq->front = 0;
	pq->rear = 0;
}

int QIsEmpty(Queue * pq) {
	if(pq->front == pq->rear)
		return TRUE;
	else
		return FALSE;
}

// Helper Function

// 다음 데이터를 삽입하기 위한 인덱스 이동 함수
// F나 R을 인자 값으로 전달
// QUE_LEN-1: 배열의 마지막 인덱스
// 데이터가 저장된 위치가 배열의 마지막 요소(QUE_LEN-1)라면 인덱스를 초기화 하기 위한 0을 반환, 아니라면 현재 F나 R의 다음 위치(인덱스)를 반환
int NextPostIdx(int pos) {
	if(pos == QUE_LEN-1)
		return 0;
	else
		return pos+1;

}

void Enqueue(Queue * pq, Data data) {
	// 데이터 삽입을 하는데 있어서, R의 다음 위치가 F의 위치와 같아진다면 데이터가 Full인 상태이므로 더 이상 데이터를 삽입할 수 없다
	if(NextPosIdx(pq->rear) == pq->front) {
		printf("Queue Memory Error!");
		exit(-1);
	}
	pq->rear = NextPosIdx(pq->rear);	// R의 위치 증가
	pq->queArr[pq->rear] = data;	// 데이터 삽입
}

void Dequeue(Queue * pq) {
	// 데이터를 반환하는데 있어서, 데이터가 모두 비워져 있다면 Empty인 상태이므로 더 이상 데이터를 반환할 수 없다
	if(QIsEmpty(pq)) {
		printf("Queue Memorry Error!")
		exit(-1)
	}
	// F가 가리키지 않는다는 사실 만으로도 그 데이터는 유효하지 않는(소멸된) 데이터로 간주되므로 굳이 삭제 코드를 구현해서 삭제를 하거나 0으로 초기화할 필요가 없다
	pq->front = NextPosIdx(pq->front);	// F의 위치 증가
	return pq->queArr[pq->front];	// 현재 F의 위치에 해당하는 데이터 반환
}

// 다음 데이터 추출
Data QPeek(Queue * pq) {
    if(QIsEmpty(pq)) {
        printf("Queue Memory Error!");
        exit(-1);
    }

    return pq->queArr[NextPosIdx(pq->front)];
}