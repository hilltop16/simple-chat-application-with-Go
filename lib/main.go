package lib

import (
	"log"
	"fmt"
	"net"
	"bufio"
	"os"
)

func RunHost(ip string) {
	listen, listenErr := net.Listen("tcp", ip)
	if listenErr != nil {
		log.Fatal("Error:",listenErr)
	}
	//Accept method returns a conn object
	conn, acceptErr := listen.Accept()
	if acceptErr != nil {
		log.Fatal("Error:", acceptErr)
	}

	for {
		handleHost(conn)
	}
}

func handleHost(conn net.Conn) {
	//create a reader interface
	reader := bufio.NewReader(conn)
	//convert to string
	message, readErr := reader.ReadString('\n') // \n = enter key
	if readErr != nil {
		log.Fatal("Error", readErr)
	}
	fmt.Println("message received:", message)


	fmt.Printf("Send message:")
	replyReader :=bufio.NewReader(os.Stdin)
	replyMessage, replyReadErr:=replyReader.ReadString('\n')
	if replyReadErr != nil {
		log.Fatal("error:", readErr)
	}
	fmt.Fprint(conn, replyMessage)
}

func RunGuest(ip string) {
	//connect to host
	conn, dialErr := net.Dial("tcp", ip)
	if dialErr != nil {
		log.Fatal("Error:",dialErr)
	}
	for {
		handleGuest(conn)
	}


}

func handleGuest(conn net.Conn) {
	fmt.Print("Send message:")
	reader:=bufio.NewReader(os.Stdin)
	message, readErr:=reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error", readErr)
	}
	fmt.Fprint(conn, message)

	replyReader:=bufio.NewReader(conn)
	replyMessage, replyErr:= replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error:", replyErr)
	}
	fmt.Println("Message received:", replyMessage)
}