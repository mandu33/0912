package main

import (
	"log"

	"pro2/user/kitex_gen/user/userservice"

	db "pro2/dal/mysql"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	//etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	db.InitDB()
	// EtcdAddress := "127.0.0.1:9000"
	// r, err := etcd.NewEtcdRegistry([]string{EtcdAddress})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	svr := userservice.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "userservice"}),
		//server.WithRegistry(r)
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
