package main

import (
	sdk_ws "git.89trillion.com/89t/server/social-server-sdk/proto/social_sdk_ws"
	"google.golang.org/protobuf/proto"
)

func genFriendList() []byte {
	req := &sdk_ws.GetFriendListReq{
		CommID: &sdk_ws.CommID{
			OpUserID:    "19999",
			OperationID: "111",
			ToUserID:    "19999",
			FromUserID:  "19999",
			ServerID:    0,
		},
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genAddFriend() []byte {
	req := &sdk_ws.AddFriendReq{
		CommID: &sdk_ws.CommID{
			OpUserID:    "19999",
			OperationID: "111",
			ToUserID:    "19999",
			FromUserID:  "19999",
			ServerID:    0,
		},
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genDeleteFriend() []byte {
	req := &sdk_ws.DeleteFriendReq{
		CommID: &sdk_ws.CommID{
			OpUserID:    "19999",
			OperationID: "111",
			ToUserID:    "19999",
			FromUserID:  "19999",
			ServerID:    0,
		},
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genGetFriendApplyList() []byte {
	req := &sdk_ws.GetFriendApplyListReq{
		CommID: &sdk_ws.CommID{
			OpUserID:    "19999",
			OperationID: "111",
			ToUserID:    "19999",
			FromUserID:  "19999",
			ServerID:    0,
		},
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genAddFriendResponse() []byte {
	req := &sdk_ws.AddFriendResponseReq{
		CommID: &sdk_ws.CommID{
			OpUserID:    "19999",
			OperationID: "111",
			ToUserID:    "19999",
			FromUserID:  "19999",
			ServerID:    0,
		},
		HandleMsg:    "可以的",
		HandleResult: 1,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genGetBlacklist() []byte {
	req := &sdk_ws.GetBlacklistReq{
		CommID: &sdk_ws.CommID{
			OpUserID:    "11",
			OperationID: "22",
			ToUserID:    "3",
			FromUserID:  "4",
			ServerID:    0,
		},
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genAddBlacklist() []byte {
	req := &sdk_ws.AddBlacklistReq{
		CommID: &sdk_ws.CommID{
			OpUserID:    "11",
			OperationID: "22",
			ToUserID:    "3",
			FromUserID:  "4",
			ServerID:    0,
		},
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genRemoveBlacklist() []byte {
	req := &sdk_ws.RemoveBlacklistReq{
		CommID: &sdk_ws.CommID{
			OpUserID:    "11",
			OperationID: "22",
			ToUserID:    "3",
			FromUserID:  "4",
			ServerID:    0,
		},
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genAddAllFriendResponse() []byte {
	req := &sdk_ws.AddFriendResponseAllReq{
		HandleResult: 2,
		HandleMsg:    "可以的啊",
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genGetSelfApplyList() []byte {
	req := &sdk_ws.GetSelfApplyListReq{
		CommID: &sdk_ws.CommID{
			OpUserID:    "1",
			OperationID: "2",
			ToUserID:    "3",
			FromUserID:  "4",
			ServerID:    0,
		},
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genRecommendUser() []byte {
	req := &sdk_ws.GetUsersReq{
		OperationID: "2007990",
		Pagination: &sdk_ws.RequestPagination{
			PageNumber: 1,
			ShowNumber: 10,
		},
		UserName: "",
		UserID:   "",
		Content:  "",
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}
