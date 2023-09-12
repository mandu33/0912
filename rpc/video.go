package rpc

import (
	"context"
	video "pro2/video/kitex_gen/video"
	"pro2/video/kitex_gen/video/videoservice"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
)

var (
	videoClient videoservice.Client
)

func InitVideo() {
	v, err := videoservice.NewClient(
		"videoservice", //video服务名称
		client.WithHostPorts("0.0.0.0:8882"),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
	)
	if err != nil {
		panic(err)
	}
	videoClient = v

}
func Feed(ctx context.Context, req *video.DouyinFeedRequest) (*video.DouyinFeedResponse, error) {
	return videoClient.Feed(ctx, req)
}
func Publish(ctx context.Context, req *video.DouyinPublishActionRequest) (*video.DouyinPublishActionResponse, error) {
	return videoClient.PublishAction(ctx, req)
}
func PublishList(ctx context.Context, req *video.DouyinPublishListRequest) (*video.DouyinPublishListResponse, error) {
	return videoClient.PublishList(ctx, req)
}
