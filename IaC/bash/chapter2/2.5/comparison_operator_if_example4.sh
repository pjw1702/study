#!/bin/bash
# 변수값의 크기를 비교하는 경우: 기호 연산자를 사용하는 경우
# 리눅스에서는 echo와 같이 터미널에 값을 출력하거나 <, >, >>, | 와 같이 터미널에 출력될 값을 파일에 저장해 주는 리다이렉션(redirection) 기호들이 있다
# 하지만 비교 연산자에도 <>와 같은 기호들이 있다
# 조건문의 기본 기호인 대괄호[]를 사용한 조건문으로 작성하면 사용한 기호가 라다이렉션인지 비교 연산자인지 구분하기가 힘들다
# 따라서 비교 연산자 기호를 이용하여 두 변수의 크기를 비교할 때는 중첩 소괄호 (())를 사용하여 아래와 같이 표현한다

VAR1=10
VAR2=20

# if (( $VAR1 > $VAR2 )); then echo True; else echo False; fi
if (( $VAR1 > $VAR2 ))
then
  echo True
else
  echo False
fi

# if (( $VAR1 >= $VAR2 )); then echo True; else echo False; fi
if (( $VAR1 >= $VAR2 ))
then
  echo True
else
  echo False
fi

# if (( $VAR1 < $VAR2 )); then echo True; else echo False; fi
if (( $VAR1 < $VAR2 ))
then
  echo True
else
  echo False
fi

# if (( $VAR1 <= $VAR2 )); then echo True; else echo False; fi
if (( $VAR1 <= $VAR2 ))
then
  echo True
else
  echo False
fi

# 출력결과
# False
# False
# True
# True