#include <stdio.h>

int BSearchRecur(int ar[], int first, int last, int target) {
    int mid;
    if(first > last)
        return -1;

    mid = (first + last) / 2;

    if(ar[mid] == target)   // 찾을 값을 찾은 경우
        return mid;
    else if(target < ar[mid])
        return BSearchRecur(ar, first, mid-1, target);  // 찾을 값이 중앙 값보다 작은 경우
    else
        return BSearchRecur(ar, mid+1, last, target);   // 찾을 값이 중앙 값보다 큰 경우
}

int main(void) {
    // ...
    return 0;
}

