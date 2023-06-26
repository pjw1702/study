#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
//#include <arpa/inet.h>
//#include <sys/socket.h>
#include <winsock.h>  // #include <winsock2.h>

void error_handling(char *message);

int main(int argc, char* argv[]) {
  int sock;
  struct sockaddr_in serv_addr;
  char message[30];
  int str_len = 0;
  int idx = 0, read_len = 0;

  if (argc != 3) {
    printf("Usage: %s <IP> <port>\n", argv[0]);
    exit(1);
  }

  sock = socket(PF_INET, SOCK_STREAM, 0);
  if (sock == -1)
    error_handling("socket() error");

  memset(&serv_addr, 0, sizeof(serv_addr)); // 클라이언트에 할당한 소켓이 아닌, 연결시킬 상대방의 소켓의 주소 정보를 담기 위한 주소 초기화
  serv_addr.sin_family = AF_INET;
  serv_addr.sin_addr.s_addr = inet_addr(argv[1]);
  serv_addr.sin_port = htons(atoi(argv[2]));

  if (connect(sock, (struct sockaddr*)&serv_addr, sizeof(serv_addr)) == -1)
    error_handling("connect() error!");
  
  // str_len = read(sock, message, sizeof(message)-1);
  // if (str_len == -1)
  //   error_handling("read() error!");
  
  // printf("Message from server: %s \n", message);

  while(read_len = read(sock, &message[idx++], 1)) {  // 서버에서는 데이터를 한 번에 전달하였지만, 클라이언트에서는 1바이트씩 쪼개서 수신받음: 전송되는 데이터의 경계(Boundary)가 존재하지 않음
    if (read_len = -1)
      error_handling("read() error!");

      str_len += read_len;
  }

  printf("Message from server: %s \n", message);
  printf("Function read call count: %d \n", str_len);
  close(sock);
  
  return 0;
}

void error_handling(char *message) {
  fputs(message, stderr);
  fputc('\n', stderr);
  exit(1);
}