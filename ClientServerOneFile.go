package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func receiveConn() net.Conn {
	for {
		xx, err := net.Dial("tcp", "localhost:9002")
		if err == nil {
			return xx
		}
		time.Sleep(1 * time.Second)
	}
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go func() {
		listen, err := net.Listen("tcp", "localhost:9001")
		if err != nil {
			log.Fatal(err)
		}
		defer listen.Close()

		for {
			conn, err := listen.Accept()
			defer conn.Close()
			if err != nil {
				log.Fatal(err)
			}

			io.Copy(os.Stdout, conn)

		}
	}()

	go func() {
		dial := receiveConn()

		defer dial.Close()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == "exit" {
				dial.Write([]byte("Your mate has left the room"))
				break
			}
			dial.Write([]byte(scanner.Text() + "\n"))
		}
	}()

	wg.Wait()
}
