package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"log"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "listen on the specified ip")
	flag.Parse()

	if isHost {
		fmt.Println("is host")
		//the third argument, after main, and the flag
		connIP := os.Args[2] + ":8080"
		runHost(connIP)
	} else {
		fmt.Println("is guest")
		//the second argument because there is no flag in the guest side
		connIP := os.Args[1] + ":8080"
		runGuest(connIP)
	}
}

func runHost(ip string) {
	listen, listenErr := net.Listen("tcp", ip)
	if listenErr != nil {
		log.Fatal("Error:",listenErr)
	}
	//Accept method returns a conn object
	conn, acceptErr := listen.Accept()
	if acceptErr != nil {
		log.Fatal("Error:", acceptErr)
	}
	//create a reader interface
	reader := bufio.NewReader(conn)
	for {
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
}

func runGuest(ip string) {
	//connect to host
	conn, dialErr := net.Dial("tcp", ip)
	if dialErr != nil {
		log.Fatal("Error:",dialErr)
	}
	for {
		fmt.Print("Send message:")
		reader:=bufio.NewReader(os.Stdin)
		message, readErr:=reader.ReadString('\n')
		if readErr != nil {
			log.Fatal("Error", readErr)
		}
		fmt.Fprint(conn, message)
	}


}
