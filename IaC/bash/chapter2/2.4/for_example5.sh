#!/bin/bash

# 중괄호를 사용할 경우 증가값이 1이 아니라 그 이상일 경우 다음 예제와 같이 {초기값..최종값..증가값}으로 표현할 수 있다

for num in {1..10..2}
do
    echo $num;
done

# 출력 결과
# $ ./for_example5.sh
# 1
# 3
# 5
# 7
# 9