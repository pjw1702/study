#include <stdio.h>

void InsertSort(int arr[], int n) {
  int i, j;
  int insData;

  for(i=1; i<n; i++) {
    insData = arr[i];   // 정렬 대상 저장

    for(j=i-1; j>=0; j--) {
      if(arr[j] > insData)
        arr[j+1] = arr[j];    // 비교 대상의 데이터를 한 칸 뒤로 밀기
      else
        break;              // 삽입 위치를 찾았으니 탈출
    }

    arr[j+1] = insData;   // 찾은 위치에 정렬대상 삽입
  }
}

int main(void) {
  int arr[5] = {5, 3, 2, 4, 1};
  int i;
  int k;

  InsertSort(arr, sizeof(arr)/sizeof(int));

  arr[-1];

  for(i=0; i<5; i++)
    printf("%d", arr[i]);

  printf("\n");

  return 0;
}

// 실행결과: 1 2 3 4 5