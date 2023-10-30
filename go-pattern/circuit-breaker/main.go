package main

import (
	"log"
	"os"
	"time"

	"github.com/longpt233/hello-golang/utils"
)

func main() {
	logger := log.New(os.Stdout, "Main\t", log.LstdFlags)
	server := utils.NewExampleServer(":8080")

	go func() {
		_ = server.ListenAndServe()
	}()

	var client utils.NotificationClient

	client = utils.NewSmsClient()

	client = utils.NewClientCircuitBreakerProxy(client)

	for {
		err := client.Send("http://127.0.0.1:8080")
		time.Sleep(1 * time.Second)
		if err != nil {
			logger.Println("caught an error", err)
		}
	}
}
