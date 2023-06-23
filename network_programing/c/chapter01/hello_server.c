#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
//#include <arpa/inet.h>
//#include <sys/socket.h>
#include <winsock.h>  // #include <winsock2.h>

void error_handling(char *message);

int main(int argc, char *argv[]) {
  int serv_sock;  // 서버 소켓에 대한 파일 디스크립터 저장
  int clnt_sock;

  struct sockaddr_in serv_addr; // IP 및 Port 정보 저장
  struct sockaddr_in clnt_addr;
  // socklen_t clnt_addr_size; // POSIX 표준에 정의된 타입, 유닉스 계열이 아닌 Windows 운영체제에서는 타입이 정의되어 있지 않으므로 socklen_t 대신 int 타입을 사용해야 함
  int clnt_addr_size;

  char message[] = "Hello World!";

  if(argc != 2) {
    printf("Usage: %s <port>\n", argv[0]);
    exit(1);
  }

  serv_sock = socket(PF_INET, SOCK_STREAM, 0);  // 소켓 생성(소켓에 대한 파일 디스크립터(소켓 정보) 반환)
  if(serv_sock == -1)
    error_handling("socket() error");
  
  memset(&serv_addr, 0, sizeof(serv_addr)); // IP 및 Port 정보 저장 전 클리어(초기 상태로 세팅: 주소 초기화)
  serv_addr.sin_family = AF_INET;
  serv_addr.sin_addr.s_addr = htonl(INADDR_ANY);
  serv_addr.sin_port = htons(atoi(argv[1]));

  if (bind(serv_sock, (struct sockaddr*) &serv_addr, sizeof(serv_addr) == -1))  // if문 실행과 동시에 bind()함수 호출: serv_sock에 저장된 소켓 정보에 해당하는 소켓에 &serv_addr를 통해 주소 정보 전달
    error_handling("bind() error");
  if (listen(serv_sock, 5) == -1) // 파일 디스크립터(소켓 정보)를 전달하면서 해당 소켓을 연결 요청
    error_handling("listen() error");

  clnt_addr_size = sizeof(clnt_addr);
  clnt_sock = accept(serv_sock, (struct sockaddr*)&clnt_addr, &clnt_addr_size); // 통신이 이루어지기 전, 서버측 OS에 요청받은 이력이 없는 지 확인 후 blocking 상태에서 머물다가 요청이 들어오면 반환을 통해 알려 줌
  if (clnt_sock == -1)
    error_handling("accept() error");
  
  write(clnt_sock, message, sizeof(message)); // 데이터 전달(Hello World! 메세지) 및 수신한 데이터를 읽음

  // 소켓 소멸
  close(clnt_sock);
  close(serv_sock);

  return 0;
}

void error_handling(char *message) {
  fputs(message, stderr);
  fputc('\n', stderr);
  exit(1);
}