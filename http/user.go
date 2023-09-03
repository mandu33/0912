package http

import (
	"context"
	"net/http"
	"pro2/rpc"
	user "pro2/user/kitex_gen/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(ctx context.Context, c *gin.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	if len(username) == 0 || len(passWord) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "用户名和密码都不能为空"})
	}
	//调用kitex/gen
	resp, err := rpc.Register(context.Background(), &user.DouyinUserRegisterRequest{
		Username: username,
		Password: password,
		Token:    token,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": err})
		return
	}
	SendResponse(c, resp)

}
func LoginHandler(ctx context.Context, c *gin.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	if len(username) == 0 || len(passWord) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "用户名和密码都不能为空"})
	}
	//调用kitex/gen
	resp, err := rpc.Login(context.Background(), &user.DouyinUserLoginRequest{
		Username: username,
		Password: password,
		Token:    token,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": err})
		return
	}
	SendResponse(c, resp)
}
func UserInfoHandler(ctx context.Context, c *gin.RequestContext) {
	userId := c.Query("user_id")
	userId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "用户ID不合法"})
	}
	//字段校验
	//调用kitex_gen
	resp, err := rpc.UserInfo(context.Background(), &user.DouyinUserRequest{
		UserId: userId,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": err})
		return
	}
	SendResponse(c, resp)

}
func SendResponse(c *gin.RequestContext, response interface{}) {
	c.JSON(https.StatusOK, response)
}
