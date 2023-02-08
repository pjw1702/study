#!/bin/bash

# 범위를 배열로 사용할 경우에는 배열 선언 시 값과 값 사이에 쉼표 ,를 사용하면 안된다
# for문에 배열의 모든 아이템을 범위로 사용할 경우에는 ${배열[@]}을 사용하여 배열의 모든 아이템을 사용한다고 명시해 주어야 한다
# 위치 매개변수 $@를 사용하면 파라미터로 넘어오는 모든 매개변수를 의미하는 것과 동일한 의미이다

array=("apple" "banana" "pineapple")

for fruit in ${array[@]}
do
 echo $fruit;
done

# 출력 결과
# $ ./for_example6.sh 
# apple
# banana   
# pineapple