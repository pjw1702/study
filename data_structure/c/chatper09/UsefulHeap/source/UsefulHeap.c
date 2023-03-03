#include "..\header\UserfulHeap.h"

void HeapInit(Heap * ph, PriorityComp pc) {
	ph->numOfData = 0;
  ph->comp = pc;
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

int GetHiPriChildIDX(Heap * ph, int idx) {

	if(GetLChildIDX(idx) > ph->numOfData)
		return 0;

	else if(GetLChildIDX(idx) == ph->numOfData)   
		return GetLChildIDX(idx);

	else {
		// if(ph->heapArr[GetLChildIDX(idx)].pr >ph->heapArr[GetRChildIDX(idx)].pr )
		if(ph->comp(ph->heapArr[GetLChildIDX(idx)], ph->heapArr[GetRChildIDX(idx)]) < 0)	// comp에 등록된 함수의 호출 결과를 통해서 우선순위를 판단
			return GetRChildIDX(idx);	
		else 
			return GetLChildIDX(idx);	
	}
}

void HInsert(Heap * ph, HData data) {

	int idx = ph->numOfData+1;	// 새 노드가 저장될 인덱스 값을 idx에 저장

	while(idx != 1) {
    
    // if(pr < (ph->heapArr[GetParentIDX(idx)].pr))
		if(ph->comp(data, ph->heapArr[GetParentIDX(idx)]) > 0) {	
			ph->heapArr[idx] = ph->heapArr[GetParentIDX(idx)];

			idx = GetParentIDX(idx);
		} else	// 새 노드의 우선순위가 높지 않다면
			break;
	}
	ph->heapArr[idx] = data;	// 새 노드를 배열에 저장
	ph->numOfData += 1;
}

HData HDelete(Heap * ph) {
	//HData retData = (ph->heapArr[1]).data;	
	//HeapElem lastElem = ph->heapArr[ph->numOfData];	
  HData retData = ph->heapArr[1];	
  HData lastElem = ph->heapArr[ph->numOfData];

	int parentIdx = 1;		
	int childIdx;

	while(childIdx = GetHiPriChildIDX(ph, parentIdx)) {
		//if(lastElem.pr <= ph->heapArr[childIdx].pr)	
    if(ph->comp(lastElem, ph->heapArr[childIdx]) >= 0)	  // comp에 등록된 함수의 호출 결과를 통해서 우선순위를 판단
			break;	
		
		ph->heapArr[parentIdx] = ph->heapArr[childIdx];
		parentIdx = childIdx;
	}

	ph->heapArr[parentIdx] = lastElem;	// 마지막 노드 최종 저장
	ph->numOfData -= 1;
	return retData;
}