package token

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
	v1 "github.com/kaifei-bianjie/irismod-parser/modules/token/v1"
	"github.com/kaifei-bianjie/irismod-parser/modules/token/v1beta1"
)

type TokenClient struct {
}

func NewClient() TokenClient {
	return TokenClient{}
}
func (token TokenClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgMintToken:
		docMsg := v1beta1.DocMsgMintToken{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgEditToken:
		docMsg := v1beta1.DocMsgEditToken{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgIssueToken:
		docMsg := v1beta1.DocMsgIssueToken{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgTransferTokenOwner:
		docMsg := v1beta1.DocMsgTransferTokenOwner{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgBurnToken:
		docMsg := v1beta1.DocMsgBurnToken{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgMintTokenV1:
		docMsg := v1.DocMsgMintTokenV1{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgEditTokenV1:
		docMsg := v1.DocMsgEditTokenV1{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgIssueTokenV1:
		docMsg := v1.DocMsgIssueTokenV1{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgTransferTokenOwnerV1:
		docMsg := v1.DocMsgTransferTokenOwnerV1{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgBurnTokenV1:
		docMsg := v1.DocMsgBurnTokenV1{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgSwapFeeTokenV1:
		docMsg := v1.DocMsgSwapFeeTokenV1{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
