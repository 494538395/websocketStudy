package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	sdk_ws "git.89trillion.com/89t/server/social-server-sdk/proto/social_sdk_ws"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var (
	token15555     = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjE1NTU1LCJleHAiOjE3MTY2Mzk2NDIsImlzcyI6Imdpbi1ibG9nIn0.eNSq0ADd2Un2AjfWtuTFGyf0jCYCkrfgiXpiXeMof28"
	token19999     = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjE5OTk5LCJleHAiOjE3MTY2Mzk2ODYsImlzcyI6Imdpbi1ibG9nIn0.ZWAkjtjuBvNWHP1wYluAc3TCcfB6JKZQIbvjHlMR_y8"
	token19999Test = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjE5OTk5LCJleHAiOjE3MTc2Njg5MjQsImlzcyI6Imdpbi1ibG9nIn0.8K8crqQWvyImJ_us0qXyUYfBQPR0lW94FDf9ruOO-kg"

	appId         = "250"
	platformAppId = "250"
	platformId    = "4"

	nacosGroupId = "social02"

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
			"token="+token19999Test+"&"+ // user的 token,必填
			"appid="+appId+"&"+ // appId 必填,填写业务appId,比如天堂就写250
			"nacosGroup="+nacosGroupId, // 写社交服注册 rpc 时的 nacos groupId,一版是 social
		header)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

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
		MsgId:      "1",
		Data: &Req_SendMsgReq{
			SendMsgReq: &SendMsgReq{
				Event: "group.create",
				Data:  genCreateGroupReq(),
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

func genCreateGroupReq() []byte {
	createGroupReq := &sdk_ws.CreateGroupReq{
		OpUserID:    "user01",
		OwnerUserID: "jerry",
	}

	bytes, err := proto.Marshal(createGroupReq)
	if err != nil {
		panic(err)
	}

	return bytes
}
