package contextlearn

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
1. context is immutable
2. children can't cancel context
3. in main we get cancel function and context
4. we passed children function the context
5. context provide a way to transport request scope data through our call graph and it is controversial because it can be easily abused because it can accept any kind of type
6. withcancel , withdeadline , withtimeout all returns cancel function we can set a deadline for the cancel in same way we can set timeout for cancellation
*/

func func1(ctx context.Context, parentWg *sync.WaitGroup, stream <-chan interface{}) {
	defer parentWg.Done()
	var wg sync.WaitGroup

	doWork := func(ctx context.Context){
		defer wg.Done()
		for {
			select {
			case <- ctx.Done():
				return
			case d, ok:=<-stream:
				if !ok {
					fmt.Println("channel closed")
					return
				}
				fmt.Println("Data..",d)
			}
		}
	}

	newCtx, cancel := context.WithTimeout(ctx,time.Second*3)
	defer cancel()

	for i:= 0 ; i<3;i++{
		wg.Add(1)
		go doWork(newCtx)
	}
	wg.Wait()
}

func genericFunc(ctx context.Context, wg *sync.WaitGroup, stream<-chan interface{}){
	defer wg.Done()
	for {
		select{
		case <- ctx.Done():
			return 
		case d,ok := <-stream:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println("Data...",d)
		}
	}
}

func Context() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	generator := func(dataItem string, stream chan interface{}) {
		for {
			select {
			case <-ctx.Done():
				return
			case stream <- dataItem:
			}
		}
	}

	infiniteApples := make(chan interface{})
	go generator("apples", infiniteApples)

	infiniteOranges := make(chan interface{})
	go generator("oranges", infiniteOranges)

	infinitePeaches := make(chan interface{})
	go generator("peaches", infinitePeaches)

	wg.Add(1)
	go func1(ctx, &wg, infiniteApples)

	func2 := genericFunc
	func3 := genericFunc

	wg.Add(1)
	go func2(ctx,&wg,infiniteOranges)
	wg.Add(1)
	go func3(ctx,&wg,infinitePeaches)

	wg.Wait()
	
}
