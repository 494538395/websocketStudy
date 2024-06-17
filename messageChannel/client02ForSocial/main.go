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
	token20000Test = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIwMDAwLCJleHAiOjE3MTgwMjcwOTYsImlzcyI6Imdpbi1ibG9nIn0.F7Kdox56cA5Q7wQYho6GH5BVlqGHHrK4lYokQQIRawg"

	appId         = "250"
	platformAppId = "250"
	platformId    = "4"

	nacosGroupId = "social"

	testHost  = "10.0.1.84:38015"
	debugHost = "127.0.0.1:8001"
)

func main() {
	// header 里面可以不设置东西
	header := http.Header{}

	// 连接社交服, path 需要是 /ws/social
	url := "ws://" + testHost + "/ws/social?"
	conn, _, err := websocket.DefaultDialer.Dial(
		url+
			"token="+token20000Test+"&"+ // user的 token,必填
			"appid="+appId+"&"+ // appId 必填,填写业务appId,比如天堂就写250
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
			sendMessage(conn)
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
		fmt.Printf("接受到消息！！  recv: %s\n", message)
	}
}

func sendMessage(conn *websocket.Conn) {
	//<-time.After(3 * time.Second)

	msg := genChatMsg()
	err := conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("write:", err)
		return
	}
	fmt.Println("发送消息成功！！！！")
}
