// Code generated by optiongen. DO NOT EDIT.
// optiongen: github.com/timestee/optiongen

package sdk

import (
	"time"

	"github.com/sandwich-go/hotswap"
)

// PluginSpec should use NewPluginSpec to initialize it
type PluginSpec struct {
	MountDir        string                           `usage:"磁盘挂载目录"`                                                                      // annotation@MountDir(comment="磁盘挂载目录")
	HotReload       bool                             `usage:"允许热更新，开启watch目录"`                                                             // annotation@HotReload(comment="允许热更新，开启watch目录")
	DirsToKeep      int                              `usage:"同一service, 磁盘保留发布的目录数。如果设为0，则不删除历史目录"`                                        // annotation@DirsToKeep(comment="同一service, 磁盘保留发布的目录数。如果设为0，则不删除历史目录")
	InternalDir     string                           `usage:"service pod内部携带的plugin目录"`                                                    // annotation@InternalDir(comment="service pod内部携带的plugin目录")
	OnFirstLoadData interface{}                      `usage:"第一次OnLoad的data参数"`                                                            // annotation@OnFirstLoadData(comment="第一次OnLoad的data参数")
	OnReloadData    interface{}                      `usage:"热更时新插件OnLoad的data参数"`                                                         // annotation@OnReloadData(comment="热更时新插件OnLoad的data参数")
	FreeDelay       time.Duration                    `usage:"the delay time of calling OnFree. The default value is 5 minutes."`           // annotation@FreeDelay(comment="the delay time of calling OnFree. The default value is 5 minutes.")
	ExtensionNewer  func() interface{}               `usage:"the function used to create a new object for PluginManager.Vault.Extension."` // annotation@ExtensionNewer(comment="the function used to create a new object for PluginManager.Vault.Extension.")
	StaticPlugins   map[string]*hotswap.StaticPlugin `usage:"the static plugins for static linking. 宿主程序直接编译的插件 用做debug和windows"`          // annotation@StaticPlugins(comment="the static plugins for static linking. 宿主程序直接编译的插件 用做debug和windows")
}

// NewPluginSpec new PluginSpec
func NewPluginSpec(opts ...PluginSpecOption) *PluginSpec {
	cc := newDefaultPluginSpec()
	for _, opt := range opts {
		opt(cc)
	}
	if watchDogPluginSpec != nil {
		watchDogPluginSpec(cc)
	}
	return cc
}

// ApplyOption apply multiple new option and return the old ones
// sample:
// old := cc.ApplyOption(WithTimeout(time.Second))
// defer cc.ApplyOption(old...)
func (cc *PluginSpec) ApplyOption(opts ...PluginSpecOption) []PluginSpecOption {
	var previous []PluginSpecOption
	for _, opt := range opts {
		previous = append(previous, opt(cc))
	}
	return previous
}

// PluginSpecOption option func
type PluginSpecOption func(cc *PluginSpec) PluginSpecOption

// WithMountDir 磁盘挂载目录
func WithMountDir(v string) PluginSpecOption {
	return func(cc *PluginSpec) PluginSpecOption {
		previous := cc.MountDir
		cc.MountDir = v
		return WithMountDir(previous)
	}
}

// WithHotReload 允许热更新，开启watch目录
func WithHotReload(v bool) PluginSpecOption {
	return func(cc *PluginSpec) PluginSpecOption {
		previous := cc.HotReload
		cc.HotReload = v
		return WithHotReload(previous)
	}
}

// WithDirsToKeep 同一service, 磁盘保留发布的目录数。如果设为0，则不删除历史目录
func WithDirsToKeep(v int) PluginSpecOption {
	return func(cc *PluginSpec) PluginSpecOption {
		previous := cc.DirsToKeep
		cc.DirsToKeep = v
		return WithDirsToKeep(previous)
	}
}

// WithInternalDir service pod内部携带的plugin目录
func WithInternalDir(v string) PluginSpecOption {
	return func(cc *PluginSpec) PluginSpecOption {
		previous := cc.InternalDir
		cc.InternalDir = v
		return WithInternalDir(previous)
	}
}

