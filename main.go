package main

import (
	"flag"
	"fmt"
	"os"
	"net"
	"bufio"
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
	listen, listenErr :=net.Listen("tcp", ip)
	if listenErr != nil {
		fmt.Println("Error:",listenErr)
		os.Exit(1)
	}
	//Accept method returns a conn object
	conn, acceptErr:=listen.Accept()
	if acceptErr != nil {
		fmt.Println("Error:",acceptErr)
		os.Exit(1)
	}
	//create a reader interface
	reader:=bufio.NewReader(conn)
	//convert to string
	message, readErr:=reader.ReadString('\n') // \n = enter key
	if readErr != nil {
		fmt.Println("Error:",readErr)
		os.Exit(1)
	}
	fmt.Println(message)

}

func runGuest(ip string) {

}