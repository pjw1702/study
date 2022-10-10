#ifndef __ARRAY_LIST_H__        // 매크로: 헤더파일의 중복 포함을 막기 위한 선언
#define __ARRAY_LIST_H__

#define TRUE    1       // 'True'를 표현하기 위한 매크로 정의
#define FALSE   0       // 'False'를 표현하기 위한 매크로 정의

#define LIST_LEN    100     
typedef int LData;      // LData에 대한 typedef 선언: LDdata 자료형에 대한 유연한 자료형 변환을 위함(여기 한 번만 자료형을 바꾸면 데이터의 자료형에 바뀐 자료형을 전체적으로 적용 가능): 저장할 대상의 자료형을 변경하기 위한 typedef 선언

typedef struct __ArrayList      // 배열기반 리스트를 정의한 구조체
{
    LData arr[LIST_LEN];        // 리스의 저장소인 배열
    int numOfData;              // 저장된 데이터의 수
    int curPosition;            // 데이터 참조위치를 기록
} ArrayList;                    // 배열 기반 리스트를 의미하는 구조체

typedef ArrayList List;         // 리스트 변경을 용이하게 하기 위한 typedef 선언

void ListInit(List * plist);                     // 초기화
void LInsert(List * plist, LData data);          // 데이터 저장

int LFirst(List * plist, LData * pdata);        // 첫 번째 데이터 참조
int LNext(List * plist, LData * pdata);        // 두 번째 이후 데이터 참조

LData LRemove(List * plist);                    // 참조한 데이터 삭제
int LCount(List * plist);                       // 저장된 데이터의 수 반환

#endif
