package g

import (
	"github.com/edwingeng/slog"
	"github.com/sandwich-go/hotswap"
)

var (
	Logger = slog.NewDevelopmentConfig().MustBuild()
)

var (
	PluginManagerSwapper *hotswap.PluginManagerSwapper
)

type VaultExtension struct {
	Meow func(repeat int)
}

func NewVaultExtension() interface{} {
	return &VaultExtension{}
}
