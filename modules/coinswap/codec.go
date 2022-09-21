package coinswap

import (
	"github.com/irisnet/irismod/modules/coinswap"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(coinswap.AppModuleBasic{})
}
