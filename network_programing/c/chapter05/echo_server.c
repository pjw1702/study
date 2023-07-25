#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
//#include <sys/socket.h>
#include <winsock.h>  //#include <winsock2.h>

#define BUF_SIZE  1024
void error_handling(char *message);

int main (int argc, char *argv[]) {
  int serv_sock, clnt_sock;
  char message[BUF_SIZE];
  int str_len, i;

  struct sockaddr_in serv_adr;
  struct sockaddr_in clnt_adr;
  int clnt_adr_sz;    //socklen_t clnt_adr_sz;

  if(argc != 2) {
    printf("Usage: %s <port>\n", argv[0]);
    exit(1);
  }

  serv_sock = socket(PF_INET, SOCK_STREAM, 0);
  if(serv_sock == -1) {
    error_handling("socket() error");
  }

  memset(&serv_adr, 0, sizeof(serv_adr));
  serv_adr.sin_family = AF_INET;
  serv_adr.sin_addr.s_addr = htonl(INADDR_ANY);
  serv_adr.sin_port = htons(atoi(argv[1]));

  if(bind(serv_sock, (struct sockaddr*)&serv_adr, sizeof(serv_adr)) == -1)
    error_handling("bind() error");

  if(listen(serv_sock, 5) == -1)  // 크기가 5인 연결요청 대기 큐 생성
    error_handling("listen() error");
  
  clnt_adr_sz = sizeof(clnt_adr);

  for (i = 0; i < 5; i++) { // accpet() 함수 반복 호출: 1:1로 클라이언트와의 통신을 반복하여 수행 
    clnt_sock = accept(serv_sock, (struct sockaddr*)&clnt_adr, &clnt_adr_sz); // OS에 의해 생성된 실제 소켓의 파일 디스크립터 반환
    if(clnt_sock == -1)
      error_handling("accept() error");
    else 
      printf("Connected client %d \n", i+1);

    while((str_len = read(clnt_sock, message, BUF_SIZE)) != 0)
    write(clnt_sock, message, strlen);

    close(clnt_sock);
  }

  close(serv_sock);
  return 0;
}

void error_handling(char *message) {
  fputs(message, stderr);
  fputc('\n', stderr);
  exit(1);
}