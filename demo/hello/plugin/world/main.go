package main

import (
	"time"

	"github.com/sandwich-go/hotswap/demo/hello/g"
	"github.com/sandwich-go/hotswap/demo/hello/plugin/world/hum"
	"github.com/sandwich-go/hotswap/vault"
)

const (
	pluginName = "world"
)

var (
	CompileTimeString string
)

var close = make(chan struct{})

func OnLoad(data interface{}) error {
	g.Logger.Infof("<%s.%s> OnLoad %s", pluginName, CompileTimeString, data)
	g.FunctionRegisterByPlugin = hum.CPUAndMemoryIntensiveFunction
	go func() {
		for {
			select {
			case <-close:
				g.Logger.Infof("<%s.%s> close", pluginName, CompileTimeString)
				return
			case <-time.Tick(time.Second * 3):
				// g.Logger.Infof("<%s.%s> tick", pluginName, CompileTimeString)
			}
		}
	}()
	return nil
}

func OnInit(sharedVault *vault.Vault) error {
	g.Logger.Infof("<%s.%s> OnInit", pluginName, CompileTimeString)
	return nil
}

func OnFree() {
	g.Logger.Infof("<%s.%s> OnFree", pluginName, CompileTimeString)
	close <- struct{}{}
}

func Export() interface{} {
	g.Logger.Infof("<%s.%s> Export", pluginName, CompileTimeString)
	return nil
}

func Import() interface{} {
	g.Logger.Infof("<%s.%s> Import", pluginName, CompileTimeString)
	return nil
}

func InvokeFunc(name string, params ...interface{}) (interface{}, error) {
	switch name {
	case "hum":
		repeat := params[0].(int)
		hum.Hum(pluginName, CompileTimeString, repeat)
	case "sort":
		return hum.CPUAndMemoryIntensiveFunction(), nil
	}
	return nil, nil
}

func Reloadable() bool {
	g.Logger.Infof("<%s.%s> Reloadable", pluginName, CompileTimeString)
	return true
}
