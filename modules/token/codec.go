package token

import (
	"github.com/irisnet/irismod/modules/token"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(token.AppModuleBasic{})
}
