package main

import (
	"fmt"
	"log"

	"github.com/diegofhdz/unifi-log-streamer/client"
)

func main() {
	fmt.Println("Creating UniFi Client")

	c, err := client.New("192.168.1.1")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully created client at ip %s", c.BaseUrl)
}
