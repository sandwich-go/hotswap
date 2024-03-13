package sdk

import (
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
		hotswap.WithLogger(NewZapLogger()), // use logbus
		hotswap.WithFreeDelay(spec.GetHotswapSpec().GetFreeDelay()),
		hotswap.WithWhitelist(spec.GetHotswapSpec().GetWhitelist()...),
		hotswap.WithExtensionNewer(spec.GetHotswapSpec().GetExtensionNewer()),
		hotswap.WithReloadCallback(spec.GetHotswapSpec().GetReloadCallback()),
		hotswap.WithStaticPlugins(spec.GetHotswapSpec().GetStaticPlugins()),
	)
	details, err := swapper.LoadPlugins(spec.GetOnFirstLoadData())
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

// InvokeEach 触发每个plugin里的name方法
// example: InvokeEach("meow", repeat)
func InvokeEach(name string, params ...interface{}) {
	GetManager().Current().InvokeEach(name, params...)
}

// Invoke 触发指定plugin里的方法
// example: Invoke("dog", "meow", repeat)
func Invoke(pluginName string, funcName string, params ...interface{}) (interface{}, error) {
	return GetManager().Current().FindPlugin(pluginName).InvokeFunc(funcName, params...)
}

// Extension 返回 HotswapSpec.ExtensionNewer 里的方法
// example: Extension().(*g.VaultExtension).Meow(repeat)
func Extension() interface{} {
	return GetManager().Current().Extension
}
