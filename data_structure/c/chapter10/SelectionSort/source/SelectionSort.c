#include <stdio.h>

void SelSort(int arr[], int n) {
  int i, j;
  int curIdx; // 현재 인덱스: swap의 기준이 되는 인덱스 (최솟값을 가지는 인덱스를 찾기 위한 인덱스)
  int temp;

  for (i=0; i<n-1; i++) {
    curIdx = i; 

    for (j=i+1; j<n; j++) {     // 최솟값 탐색
      if(arr[j] < arr[curIdx])   // 현재 위치한 인덱스의 데이터보다 작은 데이터를 가지는 인덱스만 탐색한다
        curIdx = j;
    }
    // 교환
    temp = arr[i];
    arr[i] = arr[curIdx];
    arr[curIdx] = temp;
  }
}

int main(void) {
  int arr[4] = {3, 4, 2, 1};
  int i;
  
  SelSort(arr, sizeof(arr)/sizeof(int));

  for(i=0; i<4; i++)
    printf("%d", arr[i]);

  printf("\n");
  return 0;
}

// 실행결과: 1 2 3 4