package main

import (
	"context"
	db "pro2/dal/mysql"
	"pro2/pkg/middleware"
	video "pro2/video/kitex_gen/video"
	"strconv"
	"time"

	"github.com/bytedance/gopkg/util/logger"
)

const limit = 30

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.DouyinFeedRequest) (resp *video.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	nextTime := time.Now().UnixMilli()
	var userID int64 = -1
	//验证token
	if req.Token != "" {
		_, claims, err := middleware.ParseToken(req.Token)
		if err != nil {
			res := &video.DouyinFeedResponse{
				StatusCode: -1,
				StatusMsg:  "token 解析错误",
			}
			return res, nil
		}
		userID, err = strconv.ParseInt(claims.Id, 10, 64)
	}
	//数据库查询视频列表
	videos, err := db.GetVideos(ctx, limit, &req.LatestTime)
	if err != nil {
		res := &video.DouyinFeedResponse{
			StatusCode: -1,
			StatusMsg:  "视频获取失败：服务器内部错误",
		}
		return res, nil
	}
	videoList := make([]*video.Video, 0)
	for _, r := range videos {
		author, err := db.GetUserByID(ctx, int64(r.Author.ID))
		if err != nil {
			return nil, err
		}
		relation, err := db.GetFollow(ctx, userID, int64(author.ID))
		if err != nil {
			res := &video.DouyinFeedResponse{
				StatusCode: -1,
				StatusMsg:  "视频获取失败：服务器内部错误",
			}
			return res, nil
		}
		favorite, err := db.GetFavoriteList(ctx, userID)
		if err != nil {
			res := &video.DouyinFeedResponse{
				StatusCode: -1,
				StatusMsg:  "视频获取失败：服务器内部错误",
			}
			return res, nil
		}
		playUrl, err := middleware.GetVideoByUrl(r.PlayUrl)
		if err != nil {
			res := &video.DouyinFeedResponse{
				StatusCode: -1,
				StatusMsg:  "服务器内部错误：视频获取失败",
			}
			return res, nil
		}
		coverUrl, err := middleware.GetCoverByUrl(r.CoverUrl)
		if err != nil {
			logger.Errorf("Minio获取链接失败：%v", err.Error())
			res := &video.DouyinFeedResponse{
				StatusCode: -1,
				StatusMsg:  "服务器内部错误：封面获取失败",
			}
			return res, nil
		}
		avatarUrl, err := middleware.GetPicByurl(author.Avatar)
		if err != nil {
			logger.Errorf("Minio获取链接失败：%v", err.Error())
			res := &video.DouyinFeedResponse{
				StatusCode: -1,
				StatusMsg:  "服务器内部错误：头像获取失败",
			}
			return res, nil
		}
		backgroundUrl, err := middleware.GetBackByUrl(author.BackgroundImage)
		if err != nil {
			logger.Errorf("Minio获取链接失败：%v", err.Error())
			res := &video.DouyinFeedResponse{
				StatusCode: -1,
				StatusMsg:  "服务器内部错误：背景图获取失败",
			}
			return res, nil
		}

		videoList = append(videoList, &video.Video{
			Id: int64(r.ID),
			Author: &video.User{
				Id:              int64(author.ID),
				Name:            author.UserName,
				FollowCount:     int64(author.Following),
				FollowerCount:   int64(author.Followed),
				IsFollow:        relation != nil,
				Avatar:          avatarUrl,
				BackgroundImage: backgroundUrl,
				Signature:       "",
				TotalFavorited:  int64(author.TotalFavorited),
				WorkCount:       0,
				FavoriteCount:   int64(author.FavoriteCount),
			},
			PlayUrl:       playUrl,
			CoverUrl:      coverUrl,
			FavoriteCount: int64(r.FavoriteCount),
			CommentCount:  int64(r.CommentCount),
			IsFavorite:    favorite != nil,
			Title:         r.Title,
		})
	}
	if len(videos) != 0 {
		nextTime = videos[len(videos)-1].UpdatedAt.UnixMilli()
	}
	res := &video.DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  videoList,
		NextTime:   nextTime,
	}
	return res, nil
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.DouyinPublishListRequest) (resp *video.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}
