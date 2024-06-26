package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var (
	token15555     = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjE1NTU1LCJleHAiOjE3MTY2Mzk2NDIsImlzcyI6Imdpbi1ibG9nIn0.eNSq0ADd2Un2AjfWtuTFGyf0jCYCkrfgiXpiXeMof28"
	token19999     = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjE5OTk5LCJleHAiOjE3MTY2Mzk2ODYsImlzcyI6Imdpbi1ibG9nIn0.ZWAkjtjuBvNWHP1wYluAc3TCcfB6JKZQIbvjHlMR_y8"
	token19999Test = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjE5OTk5LCJleHAiOjE3MTgwMjM3MDUsImlzcyI6Imdpbi1ibG9nIn0.JBUc8rKumsKJtsOjysMhFahaj_tpWr5b8fYj8ghiCs8"

	token2011642Test = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIwMTE2NDIsImV4cCI6MTcxODI3MjQwNSwiaXNzIjoiZ2luLWJsb2cifQ.rHufYrnwf-CFV1IURq7LC7rIEL1lkvQfLuRWxkTvW6s"
	token2011686Test = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIwMTE2ODYsImV4cCI6MTcxODM3MDg0MSwiaXNzIjoiZ2luLWJsb2cifQ.GZvX_rN4lswThcZhLJDL_7Pm1RKTMTWXLNGNbR72eHw"

	token3001Release = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjMwMDEsImV4cCI6MTcxOTA2OTQxMywiaXNzIjoiZ2luLWJsb2cifQ.feVxljgxrHJDkW8uHH9V5Wa4y5oqkHg3x0BHHNCt-2Q"

	//tokenTC = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjI2ODUzODI3NzMsImV4cCI6MTcxOTEyNzQ2NH0.HUQPrSc-PHXMfPlul-FFBSvPU2Q7YRPSpp9YeZec6EQ"
	tokenTC     = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjMwMDEsImV4cCI6MTcxOTE0MzUwNCwiaXNzIjoiZ2luLWJsb2cifQ.eMa8G0JXVXZcHPrhLn6mGBwwzfnYtH3OiFaTLTfSqcM"
	tokenTc3002 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjMwMDIsImV4cCI6MTcxOTM5MzE1MywiaXNzIjoiZ2luLWJsb2cifQ.0bz0lnu7lA0xFdy5NO4Rd6a9qRnBGUZOkh3gcYFmrd0"
	tokenTc3003 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjMwMDMsImV4cCI6MTcxOTU1OTA3MiwiaXNzIjoiZ2luLWJsb2cifQ.PSJrwYaqAPd6jVIAWfRkFIGHLryt2zeXCVEG9cJlyCQ"

	tokenIW2786448 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjI3ODY0NDgsImlzcyI6Imdpbi1ibG9nIn0.5eZbFv85K2t6Iazw-esGNpfqtjnjw-t8TkaAGFaJv_A"
	tokenIw2001    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIwMDEsImV4cCI6MTcxOTU0NTkzNCwiaXNzIjoiZ2luLWJsb2cifQ.ePMYna1r_vV_0IZG3ezQMvQ5bhQ4Fv2WfDsCHC2hV-s"
	tokenIW2786999 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjI3ODY5OTksImV4cCI6MTcxOTU0NTkzNCwiaXNzIjoiZ2luLWJsb2cifQ.9U-DNF5dEDYGnTd8jfYKJyKmUsmbyxVPQTZSRKcsqhg"
	tokenIW9001    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjkwMDEsImV4cCI6MTcxOTU2OTMzOCwiaXNzIjoiZ2luLWJsb2cifQ.kEyMr2PVIeeJNf6GV6mtrmoNUTCXuvRFPw40cew6c_Q"
	tokenIW9004    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjkwMDQsImV4cCI6MTcxOTcyOTMzNCwiaXNzIjoiZ2luLWJsb2cifQ.QRUMuk8wR_FT_KGVVI0AbNailbPV2PSbHuc8qgJYOHs"
	tokenIW9003    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjkwMDMsImV4cCI6MTcyMDAxNTg2MywiaXNzIjoiZ2luLWJsb2cifQ.qAsnwr0jh9a9jJsHysTxTaBkgMem8pkwbCGYZG_bxhg"

	tokenBH1001 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEwMDEsImV4cCI6MTcyMDM0ODczMSwiaXNzIjoiZ2luLWJsb2cifQ.wnF3sP85GA--UVThJ-OWn5bemUkq9mD2Pftl8tPc7V8"
	tokenBH2001 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIwMDEsImV4cCI6MTcyMDM0OTI1MiwiaXNzIjoiZ2luLWJsb2cifQ.0_UYiMKNcd7_aRshXNu7RsDyuYGIjHhTDfpZ90GGSFk"

	paradiseAppId = "250"
	tcAppId       = "234"
	iwAppId       = "246"
	beehiveAppId  = "260"

	platformAppId = "250"
	platformId    = "4"

	//nacosGroupId = "social-local"
	nacosGroupId = "social"
	//nacosGroupId = "social-feichuan"

	testHost = "10.0.1.84:38015"
	//debugHost = "10.11.14.7:8001"

	releaseHost = "gate.paradise89.com"
)

