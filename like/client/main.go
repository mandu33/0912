package main

import (
	"context"
	"log"
	"pro2/like/kitex_gen/likes"
	"pro2/like/kitex_gen/likes/likeservice"

	"github.com/cloudwego/kitex/client"
)

func main() {
	c, err := likeservice.NewClient("echo", client.WithHostPorts("127.0.0.1:8891"))
	if err != nil {
		log.Fatal(err)
	}
	req := &likes.DouyinFavoriteListRequest{
		UserId: int64(1),
		Token:  "12345",
	}
	resp, err := c.LikeList(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
