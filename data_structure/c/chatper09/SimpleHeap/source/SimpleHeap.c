#include "..\header\SimpleHeap.h"

void HeapInit(Heap * ph) {
	ph->numOfData = 0;
}

int HIsEmpty(Heap * ph) {
	if(ph->numOfData == 0)
		return TRUE;
	else
		return FALSE;
}

int GetParentIDX(int idx) {
	return idx/2;
}

int GetLChildIDX(int idx) {
	return idx*2;
}

int GetRChildIDX(int idx) {
	return GetLChildIDX(idx)+1;
}

// 두 개의 자식 노드 중 우선순위가 높은 자식 노드의 인덱스 값 반환
int GetHiPriChildIDX(Heap * ph, int idx) {

	// 자식 노드가 존재하지 않는경우
	// numOfData는 마지막 노드의 고유번호이므로, 새롭게 추가되는 자식 노드의 값이 numOfData보다 크면 존재하지 않는 자식노드이다
	if(GetLChildIDX(idx) > ph->numOfData)
		return 0;

	// 자식 노드가 왼쪽 자식 노드 하나만 존재할 경우
	// 힙은 완전 이진트리이므로, 자식 노드가 하나만 존재하면 이는 왼쪽 자식노드이고, 마지막 노드이다
	else if(GetLChildIDX(idx) == ph->numOfData)   // 마지막 노드의 고유번호이므로, 마지막 노드인지 아닌지 검사할 수 있다
		return GetLChildIDX(idx);

	// 자식 노드가 둘 다 존재한다면,
	else {
		// 오른쪽 자식 노드의 우선순위가 높다면,
		if(ph->heapArr[GetLChildIDX(idx)].pr >ph->heapArr[GetRChildIDX(idx)].pr )
			return GetRChildIDX(idx);	// 오른쪽 자식 노드의 인덱스 값 반환
		
		// 왼쪽 자식 노드의 우선순위가 높다면,
		else 
			return GetLChildIDX(idx);	// 왼쪽 자식 노드의 인덱스 값 반환
	}
}

void HInsert(Heap * ph, HData data, Priority pr) {

	// 노드는 우선순위에 따라 어차피 이동해야 하므로 새 노드의 인덱스 정보를 갱신만 하면 된다

	int idx = ph->numOfData+1;	// 새 노드가 저장될 인덱스 값을 idx에 저장
	HeapElem nelem = {pr, data};	// 새 노드의 생성 및 초기화

	// 새 노드가 저장될 위치가 루트 노드의 위치가 아니라면 while문 반복
	while(idx != 1) {
		// 새 노드와 부모 노드의 우선순위 비교
		if(pr < (ph->heapArr[GetParentIDX(idx)].pr)) {	// 새 노드의 우선순위가 높다면
			// 부모 노드를 한 레벨 내림, 실제로 내림
			ph->heapArr[idx] = ph->heapArr[GetParentIDX(idx)];

			// 새 노드를 한 레벨 내림, 실제로 올리지는 않고 인덱스 값만 갱신
			idx = GetParentIDX(idx);
		} else	// 새 노드의 우선순위가 높지 않다면
			break;
	}
	ph->heapArr[idx] = nelem;	// 새 노드를 배열에 저장
	ph->numOfData += 1;
}

HData HDelete(Heap * ph) {
	HData retData = (ph->heapArr[1]).data;	// 반환을 위해서 삭제할 데이터 저장(백업)
	HeapElem lastElem = ph->heapArr[ph->numOfData];	// 힙의 마지막 노드 저장

  // 힙의 마지막 노드를 별도로 임시 저장 후, 비교 하면서 맞는 자리를 찾아나간다
  // 굳이 루트 노드의 자리로 옮기지 않아도 된다: swapping이 아닌, 임시로 저장한 데이터와의 비교 및 대입 연산으로 해결할 수 있다

	// 아래의 변수 parentIdx에는 마지막 노드가 저장될 위치 정보가 담긴다
	int parentIdx = 1;		// 루트 노드가 위치해야 할 인덱스 값 저장
	int childIdx;

	// 우선순위가 높은 루트 노드의 자식 노드를 시작으로 반복문 시작
	while(childIdx = GetHiPriChildIDX(ph, parentIdx)) {
		if(lastElem.pr <= ph->heapArr[childIdx].pr)	// 마지막 노드와 루트의 우선순위가 높은 자식 노드와의 우선순위 비교
			break;	// 마지막 노드의 우선순위가 높으면 반복문 탈출
		
		// 마지막 노드보다 우선순위가 높으니, 비교대상 노드의 위치를 한 레벨 올림
		ph->heapArr[parentIdx] = ph->heapArr[childIdx];
		parentIdx = childIdx;	// 마지막 노드가 저장될 위치 정보를 한 레벨 내림
	}	// 반복문을 탈출하면 parentIdx에는 마지막 노드의 위치 정보가 저장됨

	ph->heapArr[parentIdx] = lastElem;	// 마지막 노드 최종 저장
	ph->numOfData -= 1;
	return retData;
}