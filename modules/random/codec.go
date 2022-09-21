package random

import (
	"github.com/irisnet/irismod/modules/random"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(random.AppModuleBasic{})
}
