package main

import (
	"context"
	user "pro2/user/kitex_gen/user"
	"pro2/dal/mysql"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	//检查新用户名称是否已存在
	usr,err := db.queryUser(ctx,req.Username)
	if usr != nil{
		res := &user.DouyinUserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "该用户名已存在",
		}
	}
	//创建user
	rand.Seed(time.Now().UnixMilli())
	usr = &db.User{
		UserName: req.Username,
		Password: req.Password,
	}
	if err := db.insertUser(ctx, usr); err != nil {
		res := &user.DouyinUserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败",
		}
		return res, nil
	res := &user.DouyinUserRegisterResponse{
			StatusCode: 0,
			StatusMsg:  "success",
			UserId:     int64(usr.ID),
			Token:      token,
	}
	return res, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	usr,err := db.queryUser(ctx,req.Username)
	if usr == nil{
		res := &user.DouyinUserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "该用户名不存在",
		}
		return res, nil
	}
	if req.Password != usr.passWord{
		res := &user.DouyinUserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "用户名或密码错误",
		}
		return res, nil
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
	usr,err := db.queryUser(ctx,req.Username)
	if usr == nil{
		res := &user.DouyinUserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "该用户名不存在",
		}
		return res, nil
	}

	//返回结果
	res := &user.DouyinUserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		User: &user.User{
			Id:int(usr.ID),
			userName:usr.UserName,
			passWord:usr.passWord
			profession:usr.profession
			age:int(usr.age)
			introduction:usr.introduction
			following:int(usr.ollowing)
			followed:int(usr.followed)
			favorite_count:int(usr.favorite_count)
			total_favorited:int(usr.total_favorited)
		},
	}
	return res, nil
}
