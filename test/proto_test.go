package test

import (
	"encoding/base64"
	"fmt"
	"testing"

	myProto "websocketStudy/message/proto"

	"google.golang.org/protobuf/proto"
)

func TestMarshalProto(t *testing.T) {

	msg := &myProto.Msg{}
	msg.Event = "battle"
	msg.Type = 1
	msg.Seq = "seq01"

	bytes, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}

	// 使用 base64 进行编码
	encodedMsg := base64.StdEncoding.EncodeToString(bytes)

	fmt.Println(encodedMsg)

}
