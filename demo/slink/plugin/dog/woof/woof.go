package woof

import (
	"strings"

	"github.com/edwingeng/live"
	"github.com/sandwich-go/hotswap/demo/slink/g"
)

func live_Woof(pluginName string, compileTimeString string, jobData live.Data) error {
	str := strings.TrimSpace(strings.Repeat("woof ", jobData.Int()))
	g.Logger.Infof("<%s.%s> %s. reloadCounter: %v",
		pluginName, compileTimeString, str, g.PluginManagerSwapper.ReloadCounter())
	return nil
}
