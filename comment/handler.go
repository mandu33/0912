package main

import (
	"context"
	comment "pro2/comment/kitex_gen/comment"
	db "pro2/dal/mysql"
	"pro2/pkg/middleware"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	_, claims, err := middleware.ParseToken(req.Token)
	if err != nil {
		res := &comment.DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	userID := claims.UserID
	actionType := req.ActionType
	v, _ := db.GetVideoById(ctx, req.VideoId)

	if actionType == 1 {
		cmt := &db.Comment{
			VideoID: uint(req.VideoId),
			Video:   v,
			UserID:  uint(userID),
			Content: req.CommentText,
		}
		err := db.AddComment(ctx, cmt)
		if err != nil {
			res := &comment.DouyinCommentActionResponse{
				StatusCode: -1,
				StatusMsg:  "评论发布失败：服务器内部错误",
			}
			return res, nil
		}
	} else if actionType == 2 {
		err := db.DeleteComment(ctx, req.CommentId, req.VideoId)
		if err != nil {
			res := &comment.DouyinCommentActionResponse{
				StatusCode: -1,
				StatusMsg:  "评论删除失败：服务器内部错误",
			}
			return res, nil
		}
	}
	res := &comment.DouyinCommentActionResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}
	return res, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	_, claims, err := middleware.ParseToken(req.Token)
	var userID int64 = -1
	userID = int64(claims.UserID)
	commentlist, err := db.GetVideoCommentList(ctx, req.VideoId)

	comments := make([]*comment.Comment, 0)
	for _, v := range commentlist {
		user, err := db.QueryUserByID(ctx, int(v.UserID))

		packUser := &comment.User{
			Id:   userID,
			Name: user.UserName,
		}
		if err != nil {
			return nil, err
		}

		comments = append(comments, &comment.Comment{
			Id:         int64(v.ID),
			User:       packUser,
			Content:    v.Content,
			CreateDate: v.CreatedAt.Format("01-02"),
		})
	}
	res := &comment.DouyinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   "success",
		CommentList: comments,
	}
	return res, nil
}
