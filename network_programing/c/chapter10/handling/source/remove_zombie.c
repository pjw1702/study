#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <signal.h>
#include <sys/wait.h>

void read_childproc(int sig) {
    int status
    // 좀비의 소멸 코드
    pid_t id = waitpid(-1, &status, WNOHANG);
    if(WIFEXITED(status)) {
        printf("Removed proc id: %d \n", id);
        printf("Child send: %d \n", WEXITSTATUS(status));
    }
}

int main(int argc, char *argv[]) {
    pid_t pid;

    // 구조체 sigaction에 대한 초기화
    struct sigaction act;
    act.sa_handler = read_childproc;
    sigemptyset(&act.sa_mask);
    act.sa_flags = 0;

    // 시그널 핸들러 등록
    // SIGCHLD에 대해서 시그널 핸들링을 등록하였다
    // 이 때 등록된 함수 내에서 좀비의 소멸을 막으면 좀비 프로세스는 생성되지 않는다
    sigaction(SIGCHLD, &act, 0);

    // 프로세스 복사
    pid = fork();
    if (pid == 0) {
        puts("Hi! I'm child process");
        sleep(10);
        return 12;
    } else {
        printf("Child proc id: %d \n", pid);
        pid = fork();
        if (pid == 0) {
            puts("Hi! I'm child process");
            sleep(10);
            eixt(24);
        } else {
            int i;
            printf("Child proc id: %d \n", pid);
            for (i = 0; i < 5; i++) {
                puts("wait...");
                sleep(5);
            }
        }
        return 0;
    }

}