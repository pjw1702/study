#ifndef __D_LINKED_LIST_H__        // 매크로: 헤더파일의 중복 포함을 막기 위한 선언
#define __D_LINKED_LIST_H__

#define TRUE    1       // 'True'를 표현하기 위한 매크로 정의
#define FALSE   0       // 'False'를 표현하기 위한 매크로 정의

typedef int LData;      // LData에 대한 typedef 선언: LDdata 자료형에 대한 유연한 자료형 변환을 위함(여기 한 번만 자료형을 바꾸면 데이터의 자료형에 바뀐 자료형을 전체적으로 적용 가능): 저장할 대상의 자료형을 변경하기 위한 typedef 선언

typedef struct _node {
	LData data;
	struct _node * next;
} Node;

typedef struct _linkedList {
	Node * head;				            // 더미 노드를 가리키는 멤버
	Node * cur;				                // 참조 및 삭제를 돕는 멤버
	Node * before;				            // 삭제를 돕는 멤버
	int numOfData;				            // 저장된 데이터의 수를 기록하기 위한 멤버
	int (*comp)(LData d1, LData d2);		// 정렬의 기준을 등록하기 위한 멤버 (comp: 정렬 기준 기능을 설치하기 위한 함수 포인터 변수)
} LinkedList;

typedef LinkedList List;

void ListInit(List * plist);                     // 초기화
void LInsert(List * plist, LData data);          // 데이터 저장

int LFirst(List * plist, LData * pdata);        // 첫 번째 데이터 참조
int LNext(List * plist, LData * pdata);        // 두 번째 이후 데이터 참조

LData LRemove(List * plist);                    // 참조한 데이터 삭제
int LCount(List * plist);                       // 저장된 데이터의 수 반환

void SetSortRule(List * plist, int(*comp)(LData d1, LData d2));

#endif