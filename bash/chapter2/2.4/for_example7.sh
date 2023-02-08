#!/bin/bash

# 기본 사용법 2는 ((초기값; 조건식; 증가값))으로 범위를 표현하며 C나 Java 문법과 비슷하다

for ((num=0; num<3; num++))
do
    echo $num;
done

# 출력 결과
# $ ./for_example7.sh
# 0
# 1
# 2