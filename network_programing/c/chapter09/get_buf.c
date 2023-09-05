#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/socket.h>
void error_handling(char *message);

int main(int argc, char *argv[]) {
    int sock;
    int snd_buf, rcv_buf, state;
    socklen_t len;

    sock = socket(PF_INET, SOCK_STREAM, 0);
    len = sizeof(snd_buf);
    state = getsockopt(sock, SOL_SOCKET, SO_SNDBUF, (void*)&snd_buf, &len);     // 입력 버퍼의 크기를 참조 및 변경할 때에는 SO_SNDBUF
    if (state)
        error_handling("getsocketopt()" error!);
    
    len = sizeof(rcv_buf);
    state = getsockopt(sock, SOL_SOCKET, SO_RCVBUF, (void*)&rcv_buf, &len);     // 출력 버퍼의 크기를 참조 및 변경할 때에는 SO_RCVBUF
    if (state)
        error_handling("getsockopt()" error!);

    // 디폴트 I/O 버퍼의 크기는 OS별, 그리고 OS 버전마다 다를 수 있다
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
// Input buffer size: 87380
// Output buffer size: 16384
