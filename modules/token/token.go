package token

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
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
		docMsg := DocMsgMintToken{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgEditToken:
		docMsg := DocMsgEditToken{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgIssueToken:
		docMsg := DocMsgIssueToken{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgTransferTokenOwner:
		docMsg := DocMsgTransferTokenOwner{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgBurnToken:
		docMsg := DocMsgBurnToken{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
