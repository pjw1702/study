#!/bin/bash

# if문에 비해 사용 빈도는 떨어지나, 변수의 값에 따라 분기를 해야 하는 경우에 주로 사용된다
# switch-case문은 어떤 쉘을 사용하느냐에 따라 기본 사용법이 조금씩 차이가 있다

# BASH 기준 switch-case문 사용법

# case $변수 in
#     조건값1)
#     수행문1 ;;
#     조건값2)
#     수행문2 ;;
#     조건값3)
#     수행문3 ;;
#     *)
#     수행문4
# esac

case $1 in
    start)
    echo "Start"
    ;;
    stop)
    echo "Stop"
    ;;
    restart)
    echo "Retstart"
    ;;
    help)
    echo "Help"
    ;;
    *)
    echo "Please input sub command"
esac

# 출력 결과
# $ ./switch-case-example.sh start
# Start