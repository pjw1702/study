#include <stdio.h>
#include <unistd.h>
#include <signal.h>

void timeout(int sig) {
    if(sig == SIGALRM)
        puts("Time out!");
    alarm(2);
}

int main (int argc, char *argv[]) {
    int i;

    // 구조체 sigaction에 대한 초기화
    struct sigaction act;
    act.sa_handler = timeout;
    // 멤버 sa_mask에 대한 비트를 모두 0으로 초기화
    sigemptyset(&act.sa_mask);
    act.sa_flags = 0;

    // 핸들러 등록
    sigaction(SIGALRM, &act, 0);

    alarm(2);

    for(i = 0; i < 3; i++) {
        puts("wait...");
        sleep(100);
    }

    return 0;
}

/** 실행결과
 * gcc sigaction.c -o sigaction
 * ./sigaction
 * wait...
 * Time out!
 * wait...
 * Time out!
 * wait...
 * Time out! 
*/