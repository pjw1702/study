#!/bin/bash
# 변수 값이 서로 같은지 체크하는 경우
# 두 개의 변수가 서로 같은지를 체크할 때는 영어 단어 equals의 줄임 말인 -eq 연산자를 사용하거나, 같다는 의미의 =기호를 사용할 수 있다
# 아래 예제는 모두 변수 VAR1의 값과 VAR2의 값이 같은지를 체크한다

VAR1=10
VAR2=10

# if [ $VAR1 -eq $VAR2 ]; then echo True; else echo False; fi
if [ $VAR1 -eq $VAR2 ]
then
  echo True
else
  echo False
fi

# 아래와 같다
# if [ $VAR1 = $VAR2 ]; then echo True; else echo False; fi
if [ $VAR1 = $VAR2 ]
then
  echo True
else
  echo False
fi

