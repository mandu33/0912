package main

import (
	"context"
	messages "pro2/message/kitex_gen/messages"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *messages.DouyinMessageChatRequest) (resp *messages.DouyinMessageChatResponse, err error) {
	// TODO: Your code here...
	return
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *messages.DouyinRelationActionRequest) (resp *messages.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}
