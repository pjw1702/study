# ADT

> 단순 정수가 아닌 구조체 변수를 비롯하여 각종 데이터들을 저장하기 위한 배열 기반의 리스트이다

<br>

### void SetPointPos(Point * ppos, int xpos, int ypos);
- Point 변수의 xpos, ypos 값 설정

<br>

### void ShowPointPos(Point * ppos);
- Point 변수의 xpos, ypos 정보 출력

<br>

### int PointComp(Point * pos1, Point * pos2);
- 두 Point 변수의 비교

<br>
<br>

## 실행을 위해 필요한 파일들

### Point.h
- 리스트 자료구조의 구조체 변수 데이터 관리 헤더파일
### ArrayList.h
- 리스트 자료구조의 헤더파일
### Point.c
- 리스트 자료구조의 구조체 변수 데이터 관리 소스파일
### PointListMain.c 
- 구조체 변수 데이터 관리 리스트 관련 main 함수가 담긴 소스파일