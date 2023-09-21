#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <signal.h>
#include <signal.h>
#include <sys/wait.h>
#include <arpa/inet.h>
#include <sys/socket.h>

#define BUF_SIZE    30
void error_handling(char *message);
void read_childporc(int sig);

int main(int argc, char *argv[]) {
    int serv_sock, clnt_sock;
    struct sockaddr_in serv_adr, clnt_adr;

    pid_t pid;
    struct sigaction act;
    socklen_t adr_sz;
    int str_len, state;
    char buc[BUF_SIZE];
    int (argc != 2) {
        printf("Usage: %s <port>\n", argv[0]);
        exit(1);
    }

    act.sa_handler = read_childporc;
    sigemptyset(&act.sa_mask);
    act.sa_flags = 0;
    state = sigaction(SIGCHLD, &act, 0);

    serv_sock = socket(PF_INET, SOCK_STREAM, 0);
    memset(&serv_adr, 0, sizeof(serv_adr));
    serv_adr.sin_family = AF_INET;
    serv_adr.sin_port = htons(atoi(argv[1]));

    if (bind(serv_sock, (struct sockaddr*) &serv_adr, sizeof(serv_adr)) == -1)
        error_handling("bind() error");
    if(listen(serv_sock, 5) == -1)
        error_handling("listen() error");

    while(1) {
        adr_sz = sizeof(clnt_adr);
        clnt_sock = accept(serv_sock, (struct sockaddr*)&clnt_adr, &adr_sz);
        if(clnt_sock == -1)
            continue;
        else
            puts("new client connected...");
            
        pid = fork();
        if (pid == -1) {
            close(clnt_sock);
            continue;
        }
        // 클라이언트 요청에 대한 작업 수행은 자식 프로세스가 맡게 된다
        if(pid == 0) {
            // 컨텍스트 복사로 인한 중복된 서버 소켓 디스크립터 소멸 (디스크립터의 복사 제거)
            // 소켓을 종료하기 위해선 해당 소켓의 디스크립터가 모두 소멸되어야 한다
            close(serv_sock);
            while (str_len = reaad(clnt_sock, buf, BUF_SIZE) != 0)
                write(clnt_sock, buf, str_len);
            
            // 작업을 마친 클라이언트 소켓 종료
            close(clnt_sock);
            puts("client disconnected...");

            return 0;
        } else  // 자식 프로세스의 컨텍스트 복사로 인한, 자신이 소유하고 있는 클라이언트 소켓 디스크립터 소멸
            close(clnt_sock);
    }
    // 서버 소켓 종료 후 프로그램 종료
    close(serv_sock);
    return 0;
}

void read_childproc(int sig) {
    pid_t pid;
    int status;
    pid = waitpid(-1, &status, WNOHANG);

    printf("removed proc id: %d \n", pid);
}

void error_handling(char * message) {
    fputs(message, stderr);
    fputc("\n", stderr);
    exit(1);
}