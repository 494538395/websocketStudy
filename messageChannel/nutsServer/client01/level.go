package main

import (
	"websocketStudy/messageChannel/nutsServer/client01/protocol/protobuf/public"
	"websocketStudy/messageChannel/nutsServer/client01/protocol/protobuf/request"

	"google.golang.org/protobuf/proto"
)

func genLevelFinish() []byte {
	bytes, err := proto.Marshal(&request.LevelFinishReq{
		Levels: []*public.LevelFinishInfo{
			{
				ChangeType: public.LevelScoreChangeType_Incr,
				//ChangeType:  request.LevelProgressChangeType_Set,
				Id:       1,
				NewScore: 1,
				Spend: []*public.Item{
					{
						Id:    100,
						Type:  1,
						Lv:    0,
						Name:  111,
						Count: 20,
					},
				},
			},
			//{
			//	ChangeType: public.LevelScoreChangeType_Set,
			//	//ChangeType:  request.LevelProgressChangeType_Set,
			//	Id:       9002,
			//	NewScore: 5,
			//},
		},
	})
	if err != nil {
		panic(err)
	}

	req := &Req{
		Identifier: WSSendMsg,
		Data: &Req_SendMsgReq{
			SendMsgReq: &SendMsgReq{
				Event: LevelProgressChange,
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
