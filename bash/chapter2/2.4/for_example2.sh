#!/bin/bash

# 변수를 지정함으로써 범위를 수정할 수 있다

numbers="1 2 3"

for num in $numbers
do
    echo $num;
done

# 출력 결과
# $ ./for_example2.sh 
# 1
# 2
# 3