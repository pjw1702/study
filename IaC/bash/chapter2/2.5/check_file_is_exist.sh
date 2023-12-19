#!/bin/bash

# 파일 연산자 중 -f 연산자를 이용하면 변수를 통해 입력한 값이 파일인지를 확인할 수 있다
# -L 연산자를 이용하면 해당 파일이 심볼릭 링크인지 아닌지를 확인할 수 있다
# 아래 예제는 /etc/localtime이라는 심볼릭 링크가 파일인지 여부와 심볼릭 링크인지 여부를 체크하는 예시이다

# ls를 이용하여 /etc/localtime의 파일 속성 확인
# ls -al /etc/localtime
# lrwxrwxrwx. 1 root root 32 Aug 25 13:53 /etc/localtime -> ../usr/share/zoneinfo/Asia/Seoul

# /etc/localtime 원파일의 속성 확인
# ls -l /usr/share/zoneinfo/Asia/Seoul 
# -rw-r--r--. 2 root root 645 Apr 30  2020 /usr/share/zoneinfo/Asia/Seoul

FILE=/etc/localtime

# 파일이면 True, 그렇지 않으면 False 출력
# if [ -f $FILE ]; then echo True; else False; fi
if [ -f $FILE ]
then
  echo True
else
  echo False
fi
# 출력 결과
# True

# 파일이면 True, 그렇지 않으면 False 출력
# if [ -L $FILE ]; then echo True; else False; fi
if [ -L $FILE ]
then
  echo True
else
  echo False
fi
# 출력 결과
# True