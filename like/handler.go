package main

import (
	"context"
	likes "pro2/like/kitex_gen/likes"
	db "pro2/dal/mysql"
)

// LikeServiceImpl implements the last service interface defined in the IDL.
type LikeServiceImpl struct{}

// LikeAction implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) LikeAction(ctx context.Context, req *likes.DouyinFavoriteActionRequest) (resp *likes.DouyinFavoriteActionResponse, err error) {
	// TODO: Your code here...
	resp := new(likes.DouyinFavoriteActionResponse)
	if len(req.Token) == 0 || req.VideoId == 0 || req.ActionType == 0 {
		resp = &likes.DouyinFavoriteActionResponse{
			StatusCode: -1, 
			StatusMsg: "error",
		}
		return resp, nil
	}

	if req.ActionType == 1 {
		return db.CreateFavorite(ctx, req.UserId, req.VideoId)
	}
	// 2-取消点赞
	if req.ActionType == 2 {
		return db.DeleteFavorite(ctx, req.UserId, req.VideoId)
	}

	return nil
}

// LikeList implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) LikeList(ctx context.Context, req *likes.DouyinFavoriteListRequest) (resp *likes.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	resp := new(likes.DouyinFavoriteListResponse)
	if len(req.Token) == 0 || req.UserId == 0 {
		resp = &likes.DouyinFavoriteListResponse{
			StatusCode: -1, 
			StatusMsg: "error",
			VideoList : nil,
		}
		return resp, nil
	}

	videos := make([]*likes.Video, 0)
	vids, err := db.GetFavoriteList(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	for _, vid := range vids {
		//get video
		video, err := db.GetVideoById(ctx, vid)
		if err != nil {
			resp = &likes.DouyinFavoriteListResponse{
				StatusCode: -1, 
				StatusMsg: "get video error",
			}
			return resp, err
		}
		//get author
		user, err := db.queryUser(ctx, video.UserId)
		if err != nil {
			resp = &likes.DouyinFavoriteListResponse{
				StatusCode: -1, 
				StatusMsg: "get author error",
			}
			return resp, err
		}

		u := favorite.User{
			Id:   user.Id,
			userName: user.Name,
		}

		v := favorite.Video{
			Id:            video.Id,
			Author:        &u,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
		}
		videos = append(videos, &v)

	}

	resp = &likes.DouyinFavoriteListResponse{
		StatusCode: 0, 
		StatusMsg: "success",
		VideoList : videos
	}
	return resp, nil
}
