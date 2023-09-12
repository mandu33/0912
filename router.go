package main

import (
	"pro2/dal/mysql"
	"pro2/http"

	"github.com/gin-gonic/gin"
	//"github.com/cloudwego/hertz/pkg/common/config"
)

func Init() *gin.Engine {
	//初始化数据库
	mysql.InitDB()

	r := gin.Default()
	// r.Static("static", config.Global.StaticSourcePath)
	r.GET("/douyin/feed/", http.FeedHandler)
	r.POST("/douyin/user/register/", http.RegisterHandler)
	r.POST("/douyin/user/login/", http.LoginHandler)
	r.GET("/douyin/user/", http.UserInfoHandler)
	r.POST("/douyin/publish/action/", http.PublishHandler)
	r.GET("/douyin/publish/list/", http.PublishListHandler)
	r.POST("/douyin/favorite/action/", http.LikeActionHandler)
	r.GET("/douyin/favorite/list/", http.LikeListHandler)
	r.POST("/douyin/comment/action/", http.CommentActionHandler)
	r.GET("/douyin/comment/list/", http.CommentListHandler)
	r.POST("/douyin/relation/action/", http.FollowHandler)
	r.GET("/douyin/relation/follow/list/", http.FollowListHandler)
	r.GET("/douyin/relation/follower/list/", http.FollowerListHandler)
	r.GET("/douyin/relation/friend/list/", http.FriendListHandler)
	r.GET("/douyin/message/chat/", http.MessageChatHandler)
	r.POST("/douyin/message/action/", http.MessageActionHandler)
	return r
}
