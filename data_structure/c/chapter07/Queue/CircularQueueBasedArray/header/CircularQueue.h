#include <stdio.h>
#ifndef __C_QUEUE_H
#define __C_QUEUE_H

#define TRUE    1
#define FALSE   0

#define QUE_LEN 100
typedef int Data;

// R과 F는 배열의 인덱스 값이 된다
// R을 이동시키면서 저장, F를 이동시키면서 반환
// index=0으로 이동시킬 때, 인덱스 값을 0으로 초기화 해주어야 한다
typedef struct _cQueue
{
    int front;
    int rear;
    Data queueArr[QUE_LEN];
} CQueue;

typedef CQueue Queue;

void QueueInit(Queue * pq);
int QIsEmpty(Queue * pq);

void Enqueue(Queue * pq, Data data);
Data Dequeue(Queue * pq);
Data QPeek(Queue * pq);


#endif __C_QUEUE_H