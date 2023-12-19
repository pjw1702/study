#!/bin/bash

# AND, OR 연산 - 기호 연산자와 중첩 대괄호[[]]를 사용하는 경우
# 조건식이 단순하여 조건식별로 단일 대괄호를 사용할 필요가 없는 경우에는 조건식 전체를 중첩 대괄호[[]]를 사용해야 한다

VAR1=10
VAR2=20
VAR3=30

# && 기호를 사용할 경우 조건식별로 대괄호[]를 사용해야 함
if [[ $VAR1 -lt $VAR2 && $VAR2 -gt $VAR3 ]]
then
  echo True
else
  echo False
fi
# 출력결과
# False

# || 기호를 사용할 경우 조건식별로 대괄호[]를 사용해야 함
if [[ $VAR1 -lt $VAR2 || $VAR2 -gt $VAR3 ]]
then
  echo True
else
  echo
fi
# 출력결과
# True