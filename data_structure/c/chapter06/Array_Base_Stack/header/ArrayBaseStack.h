#ifndef __AB_STCK_H__
#define __AB_STACK_H__

#define TRUE    1
#define FALSE   0
#define STACK_LEN   100

typedef int Data;

// 배열 기반을 고려하여 정의된 스택의 구조체
typedef struct _arrayStack {
    Data stackArr[STACK_LEN];       // 배열 기반 스택
    int topIndex;
} ArrayStack;


typedef ArrayStack Stack;

void StackInit(Stack * pstack);             // 스택의 초기화(topIndex = -1로 초기화)
int SIsEmpty(Stack * pstack);               // Stack이 비지 않았음을 보장(비었는지 확인): T/F

void SPush(Stack * pstack, Data data);      // 스택의 push 연산
void SPop(Stack * pstack);                  // 스택의 pop 연산
Data SPeek(Stack * pstack);                 // 스택의 peek 연산

#endif