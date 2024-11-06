// hello.go
package hello

import "fmt"

// Hello returns hello message
func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	//func Sprintf(format string, a ...any) string
	//返回值是一个格式化后的字符串
	return fmt.Sprintf("hello %s!", name)
}
