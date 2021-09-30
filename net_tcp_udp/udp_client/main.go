package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:3303")
	if err != nil {
		fmt.Printf("Cannot resolve server addr : %v", err)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Printf("Cannot connect to the server addr : %v", err)
	}

	order := 0
	for {
		msg := ""
		switch order {
			case 0:{
				msg = "qwerty"
				order++
			}
			case 1:{
				msg = "asdfgh"
				order++
			}
			case 2 : {
				msg = "zxcvbn"
				order = 0
			}
		}

		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Fatalf("Cannot send %v", err)
		}
		time.Sleep(2 * time.Second)
	}
}
