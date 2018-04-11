package main

import "flag"

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "listen on the specified ip")
	flag.Parse()
}
