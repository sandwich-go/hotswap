package g

import (
	"github.com/edwingeng/slog"
	"github.com/sandwich-go/hotswap"
	"math/rand"
)

var (
	Logger = slog.NewDevelopmentConfig().MustBuild()
)

var (
	PluginManagerSwapper *hotswap.PluginManagerSwapper
)

// CPUAndMemoryIntensiveFunction 是一个既耗费CPU又耗费内存的函数。
// 它将执行一定数量的随机数学计算，并分配一个较大的内存空间来存储结果。
func CPUAndMemoryIntensiveFunction() int {
	const size = 100 // 定义切片大小为100来模拟较高的内存使用。
	nums := make([]int, size)

	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(size)
	}

	// 执行计算密集型操作：简单排序算法（冒泡排序，仅用于示例）
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	// 执行一些计算密集型操作。
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum // 返回计算结果
}

var FunctionRegisterByPlugin func() int
