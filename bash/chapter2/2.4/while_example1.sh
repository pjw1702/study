#!/bin/bash

num=0

while [ $num -lt 3 ]
do
    echo $num
    num=$((num+1)) # while문을 사용할 경우에는 변수가 조건에 맞도록 증가식을 넣어주어야 함
done

# 출력 결과
# $ ./while_example1.sh 
# 0
# 1
# 2