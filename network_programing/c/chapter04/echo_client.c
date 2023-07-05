#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
//#include <sys/socket.h>
#include <winsock.h>  //#include <winsock2.h>

#define BUF_SIZE  1024
void error_handling(char *message);

int main (int argc, char *argv[]) {
  int sock;
  char message[BUF_SIZE];
  int str_len;
  struct sockaddr_in serv_adr;

  if(argc != 3) {
    printf("Usage: %s <IP> <port>\n", argv[0]);
    exit(1);
  }

  sock = socket(PF_INET, SOCK_STREAM, 0);
  if(sock == -1) {
    error_handling("socket() error");
  }

   memset(&serv_adr, 0, sizeof(serv_adr));
  serv_adr.sin_family = AF_INET;
  serv_adr.sin_addr.s_addr = htonl(INADDR_ANY);
  serv_adr.sin_port = htons(atoi(argv[1]));

  if(connect(sock, (struct sockaddr*)&serv_adr, sizeof(serv_adr)) == -1)  // 연결 요청
    error_handling("connect() error!");
  else 
    puts("Connected..............");

  while(1) {
    fputs("Input message(Q to quit): ", stdout);
    fgets(message, BUF_SIZE, stdin);

    if(!strcmp(message, "q\n") || !strcmp(message, "Q\n"))  // q또는 Q를 입력할 때까지 데이터 입력
      break;
    
    write(sock, message, strlen(message));
    str_len = read(sock, message, BUF_SIZE-1);
    message[str_len] = 0;
    printf("Message from server: %s", message);
  }

  close(sock);
  return 0;
}

void error_handling(char *message) {
  fputs(message, stderr);
  fputc('\n', stderr);
  exit(1);
}