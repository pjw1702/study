#include <stdio.h>

int BSearch(int ar[], int len, int target) {
    int first = 0;
    int last = len-1;
    int mid;

    while (first <= last)
    {
        mid = (first+last) / 2;     // 탐색 대상의 중앙을 찾는다
        if (target == ar[mid]) {      // 중앙에 저장된 값이 타겟 값이라면
            return mid;             // 탐색 완료!
        } else {                    // 타겟이 아니라면 탐색 대상을 반으로 줄인다
            /** +1이나 -1을 하지 않으면 first <= mid <= last가 항상 성립이 되어, 탐색 대상이 존재하는 경우 first와 last의 역전 현상이 발생하지 않는다! */
            if (target < ar[mid]) {
                last = mid-1;       // -1을 한 이유: 타겟이 중앙의 값 보다 작으므로 오른쪽을 버림(탐색할 데이터의 마지막 인덱스를 처음 중앙 인덱스보다 작게 설정)
            } else {
                first = mid+1;      // +1을 한 이유: 타겟이 중앙의 값 보다 크므로 왼쪽을 버림(탐색할 데이터의 첫 번째 인덱스를 처음 중앙 인덱스보다 크게 설정)
            }

        }
        
    }

    return -1; 
    
}

int main(void) {
    int arr[] = {1, 3, 5, 7, 9};
    int idx;

    idx = BSearch(arr, sizeof(arr)/sizeof(int), 7);
    if(idx == -1)
        printf("탐색 실패!");
    else
        printf("타겟 저장 인덱스: %d \n", idx);

    idx = BSearch(arr, sizeof(arr)/sizeof(int), 4);
    if(idx == -1)
        printf("탐색 실패!");
    else
        printf("타겟 저장 인덱스: %d \n", idx);

    return 0;

} 