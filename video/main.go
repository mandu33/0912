package main

import (
	"log"
	video "pro2/video/kitex_gen/video/videoservice"
)

func main() {
	svr := video.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
