package main

import (
	"context"
	"fmt"
	chatpb "github.com/DimKush/go_sandbox/tree/main/grpc/grpc_simple/pkg/chatpb"
	"google.golang.org/grpc"
	"sync"
	"time"
)

const (
	host = "localhost"
	port = ":8082"
	full_address = host + port
)

func main() {
	fmt.Println(runClient(10))
}

type serverAnswer struct {
	 *chatpb.ChatAnswer
	 err error
}

func runClient(requestCounts int) error{
	connect, err := grpc.Dial(full_address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil{
		return err
	}

	defer connect.Close()

	clnt := chatpb.NewGreetingSampleClient(connect)

	//ch := make(chan )
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	defer cancel()

	chAnswer := make(chan *serverAnswer, 10)

	wg := sync.WaitGroup{}
	wg.Add(requestCounts)

	var answers []*serverAnswer
	for i := 0 ; i < requestCounts ; i++ {
		time.Sleep(2 * time.Second)
		go func(i int){
			defer wg.Done()
			r, err := clnt.SayHello(ctx, &chatpb.ChatMessage{Id: uint64(i), Text: fmt.Sprintf("Hello number %d", i)})
			if err != nil {
				chAnswer <- &serverAnswer{nil, err}
			}
			chAnswer <- &serverAnswer{r, err}
		}(i)
	}

	wg.Wait()

	for x := range chAnswer {
		answers = append(answers, x)
	}

	if err != nil {
		return fmt.Errorf("error during execute grpc : %v", err)
	}

	var countErrs, countAnswers int
	for i := 0 ; i < len(answers); i++ {
		if answers[i].err != nil {
			fmt.Printf("Error : %v\n", answers[i].err)
			countErrs++
		} else {
			fmt.Printf("Answer is : %v\n", answers[i].GetAnswer())
			countAnswers++
		}
	}

	if countErrs > 0 {
		return fmt.Errorf("Errors during the server and the client talking. Count : %d\n", countErrs)
	} else {
		return fmt.Errorf("Server answered on all of messages : %d\n", countAnswers)
	}
}