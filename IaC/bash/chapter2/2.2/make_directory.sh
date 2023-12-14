#!/bin/bash

# 쉘 스크립트에서 변수는 특별한 타입을 요구하지 않는다
language="Korea English Japan"

# 문자열 변수를 이용해 디렉토리를 여러 개 생성할 수 있다
mkdir $language

# 출력 결과
# $ ls -al
# total 6
# drwxr-xr-x 1 박정우 197121  0  1월 26 10:45 ./      
# drwxr-xr-x 1 박정우 197121  0  1월 26 10:35 ../     
# drwxr-xr-x 1 박정우 197121  0  1월 26 10:45 English/
# drwxr-xr-x 1 박정우 197121  0  1월 26 10:45 Japan/  
# drwxr-xr-x 1 박정우 197121  0  1월 26 10:45 Korea/  
# -rwxr-xr-x 1 박정우 197121 60  1월 26 10:45 make_directory.sh*
# -rwxr-xr-x 1 박정우 197121 60  1월 26 10:43 myshell_use_variable.sh