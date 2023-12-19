 #!/bin/bash
 # 문자열을 비교하는 경우
 # 아래 예제는 두 개의 변수에 저장된 값이 문자열일 경우 문자열을 비교하는 예제이다
 # 문자열의 대소를 비교할 때의 연산자는 =, ==, != 이므로 일반 if문 사용법과 동일하게 사용한다
 # <, > 은 리다이렉션 기호와 동일한 연산자를 사용할 경우에는 중첩 대괄호[[]]를 사용하여 리다이렉션 기호가 아닌 비교 연산자임을 시스템에 알려주어야 한다 

 VAR1="abc"
 VAR2="def"

# if [ $VAR1 = $VAR2 ]; then echo True; else False; fi
 if [ $VAR1 = $VAR2 ]
 then
  echo True
else
  echo False
fi

# if [ $VAR1 == $VAR2 ]; then echo True; else False; fi
 if [ $VAR1 == $VAR2 ]
 then
  echo True
else
  echo False
fi

# if [ $VAR1 != $VAR2 ]; then echo True; else False; fi
 if [ $VAR1 != $VAR2 ]
 then
  echo True
else
  echo False
fi

# if [[ $VAR1 > $VAR2 ]]; then echo True; else False; fi
 if [[ $VAR1 > $VAR2 ]]
 then
  echo True
else
  echo False
fi

# if [[ $VAR1 < $VAR2 ]]; then echo True; else False; fi
 if [[ $VAR1 < $VAR2 ]]
 then
  echo True
else
  echo False
fi


# 출력 결과
# False
# False
# True
# True
# True