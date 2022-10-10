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

    // 추가된 문장, 더미 노드 추가
    head = (Node*)malloc(sizeof(Node));
    tail = head;
	
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

        /*
        // head 및 tail 초기화
        if(head == NULL)
            head = newNode;		// head는 초기 상태(NULL)일 때 빼고는 항상 처음 노드를 가리키도록 구성
        else
            tail->next = newNode;	// tail이 가리키는 곳은 newNode 이므로 newNode의 멤버 next = newNode
        */
        tail->next = newNode;

        tail = newNode;		// tail이 newNode를 가리킴
	}

	// 전체 데이터의 출력 과정
	if(head == NULL)
		printf("저장된 자연수가 존재하지 않습니다 \n");
	else {
		cur = head;
		// printf("%d", cur->data);        // 첫 번째 데이터 출력
		while(cur->next != NULL) {		// 두 번째 이후의 데이터 출력 (tail 일때 탈출)
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

		// printf("%d을 삭제\n", head->data);
		// free(delNode);      // 메모리 해제
	

        while(delNextNode != NULL) {
            delNode = delNextNode;
            delNextNode = delNextNode->next;

            printf("%d을 삭제:\n", delNode->data);
            free(delNode);
        }
    }
    return 0;
}