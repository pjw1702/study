#include <stdio.h>
#include <unistd.h>
#define BUF_SIZE    30

// 파이프 기간의 프로세스간 양방향 통신: 잘못된 방식
// 하나의 파이프를 이용해서 양방향 통신을 하는 경우, 데이터를 읽고 쓰는 타이밍이 매우 중요해진다
// 타이밍을 컨트롤하는 것은 사실상 불가능하므로 적절한 방법이 될 수 없다
// sleep() 함수를 사용하지 않으면 문제가 발생
// sleep() 함수를 사용해서 컨트롤하는 것도 예제 수준에서의 간단한 프로그램에서나 가능한거지, 실제 프로그램에선 문제가 발생한다
// 부모 프로세스가 write 한 것을 부모 프로세스가 read하는 상황이 발생할 수 있다
// 자식 프로세스는 read하지 못하고 계속 블로킹 상태에 놓이게 된다
int main(int argc, char *argv[]) {
    int fds[2];
    char str1[] = "Who are you?";
    char str2[] = "Thank you for your message";
    char buf[BUF_SIZE];
    pid_t pid;

    pipe(fds);
    pid = fork();
    if(pid == 0) {
        write(fds[1], str1, sizeof(str1));
        // 하나의 파이프에서의 read, write 실행 흐름 제어를 위함
        sleep(2);
        read(fds[0], buf, BUF_SIZE)
        printf("Child proc output: %s \n", buf);
    }
    else {
        read(fds[0], buf, BUF_SIZE);
        printf("Parent proc output: %s \n", buf);
        write(fds[1], str2, sizeof(str2));
        // 부모 프로세스가 먼저 종료되면 명령 프롬프트가 떠버린다
        // 자식 프로세스가 끝나기 전에 부모 프로세스가 먼저 종료되어 프롬프트가 떠버리는 상황 방지
        sleep(3);
    }
    return 0;
}

/** 실행결과
 * gcc pipe2.c -o pipe2
 * ./pipe2
 * Parent proc output: Who are you?
 * Child proc output: Thank you for your message
*/