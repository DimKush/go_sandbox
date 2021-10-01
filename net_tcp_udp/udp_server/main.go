package main
 
import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:3303")
	if err != nil {
		log.Fatalf("Cannot listen create UDP. Reason %s", err.Error())
	}

	l, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Cannot listen UDP. Reason %s", err.Error())
	}

	defer l.Close()

	msg := make([]byte, 10, 10)
	for {
		length, fromAddr, err := l.ReadFromUDP(msg)
		if err != nil {
			log.Fatalf("Error during read. Reason : %s", err.Error())
		}

		fmt.Printf("Message from %s with length %d : %s \n", fromAddr.String(), length, string(msg))
	}
}
