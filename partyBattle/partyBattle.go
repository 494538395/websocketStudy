package partyBattle

import (
	"fmt"
	"time"

	myProto "websocketStudy/message/proto"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func parsePartyBattle(data []byte) {
	// 1.解析最外层
	var t myProto.Packet

	err := proto.Unmarshal(data, &t)
	if err != nil {
		panic(err)
	}

	// 2.解析里层

	var biz myProto.HamsterChangeDirection
	err = proto.Unmarshal(t.Data, &biz)
	if err != nil {
		panic(err)
	}

	fmt.Println(biz)
}

func GenPartyBattleData(direction float32) []byte {
	var biz myProto.HamsterChangeDirection
	biz.NewDirection = direction
	biz.UserID = "user01"

	bytes, err := proto.Marshal(&biz)
	if err != nil {
		panic(err)
	}

	// 1.解析最外层
	var t myProto.Packet
	t.Id = 1001
	t.Event = "hamster.change.direction"
	t.Data = bytes

	data, err := proto.Marshal(&t)
	if err != nil {
		panic(err)
	}

	return data
}

func GenRewardData() []byte {
	var biz myProto.HamsterGetReward
	biz.UserID = "user01"
	biz.ItemId = 101
	biz.Count = 20

	bytes, err := proto.Marshal(&biz)
	if err != nil {
		panic(err)
	}

	// 1.解析最外层
	var t myProto.Packet
	t.Id = 1001
	t.Event = "hamster.get.reward"
	t.Data = bytes

	data, err := proto.Marshal(&t)
	if err != nil {
		panic(err)
	}

	return data
}

func GenGameEndData() []byte {

	var rank []*myProto.HamsterUserRank
	for i := 0; i < 3; i++ {
		rank = append(rank, &myProto.HamsterUserRank{
			UserID: fmt.Sprintf("user0%d", i+1),
			Score:  int32(i + 10),
			Rank:   int32(i + 1),
		})
	}

	var biz myProto.HamsterGameEnd
	biz.Winner = "user01"
	biz.Rank = rank

	bytes, err := proto.Marshal(&biz)
	if err != nil {
		panic(err)
	}

	// 1.解析最外层
	var t myProto.Packet
	t.Id = 1001
	t.Event = "hamster.game.end"
	t.Data = bytes

	data, err := proto.Marshal(&t)
	if err != nil {
		panic(err)
	}

	return data
}

func GenUserRankData(user string, score int32, rank int32) []byte {
	var biz myProto.HamsterUserRank
	biz.UserID = user
	biz.Score = score
	biz.Rank = rank

	bytes, err := proto.Marshal(&biz)
	if err != nil {
		panic(err)
	}

	// 1.解析最外层
	var t myProto.Packet
	t.Id = 1001
	t.Event = "hamster.user.rank"
	t.Data = bytes

	data, err := proto.Marshal(&t)
	if err != nil {
		panic(err)
	}

	return data
}

func TimerTask(conn *websocket.Conn) {
	go func() {
		for {
			select {
			case <-time.After(2 * time.Second):
				var err error
				err = conn.WriteMessage(websocket.BinaryMessage, GenRewardData())
				if err != nil {
					return
				}
				err = conn.WriteMessage(websocket.BinaryMessage, GenPartyBattleData(6.99))
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
