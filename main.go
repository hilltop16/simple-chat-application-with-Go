package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
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
		fmt.Println("Error:", listenErr)
		os.Exit(1)
	}
	//Accept method returns a conn object
	conn, acceptErr := listen.Accept()
	if acceptErr != nil {
		fmt.Println("Error:", acceptErr)
		os.Exit(1)
	}
	//create a reader interface
	reader := bufio.NewReader(conn)
	//convert to string
	message, readErr := reader.ReadString('\n') // \n = enter key
	if readErr != nil {
		fmt.Println("Error:", readErr)
		os.Exit(1)
	}
	fmt.Println("message received:", message)
}

func runGuest(ip string) {
	//connect to host
	conn, dialErr := net.Dial("tcp", ip)
	if dialErr != nil {
		fmt.Println("Error:",dialErr)
		os.Exit(1)
	}
	fmt.Print("Send message:")
	reader:=bufio.NewReader(os.Stdin)
	message, readErr:=reader.ReadString('\n')
	if readErr != nil {
		fmt.Println("Error:",readErr)
		os.Exit(1)
	}
	fmt.Fprint(conn, message)

}
