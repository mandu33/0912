package main

import (
	"log"
	"net"

	// handler "pro2/like"
	db "pro2/dal/mysql"
	"pro2/like/kitex_gen/likes/likeservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	//etcd "github.com/kitex-contrib/registry-etcd"
	//trace "github.com/kitex-contrib/tracer-opentracing"
)

// func main() {
// 	db.InitDB()
// 	// EtcdAddress := "127.0.0.1:8888"
// 	// r, err := etcd.NewEtcdRegistry([]string{EtcdAddress})
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8891")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	svr := likeservice.NewServer(new(LikeServiceImpl),
// 		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "likeservice"}),
// 		server.WithServiceAddr(addr),
// 		// server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
// 		// server.WithMuxTransport(),                                          // Multiplex
// 		// server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
// 		//server.WithBoundHandler(bound.NewCpuLimitHandler()),
// 		//server.WithRegistry(r),
// 	)

// 	err = svr.Run()

// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// }

func main() {
	db.InitDB()
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8891")
	svr := likeservice.NewServer(new(LikeServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "likeservice"}),
		server.WithServiceAddr(addr),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
