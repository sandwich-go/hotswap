package sdk

import (
	"time"

	"github.com/sandwich-go/hotswap"
	"github.com/sandwich-go/logbus"
)

var swapperManager *hotswap.PluginManagerSwapper

// MustInit initializes the plugin manager.
// MountDir 是在pmt里设置的磁盘挂载目录, 只对k8s环境生效
// 如果不是k8s环境, 忽略mountDir, 使用相对路径 bin/plugin 加载plugin
func MustInit(spec *PluginSpec) {
	pluginDir := initWatchDir(spec)
	swapper := hotswap.NewPluginManagerSwapper(pluginDir,
		hotswap.WithLogger(NewZapLogger()),
		hotswap.WithFreeDelay(time.Second*15),
	)
	details, err := swapper.LoadPlugins(nil)
	if err != nil {
		panic(err)
	} else if len(details) == 0 {
		panic("no plugin is found in " + pluginDir)
	} else {
		logbus.Info("hotswap first load plugin", logbus.String("pluginDir", pluginDir), logbus.Int("len", len(details)), logbus.Any("details", details))
	}
	swapperManager = swapper
}

func GetManager() *hotswap.PluginManagerSwapper {
	return swapperManager
}
