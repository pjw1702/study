#include <stdio.h>

typedef int Key;
typedef double Data;

typedef struct item {
    Key searchKey;
    Data searchData;
} Item;

int ISearch(int ar[], int low, int high, int target) {
    // 이진 탐색과 보간 탐색의 차이점 1: 탐색대상의 인덱스 값 수정
    // 이진 탐색: mid = (first + last) / 2
    // 보간 탐색: mid = ((double)(target-ar[last])/ar[first]-ar[last] * (last-first)) + first
    // int mid;
    // mid = ((double)(target-ar[last])/(ar[first]-ar[last]) * (last-first)) + first;

    Item item;

    item.searchKey = ((double)(target-ar[high])/(ar[high]-ar[low]) * (high-low)) + high;
    item.searchData = target;

    // 이진 탐색과 보간 탐색의 차이점 2: 탈출조건 수정
    // 아래 재귀함수 호출과정에서, else문이 무한 반복된다
    // if(high > low)
    //     return -1;
    // 탐색대상이 존재하지 않는 경우, 탐색대상의 값은 탐색 범위를 넘어선다
    if(ar[low]>item.searchData || ar[high]<item.searchData)
        return -1;

   if(ar[item.searchKey] == item.searchData)
        return item.searchKey;   // 탐색된 타겟의 키(인덱스 값) 반환
    else if(item.searchData < ar[item.searchKey])
        return ISearch(ar, low, item.searchKey-1, item.searchData);
    else
        return ISearch(ar, item.searchKey+1, high, item.searchData);
}

int main(void) {
    int arr[] = {1, 3, 5, 7, 9};
    int idx;

    // 대상 배열, 배열의 low, high, target을 인수로 전달
    // 탐색할 데이터: 7
    idx = ISearch(arr, 0, sizeof(arr)/sizeof(int)-1, 7);
    if(idx == -1)
        printf("탐색 실패 \n");
    else
        printf("타겟 저장 인덱스: %d \n", idx);

    // 탐색할 데이터: 10
    idx = ISearch(arr, 0, sizeof(arr)/sizeof(int)-1, 10);
    if(idx == -1)
        printf("탐색 실패 \n");
    else
        printf("타겟 저장 인덱스: %d \n", idx);

    return 0;
}

// 실행결과
// 타겟 저장 인덱스: 3
// 탐색 실패