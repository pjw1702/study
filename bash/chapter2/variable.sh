#!/bin/bash

$variable       #현재 문자열에서 해당 변수를 파라미터 값으로 치환
${variable}     # 위와 같지만 {}를 사용함으로써 뒤에 오는 문자열과 구분 가능

# 실행 및 출력 결과

# 변수 AUTH_URL에 "www.example.com/"을 저장
# $ AUTH_URL="www.examle.com/"

# 다음과 같은 경우 어디까지가 변수 명인지 알 수 없어, 변수명을 AUTH_URLlogin으로 인식함
# echo "http://$AUTH_URLlogin.html"
# http://.html

# {}를 사용하여 $AUTH_URL이 변수인지를 구분할 수 있음
# $ echo "http://${AUTH_URL}login.html"
# http://www.examle.com/login.html