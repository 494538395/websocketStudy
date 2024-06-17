package main

import (
	group_sdk "git.89trillion.com/89t/server/social-server-sdk/proto/social_group"
	sdk_ws "git.89trillion.com/89t/server/social-server-sdk/proto/social_sdk_ws"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (

	// caht
	ChatSend = "chat.send"

	CreateGroup               = "group.create"
	JoinGroup                 = "group.join"
	SetGroupInfo              = "group.set.group.info"
	QuitGroup                 = "group.quit.group"
	TransferGroupOwner        = "group.transfer"
	CheckGroupAuth            = "group.check"
	GroupApplicationResponse  = "group.application.response"
	GetUserReqApplicationList = "group.req.application.list"
	GetJoinedGroupList        = "group.joined.list"
	GetGroupMembersInfo       = "group.member.list"
	InviteUserToGroup         = "group.invite.users"
	KickGroupMember           = "group.kick.user"
	GetGroupsInfo             = "group.info"
	SetGroupMemberInfo        = "group.member.info.set"
	GetGroupMemberList        = "group.all.member.list"
	RecommendGroupsBt         = "group.recommend.bt"
	RecommendGroupsClan       = "group.recommend.clan"
	MuteSwitchGroup           = "group.mute.switch"
	InviteSwitchGroup         = "group.invite.switch"
	RecommendSwitchGroup      = "group.recommend.switch"
	AddUserToGroup            = "group.add.users"
	GroupInviteResponse       = "group.invite.response"
	AddGroupRole              = "group.role.add"
	UpdateGroupRole           = "group.role.update"
	SearchGroups              = "group.search"
	AutoTransferGroup         = ""

	// friend
	FriendList           = "friend.list"
	AddFriend            = "friend.add"
	DeleteFriend         = "friend.delete"
	GetFriendApplyList   = "friend.apply.list"
	AddFriendResponse    = "friend.response"
	GetBlacklist         = "friend.black.list"
	AddBlacklist         = "friend.black.add"
	RemoveBlacklist      = "friend.black.delete"
	AddAllFriendResponse = "friend.response.all"
	GetSelfApplyList     = "friend.self.apply.list"
	RecommendUser        = "user.recommend"
)

func genCreateGroupReq() []byte {
	createGroupReq := &sdk_ws.CreateGroupReq{
		OpUserID:    "19999",
		OwnerUserID: "jerry",
		GroupInfo: &sdk_ws.GroupInfo{
			GroupID:     "1001",
			GroupName:   "jerry-group",
			OwnerUserID: "jerry",
		},
		InitMemberList: []*sdk_ws.GroupAddMemberInfo{
			{
				UserID:    "19998",
				RoleLevel: 0,
			},
			{
				UserID:    "19997",
				RoleLevel: 0,
			},
		},
	}
	bytes, err := proto.Marshal(createGroupReq)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genJoinGroupReq() []byte {
	joinReq := &sdk_ws.JoinGroupReq{
		GroupID:       "13",
		ReqMessage:    "我要加群",
		OpUserID:      "19999",
		JoinSource:    0,
		InviterUserID: "19999",
		GroupType:     0,
	}
	bytes, err := proto.Marshal(joinReq)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genSetGroupInfoReq() []byte {
	setGroupReq := &sdk_ws.SetGroupInfoReq{
		GroupInfoForSet: &sdk_ws.GroupInfoForSet{
			GroupID:   "13",
			GroupName: "jerry的group",
		},
		OpUserID: "1999",
	}
	bytes, err := proto.Marshal(setGroupReq)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genQuitGroup() []byte {
	quitReq := &sdk_ws.QuitGroupReq{
		GroupID:   "13",
		OpUserID:  "19997",
		GroupType: 0,
	}
	bytes, err := proto.Marshal(quitReq)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genTransferGroupOwner() []byte {
	setGroupReq := &sdk_ws.TransferGroupOwnerReq{
		GroupID:        "13",
		OldOwnerUserID: "19999",
		NewOwnerUserID: "19997",
		GroupType:      0,
	}
	bytes, err := proto.Marshal(setGroupReq)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genCheckGroupAuth() []byte {
	setGroupReq := &group_sdk.GroupCheckAuthReq{
		GroupID:   "13",
		UserId:    "19998",
		RoleType:  1,
		GroupType: 2,
	}
	bytes, err := proto.Marshal(setGroupReq)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genGroupApplicationResponse() []byte {
	setGroupReq := &sdk_ws.GroupApplicationResponseReq{
		OpUserID:     "19999",
		GroupID:      "13",
		FromUserID:   "19996",
		HandledMsg:   "拒绝",
		HandleResult: 1,
		GroupType:    2,
	}
	bytes, err := proto.Marshal(setGroupReq)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genGetUserReqApplicationList() []byte {
	setGroupReq := &sdk_ws.GetUserReqApplicationListReq{
		UserID:    "19999",
		GroupType: 2,
	}
	bytes, err := proto.Marshal(setGroupReq)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genGetJoinedGroupList() []byte {
	req := &sdk_ws.GetJoinedGroupListReq{
		FromUserID:  "19999",
		OperationID: "101",
		OpUserID:    "19999",
		GroupType:   2,
		ServerID:    250,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genGetGroupMembersInfo() []byte {
	req := &sdk_ws.GetGroupMembersInfoReq{
		GroupID:     "13",
		MemberList:  []string{"19998"},
		OpUserID:    "19999",
		OperationID: "101",
		NoCache:     true,
		GroupType:   2,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genInviteUserToGroup() []byte {
	req := &sdk_ws.InviteUserToGroupReq{
		OperationID:       "101",
		GroupID:           "13",
		InvitedUserIDList: []string{"19995"},
		GroupType:         2,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genKickGroupMember() []byte {
	req := &sdk_ws.KickGroupMemberReq{
		GroupID:          "13",
		KickedUserIDList: []string{"19996"},
		Reason:           "要踢你",
		OperationID:      "101",
		OpUserID:         "19999",
		GroupType:        2,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genGetGroupsInfo() []byte {
	req := &sdk_ws.GetGroupsInfoReq{
		GroupID:     "3",
		OperationID: "1717084792322",
		OpUserID:    "2007990",
		GroupType:   2,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genSetGroupMemberInfo() []byte {
	req := &sdk_ws.SetGroupMemberInfoReq{
		GroupID:     "13",
		UserID:      "19996",
		OpUserID:    "19999",
		OperationID: "11",
		Nickname:    wrapperspb.String("我的昵称"),
		GroupType:   2,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genGetGroupMemberList() []byte {
	req := &sdk_ws.GetGroupMemberListReq{
		GroupID:     "13",
		OpUserID:    "19999",
		OperationID: "111",
		Filter:      2,
		NextSeq:     1,
		GroupType:   2,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genRecommendGroupsBt() []byte {
	req := &sdk_ws.RecommendGroupsReq{
		OpUserID:    "19999",
		OperationID: "11",
		GroupType:   2,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genRecommendGroupsClan() []byte {
	req := &sdk_ws.RecommendGroupsReq{
		OpUserID:    "19999",
		OperationID: "11",
		GroupType:   2,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genMuteSwitchGroup() []byte {
	req := &sdk_ws.GroupSwitchReq{
		OpUserID:    "19999",
		OperationID: "1112",
		GroupID:     "13",
		Switch:      1,
		GroupType:   2,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genInviteSwitchGroup() []byte {
	req := &sdk_ws.GroupSwitchReq{
		OpUserID:    "19999",
		OperationID: "1112",
		GroupID:     "13",
		Switch:      1,
		GroupType:   2,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genRecommendSwitchGroup() []byte {
	req := &sdk_ws.GroupSwitchReq{
		OpUserID:    "19999",
		OperationID: "2222",
		GroupID:     "13",
		Switch:      1,
		GroupType:   2,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genAddUserToGroup() []byte {
	req := &sdk_ws.AddUserToGroupReq{
		OperationID:       "123",
		GroupID:           "13",
		Reason:            "可以的",
		InvitedUserIDList: []string{"19995"},
		OpUserID:          "19999",
		GroupType:         2,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genGroupInviteResponse() []byte {
	req := &sdk_ws.ResponseBeInvitedToGroupReq{
		OperationID:  "111",
		GroupID:      "13",
		GroupType:    2,
		HandleResult: 2,
		FromUser:     "19999",
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genAddGroupRole() []byte {
	req := &sdk_ws.GroupRoleAddReq{
		OperationID: "111",
		GroupID:     "13",
		RoleInfo: &sdk_ws.RoleInfo{
			RoleId:      1,
			RoleName:    "管理员",
			RolePower:   99,
			NotEditable: 0,
		},
		GroupType: 2,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genUpdateGroupRole() []byte {
	req := &sdk_ws.GroupRoleUpdateReq{
		OperationID: "111",
		GroupID:     "13",
		RoleInfo: &sdk_ws.RoleInfo{
			RoleId:      1,
			RoleName:    "管理员",
			RolePower:   12,
			NotEditable: 0,
		},
		GroupType: 2,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}

func genSearchGroups() []byte {
	req := &sdk_ws.SearchGroupsReq{
		KeyWord:   "1群",
		GroupType: 2,
	}
	bytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes
}
