#!/bin/bash

# 파일 연산자는 파일이 가지고 있는 다양한 속성들을 체크하는 연산자이다
# 예를 들어 파일에 읽기 권한이 있는지, 쓰기 권한이 있는지 등을 파일 연산자를 통해 확인할 수 있다

# 연산자      # 사용법                      # 설명
-f           if[ -f $변수 ]                변수 유형이 파일이면 참
-L           if[ -L $변수 ]                변수 유형이 파일이면서 심볼릭 링크이면 참
-r           if[ -r $변수 ]                변수 유형이 파일이거나 디렉토리이면서 읽기 권한이 있으면 참
-w           if[ -w $변수 ]                변수 유형이 파일이거나 디렉토리이면서 쓰기 권한이 있으면 참
-x           if[ -x $변수 ]                변수 유형이 파일이거나 디렉토리이면서 실행 권한이 있으면 참
-s           if[ -s $변수 ]                변수 유형이 파일이거나 디렉토리이면서 사이즈가 0보다 크면 참
-O           if[ -O $변수 ]                변수 유형이 파일이거나 디렉토리이면서 스크립트 실행 소유자와 동일하면 참
-G           if[ -G $변수 ]                변수 유형이 파일이거나 디렉토리이면서 스크립트 실행 그룹과 동일하면 참