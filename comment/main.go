package main

import (
	"log"
	"pro2/comment/kitex_gen/comment/commentservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	EtcdAddress := "127.0.0.1:9002"
	r, err := etcd.NewEtcdRegistry([]string{EtcdAddress})
	if err != nil {
		log.Fatal(err)
	}
	svr := commentservice.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "commentservice"}),
		server.WithRegistry(r))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
