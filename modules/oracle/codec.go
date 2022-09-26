package oracle

import (
	"github.com/irisnet/irismod/modules/oracle"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(oracle.AppModuleBasic{})
}
