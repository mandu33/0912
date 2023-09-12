package rpc

import (
	"context"
	"log"
	user "pro2/user/kitex_gen/user"
	"pro2/user/kitex_gen/user/userservice"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	userClient userservice.Client
)

// User RPC客户端初始化
func InitUser() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"}) // r不应重复使用。
	if err != nil {
		log.Fatal(err)
	}
	c, err := userservice.NewClient(
		"userservice", //user服务名称
		client.WithHostPorts("0.0.0.0:8888"),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}
func Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (*user.DouyinUserRegisterResponse, error) {
	return userClient.Register(ctx, req)
}
func Login(ctx context.Context, req *user.DouyinUserLoginRequest) (*user.DouyinUserLoginResponse, error) {
	return userClient.Login(ctx, req)
}
func UserInfo(ctx context.Context, req *user.DouyinUserRequest) (*user.DouyinUserResponse, error) {
	return userClient.UserInfo(ctx, req)
}
