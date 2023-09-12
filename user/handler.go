package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	db "pro2/dal/mysql"
	jwt "pro2/pkg/middleware"
	user "pro2/user/kitex_gen/user"
	"time"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	//检查新用户名称是否已存在
	fmt.Printf(req.Username)
	usr, err := db.QueryUser(ctx, req.Username)
	if usr != nil {
		res := &user.DouyinUserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "该用户名已存在",
		}
		return res, nil
	}
	if err != nil {
		log.Fatal(err)
	}
	//创建user
	u := db.User{
		UserName: req.Username,
		PassWord: req.Password,
	}

	if err := db.InsertUser(ctx, u); err != nil {
		res := &user.DouyinUserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败",
		}
		return res, nil
	}
	rand.Seed(time.Now().UnixMilli())
	//生成token
	token, err := jwt.CreateToken(int64(usr.ID))
	if err != nil {
		log.Fatal(err)
	}
	res := &user.DouyinUserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     int64(u.ID),
		Token:      token,
	}
	return res, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	usr, err := db.QueryUser(ctx, req.Username)
	if err != nil {
		log.Fatal(err)
	}
	if usr == nil {
		res := &user.DouyinUserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "该用户名不存在",
		}
		return res, nil
	}
	if req.Password != usr.PassWord {
		res := &user.DouyinUserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "用户名或密码错误",
		}
		return res, nil
	}
	rand.Seed(time.Now().UnixMilli())
	//生成token
	token, err := jwt.CreateToken(int64(usr.ID))
	if err != nil {
		log.Fatal(err)
	}
	res := &user.DouyinUserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     int64(usr.ID),
		Token:      token,
	}
	return res, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	//return &user.DouyinUserResponse{StatusCode: -1, StatusMsg: "该用户名不存在"}, nil

	usr, err := db.QueryUserByID(ctx, int(req.UserId))
	if usr == nil {
		res := &user.DouyinUserResponse{
			StatusCode: -1,
			StatusMsg:  "aaaa该用户名不存在",
		}
		return res, nil
	}

	//返回结果
	res := &user.DouyinUserResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		User: &user.User{
			Id:              int64(usr.ID),
			Name:            usr.UserName,
			FollowCount:     int64(usr.Following),
			FollowerCount:   int64(usr.Followed),
			IsFollow:        false,
			Avatar:          "",
			BackgroundImage: "",
			Signature:       "",
			TotalFavorited:  int64(usr.TotalFavorited),
			WorkCount:       0,
			FavoriteCount:   int64(usr.FavoriteCount),
		},
	}
	return res, nil
}
