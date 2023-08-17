package main

import (
	"context"
	chat "github.com/ozline/tiktok/kitex_gen/chat"
	service "github.com/ozline/tiktok/cmd/chat/service"
	pack "github.com/ozline/tiktok/cmd/chat/pack"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessagePost implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessagePost(ctx context.Context, req *chat.MessagePostRequest) (resp *chat.MessagePostReponse, err error) {
	// TODO: Your code here...
	return
}

// MessageList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageList(ctx context.Context, req *chat.MessageListRequest) (resp *chat.MessageListResponse, err error) {
	// TODO: Your code here...
	resp = new(chat.MessageListResponse)
	claim,err :=utils.CheckToken(req.Token)
	if err!=nil||claim==nil{
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp,nil
	}
	//获取消息列表
	
	//redis->mysql
	//redis中存在则返回，不存在查询mysql,
	messageList,err :=service.NewChatService(ctx).GetMessages(req,claim.UserId)
	if err!=nil{
		resp.Base = pack.BuildBaseResp(err)
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.MessageList = pack.BuildMessage(messageList)
	resp.Total = int64(len(messageList))
	return
}