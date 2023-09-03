package rpc
import(
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	user "pro2/user/kitex_gen/user"
	"pro2/user/kitex_gen/user/userservice"
	"fmt"
	"context"
	"github.com/cloudwego/kitex/pkg/retry"
)
var (
	userClient userservice.Client
)
//User RPC客户端初始化
func InitUser(){
	c, err := userservice.NewClient(
		"userservice",//user服务名称
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}
func Register(ctx context.Context, req *user.DouyinUserRegisterRequest){
	return userClient.Register(ctx, req)
}
func Login(ctx context.Context, req *user.DouyinUserLoginRequest){
	return userClient.Login(ctx, req)
}
func UserInfo(ctx context.Context, req *user.DouyinUserInfoRequest){
	return userClient.UserInfo(ctx, req)
}