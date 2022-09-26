package nft

import (
	"github.com/irisnet/irismod/modules/nft"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(nft.AppModuleBasic{})
}
