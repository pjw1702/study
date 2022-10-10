#include <stdio.h>
#include "ArrayList.h"

// 배열 기반 리스트의 초기화
// 실제로 초기화할 대상은 구조체 변수의 멤버이다 따라서 초기화 함수의 구성은 구조체 정의를 기반으로 한다
void ListInit(List * plist) {
	(plist -> numOfData) = 0;		// 저장된 데이터 수
	(plist -> curPosition) = -1;		// 아무런 위치도 참조하지 않았음을 의미(초기 상태이므로)
}

// 배열 기반 리스트의 삽입
void LInsert(List * plist, LData data) {
	if(plist -> numOfData > LIST_LEN) {
		puts("저장이 불가능합니다.");
	}

	plist -> arr[plist->numOfData] = data;		// 데이터 저장
	(plist -> numOfData)++;			// 저장된 데이터 수 증가
} 

//배열 기반 리스트의 조회
int LFirst(List * plist, LData * pdata) {
	if(plist->numOfData == 0)		// 저장된 데이터가 하나도 없을 경우(초기 상태)
		return FALSE;

	(plist->curPosition) = 0;		// 참조 위치 초기화! 첫 번째 데이터의 참조를 의미!
	*pdata = plist->arr[0];		// pdata가 가리키는 공간(data)에 데이터 저장
	
	return TRUE;
}

int LNext(List * plist, LData * pdata) {
	if(plist->curPosition >= (plist->numOfData)-1)		// 더 이상 참조할 데이터가 없다면
		return FALSE;

	(plist->curPosition)++;		// 참조 위치 초기화! 첫 번째 데이터의 참조를 의미!
	*pdata = plist->arr[plist->curPosition];		// pdata가 가리키는 공간(data)에 데이터 저장
	
	return TRUE;
}

// 값의 반환은 매개변수를 통해서! 함수의 반환은 성공 여부를 알리기 위해서!

// 배열 기반 리스트의 삭제
LData LRemove(List * plist) {
	int rpos = plist->curPosition;	// 삭제할 데이터의 인덱스 값 참조
	int num = plist->numOfData;
	int i;
	LData rdata = plist->arr[rpos];	// 삭제할 데이터를 임시로 저장

	// 삭제를 위한 데이터의 이동을 진행하는 반복문
	for(i=rpos; i<num-1; i++)
		plist->arr[i] = plist->arr[i+1];

	(plist->numOfData)--;		// 데이터의 수 감소
	(plist->curPosition)--;		// 참조위치를 하나 되돌린다
	return rdata;			// 삭제된 데이터의 반환 (코드를 작성함에 있어서 삭제되는 데이터는 반환해주도록 함수를 정의하는 것이 원칙이다)
}

// 배열 기반 리스트에 저장된 데이터 갯수 조회
int LCount(List * plist) {
	return plist->numOfData;
}