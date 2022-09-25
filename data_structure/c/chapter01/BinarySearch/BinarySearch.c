#include <stdio.h>

int BSearch(int ar[], int len, int target) {
    int first = 0;
    int last = len-1;
    int mid;

    while (first <= last)
    {
        mid = (first+last) / 2;     // Ž�� ����� �߾��� ã�´�
        if (target == ar[mid]) {      // �߾ӿ� ����� ���� Ÿ�� ���̶��
            return mid;             // Ž�� �Ϸ�!
        } else {                    // Ÿ���� �ƴ϶�� Ž�� ����� ������ ���δ�
            /** +1�̳� -1�� ���� ������ first <= mid <= last�� �׻� ������ �Ǿ�, Ž�� ����� �����ϴ� ��� first�� last�� ���� ������ �߻����� �ʴ´�! */
            if (target < ar[mid]) {
                last = mid-1;       // -1�� �� ����: Ÿ���� �߾��� �� ���� �����Ƿ� �������� ����(Ž���� �������� ������ �ε����� ó�� �߾� �ε������� �۰� ����)
            } else {
                first = mid+1;      // +1�� �� ����: Ÿ���� �߾��� �� ���� ũ�Ƿ� ������ ����(Ž���� �������� ù ��° �ε����� ó�� �߾� �ε������� ũ�� ����)
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
        printf("Ž�� ����!");
    else
        printf("Ÿ�� ���� �ε���: %d \n", idx);

    idx = BSearch(arr, sizeof(arr)/sizeof(int), 4);
    if(idx == -1)
        printf("Ž�� ����!");
    else
        printf("Ÿ�� ���� �ε���: %d \n", idx);

    return 0;

} 