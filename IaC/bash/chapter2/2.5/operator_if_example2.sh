#!/bin/bash
# 문자열 변수에 문자열이 저장되었는지 체크할 경우
# -z 연산자와는 달리, -n 연산자 문자열 길이가 0보다 크면 True를 리턴하고, 문자열 길이가 0이면 False를 리턴한다
# 아래 예제를 통해 파라미터가 입력되었으면 True를 출력하고, 파라미터가 없으면 False를 출력하는 것을 확인할 수 있다

if [ -n "$1" ]
then
  echo True
else
  echo False
fi

# 출력 결과
# $ sh operator_if_example2.sh 
# False

# $ sh operator_if_example2.sh test
# True

