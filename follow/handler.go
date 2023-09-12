package main

import (
	"context"
	db "pro2/dal/mysql"
	"pro2/follow/kitex_gen/follow"
	"pro2/pkg/middleware"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct{}

// Follow implements the FollowServiceImpl interface.
// 是否关注
func (s *FollowServiceImpl) Follow(ctx context.Context, req *follow.DouyinRelationActionRequest) (resp *follow.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	_, claims, err := middleware.ParseToken(req.Token)
	UserId := claims.UserID
	ToUserId := req.ToUserId
	if err != nil {
		res := &follow.DouyinRelationActionResponse{
			StatusCode: -1,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}

	if req.ActionType == 1 {
		err := db.GetFollowing(ctx, UserId, ToUserId)
		if err != nil {
			res := &follow.DouyinRelationActionResponse{
				StatusCode: -1,
				StatusMsg:  "关注失败",
			}
			return res, nil
		}
	}
	// 2-取消关注
	if req.ActionType == 2 {
		err := db.DeleteFollowing(ctx, UserId, ToUserId)
		if err != nil {
			res := &follow.DouyinRelationActionResponse{
				StatusCode: -1,
				StatusMsg:  "取关失败",
			}
			return res, nil
		}
	}
	res := &follow.DouyinRelationActionResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}
	return res, nil
}

// FollowList implements the FollowServiceImpl interface.
// 关注的人
func (s *FollowServiceImpl) FollowList(ctx context.Context, req *follow.DouyinRelationFollowListRequest) (resp *follow.DouyinRelationFollowListResponse, err error) {
	// TODO: Your code here...
	UserId := req.UserId
	_, chaims, err := middleware.ParseToken(req.Token)
	if err != nil {
		res := &follow.DouyinRelationFollowListResponse{
			StatusCode: -1,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	if UserId != chaims.UserID {
		if err != nil {
			res := &follow.DouyinRelationFollowListResponse{
				StatusCode: -1,
				StatusMsg:  "当前登录用户无法访问其他用户的关注列表",
			}
			return res, nil
		}
	}
	followlist, err := db.GetFollowingList(ctx, UserId)
	if err != nil {
		res := &follow.DouyinRelationFollowListResponse{
			StatusCode: -1,
			StatusMsg:  "关注列表获取失败",
		}
		return res, nil
	}
	users := make([]*follow.User, 0)
	for _, fo := range followlist {
		user, err := db.GetUserByID(ctx, int64(fo.UserID))
		if err != nil {
			return nil, err
		}
		users = append(users, &follow.User{
			Id:             int64(user.ID),
			Name:           user.UserName,
			FollowCount:    int64(user.Following),
			FollowerCount:  int64(user.Followed),
			IsFollow:       true,
			TotalFavorited: int64(user.TotalFavorited),
			FavoriteCount:  int64(user.FavoriteCount),
		})

	}
	res := &follow.DouyinRelationFollowListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   users,
	}
	return res, nil

}

// FollowerList implements the FollowServiceImpl interface.
// 被关注的人
func (s *FollowServiceImpl) FollowerList(ctx context.Context, req *follow.DouyinRelationFollowerListRequest) (resp *follow.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	UserId := req.UserId
	_, claims, err := middleware.ParseToken(req.Token)
	if err != nil {
		res := &follow.DouyinRelationFollowerListResponse{
			StatusCode: -1,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	if UserId != claims.UserID {
		if err != nil {
			res := &follow.DouyinRelationFollowerListResponse{
				StatusCode: -1,
				StatusMsg:  "当前登录用户无法访问其他用户的粉丝列表",
			}
			return res, nil
		}
	}
	followers, err := db.GetFollowingList(ctx, UserId)
	if err != nil {
		res := &follow.DouyinRelationFollowerListResponse{
			StatusCode: -1,
			StatusMsg:  "粉丝列表获取失败",
		}
		return res, nil
	}

	users := make([]*follow.User, 0)

	for _, fo := range followers {
		user, err := db.GetUserByID(ctx, int64(fo.UserID))
		if err != nil {
			return nil, err
		}
		users = append(users, &follow.User{
			Id:             int64(user.ID),
			Name:           user.UserName,
			FollowCount:    int64(user.Following),
			FollowerCount:  int64(user.Followed),
			IsFollow:       true,
			TotalFavorited: int64(user.TotalFavorited),
			FavoriteCount:  int64(user.FavoriteCount),
		})

	}
	res := &follow.DouyinRelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   users,
	}
	return res, nil

}

// FriendList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FriendList(ctx context.Context, req *follow.DouyinRelationFriendListRequest) (resp *follow.DouyinRelationFriendListResponse, err error) {
	// TODO: Your code here...
	UserId := req.UserId
	_, claims, err := middleware.ParseToken(req.Token)
	if err != nil {
		res := &follow.DouyinRelationFriendListResponse{
			StatusCode: -1,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	if UserId != claims.UserID {
		if err != nil {
			res := &follow.DouyinRelationFriendListResponse{
				StatusCode: -1,
				StatusMsg:  "当前登录用户无法访问其他用户的朋友列表",
			}
			return res, nil
		}
	}

	return
}
