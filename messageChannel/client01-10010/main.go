package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var (
	token10010Test = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEwMDEwLCJleHAiOjE3MTc1NzM3NDksImlzcyI6Imdpbi1ibG9nIn0.SafbtL-lBeFEis2Mfccweav2BC0JrYvFivY710a87k0"

	appId = "223"
	//appId = "666"

	platformAppId = "223"
	//platformAppId = "6001"

	platformId = "4"

	nacosGroupId = "jack-test"

	testHost  = "10.0.1.84:38015"
	debugHost = "127.0.0.1:8001"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// 创建一个 http.Header 对象，并设置需要的头部信息
	header := http.Header{}
	header.Set("appId", platformAppId)

	url := "ws://" + testHost + "/ws?"

	conn, _, err := websocket.DefaultDialer.Dial(
		url+
			"token="+token10010Test+"&"+
			"platformID=pc&"+
			"operationID="+platformId+"&"+
			"appId="+appId+"&"+
			"nacosGroup="+nacosGroupId,
		header)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			fmt.Printf("接受到消息！！  recv: %s\n", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	msg := genProtoMsg()
	err = conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("write:", err)
		return
	}

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			fmt.Println(t)
			//msg := genProtoMsg()
			//err := conn.WriteMessage(websocket.TextMessage, msg)
			//if err != nil {
			//	log.Println("write:", err)
			//	fmt.Println(t)
			//	return
			//}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func genProtoMsg() []byte {
	req := &Req{
		Identifier: WSSendMsg,
		Operation:  "1001",
		MsgId:      "1",
		Data: &Req_SendMsgReq{
			SendMsgReq: &SendMsgReq{
				Event: "test",
				Data:  []byte("hello world"),
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
