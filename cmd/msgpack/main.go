package main

import (
	"fmt"
	"reflect"

	"github.com/vmihailenco/msgpack"
)

func main() {
	countryCapitalMap := make(map[string]string)

	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"
	fmt.Println("原数据-", countryCapitalMap)

	in := countryCapitalMap
	res, err := msgpack.Marshal(in)
	if err != nil {
		fmt.Printf("序列化失败")
	}
	//fmt.Sprintf("数据类型%T", b)
	fmt.Println()
	fmt.Println("序列化数据类型", reflect.TypeOf(res), "\n序列化数据--", res)

	var out map[string]string
	msgpack.Unmarshal(res, &out)

	fmt.Println("反序列化数据--", out)

}
