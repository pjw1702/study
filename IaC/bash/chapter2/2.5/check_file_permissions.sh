#!/bin/bash

# 파일 권한을 체크하는 경우
# -r 연산자는 파일에 읽기 권한이 있는지 여부를 체크한다
# -w 연산자는 쓰기 권한이 있는지 여부를 체크한다
# -x 연산자는 실행 권한이 있는지 여부를 체크한다
# -s 연산자는 파일의 크기를 체크하여 해당 파일이 빈 파일인지 여부를 확인한다

# ls를 이용하여 /etc/localtime의 파일 속성 확인
# ls -al /etc/localtime
# lrwxrwxrwx. 1 root root 32 Aug 25 13:53 /etc/localtime -> ../usr/share/zoneinfo/Asia/Seoul

# /etc/localtime 원파일의 속성 확인
# ls -l /usr/share/zoneinfo/Asia/Seoul 
# -rw-r--r--. 2 root root 645 Apr 30  2020 /usr/share/zoneinfo/Asia/Seoul

FILE=/etc/localtime

# 파일에 읽기 권한이 있으면 True, 그렇지 않으면 False
# if [ -r $FILE ]; then echo True; else echo False; fi
if [ -r $FILE ]
then
  echo True
else
  echo False
fi

# 파일에 읽기 권한이 있으면 True, 그렇지 않으면 False
# if [ -r $FILE ]; then echo True; else echo False; fi
if [ -r $FILE ]
then
  echo True
else
  echo False
fi

# 파일에 읽기 권한이 있으면 True, 그렇지 않으면 False
# if [ -r $FILE ]; then echo True; else echo False; fi
if [ -r $FILE ]
then
  echo True
else
  echo False
fi
# 출력 결과
# True

# 파일에 쓰기 권한이 있으면 True, 그렇지 않으면 False
# if [ -w $FILE ]; then echo True; else echo False; fi
if [ -w $FILE ]
then
  echo True
else
  echo False
fi
# 출력 결과
# True

# 파일에 실행 권한이 있으면 True, 그렇지 않으면 False
# if [ -x $FILE ]; then echo True; else echo False; fi
if [ -x $FILE ]
then
  echo True
else
  echo False
fi
# 출력 결과
# True

# 파일의 크기가 0보다 크면 True, 그렇지 않으면 False
# if [ -s $FILE ]; then echo True; else echo False; fi
if [ -s $FILE ]
then
  echo True
else
  echo False
fi
# 출력 결과
# True