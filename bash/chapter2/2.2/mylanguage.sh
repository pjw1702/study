#!/bin/bash

# sh mylanguage.sh Korean English 를 실행 시, $0 = mylanguage.sh, $1 = Korean, $2 = English 가 된다
echo "This is shell script name is $0"
echo "I can speak $1 and $2"

echo "This shell script parameter are $*"
echo "This shell script parameter are $@"
echo "This parameter count is $#"

# 출력 결과
# $ sh mylanguage.sh Korean English
# This is shell script name is mylanguage.sh
# I can speak Korean and English
# This shell script parameter are Korean English      /github/study/bash/chapter2 
# This shell script parameter are Korean English
# This parameter count is 2