package main

import (
	"log"
	likes "pro2/like/kitex_gen/likes/likeservice"
)

func main() {
	svr := likes.NewServer(new(LikeServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
