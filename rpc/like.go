package rpc

import (
	"context"
	"pro2/like/kitex_gen/likes"
	"pro2/like/kitex_gen/likes/likeservice"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
)

var likeClient likeservice.Client

func InitLikeRpc() {
	// EtcdAddress := "127.0.0.1:2379"
	// // EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	// r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
	// if err != nil {
	// 	panic(err)
	// }

	c, err := likeservice.NewClient(
		"likeservice", //服务名称
		// client.WithMiddleware(middleware.CommonMiddleware),
		// client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		// client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		//client.WithResolver(r), // resolver
	)
	if err != nil {
		panic(err)
	}
	likeClient = c
}

func LikeAction(ctx context.Context, req *likes.DouyinFavoriteActionRequest) (resp *likes.DouyinFavoriteActionResponse, err error) {
	return likeClient.LikeAction(ctx, req)
}

// 传递 获取点赞列表操作 的上下文, 并获取 RPC Server 端的响应.
func LikeList(ctx context.Context, req *likes.DouyinFavoriteListRequest) (resp *likes.DouyinFavoriteListResponse, err error) {
	return likeClient.LikeList(ctx, req)

}
