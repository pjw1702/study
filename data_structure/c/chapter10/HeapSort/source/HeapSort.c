#include <stdio.h>
#include "..\..\..\chatper09\UsefulHeap\header\UserfulHeap.h"

// 첫 번째 인자의 우선순위가 높으면 양수, 두 번째 인자의 우선순위가 높으면 음수 반환
int PriComp(int n1, int n2) {
  return n2-n1;   // 오름차순 정렬을 위한 문장
  //return n1-n2;
}

// 힙에서 요구하는 정렬 기준 함수와 동일한 기준을 적용한다
void HeapSort(int arr[], int n, PriorityComp pc) {
  Heap heap;
  int i;

  HeapInit(&heap, pc);

  // 정렬 대상을 가지고 힙을 구성한다
  for(i=0; i<n; i++)
    HInsert(&heap, arr[i]);

  // 순서대로 하나씩 꺼내어 정렬을 완성한다
  for(i=0; i<n; i++)
    HDelete(&heap);
}

int main(void) {
  int arr[4] = {1, 2, 3, 4};
  int i;

  HeapSort(arr, sizeof(arr)/sizeof(int), PriComp);

  for(i=0; i<4; i++)
    printf("%d", arr[i]);

  printf("\n");
  return 0;
}