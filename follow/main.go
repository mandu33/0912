package main

import (
	"log"
	follow "pro2/follow/kitex_gen/follow/followservice"
)

func main() {
	svr := follow.NewServer(new(FollowServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
