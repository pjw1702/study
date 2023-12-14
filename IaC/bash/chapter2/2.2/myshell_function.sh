#!/bin/bash

# 변수의 종류
# 변수에는 단순한게 선언만 해서 쓰는 변수도 있지만, 함수를 함께 사용하면서 함수 내에서만 사용할 수 있는 변수, 함수 밖에서도 사용할 수 있는 변수, 함수에 파라미털 변수를 넘길 때 사용하는 변수 등 여러 종류들의 변수들이 많다
# 시스템에서 시스템을 위해 미리 할당된 변수: 예약변수, 환경변수

function print() {
    # 전달된 "I can speak Korean" 문자열을 넘겨받아 출력
    echo $1
}

# print() 함수에 "I can speak Korean" 문자열을 함수의 인자로 전달
print "I can speak Korean"