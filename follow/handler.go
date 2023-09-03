package main

import (
	"context"
	follow "pro2/follow/kitex_gen/follow"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct{}

// Follow implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) Follow(ctx context.Context, req *follow.DouyinRelationActionRequest) (resp *follow.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowList(ctx context.Context, req *follow.DouyinRelationFollowListRequest) (resp *follow.DouyinRelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowerList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowerList(ctx context.Context, req *follow.DouyinRelationFollowerListRequest) (resp *follow.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// FriendList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FriendList(ctx context.Context, req *follow.DouyinRelationFriendListRequest) (resp *follow.DouyinRelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}
