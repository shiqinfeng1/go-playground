package main

import (
	"fmt"

	"github.com/axgle/mahonia"
)

type strBytes []byte

func (s strBytes) String() string {
	r := "["
	for _, v := range s {
		r += fmt.Sprintf(" 0x%x ", v)
	}
	r += "]"
	return r
}
func main() {
	str := "中国"
	fmt.Println("中文字符串：", str, " 长度:", len(str))
	strByte := strBytes(str)
	fmt.Printf("对应的utf-8内存编码：%s\n", strByte)
	enc := mahonia.NewEncoder("gbk")
	fmt.Printf("对应的gbk码点：%s\n", strBytes(enc.ConvertString(str))) //converts a  string from UTF-8 to gbk encoding.
	enc2 := mahonia.NewEncoder("gb18030")
	fmt.Printf("对应的gb18030码点：%s\n", strBytes(enc2.ConvertString(str))) //converts a  string from UTF-8 to gbk encoding.

}
