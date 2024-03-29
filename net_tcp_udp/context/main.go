package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main(){
	wg := sync.WaitGroup{}
	ctx := context.Background()

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	randTimer := time.Duration(rnd.Intn(5000)) * time.Microsecond

	fmt.Printf("random time main : %s\n", randTimer)

	ctx, cancel := context.WithTimeout(ctx,time.Second*60)
	wg.Add(1)

	go dealLongWithCtx(&wg,ctx)
	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}

func dealLongWithCtx(wg * sync.WaitGroup, ctx context.Context){
	defer wg.Done()

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	randTimer := time.Duration(rnd.Intn(5000)) * time.Microsecond

	fmt.Printf("random time : %s\n", randTimer)

	timer := time.NewTimer(randTimer)
	select{
		case <-ctx.Done(): {
			fmt.Println("Rejected")
		}
		case <-timer.C: {
			fmt.Println("Done")
		}
	}
}