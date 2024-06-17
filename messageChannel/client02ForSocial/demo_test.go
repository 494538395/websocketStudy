package main

import (
	"fmt"
	"testing"
	"time"

	open_im_sdk "git.89trillion.com/89t/server/social-server-sdk/proto/social_sdk_ws"
	"github.com/jinzhu/copier"
)

func WithPrefixes(str string, prefixes ...string) string {
	for i := len(prefixes) - 1; i >= 0; i-- {
		str = prefixes[i] + "_" + str
	}

	return str
}

func TestFunc(t *testing.T) {
	res := WithPrefixes("user01", "1", "2")
	fmt.Println(res)
}

func CopyStructFields(a interface{}, b interface{}, fields ...string) (err error) {
	return copier.Copy(a, b)
}

type Group struct {
	GroupID       string    `gorm:"column:group_id;primary_key;size:64" json:"groupID" bson:"group_id" binding:"required"`
	GroupName     string    `gorm:"column:name;index:name;size:255" json:"groupName" bson:"group_name"`
	Introduction  string    `gorm:"column:introduction;size:255" json:"introduction" bson:"introduction"`
	FaceURL       string    `gorm:"column:face_url;size:255" json:"faceURL" bson:"face_url"`
	CreateTime    time.Time `gorm:"column:create_time;index:create_time" bson:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time" bson:"update_time"`
	Ex            string    `gorm:"column:ex;size:1024" json:"ex" bson:"ex"`
	Status        int32     `gorm:"column:status" bson:"status"`
	CreatorUserID string    `gorm:"column:creator_user_id;size:64" bson:"creator_user_id"`
	GroupType     int32     `gorm:"column:group_type" bson:"group_type"`
}

type Person01 struct {
	id   string
	name string
}

type Person02 struct {
	id   int32
	name string
}

func TestCopy(t *testing.T) {
	origin := &Group{GroupID: "1", GroupType: 2}

	//origin := &Person02{id: 2, name: "jerry"}

	//var after Person01

	var groupInfo open_im_sdk.GroupInfo

	err := CopyStructFields(&groupInfo, origin)
	if err != nil {
		panic(err)
	}

	fmt.Println(groupInfo)

}
