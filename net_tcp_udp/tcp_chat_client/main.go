package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func readRoutine(ctx context.Context, conn net.Conn){
	scanner := bufio.NewScanner(conn)
	OUTER: for {
		select {
			case <- ctx.Done():
				{
					break OUTER
				}
		default:
			{
				if !scanner.Scan(){
					log.Printf("Cannot scan.")
					break OUTER
				}
				text := scanner.Text()
				log.Printf("From server: %s", text)
			}
		}
	}
	log.Printf("Finished readRoutine")
}

func writeRoutine(ctx context.Context, cancel context.CancelFunc, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	OUTER : for {
		select {
			case <-ctx.Done(): {
				break OUTER
			}
			default: {
				if !scanner.Scan(){
					cancel()
					break OUTER
				}
			}
			str := scanner.Text()
			log.Printf("Send to the server %s\n", str)

			conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}
	}
	log.Printf("Finished write.")
}

func main(){
	dialer := &net.Dialer{}
	ctx, cancel := context.WithTimeout(context.Background(), 5 * 60 * time.Second)

	conn, err := dialer.DialContext(ctx, "tcp", "127.0.0.1:3302")
	if err != nil {
		log.Fatalf("Cannot connect %v", err)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(){
		readRoutine(ctx, conn)
		wg.Done()
	}()

	wg.Add(1)
	go func(){
		writeRoutine(ctx, cancel, conn)
		wg.Done()
	}()

	wg.Wait()
	conn.Close()
}
