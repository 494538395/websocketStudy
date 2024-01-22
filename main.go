package main

import (
	"fmt"

	myProto "websocketStudy/message/proto"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var upgrader = websocket.Upgrader{}

func main() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		for {
			messageType, p, err := conn.ReadMessage()
			fmt.Println("收到消息-->", string(p))

			msg := myProto.Msg{}

			err = proto.Unmarshal(p, &msg)
			if err != nil {
				panic(err)
			}
			fmt.Println(msg)

			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			// 返回消息
			resp := myProto.PaperPlayerRank{
				PlayerID: "111",
				Score:    50,
				Rank:     1,
			}

			bytes, err := proto.Marshal(&resp)
			if err != nil {
				panic(err)
			}

			if err := conn.WriteMessage(messageType, bytes); err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		}
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, "hello world")
	})

	r.Run(":8002")
}
