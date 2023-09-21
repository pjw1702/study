#include <stdio.h>
#include <unistd.h>
#include <sys/wait.h>

int main(int argc, char *argv[]) {
    int status;
    pid_t pid = fork();

    if(pid == 0) {
        sleep(15);
        return 24;
    } else {
        // 좀비 프로세스의 소멸
        // waitpid 함수 호출 시 첫 번째 인자로 -1, 세 번째 인자로 WNOHANG가 전달되었다
        // 종료된 자식 프로세스가 없으면 0을 반환하면서 함수를 빠져나온다
        while (!waitpid(-1, &status, WNOHANG)) {
            // 자식 프로세스의 종료를 하기까지 15초간 시켰으므로 3초 간격으로 총 5번 출력하게 된다
            sleep(3);
            puts("sleep 3sec");
        }
        
        if(WIFEXITED(status))
            printf("Child send %d \n", WIFEXITSTATUS(status));
    }
}
/** 실행결과
 * gcc waitpid.c -o waitpid
 * ./waitpid
 * sleep 3sec.
 * sleep 3sec.
 * sleep 3sec.
 * sleep 3sec.
 * sleep 3sec.
 * Child send 24
*/