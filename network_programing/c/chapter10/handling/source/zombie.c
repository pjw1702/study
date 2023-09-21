#include <stdio.h>
#include <unistd.h>

int main(int argc, char *argv[]) {
    pid_t pid = fork();

    if (pid == 0)   // if child Process
        puts("Hi, I am a child Process");
    else {  // Parent Process
        printf("Chlid Process ID: %d \n", pid);
        // 자식 프로세스의 종료 값을 반환 받을 부모 프로세스가 소멸되면, 좀비의 상태로 있던 자식 프로세스도 함께 소멸
        // 부모 프로세스가 소멸되기 전에 좀비의 생성을 확인을 위한 30초간 지연
        sleep(30);  // Sleep 30 sec
    }

    // 부모 프로세스는 30초간 지연 상태인데, 자식 프로세스가 먼저 종료되었으므로 30초 간은 자식 프로세스가 좀비 프로세스로 남게된다
    // ps au 등의 명령어를 수행했을때, 상태가 Z+로 남아있는 것을 확인할 수 있다
    if (pid == 0)
        puts("End Child Process");
    else
        puts("End Parent process");
    
    return 0;
}