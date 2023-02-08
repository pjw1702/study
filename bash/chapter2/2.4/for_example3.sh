#!/bin/bash

# 숫자나 문자열 이외에 디렉토리나 파일과 같은 객체를 범위로 사용할 수 있다

for file in $HOME/*
do
    echo $file;
done

# 출력 결과
# $ ./for_example3.sh 
# /c/Users/박정우/Apple
# /c/Users/박정우/Application Data
# /c/Users/박정우/Contacts
# /c/Users/박정우/Cookies
# /c/Users/박정우/Desktop
# /c/Users/박정우/Documents
# /c/Users/박정우/Downloads
# /c/Users/박정우/eclipse-workspace
# /c/Users/박정우/Favorites
# /c/Users/박정우/git
# /c/Users/박정우/go
# /c/Users/박정우/IntelGraphicsProfiles
# /c/Users/박정우/Links
# /c/Users/박정우/Local Settings
# /c/Users/박정우/MicrosoftEdgeBackups
# ...