// WithOnFirstLoadData 第一次OnLoad的data参数
func WithOnFirstLoadData(v interface{}) PluginSpecOption {
	return func(cc *PluginSpec) PluginSpecOption {
		previous := cc.OnFirstLoadData
		cc.OnFirstLoadData = v
		return WithOnFirstLoadData(previous)
	}
}

// WithOnReloadData 热更时新插件OnLoad的data参数
func WithOnReloadData(v interface{}) PluginSpecOption {
	return func(cc *PluginSpec) PluginSpecOption {
		previous := cc.OnReloadData
		cc.OnReloadData = v
		return WithOnReloadData(previous)
	}
}

// WithFreeDelay the delay time of calling OnFree. The default value is 5 minutes.
func WithFreeDelay(v time.Duration) PluginSpecOption {
	return func(cc *PluginSpec) PluginSpecOption {
		previous := cc.FreeDelay
		cc.FreeDelay = v
		return WithFreeDelay(previous)
	}
}

// WithExtensionNewer the function used to create a new object for PluginManager.Vault.Extension.
func WithExtensionNewer(v func() interface{}) PluginSpecOption {
	return func(cc *PluginSpec) PluginSpecOption {
		previous := cc.ExtensionNewer
		cc.ExtensionNewer = v
		return WithExtensionNewer(previous)
	}
}

// WithStaticPlugins the static plugins for static linking. 宿主程序直接编译的插件 用做debug和windows
func WithStaticPlugins(v map[string]*hotswap.StaticPlugin) PluginSpecOption {
	return func(cc *PluginSpec) PluginSpecOption {
		previous := cc.StaticPlugins
		cc.StaticPlugins = v
		return WithStaticPlugins(previous)
	}
}

// InstallPluginSpecWatchDog the installed func will called when NewPluginSpec  called
func InstallPluginSpecWatchDog(dog func(cc *PluginSpec)) { watchDogPluginSpec = dog }

// watchDogPluginSpec global watch dog
var watchDogPluginSpec func(cc *PluginSpec)

// newDefaultPluginSpec new default PluginSpec
func newDefaultPluginSpec() *PluginSpec {
	cc := &PluginSpec{}

	for _, opt := range [...]PluginSpecOption{
		WithMountDir("/mount/data"),
		WithHotReload(true),
		WithDirsToKeep(10),
		WithInternalDir("bin/plugin"),
		WithOnFirstLoadData(nil),
		WithOnReloadData(nil),
		WithFreeDelay(15 * time.Second),
		WithExtensionNewer(nil),
		WithStaticPlugins(nil),
	} {
		opt(cc)
	}

	return cc
}

// all getter func
func (cc *PluginSpec) GetMountDir() string                                { return cc.MountDir }
func (cc *PluginSpec) GetHotReload() bool                                 { return cc.HotReload }
func (cc *PluginSpec) GetDirsToKeep() int                                 { return cc.DirsToKeep }
func (cc *PluginSpec) GetInternalDir() string                             { return cc.InternalDir }
func (cc *PluginSpec) GetOnFirstLoadData() interface{}                    { return cc.OnFirstLoadData }
func (cc *PluginSpec) GetOnReloadData() interface{}                       { return cc.OnReloadData }
func (cc *PluginSpec) GetFreeDelay() time.Duration                        { return cc.FreeDelay }
func (cc *PluginSpec) GetExtensionNewer() func() interface{}              { return cc.ExtensionNewer }
func (cc *PluginSpec) GetStaticPlugins() map[string]*hotswap.StaticPlugin { return cc.StaticPlugins }

// PluginSpecVisitor visitor interface for PluginSpec
type PluginSpecVisitor interface {
	GetMountDir() string
	GetHotReload() bool
	GetDirsToKeep() int
	GetInternalDir() string
	GetOnFirstLoadData() interface{}
	GetOnReloadData() interface{}
	GetFreeDelay() time.Duration
	GetExtensionNewer() func() interface{}
	GetStaticPlugins() map[string]*hotswap.StaticPlugin
}

// PluginSpecInterface visitor + ApplyOption interface for PluginSpec
type PluginSpecInterface interface {
	PluginSpecVisitor
	ApplyOption(...PluginSpecOption) []PluginSpecOption
}
