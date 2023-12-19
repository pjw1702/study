#!/bin/bash

# AND, OR 연산 - 문자형 연산자를 사용하는 경우
# AND와 OR 연산 역시 문자형 연산자를 사용할 경우에는 아래와 같이 모두 문자형 연산자를 사용하여 기본 사용법으로 사용할 수 있다

VAR1=10
VAR2=20
VAR3=30

# VAR1이 VAR2보다 작고, VAR2가 VAR3보다 크면 True, 그렇지 않으면 False 출력
if [ $VAR1 -lt $VAR2 -a $VAR2 -gt $VAR3 ]
then
  echo True
else
  echo False
fi
# 출력결과
# False

# VAR1이 VAR2보다 작거나, VAR2가 VAR3보다 크면 True, 그렇지 않으면 False 출력
if [ $VAR1 -lt $VAR2 -o $VAR2 -gt $VAR3 ]
then
  echo True
else
  echo
fi
# 출력결과
# True