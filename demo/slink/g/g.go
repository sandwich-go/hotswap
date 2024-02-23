package g

import (
	"github.com/edwingeng/slog"
	"github.com/edwingeng/tickque"
	"github.com/sandwich-go/hotswap"
)

var (
	Logger = slog.NewDevelopmentConfig().MustBuild()
)

var (
	PluginManagerSwapper *hotswap.PluginManagerSwapper
)

type VaultExtension struct {
	OnJob func(job *tickque.Job) error
}

func NewVaultExtension() interface{} {
	return &VaultExtension{}
}
