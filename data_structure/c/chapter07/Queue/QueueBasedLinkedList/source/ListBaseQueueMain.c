#include <stdio.h>
#include "ListBaseQueue.h"

int main(void) {
    // Queue의 생성 및 초기화
    Queue q;
    QueueInit(&q);

    // 데이터 삽입
    Enqueue(&q, 1);     Enqueue(&q, 2);
    Enqueue(&q, 3);     Enqueue(&q, 4);
    Enqueue(&q, 5);

    // 데이터 꺼내기
    while (!QIsEmpty(&q))
        printf("%d", Dequeue(&q));

    return 0;
}

// 실행결과: 1 2 3 4 5
