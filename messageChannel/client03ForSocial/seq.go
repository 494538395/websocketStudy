package main

import (
	sdk_ws "git.89trillion.com/89t/server/social-server-sdk/proto/social_sdk_ws"
	"google.golang.org/protobuf/proto"
)

func genGetTopicSeqReq() []byte {
	joinReq := &sdk_ws.GetMaxAndMinSeqReq{
		GroupIDList: []string{"50033"},
		UserID:      "2011686",
	}
	bytes, err := proto.Marshal(joinReq)
	if err != nil {
		panic(err)
	}
	req := &Req{
		Identifier: WSSendMsg,
		Data: &Req_SendMsgReq{
			SendMsgReq: &SendMsgReq{
				Event: TopicGetSeq,
				Data:  bytes,
			},
		},
	}
	bytes, err = proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genUploadTopicReq() []byte {
	topicUploadReq := &sdk_ws.PullMsgBySeqListReq{
		TopicSeqList: map[int32]*sdk_ws.SeqList{
			2: &sdk_ws.SeqList{SeqList: []uint32{1763, 1764, 1765, 1766, 1767}},
		},
	}
	bytes, err := proto.Marshal(topicUploadReq)
	if err != nil {
		panic(err)
	}
	req := &Req{
		Identifier: WSSendMsg,
		Data: &Req_SendMsgReq{
			SendMsgReq: &SendMsgReq{
				Event: UploadTopic,
				Data:  bytes,
			},
		},
	}
	bytes, err = proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}
