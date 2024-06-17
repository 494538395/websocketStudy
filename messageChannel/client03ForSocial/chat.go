package main

import (
	"git.89trillion.com/89t/server/social-server-sdk/constant"
	sdk_ws "git.89trillion.com/89t/server/social-server-sdk/proto/social_sdk_ws"
	"google.golang.org/protobuf/proto"
)

func genChatSendReq() []byte {
	createGroupReq := &sdk_ws.MsgData{
		SendID:               "19999",
		RecvID:               "20000",
		GroupID:              "",
		ClientMsgID:          "2007990-eed12614-f857-47ea-ba56-562726adab2c",
		ServerMsgID:          "",
		SenderPlatformID:     0,
		SenderNickname:       "",
		SenderFaceURL:        "",
		SessionType:          constant.SingleChatType,
		MsgFrom:              0,
		ContentType:          0,
		Content:              []byte("我是19999,嗷呜嗷呜嗷呜嗷呜"),
		Seq:                  0,
		SendTime:             1715432097000,
		CreateTime:           1715432097000,
		Status:               1,
		Options:              nil,
		OfflinePushInfo:      nil,
		AtUserIDList:         nil,
		MsgDataList:          nil,
		AttachedInfo:         "",
		Ex:                   "",
		Topic:                0,
		GroupType:            0,
		ServerID:             0,
		IsReact:              false,
		IsExternalExtensions: false,
		MsgFirstModifyTime:   0,
	}
	bytes, err := proto.Marshal(createGroupReq)
	if err != nil {
		panic(err)
	}
	return bytes
}
