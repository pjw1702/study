#!/bin/bash

# AND, OR 연산 - 기호 연산자와 단일 대괄호[]를 사용하는 경우
# AND와 OR 연산 조건과 조건식이 모두 참인지 둘 중 하나만 참인지를 판단할 때 사용
# 조건식은 모두 문자형 연산자를 사용하고 AND, OR 연산만 기호 연산자를 사용할 경우에는 아래와 같이 조건식별로 단일 대괄호를 사용해야 한다

VAR1=10
VAR2=20
VAR3=30

# && 기호를 사용할 경우 조건식별로 대괄호[]를 사용해야 함
if [ $VAR1 -lt $VAR2] && [ $VAR2 -gt $VAR3 ]
then
  echo True
else
  echo False
fi
# 출력결과
# False

# || 기호를 사용할 경우 조건식별로 대괄호[]를 사용해야 함
if [ $VAR1 -lt $VAR2 ] || [ $VAR2 -gt $VAR3 ]
then
  echo True
else
  echo False
fi
# 출력결과
# True