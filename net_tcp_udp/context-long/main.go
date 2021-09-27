package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func emulateLongOperation(ctx context.Context, id int){
	randVal := rand.Intn(5000)
	randTime := time.Duration(randVal) * time.Microsecond
	log.Printf("Id %d will be evalutated for %d", id, randVal)
	timer := time.NewTimer(randTime)

	select {
		case <-timer.C: {
			fmt.Printf("Successfully finished job %d\n", id)
		}
		case <-ctx.Done() : {
			fmt.Printf("id %d timed out", id)
		}
	}
}

func main(){
	wg := sync.WaitGroup{}
	ctx, _ := context.WithTimeout(context.Background(), 2000 *time.Millisecond)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int){
			emulateLongOperation(ctx, id)
			wg.Done()
		}(i)
	}

	for i:= 0 ; i < 10; i++ {
		wg.Wait()
	}

}
