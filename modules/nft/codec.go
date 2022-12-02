package nft

import (
	nft "github.com/irisnet/irismod/modules/nft/module"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(nft.AppModuleBasic{})
}
