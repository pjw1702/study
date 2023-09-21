#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>

int main(int argc, char *argv[]) {
    int status;
    pid_t pid = fork();

    // 자식 프로세스 생성 종료
    // OS에 종료 코드 3이 등록된 이후 좀비로 남게 됨
    if (pid == 0)
        return 3;
    else {
        printf("Child PID: %d \n", pid);
        pid = fork();
        // 자식 프로세스 생성 종료
        if (pid == 0)
            exit(7);
        else {  // 부모 프로세스 실행 영역
        
            printf("Child PID: %d \n", pid);
            // 좀비 프로세스의 소멸
            // wait 함수는 OS에 등록된 종료 코드를 획득하는 기능의 함수이다
            // wait 함수의 인자 값으로 전달된 주소값을 가지는 변수에, 반환에 관련된 각종 정보가 저장
            // wait 함수의 경우 자식 프로세스가 종료되지 않은 상황에서는 반환하지 않고 블로킹 상태에 놓인다는 단점이 있다
            // 부모 프로세스의 필요에 의해 wait 함수를 호출했는데, 자식 프로세스가 실행 중인 경우 그 다음 호출을 할 수 없다
            wait(&status);
            // wait 함수로 부터 전달받은 반환에 관련된 각종 정보가 담긴 status 변수에서 정보를 얻기 위해서는 아래 메크로 함수 호출이 필요하다
            // WIFEXITED: 자식 프로세스가 정상 종료한 경우 참(True)을 반환
            // WEXITSTATUS: 자식 프로세스의 전달 값을 반환
            if(WIFEXITED(status))
                printf("Child send one: %d \n", WEXITSTATUS(status));
            wait(&status);
            if(WIFEXITED(status))
                printf("Child send two: %d \n", WEXITSTATUS(status));
            sleep(30);
        }
    }
}

/** 실행 결과
 * gcc wait.c -o wait
 * ./ wait
 * Child PID: 12337
 * Child PID: 12338
 * Child send one: 3
 * Child send two: 7
*/