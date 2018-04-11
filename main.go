package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "listen on the specified ip")
	flag.Parse()

	if isHost {
		fmt.Println("is host")
		//the third argument, after main, and the flag
		connIP := os.Args[2]
		runHost(connIP)
	} else {
		fmt.Println("is guest")
		//the second argument because there is no flag in the guest side
		connIP := os.Args[1]
		runGuest(connIP)
	}
}


func runHost(ip string) {

}

func runGuest(ip string) {

}