package main

import (
	"context"
	db "pro2/dal/mysql"
	likes "pro2/like/kitex_gen/likes"
	"pro2/pkg/middleware"
	Video "pro2/video/kitex_gen/video"
)

// LikeServiceImpl implements the last service interface defined in the IDL.
type LikeServiceImpl struct{}

// LikeAction implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) LikeAction(ctx context.Context, req *likes.DouyinFavoriteActionRequest) (resp *likes.DouyinFavoriteActionResponse, err error) {
	// TODO: Your code here...
	//得到token
	_, claims, err := middleware.ParseToken(req.Token)
	UserId := claims.UserID
	resp = new(likes.DouyinFavoriteActionResponse)

	if len(req.Token) == 0 || req.VideoId == 0 || req.ActionType == 0 {
		resp = &likes.DouyinFavoriteActionResponse{
			StatusCode: -1,
			StatusMsg:  "error",
		}
		return resp, nil
	}

	if req.ActionType == 1 {

		db.CreateFavorite(ctx, UserId, req.VideoId)

		resp = &likes.DouyinFavoriteActionResponse{
			StatusCode: 0,
			StatusMsg:  "success",
		}
		return resp, nil
	}
	// 2-取消点赞
	if req.ActionType == 2 {
		db.DeleteFavorite(ctx, UserId, req.VideoId)

		resp = &likes.DouyinFavoriteActionResponse{
			StatusCode: 0,
			StatusMsg:  "success",
		}
		return resp, nil
	}

	return resp, nil
}

// LikeList implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) LikeList(ctx context.Context, req *likes.DouyinFavoriteListRequest) (resp *likes.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	// if len(req.Token) == 0 || req.UserId == 0 {
	// 	resp = &likes.DouyinFavoriteListResponse{
	// 		StatusCode: -1,
	// 		StatusMsg:  "error",
	// 		VideoList:  nil,
	// 	}
	// 	return nil
	// }

	videos := make([]*Video.Video, 0)
	vids, err := db.GetFavoriteList(ctx, req.UserId)
	if err != nil {
		res := &likes.DouyinFavoriteListResponse{
			StatusCode: -1,
			StatusMsg:  "获取喜欢列表失败：服务器内部错误",
		}
		return res, nil
	}

	for _, vid := range vids {
		//get video
		video, err := db.GetVideoById(ctx, int64(vid.ID))
		if err != nil {
			res := &likes.DouyinFavoriteListResponse{
				StatusCode: -1,
				StatusMsg:  "获取喜欢列表失败：服务器内部错误",
			}
			return res, nil
		}

		//get author
		user, err := db.QueryUserByID(ctx, video.UserID)
		if err != nil {
			res := &likes.DouyinFavoriteListResponse{
				StatusCode: -1,
				StatusMsg:  "获取喜欢列表失败：服务器内部错误",
			}
			return res, nil
		}

		u := Video.User{
			Id:   int64(user.ID),
			Name: user.UserName,
		}

		v := Video.Video{
			Id:            int64(video.ID),
			Author:        &u,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: int64(video.FavoriteCount),
			CommentCount:  int64(video.CommentCount),
			Title:         video.Title,
		}
		videos = append(videos, &v)

	}
	res := &likes.DouyinFavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  videos,
	}
	return res, nil
}
