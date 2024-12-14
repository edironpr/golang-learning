package benchmark

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"
)

// 函数名必须以 Benchmark 开头，后面一般跟待测试的函数名
// 参数为 b *testing.B
// 执行基准测试时，需要添加 -bench 参数
func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

// 基准测试报告每一列值对应的含义如下
//type BenchmarkResult struct {
//	N         int           // 迭代次数
//	T         time.Duration // 基准测试花费的时间
//	Bytes     int64         // 一次迭代处理的字节数
//	MemAllocs uint64        // 总的分配内存的次数
//	MemBytes  uint64        // 总的分配内存的字节数
//}

// 如果在运行前基准测试需要一些耗时的配置，则可以使用 b.ResetTimer() 先重置定时器
func BenchmarkHello(b *testing.B) {
	//... // 耗时操作

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

// 使用 RunParallel 测试并发性能
func BenchmarkParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			// 所有 goroutine 一起，循环一共执行 b.N 次
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}
