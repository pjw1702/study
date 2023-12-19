#!/bin/bash
# 변수값이 서로 다른지 체크하는 경우
# -eq 연산자와는 반대로 두 개의 변수 값이 서로 다른지를 체크할 때는 not equal to의 줄임말인 -ne 연산자와 같지 않다는 의미를 지닌 != 기호를 사용하여 아래와 같이 사용할 수 있다
# 아래 예제는 모두 두 변수의 값이 다르면 True를 출력하고, 같으면 False를 출력한다

VAR1=10
VAR2=20

# if [ $VAR1 -ne $VAR2 ]; then echo True; else echo False; fi
if [ $VAR1 -ne $VAR2 ]
then
  echo True
else
  echo False
fi

# if [ $VAR1 != $VAR2 ]; then echo True; else echo False; fi
if [ $VAR1 != $VAR2 ]
then
  echo True
else
  echo False
fi