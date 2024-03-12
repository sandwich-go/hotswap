package hotswap

import (
	"time"

	"github.com/edwingeng/slog"
)

//go:generate optionGen --v=true --xconf=false --usage_tag_name=usage
func SpecOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"Logger":         slog.Logger(slog.NewDevelopmentConfig().MustBuild()), // annotation@Logger(comment="replaces the default logger with your own.")
		"FreeDelay":      time.Duration(5 * time.Minute),                       // annotation@FreeDelay(comment="the delay time of calling OnFree. The default value is 5 minutes.")
		"ReloadCallback": ReloadCallback(nil),                                  // annotation@ReloadCallback(comment="the callback function of reloading.")
		"ExtensionNewer": (func() interface{})(nil),                            // annotation@ExtensionNewer(comment="the function used to create a new object for PluginManager.Vault.Extension.")
		"StaticPlugins":  map[string]*StaticPlugin(nil),                        // annotation@StaticPlugins(comment="the static plugins for static linking. 宿主程序直接编译的插件 用做debug和windows")
		"Whitelist":      []string(nil),                                        // annotation@Whitelist(comment="the plugins to load explicitly 若不为空 只加载白名单里的插件")
	}
}
