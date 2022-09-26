package irismod_parser

import (
	. "github.com/kaifei-bianjie/common-parser"
	"github.com/kaifei-bianjie/irismod-parser/codec"
	"github.com/kaifei-bianjie/irismod-parser/modules/coinswap"
	"github.com/kaifei-bianjie/irismod-parser/modules/farm"
	"github.com/kaifei-bianjie/irismod-parser/modules/htlc"
	"github.com/kaifei-bianjie/irismod-parser/modules/mt"
	"github.com/kaifei-bianjie/irismod-parser/modules/nft"
	"github.com/kaifei-bianjie/irismod-parser/modules/oracle"
	"github.com/kaifei-bianjie/irismod-parser/modules/random"
	"github.com/kaifei-bianjie/irismod-parser/modules/record"
	"github.com/kaifei-bianjie/irismod-parser/modules/service"
	"github.com/kaifei-bianjie/irismod-parser/modules/token"
)

type MsgClient struct {
	Coinswap Client
	Farm     Client
	Htlc     Client
	Mt       Client
	Nft      Client
	Oracle   Client
	Random   Client
	Record   Client
	Service  Client
	Token    Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Coinswap: coinswap.NewClient(),
		Farm:     farm.NewClient(),
		Htlc:     htlc.NewClient(),
		Mt:       mt.NewClient(),
		Nft:      nft.NewClient(),
		Oracle:   oracle.NewClient(),
		Random:   random.NewClient(),
		Record:   record.NewClient(),
		Service:  service.NewClient(),
		Token:    token.NewClient(),
	}
}
