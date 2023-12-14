#!/bin/bash

# $*과 $@를 for문과 함께 사용할 때 큰 따옴표 ""를 앞 뒤로 사용했을 경우와 그렇지 않을 경우의 차이점을 확실하게 알 수 있다

# $*을 이용해 for문을 사용하면 다음과 출력과 같이 큰 따옴표 ""와 상관없이 스페이스를 기준으로 문자열을 파라미터로 인식한다
# 실행: sh mylanguage_2.sh Korean English "Japanese Chinese"
for language in $*
do
    echo "I can speak $language"
done

# 출력 결과
# I can speak Korean
# I can speak English 
# I can speak Japanese
# I can speak Chinese 

# $@를 사용했을 때도 $*를 사용했을 때와 같이 결과값이 동일함을 다음 출력에서 알 수 있다
for language in $@
do
    echo "I can speak $language"
done

# 출력 결과
# I can speak Korean
# I can speak English 
# I can speak Japanese
# I can speak Chinese 

# 큰 따옴표와 함께 $*을 사용하면 매개변수를 개별로 인식하는 것이 아니라 하나의 문자열로 인식한다
for language in "$*"
do
    echo "I can speak $language"
done

# 출력 결과
# I can speak Korean English Japanese Chinese

# 큰 따옴표와 함께 $@를 사용하면 큰 따옴표 사이의 문자열 인자를 하나의 매개변수로 인식하고 있음을 알 수 있다
for language in "$@"
do
    echo "I can speak $language"
done

# 출력 결과
# I can speak Korean
# I can speak English
# I can speak Japanese Chinese


