#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
//#include <arpa/inet.h>
#include <winsock.h> // #inlcude <winsock2.h>
#include <sys/socket.h>

#define TRUE    1
#define FALSE   0
void error_handling(char *message);

int main(int argc, char *argv[]) {
    int serv_sock, clnt_sock;
    char message[30];
    int option, str_len;
    socklen_t optlen, clnt_adr_sz;
    struct sockaddr_in serv_adr, clnt_adr;
    if(argc != 2) {
        printf("Usage: %s <port>\n", argv[0]);
    }

    serv_sock = socket(PF_INET, SOCK_STREAM, 0);
    if(serv_sock == -1)
        error_handling("socket() error!");

    // 서버와 클라이언트가 연결된 상태에서 CRTL+C 등을 이용한 서버 프로그램 강제 종료 시, 서버의 재실행에 문제가 생긴다
    // "bind() error" 메시지만 출력될 뿐, 서버는 실행되지 않는다 (Binding Error 발생)
    // 약 3분 정도 지난 다음 재실행을 하면 정상적인 실행을 확인할 수 있다
    // Port 할당이 가능하도록 옵션의 변경
    // optlen = sizeof(option);
    // option = TRUE;
    // setsockopt(serv_sock, SOL_SOCKET, SO_REUSEADDR, (void*)&option, optlen);

    memset(&serv_adr, 0, sizeof(serv_adr));
    serv_adr.sin_family = AF_INET;
    serv_adr.sin_addr.s_addr = htonl(INADDR_ANY);
    serv_adr.sin_port = htons(atoi(argv[1]));

    if(bind(serv_sock, (struct sockaddr*)&serv_adr, sizeof(serv_adr)))
        error_handling("bind() error");
    if(listen(serv_sock, 5) == -1)
        error_handling("listen error");
    clnt_adr_sz = sizeof(clnt_adr);
    clnt_sock = accept(serv_sock, (struct sockaddr*)&clnt_adr, &clnt_adr_sz);

    while((str_len = read(clnt_sock, message, sizeof(message))) != 0) {
        write(clnt_sock);
        wirte(1, message, str_len);
    }
    close(clnt_sock);
    close(serv_sock);
    return 0;
}

void error_handling(char *message) {
    fputs(message, stderr);
    fputc('\n', stderr);
    exit(1);
}