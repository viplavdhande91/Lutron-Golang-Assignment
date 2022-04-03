package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	var ch chan bool
	conn, err := net.Dial("tcp", ":8111")

	if err != nil {
		panic(err)
	}

	fmt.Println("Enter name: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	_, err = conn.Write([]byte(text))

	if err != nil {
		panic(err)
	}
	go readMsg(conn)
	go writeMsg(conn)
	<-ch

}

func writeMsg(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		_, err := conn.Write([]byte(text))

		if err != nil {
			panic(err)
		}
	}
}

func readMsg(conn net.Conn) {

	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')

		if err == io.EOF {
			fmt.Println("Connection close, Bye!")
			conn.Close()
		} else if err != nil {
			fmt.Println(err.Error())
			conn.Close()
		}

		msg = msg[:len(msg)-1]
		fmt.Println(string(msg))
	}
}