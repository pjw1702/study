#!/bin/bash

# 변수의 값이 문자열로 선언되었을 경우 패턴을 통해 문자열을 변경할 경우에 사용할 수 있는 확장자들이다

#확장자                                 # 설명
${변수#*패턴}                            # 변수에 설정된 문자열 앞에서 부터 처음 찾은 패턴과 일치하는 패턴 앞의 모든 문자열 제거   
${변수##*패턴}                           # 변수에 설정된 문자열 뒤에서 부터 마지막으로 찾은 패턴과 일치하는 패턴 앞의 모든 문자열 제거
${변수%패턴*}                            # 변수에 설정된 문자열 뒤에서 처음 찾은 패턴과 일치하는 패턴 뒤의 모든 문자열 제거(뒤에서부터 처음 찾은 .뒤의 모든 문자열 제거)
${변수%%패턴*}                           # 변수에 설정된 문자열 앞에서부터 마지막으로 찾은 패턴과 일치하는 패턴 뒤의 모든 문자열 제거(뒤에서부터 마지막으로 찾은 . 뒤의 모든 문자열 제거)
${#변수}                                # 변수에 저장된 문자열의 길이 리턴
${변수/찾을문자열/바꿀문자열}             # 변수에 설정된 문자열에서 첫 번째 리턴에 해당하는 부분을 문자열로 변경, 문자열을 지정하지 않으면 해당 문자열을 제거  
${변수/#찾을문자열/바꿀문자열}            # 변수에 설정된 문자열의 시작 문자열이 패턴과 맞는 경우 문자열로 변경
${변수/%찾을문자열/바꿀문자열}            # 변수에 설정된 문자열 마지막 문자열이 패턴과 맞는 경우 문자열로 변경

# 출력 결과
# $ FILE_NAME="myvm_container-repo.tar.gz"

# $ echo ${FILE_NAME#*_}
# container-repo.tar.gz

# $ echo ${FILE_NAME##*-}
# repo.tar.gz

# $ echo ${FILE_NAME%.*}
# myvm_container-repo.tar

# $ echo ${FILE_NAME%%.*}
# myvm_container-repo

# # 파일명과 파일 경로 추출
# $ FILE_PATH="/etc/nova/nava.conf"

# $ echo ${FILE_PATH%/*}
# /etc/nova

# $ echo ${FILE_PATH##*/}
# nava.conf

# $ echo ${#FILE_PATH}
# 19

# ${변수/찾을문자열/바꿀문자열}과 ${변수//찾을문자열/바꿀문자열}

# 변수에 설정된 문자열에서 특정 문자열을 찾아 다른 문자열로 치환하여 리턴해 준다
# ${변수/찾을문자열/바꿀문자열}을 사용하면 말 그대로 변수에 설정된 문자열 앞에서부터 처음으로 찾은 문자열을 바꿀문자열로 바꿔준다
# ${변수//찾을문자열/바꿀문자열}을 이용하면 문자열 전체에서 해당 문자열을 찾아 바꿀 문자열로 모두 바꿔준다
# ${변수/#찾을문자열/바꿀문자열}을 사용하면 찾을 문자열로 시작하는 문자열을 변경할 수 있다
# ${변수/%찾을문자열/바꿀문자열}을 사용하면 찾을 문자열로 끝나는 문자열을 변경할 수 있다

# ${변수/찾을문자열}을 사용하면 변수에 설정된 문자열 앞에서부터 처음으로 찾은 문자열을 삭제해준다
# ${변수//찾을문자열}을 사용하면 변수에 설정된 문자열 처음부터 찾은 모든 해당 문자열을 삭제해준다

# 출력 결과
# $ OS_TYPE="Redhat Linux Ubuntu Linux Fedora Linux"

# $ echo ${OS_TYPE/Linux/OS}
# Redhat OS Ubuntu Linux Fedora Linux

# $ echo ${OS_TYPE//Linux/OS}
# Redhat OS Ubuntu OS Fedora OS

# $ echo ${OS_TYPE/Linux}
# Redhat Ubuntu Linux Fedora Linux

# $ echo ${OS_TYPE//Linux}
# Redhat Ubuntu Fedora

# $ echo ${OS_TYPE/#Redhat/Unknown}
# Unknown Linux Ubuntu Linux Fedora Linux

# $ echo ${OS_TYPE/%Linux/Unknown}
# Redhat Linux Ubuntu Linux Fedora Unknown