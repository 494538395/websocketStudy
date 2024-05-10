# 通用协议结构使用说明

### Proto 注释规范
最外层的proto结构需要增加规范注释 表示该结构体被哪些事件直接引用

```
//@{type} {event1} {event2} ...
```
- {type}: 消息类型, send是客户端发送到服务器的事件，push是服务器推送给客户端的
- {event}: 事件名，多个事件用空格分割

例如:
```protobuf
// @send apvp.match
message ApvpMatchReq {
  string otherId = 1;
  string attackId = 2;  // 复仇时用，传入复仇的战斗记录的attackId
  uint32 battleVer = 3; // 战斗版本号
  string battleCfgVer = 4; // 战斗配置版本号
}
```

如果既有send 又有push:
```protobuf
// @send barrack.item.clean
// @push game.barrack.item.change
message BarrackInfo {
  int32 id = 1;
  string name = 2;
  repeated BarrackItem items = 3;
  BarrackPop pop = 4;
  ItemSourceType reason = 5;
  int32 Cbt = 6; // 战力
}
```


### ItemChangeResp

```protobuf
syntax = "proto3";
message ItemChangeResp {
  repeated RewardItem changes = 1;
  repeated TransformItem transformItems = 2;
}

message RewardItem {
  int32 itemId = 1;
  int64 itemCnt = 2;
}

message TransformItem {
  RewardItem before = 1;
  RewardItem after = 2;
}
```

**如何使用RewardItem表示奖励信息**

- itemId表示奖励的[道具id](itemId.md)(带大类)，如卡牌奖励id：type*10000+subtype*100+奖励大类30
- itemCnt表示奖励的物品数量，如获得一张卡牌，itemCnt=1

**如何使用ItemChangeResp返回奖励信息的变化**

- 所有获得的最终奖励都放到changes中，但是不包含转换得到的奖励
- 所有转换后的奖励都放到transformItems中，包含转换前后的奖励



