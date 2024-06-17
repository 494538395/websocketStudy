package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type Param struct {
	Desc string      `json:"desc"`
	Data interface{} `json:"data"`
}

func main() {
	p := &Person{
		Name: "jerry",
		Age:  21,
	}
	bytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println("bytes-->", bytes)

	parm := &Param{
		Data: &Person{},
	}

	// 获取 Data 字段的类型
	dataType := reflect.TypeOf(parm.Data)
	fmt.Println("Data type:", dataType)

	// 利用反射创建 Data 字段的指针实例
	dataPtr := reflect.New(dataType.Elem()).Interface()

	// 反序列化 bytes 到 dataPtr 中
	err = json.Unmarshal(bytes, dataPtr)
	if err != nil {
		panic(err)
	}

	// 将反序列化后的数据设置到 Data 字段中
	parm.Data = dataPtr

	// 打印解码后的数据
	fmt.Println("Decoded data:", parm.Data)
}
