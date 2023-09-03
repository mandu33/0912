package main

import (
	"log"
	messages "pro2/message/kitex_gen/messages/messageservice"
)

func main() {
	svr := messages.NewServer(new(MessageServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
