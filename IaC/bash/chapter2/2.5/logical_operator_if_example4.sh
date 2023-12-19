#!/bin/bash

# AND, OR 연산 - 조건식도 기호 연산자를 사용할 경우
# 조건식과 AND, OR 연산 모두 기호 연산자를 사용할 경우에는 대괄호[]가 아닌 중첩 소괄호 (())를 사용하여 다음과 같이 표현할 수 있다

VAR1=10
VAR2=20
VAR3=30

# && 기호를 사용할 경우 조건식별로 대괄호[]를 사용해야 함
if (( $VAR1 < $VAR2 )) && (( $VAR2 > $VAR3 ))
then
  echo True
else
  echo False
fi
# 출력결과
# False

# || 기호를 사용할 경우 조건식별로 대괄호[]를 사용해야 함
if (( $VAR1 < $VAR2 || $VAR2 > $VAR3 ))
then
  echo True
else
  echo
fi
# 출력결과
# True