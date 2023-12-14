#!/bin/bash

# 전역변수
# 스크립트 전체에서 변수에 저장한 값을 사용할 수 있는 변수
# 함수 밖에서 선언된 변수는 함수 내에서도 그 값이 유효하다
language="Korean"

function print() {
    echo "I can speak $language"
}

print