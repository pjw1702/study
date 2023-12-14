#!/bin/bash

# 조건식 타입 - if [ 조건식 ] 연산자 [ 조건식 ]; then
# 여러 개의 조건식을 AND나 OR과 같은 논리 연산자에 의해 복합 연산을 한 경우에 사용된다
# 연산자 -gt는 A가 B보다 큰지를 비교하는 연산자이며, -lt는 A가 B보다 작은지를 비교하는 연산자이다

value=5
# if [ $value -gt 0 ] && [ $value -lt 10 ]; then echo True; else echo False; fi 
if [ $value -gt 0 ] && [ $value -lt 10 ]
then
    echo True
else
    echo False
fi

# 출력 결과
# $ ./if_example4.sh
# True