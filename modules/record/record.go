package record

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
)

type RecordClient struct {
}

func NewClient() RecordClient {
	return RecordClient{}
}

func (record RecordClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgRecordCreate:
		docMsg := DocMsgRecordCreate{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
