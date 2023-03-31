package evm

import (
	"github.com/evmos/ethermint/x/evm"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		evm.AppModuleBasic{},
	)
}
