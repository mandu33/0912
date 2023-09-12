package http

import (
	"context"
	"net/http"
	"pro2/like/kitex_gen/likes"
	"pro2/rpc"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SendLikeResponse(c *gin.Context, err error) {
	// always return http.StatusOK
	c.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": "error",
	})
}
func LikeActionHandler(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")

	vid, err := strconv.Atoi(video_id)
	if err != nil {
		SendLikeResponse(c, err)
		return
	}
	act, err := strconv.Atoi(action_type)
	if err != nil {
		SendLikeResponse(c, err)
		return
	}
	//
	resp, err := rpc.LikeAction(context.Background(), &likes.DouyinFavoriteActionRequest{
		VideoId:    int64(vid),
		Token:      token,
		ActionType: int32(act),
	})

	if err != nil {
		SendLikeResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "success",
		"Data":    resp,
	})
}

func LikeListHandler(c *gin.Context) {
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendLikeResponse(c, err)
		return
	}

	Token := c.Query("token")
	if len(Token) == 0 || userid < 0 {
		SendLikeResponse(c, err)
		return
	}

	resp, err := rpc.LikeList(context.Background(), &likes.DouyinFavoriteListRequest{
		UserId: int64(userid),
		Token:  Token,
	})

	if err != nil {
		SendLikeResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"Data":    resp,
	})
}
