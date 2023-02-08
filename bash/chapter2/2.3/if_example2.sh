#!/bin/bash

# 조건식 타입 - if [ 변수 연산자 조건값 ]; then
# 조건값이 고정되어 있을 경우 사용하는 예제로 변수와 조건값을 비교할 때 주로 사용한다
# 조건식을 표현할 때는 if문의 앞뒤로 반드시 공백을 두어야 한다

# value값이 0이면 True, 그렇지 않으면 False를 출력

# if [ $value=0 ]; then echo True; else echo False; fi
value=0
if [ $value=0 ]
then
    echo True
else
    echo False
fi

# 출력 결과
# $ ./if_example2.sh 
# True