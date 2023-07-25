#include <stdio.h>
#include <stdlib.h>
#include <string.h> 
#include <unistd.h>
//#include <sys/socket.h>
#include <winsock.h>  //#include <winsock2.h>

#define BUF_SIZE  1024
#define RLT_SIZE  4
#define OPSZ  4
void error_handling(char *message);

int main(int argc, char *argv[]) {
  int sock;
  char opmsg[BUF_SIZE];
  int result, opnd_cnt, i;
  struct sockaddr_in serv_adr;
  if (argc != 3) {
    printf("Usage: %s <IP> <port>\n", argv[0]);
    exit(1);
  }

  sock = socket(PF_INET, SOCK_STREAM, 0);
  if(sock == -1)
    error_handling("socket() error");
  
  memset(&serv_adr, 0, sizeof(serv_adr));
  serv_adr.sin_family = AF_INET;
  serv_adr.sin_addr.s_addr = inet_addr(argv[1]);
  serv_adr.sin_port = htons(atoi(argv[2]));

  if(connect(sock, (struct sockaddr*)&serv_adr, sizeof(serv_adr)) == -1)
    error_handling("connect() error!");
  else
    puts("Connected..........");
  
  fputs("Operand count: ", stdout);
  scanf("%d", &opnd_cnt);
  opmsg[0] = (char)opnd_cnt;

  for(i = 0; i < opnd_cnt; i++) {
    printf("Operand: %d", i+1);
    scanf("%d", (int*)&opmsg[i*OPSZ+1]);
  }
  fgetc(stdin); // 문자 입력을 받기 위해, 버퍼에 남아있는 \n 문자의 삭제 수행
  fputs("Operator: ", stdout);
  scanf("%c", &opmsg[opnd_cnt*OPSZ+1]);
  write(sock, opmsg, opnd_cnt*OPSZ+2);  // opmsg에 저장되어 있는 연산과 관련된 정보를 한 번에 전송
  read(sock, &result, RLT_SIZE);  // 서버가 전송해주는 연산결과를 저장

  printf("Operation result: %d \n", result);
  close(sock);
  return 0;
}

void error_handling(char *message) {
  fputs(message, stderr);
  fputc('\n', stderr);
  exit(1);
}