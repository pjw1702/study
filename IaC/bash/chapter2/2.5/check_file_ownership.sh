#!/bin/bash

# 파일의 소유권을 체크하는 경우
# -O 연산자와 -G 연산자를 이용하여 스크립트가 수행되는 프롬프트의 소유자 및 그룹의 변수에 정의된 파일의 소유자 및 그룹과 동일한지 여부를 확인할 수 있다

# ls를 이용하여 /etc/localtime의 파일 속성 확인
# ls -al /etc/localtime
# lrwxrwxrwx. 1 root root 32 Aug 25 13:53 /etc/localtime -> ../usr/share/zoneinfo/Asia/Seoul

# /etc/localtime 원파일의 속성 확인
# ls -l /usr/share/zoneinfo/Asia/Seoul 
# -rw-r--r--. 2 root root 645 Apr 30  2020 /usr/share/zoneinfo/Asia/Seoul

FILE=/etc/localtime

# user1 계정으로 실행하는 경우
# if문을 실행하는 소유자와 파일의 소유자가 다르므로 False 출력
# [user1@localhost ~]$ if [ -O $FILE ]; then echo True; else echo False; fi
if [ -O $FILE ]
then
  echo True
else
  echo False
fi
# 실행 결과
# False

# user1 계정으로 실행하는 경우
# if문을 실행하는 소유자와 파일의 그룹이 다르므로 False 출력
# [user1@localhost ~]$ if [ -O $FILE ]; then echo True; else echo False; fi
if [ -G $FILE ]
then
  echo True
else
  echo False
fi
# 실행 결과
# False

# root 계정으로 실행하는 경우
# if문을 실행하는 소유자와 파일의 소유자가 같으므로 False 출력
# [root@localhost ~]# if [ -O $FILE ]; then echo True; else echo False; fi
if [ -O $FILE ]
then
  echo True
else
  echo False
fi
# 실행 결과
# True

# root 계정으로 실행하는 경우
# if문을 실행하는 소유자와 파일의 그룹이 같으므로 False 출력
# [root@localhost ~]# if [ -O $FILE ]; then echo True; else echo False; fi
if [ -G $FILE ]
then
  echo True
else
  echo False
fi
# 실행 결과
# True