#include <stdio.h>
//#include <arpa/inet.h>
#include <winsock.h> // #include <winsock2.h>

int main(int argc, char *argv[]) {
  unsigned short host_port = 0x1234;  // 호스트 바이트 순서로 저장
  unsigned short net_port;
  unsigned long host_addr = 0x12345678; // 호스트 바이트 순서로 저장
  unsigned long net_addr;

  net_port = htons(host_port);  // 저장된 데이터를 네트워크 바이트 순서로 변환 
  net_addr = htonl(host_addr);  // 저장된 데이터를 네트워크 바이트 순서로 변환
  
  // 출력 값에서, host_port 및 host_addr과 net_port 및 net_addr과의 값이 서로 다르다면, 해당 컴퓨터의 CPU의 호스트 바이트 순서는 리틀 엔디안 타입이다
  printf("Host orderd port: %#x \n", host_port);
  printf("Network orderd port: %#x \n", net_port);
  printf("Host orderd address: %#lx \n", host_addr);
  printf("Network orderd address: %#lx \n", net_addr);

  return 0;
}

/** 실행 결과
Host orderd port: 0x1234
Network orderd port: 0x3412
Host orderd address: 0x12345678
Network orderd address: 0x78563412
*/