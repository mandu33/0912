package main
import(
	"github.com/gin-gonic/gin"
	"./dal/mysql"
	"./http"
)
func Init() *gin.Engine {
	//初始化数据库
	mysql.InitDB()
	r := gin.Default()
	r.Static("static", config.Global.StaticSourcePath)
	r.GET("/douyin/feed/",FeedHandler())
	r.POST("/douyin/user/register/",RegisterHandler())
	r.POST("/douyin/user/login/",LoginHandler())
	r.GET("/douyin/user/",UserInfoHandler())
	r.POST("/douyin/publish/action/",PublishHandler())
	r.GET("/douyin/publish/list/",PublishListHandler())
	r.POST("/douyin/favorite/action/",LikeActionHandler())
	r.GET("/douyin/favorite/list/",LikeListHandler())
	r.POST("/douyin/comment/action/",CommentActionHandler())
	r.GET("/douyin/comment/list/",CommentListHandler())
	r.POST("/douyin/relation/action/",FollowHandler())
	r.GET("/douyin/relation/follow/list/",FollowListHandler())
	r.GET("/douyin/relation/follower/list/",FollowerListHandler())
	r.GET("/douyin/relation/friend/list/",FriendListHandler())
	r.GET("/douyin/message/chat/",MessageChatHandler())
	r.POST("/douyin/message/action/",MessageActionHandler())
	return r
}