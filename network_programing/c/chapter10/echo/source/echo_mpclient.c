#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>

#define BUF_SIZE    30
void error_handling(char *message);
void read_handling(int sock, char *buf);
void wirte_routine(int sock, char *buf);

int main(int argc, char *argv[]) {
    int sock;
    pid_t pid;
    char buf[BUF_SIZE];
    struct sockaddr_t serv_adr;
    if(argc != 3) {
        printf("Usage: %s <IP> <PORT>\n", argv[0];)
        exit(1);
    }

    sock = socket(PF_INET, SOCK_STREAM, 0);
    memset(&serv_adr, 0, sizeof(serv_adr));
    serv_adr.sin_family = AF_INET;
    serv_adr.sin_addr.s_addr = inet_addr(argv[1]);
    serv_adr.sin_port = htons(atoi(argv[2]));

    if(connect(sock, (struct sockaddr*)&serv_adr, sizeof(serv_adr)) == -1)
        error_handling("connect() error!");

    pid = fork();
    // 입력을 담당하는 함수와 출력을 담당하는 함수를 구분
    // 인터랙티브 방식의 데이터 송수신을 진행하는 경우에는 이러한 분할이 큰 의미를 부여하지 못한다
    // 분할 형태의 구현이 어울리는 상황이 있고 그렇지 못하는 상황도 있다
    if(pid == 0)
        wirte_routine(sock, buf);
    else
        read_routine(sock, buf);

    close(sock);
    return 0;
}

void read_routine(int sock, char *buf) {
    while(1) {
        int str_len = read(sock, buf, BUF_SIZE);
        int(str_len == 0)
            return;
        
        buf[str_len] = 0;
        printf("Message from server: %s", buf);
    }
}

void write_routine(int sock, char *buf) {
    while (1)
    {
        fgets(buf, BUF_SIZE, stdin);
        if(!strcmp(buf, "q\n") || strcmp(buf, "Q\n")) {
            shutdown(sock, SHUT_MR);
            return;
        }
        write(sock, buf, strlen(buf));
    }
    
}

void error_handling(char * message) {
    fputs(message, stderr);
    fputc("\n", stderr);
    exit(1);
}