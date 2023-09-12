package main

import (
	"context"
	"log"
	"pro2/user/kitex_gen/user"
	"pro2/user/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/client"
)

func main() {
	c, err := userservice.NewClient("echo", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	// req := &user.DouyinUserRequest{
	// 	UserId: 1,
	// }
	// resp, err := c.UserInfo(context.Background(), req)
	req := &user.DouyinUserRegisterRequest{
		Username: "admin",
		Password: "123456",
	}
	resp, err := c.Register(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
