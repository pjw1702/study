#!/bin/bash

# 기본 사용법
# 조건식 앞 뒤로는 반드시 대괄호[]를 사용해야 하며, 대괄호[]와 조건식 사이에는 반드시 한 칸 스페이스를 두어야 한다

if [ 첫 번째 조건식 ]
then
    수행문
elif [ 두 번째 조건식 ]
then
    수행문
else
    수행문
fi

# 조건식 타입
# 조건식은 변수의 타입이 숫자나 문자열 또는 파일과 같은 객체형이냐에 따라 사용되는 타입에 차이가 있으며,
# 연산의 종류에 따라서도 차이가 있을 수 있다

# 조건식 타입
if [ $변수 연산자 $변수 ]; then                 # 일반적인 조건식 타입으로, 두 변수의 값을 비교할 때 쓰임
if [ $변수 연산자 조건값 ]; then                # 조건값이 고정되어 있을 경우 변수와 조건값을 비교할 때 사용
if [ 연산자 $변수 ]; then                       # 변수의 값이 문자열이거나 디렉토리와 같은 경우일 때 주로 사용
if [ 조건식 ]  연산자 [ 조건식 ]; then           # 여러 개의 조건식을 AND나 OR로 복합 연산할 때 사용