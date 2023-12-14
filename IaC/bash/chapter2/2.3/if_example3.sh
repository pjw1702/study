#!/bin/bash

# 조건식 타입 - if [ 연산자 $변수 ]; then
# 변수가 문자열이거나 디렉토리 또는 파일과 같이 객체형일 때 주로 사용된다

# 연산자 -z는 변수의 저장된 값의 길이가 0인지를 비교하여 9이면 True, 그렇지 않으면 False를 리턴하는 문자열 연산자이다


value=""
# if [ -z $value ]; then echo True; else echo False; fi
if [ -z $value ]
then
    echo True
else
    echo False
fi

# 출력 결과
# $ ./if_example3.sh 
# True