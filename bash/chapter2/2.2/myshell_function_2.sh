#!/bin/bash

language="Korean"

function learn() {
    # 지역변수 선언
    local learn_language="English"
    echo "I am learning $learn_language"
}

function print() {
    # 전달받은 인자들 중 첫 번째 인자
    echo "I can speak $1"
}

learn

print $language
print $learn_language