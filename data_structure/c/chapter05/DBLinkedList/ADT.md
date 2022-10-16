# ADT
> 단순 연결 리스트가 아닌, 양방향 연결 리스트이다
<br>
<br>
단, 더미노드를 활용하지 않았으며 데이터 삭제 기능을 구현하는 함수를 정의하지 않았다

<br>

### void ListInit(List * plist);
- 초기화할 리스트의 주소 값을 인자로 전달한다
- 리스트 생성 후 제일 먼저 호출되어야 하는 함수이다

<br>

### void LInsert(List * plist, LData data);
- 리스트에 데이터를 저장한다
- 매개변수 data에 전달된 값을 저장한다

<br>

### int LFirst(List * plist, LData * pdata);
- 첫 번째 데이터가 pdata가 가리키는 메모리에 저장된다
- 데이터의 참조를 위한 초기화가 진행된다
- 참조 성공 시 TRUE(1), 실패 시 FALSE(0) 반환	

<br>

### int LNext(List * plist, LData * pdata);
- 참조된 데이터의 다음 데이터가 pdata가 가리키는 메모리에 저장된다
- 정방향으로 순차적인 참조를 위해서 반복 호출이 가능하다
- 참조를 새로 시작하려면 먼저 LFirst 함수를 호출해야 한다
- 참조 성공시 TRUE(1), 실패 시 FALSE(0) 반환

<br>

### int LPrevious(List * plist, Data * pdata);
- 참조된 데이터의 다음 데이터가 pdata가 가리키는 메모리에 저장된다
- 역방향으로 순차적인 참조를 위해서 반복 호출이 가능하다
- 참조를 새로 시작하려면 먼저 LFirst 함수를 호출해야 한다
- 참조 성공시 TRUE(1), 실패 시 FALSE(0) 반환

<br>

### int LCount(List * plist);
- 리스트에 저장되어 있는 데이터 수를 반환한다

<br>

## 실행을 위해 필요한 파일들

### DBLinkedList.h
- 더미노드 활용과 데이터 삭제 기능을 구현하지 않은 양방향 연결 리스트 자료구조의 헤더파일

### DBLinkedList.c
- 더미노드 활용과 데이터 삭제 기능을 구현하지 않은 양방향 연결 리스트 자료구조의 기능 정의 소스 파일

### DBLinkedListSortMain.c
- 더미노드 활용과 데이터 삭제 기능을 구현하지 않은 양방향 연결 리스트 관련 main 함수가 담긴 소스파일