package test

import (
	"fmt"
	"github.com/sandwich-go/hotswap"
	"github.com/sandwich-go/hotswap/demo/hello/g"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	g.PluginManagerSwapper = hotswap.NewPluginManagerSwapper("../bin/darwin/plugin/hello",
		hotswap.WithLogger(g.Logger),
		hotswap.WithFreeDelay(time.Second*15),
	)
	swapper := g.PluginManagerSwapper
	details, err := swapper.LoadPlugins("hello onload")
	if err != nil {
		panic(err)
	}
	fmt.Println(details)
	os.Exit(m.Run())
}

func BenchmarkInvokeMain(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		// 确保使用函数的返回值，以防止编译器优化。
		result = g.CPUAndMemoryIntensiveFunction()
	}
	// 使用测试结果来防止编译器优化。
	if result == -1 {
		b.Fatal("just to use result")
	}
}

func BenchmarkInvokeRegisterPlugin(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = g.FunctionRegisterByPlugin()
	}
	// 使用测试结果来防止编译器优化。
	if result == -1 {
		b.Fatal("just to use result")
	}
}

func BenchmarkInvokeDirectPlugin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g.PluginManagerSwapper.Current().InvokeEach("sort")
	}
}
