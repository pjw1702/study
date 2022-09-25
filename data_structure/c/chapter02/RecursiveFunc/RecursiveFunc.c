# include<stdio.h>

/** 함수를 호출하면 함수에 대한 메모리 공간(스택)이 할당된다
    지역변수(매개변수)를 선언하지 않아도 별도의 메모리 공간(스택)이 요구된다: 실행 종료 후 돌아갈(back) 메모리 주소값을 저장하기 위함
    탈출 조건을 정의하지 않으면 제한된 리소스인 스택이라는 메모리는 오버플로우 되므로 문제가 된다 
 */
void Recursive(int num) {
    if (num <= 0)   // 재귀의 탈출조건(재귀함수를 사용할 때, 탈출조건을 반드시 정의하여야 한다)
        return;     // 재귀의 탈출!
    printf("Recursive call! %d \n", num);
    Recursive(num-1);
}

int main(void) {
    Recursive(3);
    return 0;
}