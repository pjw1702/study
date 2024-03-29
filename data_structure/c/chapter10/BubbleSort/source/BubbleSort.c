#include <stdio.h>

void BubbleSort(int arr[], int n) {
  int i, j;
  int temp;

  // 버블 정렬의 실질적 구현 코드
  for (i=0; i < n-1; i++) {
    for (j=0; j<(n-i)-1; j++) {
      if (arr[j] > arr[j+1]) {
        // 데이터 교환
        temp = arr[j];
        arr[j] = arr[j+1];
        arr[j+1] = temp;
      }
    }
  }
}

int main(void) {
  int arr[4] = {3, 2, 4, 1};
  int i;

  BubbleSort(arr, sizeof(arr)/sizeof(int));

  for (i=0; i<4; i++)
    printf("%d", arr[i]);

  printf("\n");
  return 0;
}

// 실행결과: 1 2 3 4