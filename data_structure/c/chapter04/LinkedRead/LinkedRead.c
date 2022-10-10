// 이 자료구조는 아직 연결 리스트의 ADT와 삽입, 삭제, 조회의 기능이 별도로 구분되어 구현되지 않은 코드이다
// 1. 자료구조의 ADT 정의, 2. 정의한 ADT 구현, 3. 구현이 완료된 자료구조의 활용 의 순서를 지켜야 한다

#include <stdio.h>
#include <stdlib.h>

typedef struct _node {
	int data;
	struct _node * next;
} Node;

int main(void) {
	Node * head = NULL;
	Node * tail = NULL;
	Node * cur = NULL;

	Node * newNode = NULL;
	int readData;
	
    // 데이터 입력 및 저장
	while(1) {
        printf("자연수 입력: ");
        scanf("%d", &readData);
        if(readData < 1)
            break;

        // 노드의 추가과정
        newNode = (Node*)malloc(sizeof(Node));
        newNode->data = readData;
        newNode->next = NULL;

        // head 및 tail 초기화
        if(head == NULL)
            head = newNode;		// head는 초기 상태(NULL)일 때 빼고는 항상 처음 노드를 가리키도록 구성
        else
            tail->next = newNode;	// tail이 가리키는 곳은 newNode 이므로 newNode의 멤버 next = newNode

        tail = newNode;		// tail이 newNode를 가리킴
	}

	// 전체 데이터의 출력 과정
	if(head == NULL)
		printf("저장된 자연수가 존재하지 않습니다 \n");
	else {
		cur = head;
		printf("%d", cur->data);
		while(cur->next != NULL) {		// tail 일때 탈출
			cur = cur->next;		// cur이 가리키는 곳의 next
			printf("%d", cur->data);
		}
	}

	// 전체 노드의 삭제 과정	
	if(head == NULL)
		return 0;
	else {
		Node * delNode = head;
		Node * delNextNode = head->next;

		printf("%d을 삭제\n", head->data);
		free(delNode);      // 메모리 해제
	

        while(delNextNode != NULL) {
            delNode = delNextNode;
            delNextNode = delNextNode->next;

            printf("%d을 삭제:\n", delNode->data);
            free(delNode);
        }
    }
	return 0;
}