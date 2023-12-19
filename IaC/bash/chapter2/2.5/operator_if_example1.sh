#!/bin/bash
# 문자열 변수가 NULL 값인지 체크할 경우
# 변수가 초기화 되었는지 그렇지 않은지를 체크하면 연산자 -z를 매우 유용하게 사용할 수 있다
# 아래는 스크립트를 실행할 때 파라미터 인자가 입력되었는지 여부를 체크하는 예제이다
# -z 연산자는 문자열 길이가 0일 때 True를 리턴한다
# 따라서 파라미터의 입력이 없으면 True를 출력하고, 파라미터 입력이 있으면 False를 출력한다

if [ -z $1 ]
then
  echo True
else
  echo False
fi

# 출력 결과
# $ sh operator_if_example1.sh 
# True

# sh operator_if_example1.sh test
# False