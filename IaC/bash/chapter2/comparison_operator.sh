#!/bin/bash

# 비교 연산자는 두 개의 변수에 저장된 값의 크기를 비교할 때 사용한다
# 변수의 유형은 1 2 3 과 같은 숫자형과 세미콜론''이나 콜론""을 함께 사용할 문자형으로 나눌 수 있다
# 쉘 스크립트에서는 위 와 같은 두 변수 유형의 비교 연산자를 다르게 표현한다

# 정수 비교 연산자
# 일반적인 프로그래밍 언어에서 정수를 비교할 때는 =, !=, >, >=, <, <= 와 같은 기호 연산자를 사용하지만, 쉘 스크립트에서는 이 외에도 아래와 같은 영문 약자를 사용한 문자형 연산자를 사용한다
# 리다이렉션 기호와의 구분을 위해 중첩 소괄호(())를 사용하여 리다이렉션 기호가 아닌 비교 연산자임을 시스템에 알려주어야 한다

# 연산자      # 사용법                      # 설명
-eq           if[ $변수1 -eq $변수2 ]       변수1과 변수2의 값이 같으면 참 (=과 동일)
-ne           if[ $변수2 -ne $변수2 ]       변수1과 변수2의 값이 다르면 참 (!=과 동일)
-gt           if[ $변수2 -gt $변수2 ]       변수1의 값이 변수2의 값보다 크면 참
-ge           if[ $변수2 -ge $변수2 ]       변수1의 값이 변수2의 값보다 크거나 같으면 참
-lt           if[ $변수2 -lt $변수2 ]       변수1의 값이 변수2의 값보다 작으면 참
-le           if[ $변수2 -le $변수2 ]       변수1의 값이 변수2의 값보다 작거나 같으면 참
>             if (( $변수1 > $변수2 ))      변수1의 값이 변수2보다 크면 참이며 > 기호를 사용할 경우에는 중첩소괄호(())를 사용해야 함
>=            if (( $변수1 >= $변수2 ))     변수1의 값이 변수2보다 크거나 같으면 참이며 >= 기호를 사용할 경우에는 중첩소괄호(())를 사용해야 함
<             if (( $변수1 < $변수2 ))      변수1의 값이 변수2보다 작으면 참이며 < 기호를 사용할 경우에는 중첩소괄호(())를 사용해야 함
<=            if (( $변수1 <= $변수2 ))     변수1의 값이 변수2보다 작거나 같으면 참이며 <= 기호를 사용할 경우에는 중첩소괄호(())를 사용해야 함