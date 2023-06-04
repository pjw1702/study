package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil { // Value(): "number"라는 키의 value를 9로 지정한 데이터를 가져올 수 있다
		n := v.(int) // v를 integer 타입으로 변환
		fmt.Printf("Square: %d", n*n)
	}
	wg.Done()
}

func main() {
	wg.Add(1)

	ctx := context.WithValue(context.Background(), "number", 9) // WithValue(): 어떠한 데이터를 지정할 수 있다 ("number"라는 키의 value를 9로 지정)
	go square(ctx)

	wg.Wait()
}
