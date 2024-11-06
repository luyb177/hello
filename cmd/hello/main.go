// cmd/hello/main.go
package main

import (
	"flag"
	"github.com/luyb177/hello"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "world", "name to say hello")
	//func StringVar(p *string, name string, value string, usage string)
	// p 是一个指向字符串的指针，用于存储解析后的命令行参数值
	//name 命令行参数的名称
	//value 参数的默认值
	//usage 参数的使用说明
}

func main() {
	flag.Parse() //解析命令行参数
	msg := hello.Hello(name)
	_, err := os.Stdout.WriteString(msg) 
	//os.Stdout 代表标准输出
	//WriteString 是用于向这个标准输出写入一个字符串
	if err != nil {
		os.Stderr.WriteString(err.Error()) //打印错误信息
	}
}