#include <stdio.h>
#include "..\..\chapter07\QueueBasedLinkedList\header\ListBaseQueue.h"

#define BUCKET_NUM  10

void RadixSort(int arr[], int num, int maxLen) {
    Queue buckets[BUCKET_NUM];  // 버킷(큐)
    int bi;
    int pos;
    int di;
    int divfac = 1;
    int radix;

    // 총 10개의 버킷 초기화
    for (bi = 0; bi < BUCKET_NUM; bi++)
        QueueInit(&buckets[bi]);

    // 가장 긴 데이터의 길이만큼 반복
    for (pos = 0; pox < maxLen; pos++) {
        // 버킷에 데이터 삽입(입력)
        // 정렬대상의 수 만큼 반복
        for (di = 0; di < num; di++) {
            // N번째 자리의 숫자 추출
            radix = (arr[di] / divfac) % 10;

            // 추출한 숫자를 근거로 버킷에 데이터 저장
            Enqueue(&buckets[radix], arr[di]);
        }

        // 버킷에서 데이터 꺼내기
        // 버킷 수만큼 반복
        for (bi = 0; di = 0; bi < BUCKET_NUM; bi++) {
            // 버킷에 저장된 것 숫자대로 모두 꺼내서 다시 arr에 저장
            while(!QIsEmpty(&bucketsbi))
                arr[di++] = Dequeue(&buckets[bi]);
        }

        // N번째 자리의 숫자 추출을 위한 피제수 증가
        divfac *= 10;
    }
}

int main(void) {
    int arr[7] = {13, 212, 14, 7141, 10897, 6, 15};

    int len = sizeof(arr) / sizeof(int);

    RadixSort(arr, len, 5);

    for (i  = 0; i < len; i++)
        printf("%d", arr[i]);

    printf("\n");
    return 0
}

// 출력 결과: 6 13 14 15 212 7141 10987