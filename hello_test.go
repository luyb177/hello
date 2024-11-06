// hello_test.go
//go测试的标准命名结构
package hello_test

import (
	"fmt"
	"github.com/luyb177/hello"
	"testing" //Go 标准库提供的包，用来编写测试。通过 testing.T 类型，可以记录测试的结果、断言、错误等。
)

func TestHello(t *testing.T) { //测试函数的名称必须以 Test 开头
	data := "jack"
	expected := fmt.Sprintf("hello %s!", data)
	result := hello.Hello(data)

	if result != expected {
		t.Fatalf("expected result %s, but got %s", expected, result)
		//t.Fatalf 用来在断言失败时输出错误信息并终止测试。Fatalf 方法接受一个格式化字符串和变量，会将错误信息输出到测试结果中，并停止执行当前的测试函数。
	}

}
