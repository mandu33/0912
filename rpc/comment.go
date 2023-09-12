package rpc

import (
	"context"
	comment "pro2/comment/kitex_gen/comment"
	"pro2/comment/kitex_gen/comment/commentservice"
	"time"

	etcd "github.com/kitex-contrib/registry-etcd"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
)

var (
	commentClient commentservice.Client
)

func InitComment() {
	EtcdAddress := "127.0.0.1:9002"
	// // EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := commentservice.NewClient(
		"commentservice", //服务名称
		// client.WithMiddleware(middleware.CommonMiddleware),
		// client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		// client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		panic(err)
	}
	commentClient = c
}

func CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (*comment.DouyinCommentActionResponse, error) {
	return commentClient.CommentAction(ctx, req)
}

func CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (*comment.DouyinCommentListResponse, error) {
	return commentClient.CommentList(ctx, req)
}
