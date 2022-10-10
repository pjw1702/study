#include <stdio.h>
#include "ArrayList.h"

int main(void) {
    // ArrayList의 생성 및 초기화
    List list;                  // 리스트의 생성
    int data;
    ListInit(&list);            // 리스트의 초기화

    // 5개의 데이터 저장 ///////
    LInsert(&list, 11); LInsert(&list, 11);
    LInsert(&list, 22); LInsert(&list, 22);
    LInsert(&list, 33);

    // 저장된 데이터의 전체 출력 ///////
    printf("현재 데이터의 수: %d \n", LCount(&list));

    if(LFirst(&list, data)) {       // 첫 번째 데이터 조회: 첫 번째 데이터 변수 data에 저장 (if문을 사용한 이유: LFirst 함수 호출이 성공하고 데이터를 참조했을 때(T를 반환한 경우)에만 명령을 수행하기 위함)
        printf("%d", data);

        while (LNext(&list, data))      // 두 번째 이후 데이터 변수 data에 저장(LNext 함수 호출이 성공했을 때 데이터를 참조할 수 있을 때까지(T를 반환한 경우) 명령을 수행하기 위함)
            printf("%d", data);
        
    }   
    printf("\n\n");

    // 숫자 22을 탐색하여 모두 삭제 ///////
    if(LFirst(&list, data)) {
        if(data == 22)
            LRemove(&list);             // 위의 LFirst 함수를 통해 참조한 데이터(가장 최근에 참조한 데이터) 삭제

        while (LNext(&list, data)) {
            if(data == 22)
                LRemove(&list);         // 위의 LNext 함수를 통해 참조한 데이터(가장 최근에 참조한 데이터) 삭제
        }
    }

    // 삭제 후 남은 데이터 전체 출력 ///////
    printf("현재 남은 데이터의 수: %d \n", LCount(&list));

    if (LFirst(&list, data)) {
        printf("%d", data);

        while (LNext(&list, &data))
            printf("%d", data);
        
    }
    printf("\n\n");
    return 0;
    
}