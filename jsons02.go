package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	var f interface{}
	json.Unmarshal(b, &f)

	m := f.(map[string]interface{})
	fmt.Println(m["Parents"])  // 读取 json 内容
	fmt.Println(m["a"] == nil) // 判断键是否存在
}

//我们可以使用 interface 接收 json.Unmarshal 的结果，
//然后利用 type assertion 特性 (把解码结果转换为 map[string]interface{} 类型) 来进行后续操作。
