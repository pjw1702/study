#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
//#include <sys/socket.h>
#include <winsock.h>  //#include <winsock2.h>

#define BUF_SIZE  30
void error_handling(char *message);

int main (int argc, char *argv[]) {
  int sock;
  char message[BUF_SIZE];
  int str_len;
  //socklen_t adr_sz;
  int adr_sz;
  struct sockaddr_in serv_adr, from_adr;

  if(argc != 3) {
    printf("Usage: %s <IP> <port>\n", argv[0]);
    exit(1);
  }

  sock = socket(PF_INET, SOCK_DGRAM, 0);
  if(sock == -1) {
    error_handling("socket() error");
  }

   memset(&serv_adr, 0, sizeof(serv_adr));
  serv_adr.sin_family = AF_INET;
  serv_adr.sin_addr.s_addr = inet_addr(argv[1]);
  serv_adr.sin_port = htons(atoi(argv[2]));

  // UDP는 데이터의 경계가 존재하므로 한 번의 recvfrom 함수호출을 통해서 하나의 메시지를 완전히 읽어 들인다
  // sendto 함수호출 시 IP와 PORT 번호가 자동으로 할당되기 때문에 일반적으로 UDP의 클라이언트 프로그램에서는 주소정보를 할당하는 별도의 과정이 필요하다
  while(1) {
    fputs("Insert message(q to quit): ", stdout);
    fgets(message, sizeof(message), stdin);
    if(!strcmp(message, "q\n") || !strcmp(message, "Q\n"))
      break;

    sendto(sock, message, strlen(message), 0, (struct sockaddr*)&serv_adr, sizeof(serv_adr));	// 바인딩(IP, PORT 할당)이 자동으로 이루어짐
    adr_sz = sizeof(from_adr);
    str_len = recvfrom(sock, message, BUF_SIZE, 0, (struct sockaddr*)&from_adr, &adr_sz);
    
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