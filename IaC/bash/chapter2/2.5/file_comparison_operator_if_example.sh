#!/bin/bash

# 두 개의 파일을 생성한 후 -nt 연산자를 이용하여 어떤파일이 최신 파일인지를 확인할 수 있다
# -ot 연산자를 사용하여 생성한지 오래된 파일인지를 확인할 수 있댜
# -ef 연산자를 사용하여 결과값이 True가 되려면, 심볼릭 링크로 연결된 두 개의 파일을 비교하면 이 둘은 동일파일이므로 True 결과값을 얻을 수 있다

# 서로 다른 문자열을 각각의 파일로 생성
# echo "AAA" > first.txt
# echo "BBB" > second.txt

FILE1=first.txt
FILE2=second.txt

# FIL1이 FILE2보다 최신 파일이라면 True, 그렇지 않으면 False
# if [ $FILE1 -nt $FILE2 ]; then echo True; else echo False; fi
if [ $FILE1 -nt $FILE2 ]
then
  echo True
else
  echo False
fi
# 출력 결과
# False

# FIL1이 FILE2보다 예전 파일이라면 True, 그렇지 않으면 False
# if [ $FILE1 -ot $FILE2 ]; then echo True; else echo False; fi
if [ $FILE1 -ot $FILE2 ]
then
  echo True
else
  echo False
fi
# 출력 결과
# True

FILE1=/etc/localtime
FILE2=/usr/share/zoneinfo/Asia/Seoul

# FIL1과 FILE2가 동일한 파일이라면 True, 그렇지 않으면 False
# if [ $FILE1 -nt $FILE2 ]; then echo True; else echo False; fi
if [ $FILE1 -ef $FILE2 ]
then
  echo True
else
  echo False
fi
# 출력 결과
# False