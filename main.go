package main

import (
	"gotlsgrpc/client"
	"gotlsgrpc/server"
	"time"
)

func main() {
	go server.NewServer()

	time.Sleep(5 * time.Second)

	client.ClientTLS()
}
