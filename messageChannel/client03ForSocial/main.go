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

	appId         = "250"
	platformAppId = "250"

	platformId = "4"

	nacosGroupId = "social"

	basicInfo = "eyJnYWlkIjoiIiwidWlkIjoiIiwiY3ZjIjowLCJzdmMiOjAsImRldmljZSI6IiIsIm5ldHdvcmsiOiIiLCJzaW1jb2RlIjoiIiwibGFuZyI6IiIsImxzIjoiIiwicGYiOiIiLCJpcCI6IiIsImNvdW50cnkiOiIiLCJhcHBpZCI6MjUwLCJhc3YiOjB9"

	testHost = "10.0.1.84:38015"
	//debugHost = "127.0.0.1:8001"
)

func main() {
	// header 里面可以不设置东西
	header := http.Header{
		"appId": []string{"250"},
	}

	// 连接社交服, path 需要是 /ws/social
	//url := "ws://" + testHost + "/ws/social?"
	url := "ws://10.0.1.84:38015/ws/social?"
	conn, _, err := websocket.DefaultDialer.Dial(
		url+
			"token="+token2011686Test+"&"+ // user的 token,必填
			"appid="+appId+"&"+ // appId 必填,填写业务appId,比如天堂就写250
			"Basic-Info="+basicInfo+"&"+ // basicInfo
			"nacosGroup="+nacosGroupId, // 写社交服注册 rpc 时的 nacos groupId,一版是 social
		header)
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
		if char == '1' {
			sendMessage(conn, genChatMsg())
			fmt.Println("消息发送成功")
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

		if char == '6' { // 获取自己加的群列表
			sendMessage(conn, genGetGroupsInfo())
			fmt.Println("消息发送成功")
		}

		if char == '7' { // 获取自己加的群列表
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

		if char == '.' { // 获取 seq
			sendMessage(conn, genGetTopicSeqReq())
			fmt.Println("消息发送成功")
		}

		if char == '+' { // upload t
			sendMessage(conn, genUploadTopicReq())
			fmt.Println("消息发送成功")
		}
	}

	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt)
	//ticker := time.NewTicker(time.Second)
	//defer ticker.Stop()
	//for {
	//	select {
	//	case <-done:
	//		fmt.Println("调试程序结束")
	//		return
	//	case <-ticker.C:
	//		fmt.Println("连接依然存在")
	//	case <-interrupt:
	//		log.Println("interrupt")
	//
	//		// Cleanly close the connection by sending a close message and then
	//		// waiting (with timeout) for the server to close the connection.
	//		err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	//		if err != nil {
	//			log.Println("write close:", err)
	//			return
	//		}
	//		return
	//	}
	//}
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

func genChatMsg() []byte {
	req := &Req{
		Identifier: WSSendMsg,
		MsgId:      "1",
		Data: &Req_SendMsgReq{
			SendMsgReq: &SendMsgReq{
				Event: ChatSend,
				Data:  genChatSendReq(),
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

		//sd := &transfer.SendData{}
		//proto.Unmarshal(message, sd)

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