func main() {
	// header 里面可以不设置东西
	header := http.Header{}

	// 连接社交服, path 需要是 /ws/social
	url := "ws://" + testHost + "/ws/social?"
	url += "token=" + tokenBH2001 + "&" +
		"appId=" + beehiveAppId + "&" +
		"nacosGroup=" + nacosGroupId
	conn, _, err := websocket.DefaultDialer.Dial(
		url,
		//"token="+token2011686Test+"&"+ // user的 token,必填
		//"appid="+paradiseAppId+"&"+ // paradiseAppId 必填,填写业务appId,比如天堂就写250
		//"nacosGroup="+nacosGroupId, // 写社交服注册 rpc 时的 nacos groupId,一版是 social
		header,
	)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	done := make(chan struct{})
	go receiveMessage(conn, done)
	//go sendMessage(conn)

	// 开始监听键盘输入
	reader := bufio.NewReader(os.Stdin)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println("Error reading from keyboard:", err)
			return
		}

		// 如果输入的字符是数字1，则打印"hello"
		if char == '0' {
			sendMessage(conn, genSingleChatMsg())
			fmt.Println("私聊")
		}
		if char == '1' {
			sendMessage(conn, genGroupChatMsg())
			fmt.Println("群聊")
		}
		if char == '2' { // 群组推荐
			sendMessage(conn, genRecommendGroupsClan())
			fmt.Println("消息发送成功")
		}
		if char == '3' { // 获取自己能管理的群组 入群申请列表
			sendMessage(conn, genGetGroupApplicationList())
			fmt.Println("消息发送成功")
		}

		if char == '4' { // 获取自己申请加群的列表
			sendMessage(conn, genGetUserReqApplicationList())
			fmt.Println("消息发送成功")
		}

		if char == '5' { // 获取自己加的群列表
			sendMessage(conn, genGetJoinedGroupList())
			fmt.Println("消息发送成功")
		}

		if char == '6' { // 设置群信息
			sendMessage(conn, genSetGroupInfoReq())
			fmt.Println("消息发送成功")
		}

		if char == '7' { // 群推荐
			sendMessage(conn, genRecommendGroupsClan())
			fmt.Println("消息发送成功")
		}

		if char == '8' { // 搜索公共
			sendMessage(conn, genSearchGroups())
			fmt.Println("消息发送成功")
		}

		if char == '9' { // 创建群组
			sendMessage(conn, genCreateGroupReq())
			fmt.Println("消息发送成功")
		}

		if char == '.' { // 加
			sendMessage(conn, genJoinGroupReq())
			fmt.Println("消息发送成功")
		}

		if char == 'q' { // 加
			sendMessage(conn, genConvListReq())
			fmt.Println("消息发送成功")
		}

		if char == 'w' { // 导出聊天记录
			sendMessage(conn, genConvListReq())
			fmt.Println("消息发送成功")
		}
	}
}

func genProtoMsg() []byte {
	req := &Req{
		Identifier: WSSendMsg,
		MsgId:      "1",
		Data: &Req_SendMsgReq{
			SendMsgReq: &SendMsgReq{
				Event: GetGroupsInfo,
				Data:  genGetGroupsInfo(),
				Topic: 11,
			},
		},
	}

	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}

	return bytes
}

func genGroupChatMsg() []byte {
	req := &Req{
		Identifier: WSSendMsg,
		MsgId:      "1",
		Data: &Req_SendMsgReq{
			SendMsgReq: &SendMsgReq{
				Event: ChatSend,
				Data:  genGroupChatSendReq(),
				Topic: 11,
			},
		},
	}

	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}

	return bytes
}

func genSingleChatMsg() []byte {
	req := &Req{
		Identifier: WSSendMsg,
		MsgId:      "1",
		Data: &Req_SendMsgReq{
			SendMsgReq: &SendMsgReq{
				Event: ChatSend,
				Data:  genSingleChatSendReq(),
				Topic: 11,
			},
		},
	}

	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}

	return bytes
}

func genSeqProtoMsg() []byte {
	req := &Req{
		Identifier: WSGetNewestSeq,
		MsgId:      "1",
		Data: &Req_MaxAndMinSeqReq{MaxAndMinSeqReq: &GetMsgMaxAndMinSeqReq{
			TopicList: []int32{1, 2},
		}},
	}

	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}

	return bytes
}

func genSeqUploadProtoMsg() []byte {
	req := &Req{
		Identifier: WSPullMsgBySeqList,
		MsgId:      "1",
		Data: &Req_MaxAndMinSeqReq{MaxAndMinSeqReq: &GetMsgMaxAndMinSeqReq{
			TopicList: []int32{1, 2},
		}},
	}

	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}

	return bytes
}

func receiveMessage(conn *websocket.Conn, done chan struct{}) {
	defer close(done)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read读到异常:", err)
			return
		}
		fmt.Printf("接受到消息！！  recv: %s\n", message)
	}
}

func sendMessage(conn *websocket.Conn, msg []byte) {
	//<-time.After(3 * time.Second)

	err := conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("write:", err)
		return
	}
	fmt.Println("发送消息成功！！！！")
}
