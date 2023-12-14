#!/bin/bash

# 조건식 타입 - if [ $변수 연산자 $변수 ]; then
# if문을 사용할 때 가장 일반적으로 사용하는 조건식 타입으로, 두 변수의 값을 비교할 때 주로 사용된다

# 쉘 스크립트에서 세미콜론(;)을 사용한다는 것은 문법이나 명령어 또는 구문이 완료되어 다음 줄로 넘길 경우에 세미콜론을 사용한다

value1=10
value2=10
# 한줄로 사용할 경우에는 if [ $value = $value ]; then으로 표현할 수 있음
if [ $value1 = $value2 ]
then
    echo True
else
    echo False
fi

# 출력 결과
# $ ./if_example1.sh 
# True