package wfl

import (
	"fmt"
	"time"

	myProto "websocketStudy/message/proto"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func ParseWfl(data []byte) {
	// 1.解析最外层
	var t myProto.Req

	err := proto.Unmarshal(data, &t)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	req := t.GetSendMsgReq()

	var bossProto myProto.BossLevelReq

	// 处理 SendMsgReq 类型的消息
	err = proto.Unmarshal(req.Data, &bossProto)
	if err != nil {
		panic(err)
	}
	fmt.Println("bossProto-->", bossProto)

	fmt.Println(req)

	switch data := t.Data.(type) {
	case *myProto.Req_SendMsgReq:
		// 2.解析里层

		fmt.Println("myProto.Req_SendMsgReq-->", data)

	default:
		// 如果不是以上任何一种类型，则进行默认处理
		fmt.Println("Received unknown message type")

	}
}

// GenWflBigData 生成很大的 数据
func GenWflBigData() []byte {
	// 内层 Proto
	respBytes, err := proto.Marshal(&myProto.Resp{
		Identifier: 1,
		Operation:  "admin",
		MsgId:      "1001",
		Data: &myProto.Resp_MaxAndMinSeqResp{
			MaxAndMinSeqResp: &myProto.GetMsgMaxAndMinSeqResp{
				ConvMaxAndMinSeq: genMinSeqRespMap(200),
				ErrCode:          200,
				ErrMsg:           "jerry",
			},
		}})
	if err != nil {
		panic(err)
	}

	// 最外层 Proto
	var respProto myProto.Resp

	err = proto.Unmarshal(respBytes, &respProto)
	if err != nil {
		panic(err)
	}
	fmt.Println("respProto-->", respProto)

	return respBytes
}

// GenWflSmallData 生成小的 数据
func GenWflSmallData() []byte {
	// 内层 Proto
	respBytes, err := proto.Marshal(&myProto.Resp{
		Identifier: 1,
		Operation:  "admin",
		MsgId:      "1001",
		Data: &myProto.Resp_MaxAndMinSeqResp{
			MaxAndMinSeqResp: &myProto.GetMsgMaxAndMinSeqResp{
				ConvMaxAndMinSeq: genMinSeqRespMap(3),
				ErrCode:          200,
				ErrMsg:           "jerry",
			},
		}})
	if err != nil {
		panic(err)
	}

	// 最外层 Proto
	var respProto myProto.Resp

	err = proto.Unmarshal(respBytes, &respProto)
	if err != nil {
		panic(err)
	}
	//fmt.Println("respProto-->", respProto)

	return respBytes
}

// GenWflEventlData 生成小的 数据
func GenWflEventlData() []byte {
	// 内层
	bytes, err := proto.Marshal(&myProto.SenkTopicSync{TableStage: 3})
	if err != nil {
		panic(err)
	}

	// 内层 Proto
	respBytes, err := proto.Marshal(&myProto.Resp{
		Identifier: 1,
		Operation:  "admin",
		MsgId:      "1001",
		Data: &myProto.Resp_PushData{
			PushData: &myProto.GeneralMsgData{
				Event:   "topic.senk.sync",
				Topic:   myProto.TopicType_EVENT,
				SeqId:   103335,
				Data:    bytes,
				ErrCode: 1,
				ErrMsg:  "ok",
			},
		}})
	if err != nil {
		panic(err)
	}

	// 最外层 Proto
	var respProto myProto.Resp

	err = proto.Unmarshal(respBytes, &respProto)
	if err != nil {
		panic(err)
	}
	//fmt.Println("respProto-->", respProto)

	return respBytes
}

// 模拟 GameServer 返回的数据
func genWflData() []byte {
	bytes, err := proto.Marshal(&myProto.RecruitQueryPoolResp{
		Id:          10001,
		CardIds:     []int32{1, 2, 3},
		LuckCardIds: []int32{9, 9, 9},
	})
	if err != nil {
		panic(err)
	}

	respBytes, err := proto.Marshal(&myProto.Resp{
		Identifier: 1,
		Operation:  "admin",
		MsgId:      "1001",
		Data: &myProto.Resp_PushData{PushData: &myProto.GeneralMsgData{
			Event: "recruit.query.pool",
			Topic: 1,
			SeqId: 500,
			Data:  bytes,
		}},
	})
	if err != nil {
		panic(err)
	}

	var tt myProto.Resp

	err = proto.Unmarshal(respBytes, &tt)
	if err != nil {
		panic(err)
	}
	fmt.Println("tt-->", tt)

	return respBytes

	return nil
}

func genMinSeqRespMap(count int) map[int32]*myProto.MaxAndMinSeq {
	// 构造 map
	m := make(map[int32]*myProto.MaxAndMinSeq)
	for i := 0; i < count; i++ {
		m[int32(i+1)] = &myProto.MaxAndMinSeq{
			MaxSeq: uint32(i + 1),
			MinSeq: uint32(i + 1),
		}
	}

	return m
}

func genWflDataWithSimple() []byte {
	bytes, err := proto.Marshal(&myProto.PullMsg{
		ModelName: "jerry",
	})

	respBytes, err := proto.Marshal(&myProto.Resp{
		Identifier: 1,
		Operation:  "admin",
		MsgId:      "1001",
		Data: &myProto.Resp_PushData{PushData: &myProto.GeneralMsgData{
			Event: "recruit.query.pool",
			Topic: 1,
			SeqId: 500,
			Data:  bytes,
		}},
	})
	if err != nil {
		panic(err)
	}

	var tt myProto.Resp

	err = proto.Unmarshal(respBytes, &tt)
	if err != nil {
		panic(err)
	}
	fmt.Println("tt-->", tt)

	return respBytes

	return nil
}

func TimerTask(conn *websocket.Conn) {
	go func() {
		for {
			select {
			case <-time.After(500 * time.Millisecond):
				var err error
				err = conn.WriteMessage(websocket.BinaryMessage, GenWflEventlData())
				if err != nil {
					return
				}
				err = conn.WriteMessage(websocket.BinaryMessage, GenWflSmallData())
				if err != nil {
					return
				}

				//err = conn.WriteMessage(websocket.BinaryMessage, GenWflBigData())
				//if err != nil {
				//	return
				//}
				//
				//conn.WriteMessage(websocket.TextMessage, []byte(jsonData))

			}
		}
	}()
}
