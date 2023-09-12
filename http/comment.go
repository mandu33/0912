package http

import (
	"context"
	"net/http"
	"pro2/comment/kitex_gen/comment"
	"pro2/rpc"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CommentActionHandler(c *gin.Context) {

	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")

	vid, err := strconv.Atoi(video_id)

	if err != nil {
		SendCommentResponse(c, err)
		return
	}

	acty, err := strconv.Atoi(action_type)

	if err != nil {
		SendCommentResponse(c, err)
		return
	}

	rpc_Req := comment.DouyinCommentActionRequest{
		VideoId:    int64(vid),
		Token:      token,
		ActionType: int32(acty),
	}
	// 1-发布评论，2-删除评论
	if acty == 1 {
		comment_text := c.Query("comment_text")
		rpc_Req.CommentText = comment_text
	} else {
		comment_id := c.Query("comment_id")
		c_id, err := strconv.Atoi(comment_id)

		if err != nil {
			SendCommentResponse(c, err)
			return
		}

		c_id64 := int64(c_id)
		rpc_Req.CommentId = c_id64
	}

	resp, err := rpc.CommentAction(context.Background(), &rpc_Req)

	if err != nil {
		SendCommentResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"Data":    resp,
	})

}

func CommentListHandler(c *gin.Context) {
	video_id, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		SendCommentResponse(c, err)
		return
	}

	VideoId := int64(video_id)
	Token := c.Query("token")

	if len(Token) == 0 || VideoId < 0 {
		SendCommentResponse(c, err)
		return
	}

	resp, err := rpc.CommentList(context.Background(), &comment.DouyinCommentListRequest{
		VideoId: VideoId,
		Token:   Token,
	})

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"Data":    resp,
	})

}

func SendCommentResponse(c *gin.Context, err error) {
	// always return http.StatusOK
	c.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": "error",
	})
}
