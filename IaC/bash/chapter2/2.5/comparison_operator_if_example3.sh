#!/bin/bash
# 변수값의 크기를 비교하는 경우: 문자형 연산자를 사용할 경우
# 두 개의 변수값의 크기를 비교할 때는 첫 번째 변수를 기준으로 두 번째 변수보다 큰지, 아니면 작은지를 비교한다
# 첫 번째 변수가 두 번째 변수보다 큰지 여부를 비교할 때는 greather than의 줄임말의 -gt 연산자를 사용한다
# 크거나 같은지를 비교할 때는 greather than or equal to의 줄임말인 -ge 연산자를 사용한다
# 첫 번째 변수가 두 번째 변수보다 작은지를 비교할 때는 less than의 약자인 -lt 연산자를 사용한다
# 작거나 같은지를 비교할 때는 less than or equal to의 약자인 -le 연산자를 사용한다

VAR1=10
VAR2=20

# if [ $VAR1 -gt $VAR2 ]; then echo True; else echo False; fi
if [ $VAR1 -gt $VAR2 ]
then
  echo True
else
  echo False
fi

# if [ $VAR1 -ge $VAR2 ]; then echo True; else echo False; fi
if [ $VAR1 -ge $VAR2 ]
then
  echo True
else
  echo False
fi
# if [ $VAR1 -gt $VAR2 ]; then echo True; else echo False; fi
if [ $VAR1 -lt $VAR2 ]
then
  echo True
else
  echo False
fi
# if [ $VAR1 -gt $VAR2 ]; then echo True; else echo False; fi
if [ $VAR1 -le $VAR2 ]
then
  echo True
else
  echo False
fi

# 출력 결과
# False
# False
# True
# True
