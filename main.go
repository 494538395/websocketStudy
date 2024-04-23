package main

import (
	"fmt"
	"net/http"

	"websocketStudy/partyBattle"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		//cookies := c.Request.Cookies()
		//query := c.Request.URL.Query()
		//fmt.Println("cookies-->", cookies)
		//fmt.Println("querys-->", query)
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		//go partyBattle.TimerTask(conn)
		var data []byte

		data = partyBattle.GenPartyBattleData(3.14)
		conn.WriteMessage(websocket.BinaryMessage, data)

		data = partyBattle.GenRewardData()
		conn.WriteMessage(websocket.BinaryMessage, data)

		data = partyBattle.GenUserRankData("user01", 10, 1)
		conn.WriteMessage(websocket.BinaryMessage, data)

		for {
			messageType, p, _ := conn.ReadMessage()
			fmt.Println("收到消息-->", string(p))
			fmt.Println("messageType-->", messageType)

			//parsePartyBattle(p)
			//wfl.ParseWfl(p)

			//parsePartyBattle(p)

			//str := "接收到了：" + string(p)
			//str := string(p)

			//data := partyBattle.GenPartyBattleData(6.66)
			////data := genPartyBattleData()
			//conn.WriteMessage(messageType, data)

			//websocket.TextMessage

			data = partyBattle.GenGameEndData()
			conn.WriteMessage(websocket.BinaryMessage, data)

			//if err := conn.WriteMessage(messageType, []byte(jsonData)); err != nil {
			//	c.JSON(500, gin.H{"error": err.Error()})
			//	return
			//}
		}
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, "hello world")
	})

	r.Run(":8009")
}
