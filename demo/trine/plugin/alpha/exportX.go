package alpha

import (
	"github.com/sandwich-go/hotswap/demo/trine/g"
	"github.com/sandwich-go/hotswap/demo/trine/plugin/alpha/aimpl"
)

type exportX struct{}

func (_ exportX) One(str1 string, v1 g.Vector) {
	aimpl.One(str1, v1, pluginName, CompileTimeString)
}
