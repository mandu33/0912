package rpc
import(
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	comment "pro2/comment/kitex_gen/comment"
	"pro2/comment/kitex_gen/comment/commentservice"
	"fmt"
	"context"
)
var (
	commentClient commentservice.Client
)
func InitComment(){

}
func CommentAction(ctx context.Context, req *comment.CommentActionRequest){
	return commentClient.CommentAction(ctx, req)
}
func CommentList(ctx context.Context, req *comment.CommentListRequest){
	return commentClient.CommentList(ctx, req)
}