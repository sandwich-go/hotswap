package sdk

import (
	"time"

	"github.com/sandwich-go/hotswap"
)

//go:generate optionGen --v=true --xconf=false --usage_tag_name=usage
func PluginSpecOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"MountDir":        "/mount/data",                         // annotation@MountDir(comment="磁盘挂载目录")
		"HotReload":       true,                                  // annotation@HotReload(comment="允许热更新，开启watch目录")
		"DirsToKeep":      10,                                    // annotation@DirsToKeep(comment="同一service, 磁盘保留发布的目录数。如果设为0，则不删除历史目录")
		"InternalDir":     "bin/plugin",                          // annotation@InternalDir(comment="service pod内部携带的plugin目录")
		"OnFirstLoadData": interface{}(nil),                      // annotation@OnFirstLoadData(comment="第一次OnLoad的data参数")
		"OnReloadData":    interface{}(nil),                      // annotation@OnReloadData(comment="热更时新插件OnLoad的data参数")
		"FreeDelay":       time.Duration(15 * time.Second),       // annotation@FreeDelay(comment="the delay time of calling OnFree. The default value && min value is 15 Seconds.")
		"ExtensionNewer":  (func() interface{})(nil),             // annotation@ExtensionNewer(comment="the function used to create a new object for PluginManager.Vault.Extension.")
		"StaticPlugins":   map[string]*hotswap.StaticPlugin(nil), // annotation@StaticPlugins(comment="the static plugins for static linking. 宿主程序直接编译的插件 用做debug和windows")
	}
}
