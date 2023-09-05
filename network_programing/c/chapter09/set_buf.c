#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
void error_handling(char *message);

int main(int argc, char *argv[]) {
    int sock;
    int snd_buf = 1024*3, rcv_buf = 1024*3;
    int state;
    socklen_t len;

    sock = socket(PF_INET, SOCK_STREAM, 0);
    state = setsockopt(sock, SOL_SOCKET, SO_RCVBUF, (void*)&rcv_buf, sizeof(rcv_buf));
    if(state)
        error_handling("setsockopt() error!");
    
    state = setsockopt(sock, SOL_SOCKET, SO_SNDBUF, (void*)&snd_buf, sizeof(snd_buf));
    if(state)
        error_handling("setsockopt() error!");

    len = sizeof(snd_buf);
    state = getsockopt(sock, SOL_SOCKET, SO_SNDBUF, (void*)&snd_buf, &len);
    if(state)
        error_handling("getsockopt() error!");

    len = sizeof(rcv_buf);
    state = getsockopt(sock, SOL_SOCKET, SO_RCVBUF, (void*)&rcv_buf, &len);
    if(state)
        error_handling("getsockopt() error!");

    // get_buf.c에서의 실행결과를 확인했을 땐, 입력 버퍼와 출력 버퍼의 크기는 각각 87380과 16384로 확인되었다
    // I/O 버퍼는 상당히 주의 깊게 다뤄져야 하는 영역이기 때문에, 실행결과에서 보이듯이 요구하는 바가 완벽히 반영되지는 않는다
    // 내가 요구하는대로 버퍼의 크기가 정확히 맞춰지지 않는다
    // 예를 들어, 출력 버퍼의 크기를 0으로 변경하려는 경우, 흐름제어와 오류 발생시의 데이터 전송을 위해서라도 최소한의 버퍼는 존재해야 한다
    // 내가 전달하는 변경 요구사항은 참조만 될 뿐이지 100퍼센트 반영은 되지 않는다
    // TCP 프로토콜의 원할한 구현을 위해 OS에서 적절하게 판단한다
    // 하지만 나름대로의 요구사항은 반영되었음을 알 수 있다
    printf("Input buffer size: %d \n", rcv_buf);
    printf("Output buffer size: %d \n", snd_buf);
    return 0;
}

void error_handling(char *message) {
    fputs(message, stderr);
    fputc('\n', stderr);
    exit(1);
}

// 실행결과
// Input buffer size: 6144
// Output buffer size: 6144