package sdk

import "github.com/sandwich-go/hotswap"

//go:generate optionGen --v=true --xconf=false --usage_tag_name=usage
func PluginSpecOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"MountDir":      "/mount/data",                         // annotation@MountDir(comment="磁盘挂载目录")
		"HotReload":     true,                                  // annotation@HotReload(comment="允许热更新，开启watch目录")
		"DirsToKeep":    10,                                    // annotation@DirsToKeep(comment="同一service, 磁盘保留发布的目录数")
		"InternalDir":   "bin/plugin",                          // annotation@InternalDir(comment="service pod内部携带的plugin目录")
		"StaticPlugins": map[string]*hotswap.StaticPlugin(nil), // annotation@StaticPlugins(comment="宿主程序直接编译的插件 用做debug和windows")
	}
}
