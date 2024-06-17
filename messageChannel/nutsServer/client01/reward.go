package main

import (
	"websocketStudy/messageChannel/nutsServer/client01/protocol/protobuf/public"
	"websocketStudy/messageChannel/nutsServer/client01/protocol/protobuf/request"

	"google.golang.org/protobuf/proto"
)

func genAddItem() []byte {
	bytes, err := proto.Marshal(&request.AddItemReq{
		Items: []*public.Item{
			//{
			//	Id:    100,
			//	Type:  1,
			//	Lv:    0,
			//	Name:  999,
			//	Count: 20,
			//},
			{
				Id:    200,
				Type:  1,
				Lv:    0,
				Name:  999,
				Count: 10,
			},
			{
				Id:    300,
				Type:  1,
				Lv:    0,
				Name:  999,
				Count: 10,
			},
			//{
			//	Id:    400,
			//	Type:  1,
			//	Lv:    0,
			//	Name:  999,
			//	Count: 5,
			//},
			//{
			//	Id:    500,
			//	Type:  1,
			//	Lv:    0,
			//	Name:  999,
			//	Count: 5,
			//},
		},
		Source: "加奖励",
	})
	if err != nil {
		panic(err)
	}

	req := &Req{
		Identifier: WSSendMsg,
		Data: &Req_SendMsgReq{
			SendMsgReq: &SendMsgReq{
				Event: AddItem,
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

func genReduceItem() []byte {
	bytes, err := proto.Marshal(&request.ReduceItemReq{
		Items: []*public.Item{
			{
				Id:    100,
				Type:  1,
				Lv:    0,
				Name:  999,
				Count: 10,
			},
			{
				Id:    200,
				Type:  1,
				Lv:    0,
				Name:  999,
				Count: 10,
			},
			{
				Id:    300,
				Type:  1,
				Lv:    2,
				Name:  555,
				Count: 10,
			},
		},
		Source: "减资源",
	})
	if err != nil {
		panic(err)
	}

	req := &Req{
		Identifier: WSSendMsg,
		Data: &Req_SendMsgReq{
			SendMsgReq: &SendMsgReq{
				Event: ReduceItem,
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
