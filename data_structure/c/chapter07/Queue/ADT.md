# ADT
> FIFO 구조의 자료구조인 큐(Queue)의 ADT이다
<br>
> ADT를 대상으로 배열 기반의 큐 또는 연결 리스트 기반의 큐를 구현할 수 있다 

<br>

### QueueInit(Queue * pq);
- 큐의 초기화를 진행한다
- 큐 생성 후 제일 먼저 호출되어야 하는 함수이다

<br>

### int QIsEmpty(Queue * pq);
- 큐가 빈 경우 TRUE(1)을, 그렇지 않은 경우 FALSE(0)을 반환한다

<br>

### void Enqueue(Queue * pq, Data data);
- 큐의 기본 연산인, enqueue를 구현하는 함수로, 큐에 데이터를 삽입하는 연산이다
- 큐에 데이터를 저장한다
- 매개변수 data로 전달된 값을 저장한다

<br>

### Data Dequeue(Queue * pq);
- 큐의 기본 연산인, dequeue를 구현하는 함수로, 큐에서 데이터를 꺼내는 연산이다
- 저장순서가 가장 앞선 데이터를 삭제한다
- 삭제된 데이터는 반환된다
- 본 함수의 호출을 위해서는 데이터가 하나 이상 존재함이 보장되어야 한다

<br>

### Data QPeek(Queue * pq);
- 저장 순서가 가장 앞선 데이터를 반환하되 삭제하지 않는다
- 본 함수의 호출을 위해서는 데이터가 하나 이상 존재함이 보장되어야 한다