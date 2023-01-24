// 영역을 나누는 방법으로 데드락 문제를 해결할 수 있다
// 포인터를 이용해 객체에 접근하여 각자 서로 다른 메모리 자원에 접근하여 고루틴을 실행한다
// Go 쓰레드 내부는 Local rum Queue가 있다: 현재 진행되고 있는 고루틴이 빠져나가면 큐에 의해 대기하고 있는 고루틴이 순차적으로 빠져나간다
// 역할을 나누는 방법은 채널(Channel)을 통해 알 수 있다
package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과 %d\n", j.index, j.index*j.index)
}

func main() {
	// 작업(Job) 10개 할당
	var jobList [10]Job
	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}

/* 출력 결과
9 작업 시작
4 작업 시작
1 작업 시작
0 작업 시작
2 작업 시작
7 작업 시작
6 작업 시작
5 작업 시작
8 작업 시작
3 작업 시작
3 작업 완료 - 결과 9
4 작업 완료 - 결과 16
9 작업 완료 - 결과 81
0 작업 완료 - 결과 0
6 작업 완료 - 결과 36
2 작업 완료 - 결과 4
5 작업 완료 - 결과 25
7 작업 완료 - 결과 49
8 작업 완료 - 결과 64
1 작업 완료 - 결과 1
*/
