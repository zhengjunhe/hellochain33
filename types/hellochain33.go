package types

import (
	"github.com/33cn/chain33/common/address"
	"github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/types"
)

var tlog = log15.New("module", ParaHellochain33)

// action type
const (
	TypeSetSelfIntroduction = iota
	TypePingPong
)

const (
	TyLogInvalid = iota + 100
	TyLogSetSelfIntroduction
	TyLogPlayPingPong
)


type paraSetSelfIntroActionTx struct {
	Fee                int64              `json:"fee"`
	SetSelfIntroAction SetSelfIntroAction `json:"SetSelfIntroAction"`
}

type paraPingPongActionTx struct {
	Fee            int64          `json:"fee"`
	PingPongAction PingPongAction `json:"PingPongAction"`
}

func createSetSelfIntroActionTx(para *SetSelfIntroAction) (*types.Transaction, error) {
	v := &SetSelfIntroAction{
		Introduction: para.Introduction,
		Age:para.Age,
		Male:para.Male,
		Married:para.Married,
		Notes:para.Notes,
	}
	action := &Hellochain33Action{
		Ty:    TypeSetSelfIntroduction,
		Value: &Hellochain33Action_SetSelfIntro{v},
	}
	tx := &types.Transaction{
		Execer:  []byte(ParaHellochain33),
		Payload: types.Encode(action),
		Fee:     0,
		To:      address.ExecAddress(ParaHellochain33),
	}
	tx, err := types.FormatTx(ParaHellochain33, tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func createPingPongActionTx(para *PingPongAction) (*types.Transaction, error) {
	v := &PingPongAction{
		Account: para.Account,
		Notes:   para.Notes,
	}
	action := &Hellochain33Action{
		Ty:    TypePingPong,
		Value: &Hellochain33Action_PingPong{v},
	}
	tx := &types.Transaction{
		Execer:  []byte(ParaHellochain33),
		Payload: types.Encode(action),
		Fee:     0,
		To:      address.ExecAddress(ParaHellochain33),
	}
	tx, err := types.FormatTx(ParaHellochain33, tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
