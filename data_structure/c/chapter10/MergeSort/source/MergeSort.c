#include <stdio.h>
#include <stdlib.h>

void MergeTwoArea(int arr[], int left, int mid, int right) {
  int fIdx = left;
  int rIdx = mid+1;
  int i;

  int * sortArr = (int *)malloc(sizeof(int)*(right+1));   // 병합 결과를 담을 메모리 공간 할당(배열 내에서 데이터를 직접 옮기기가 번거로우므로 임시 메모리 공간을 할당하여 정렬 후 배열에 데이터 저장)
  int sIdx = left;

  while (fIdx <= mid && rIdx <= right) {    // 병합할 두 영역의 데이터를 비교하여 배열 sortArr에 담는다
    if(arr[fIdx] <= arr[rIdx])
      sortArr[sIdx] = arr[fIdx++];        // fIdx가 rIdx보다 작거나 같으면 fIdx의 데이터를 담는다
    else
      sortArr[sIdx] = arr[rIdx++];        // fIdx가 rIdx보다 크면 rIdx의 데이터를 담는다
    sIdx++;
  }

  if(fIdx > mid) {      // 배열의 앞 부분이 sortArr로 모두 옮겨졌다면 배열의 뒷부분에 남은 모든 데이터들을 sortArr에 그대로 옮긴다
    for(i=rIdx; i <= right; i++, sIdx++)
      sortArr[sIdx] = arr[i];
  } else {              // 배열의 뒷 부분이 sortArr로 모두 옮겨졌다면 배열의 앞부분에 남은 모든 데이터들을 sortArr에 그대로 옮긴다
    for(i=fIdx; i <= mid; i++, sIdx++)
      sortArr[sIdx] = arr[i];
  }

  for(i=left; i <= right; i++)    // 병합 결과를 옮겨 담는다
    arr[i] = sortArr[i];

  free(sortArr);
}

// left는 분할을 할 시작 인덱스, right는 분할을 할 끝 인덱스를 인자로서 전달한다
void MergeSort(int arr[], int left, int right) {
  int mid;

  // if문의 조건이 참이되지 않을 시 함수 종료
  if(left < right) {      // left가 더 작다는 것은 더 나눌 수 있다는 뜻이다 (=데이터가 하나만 남았을 경우)
    // 중간 지점을 계산한다
    mid = (left+right) / 2;

    // MergeSort 함수는 둘로 나눌 수 없을 때까지 재귀적으로 호출된다

    // 둘로 나눠서 각각을 정렬한다
    MergeSort(arr, left, mid);    // left~mid에 위치한 데이터 정렬(왼쪽 분할)
    MergeSort(arr, mid+1, right); // mid+1 ~ right에 위치한 데이터 정렬(오른쪽 분할)

    // 정렬된 두 배열을 병합한다(병합을 함과 동시에 정렬 진행)
    MergeTwoArea(arr, left, mid, right);
  }
}