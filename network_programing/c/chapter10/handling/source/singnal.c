#include <stdio.h>
#include <unistd.h>
#include <signal.h>

void timeout(int sig) {
    if(sig == SIGALRM)
        puts("Time out!");
    // 2초 간격으로 시그널을 발생
    alarm(2);
}

void keycontrol(int sig) {
    if (sig == SIGINT)
        puts("CTRL+C pressed");
}

int main(int argc, char *argv[]) {
    int i;
    // signal 함수는 OS 별로 동작방식의 차이를 보이므로 sigaction 함수를 대신 사용
    // signal 함수는 과거 프로그램과의 호환성을 유지하기 위해서 제공
    // 시그널 핸들러 등록
    // 핸들러를 등록하는 작업이지, signal 함수 자체가 핸들링을 수행하는 기능을 하는 것이 아니다
    signal(SIGALRM, timeout);
    signal(SIGINT, keycontrol);
    // 2초가 지나면 시그널이 발생 (SIGALRM 발생을 2초 뒤로 예약)
    // 2초마다 signal 함수를 통해 등록된 핸들러 함수가 호출되어 핸들링을 수행
    alarm(2);

    for (i = 0; i < 3; i++) {
        puts("wait...");
        // 시그널이 발생하면, sleep함수의 호출을 통해 블로킹 상태에 있던 프로세스가 깨어난다
        // 따라서 코드의 내용대로 300초의 sleep 시간을 갖지 않는다
        // sleep함수를 통해 프로세스가 잠들게 되는데, 시그널 발생으로 인해 시그널이 프로세스를 깨워주는 꼴이 된다
        // 실행 시 1초도 채 걸리지 않는다는 것을 알 수 있다
        sleep(100);
    }

    return 0;
}

/** 실행결과
 * gcc signal.c -o signal
 * ./signal
 * wait...
 * Time out!
 * wait...
 * Time out!
 * wait...
 * Time out!
*/