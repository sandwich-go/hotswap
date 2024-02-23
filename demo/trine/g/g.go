package g

import (
	"fmt"

	"github.com/edwingeng/slog"
	"github.com/sandwich-go/hotswap"
)

var (
	Logger = slog.NewDevelopmentConfig().MustBuild()
)

var (
	PluginManagerSwapper *hotswap.PluginManagerSwapper
)

type Vector struct {
	X, Y int
}

func (v Vector) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}
