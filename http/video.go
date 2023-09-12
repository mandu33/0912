package http

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"pro2/rpc"
	video "pro2/video/kitex_gen/video"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Base struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}
type RelationAction struct {
	Base
}
type PublishAction struct {
	Base
}
type PublishList struct {
	Base
	VideoList []*video.Video `json:"video_list"`
}
type Feed struct {
	Base
	NextTime  int64          `json:"next_time"`
	VideoList []*video.Video `json:"video_list"`
}

func FeedHandler(c *gin.Context) {
	token := c.Query(("token"))
	latestTime := c.Query("latest_time")
	var timestamp int64 = 0
	if latestTime != "" {
		timestamp, _ = strconv.ParseInt(latestTime, 10, 64)
	} else {
		timestamp = time.Now().UnixMilli()
	}
	req := &video.DouyinFeedRequest{
		LatestTime: timestamp,
		Token:      token,
	}
	res, _ := rpc.Feed(context.Background(), req)
	if res.StatusCode == -1 {
		c.JSON(http.StatusOK, Feed{
			Base: Base{
				StatusCode: -1,
				StatusMsg:  res.StatusMsg,
			},
		})
		return
	}
	c.JSON(http.StatusOK, Feed{
		Base: Base{
			StatusCode: 0,
			StatusMsg:  res.StatusMsg,
		},
		VideoList: res.VideoList,
	})
}
func PublishListHandler(c *gin.Context) {
	token := c.GetString("token")
	uid, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, PublishList{
			Base: Base{
				StatusCode: -1,
				StatusMsg:  "user_id 不合法",
			},
		})
		return
	}
	req := &video.DouyinPublishListRequest{
		Token:  token,
		UserId: uid,
	}
	res, _ := rpc.PublishList(context.Background(), req)
	if res.StatusCode == -1 {
		c.JSON(http.StatusOK, PublishList{
			Base: Base{
				StatusCode: -1,
				StatusMsg:  res.StatusMsg,
			},
		})
		return
	}
	c.JSON(http.StatusOK, PublishList{
		Base: Base{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: res.VideoList,
	})
}
func PublishHandler(c *gin.Context) {
	token := c.PostForm(("token"))
	if token == "" {
		c.JSON(http.StatusOK, PublishAction{
			Base: Base{
				StatusCode: -1,
				StatusMsg:  "用户鉴权失败，token为空",
			},
		})
		return
	}
	title := c.PostForm(("title"))
	if title == "" {
		c.JSON(http.StatusOK, PublishAction{
			Base: Base{
				StatusCode: -1,
				StatusMsg:  "视频title不能为空",
			},
		})
		return
	}
	//加载视频数据
	file := []byte(c.PostForm(("data")))
	reader := bytes.NewReader(file)
	// buffer := make([]byte, len(file))
	buf := bytes.NewBuffer(make([]byte, len(file)))
	if _, err := io.Copy(buf, reader); err != nil {
		c.JSON(http.StatusBadRequest, RelationAction{
			Base: Base{
				StatusCode: -1,
				StatusMsg:  "视频上传失败",
			},
		})
		return
	}
	req := &video.DouyinPublishActionRequest{
		Token: token,
		Title: title,
		Data:  buf.Bytes(),
	}
	res, _ := rpc.Publish(context.Background(), req)
	if res.StatusCode == -1 {
		c.JSON(http.StatusOK, PublishAction{
			Base: Base{
				StatusCode: -1,
				StatusMsg:  res.StatusMsg,
			},
		})
		return
	}
	c.JSON(http.StatusOK, PublishAction{
		Base: Base{
			StatusCode: 0,
			StatusMsg:  res.StatusMsg,
		},
	})
}
