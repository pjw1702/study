#!/bin/bash

# 환경변수 HOME을 이용하여 HOME의 값이 디렉토리인지
# 디렉토리라면 해당 디렉토리 내에 또 다른 디렉토리나 파일이 있는지를 -d 연산자와 -e 연산자를 이용하여 체크할 수 있다

# if [ -d $HOME ]; then echo True; else echo False; fi
if [ -d $HOME ]
then
  echo True
else
  echo False
fi

# if [ -d $HOME ]; then echo True; else echo False; fi
if [ -e $HOME ]
then
  echo True
else
  echo False
fi