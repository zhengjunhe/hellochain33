// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"github.com/33cn/chain33/account"
	dbm "github.com/33cn/chain33/common/db"
	"github.com/33cn/chain33/system/dapp"
	log "github.com/33cn/chain33/common/log/log15"
	drivers "github.com/33cn/chain33/system/dapp"
	"github.com/33cn/chain33/types"
	pt "github.com/33cn/plugin/plugin/dapp/hellochain33/types"
	"github.com/33cn/chain33/client"
)

var (
	clog                    = log.New("module", "execs.hellochain33")
	driverName              = pt.ParaHellochain33
)

// Hellochain33 exec
type Hellochain33 struct {
	drivers.DriverBase
}

type action struct {
	coinsAccount *account.DB
	db           dbm.KV
	localdb      dbm.KVDB
	txhash       []byte
	fromaddr     string
	blocktime    int64
	height       int64
	execaddr     string
	api          client.QueueProtocolAPI
	tx           *types.Transaction
	exec         *Hellochain33
}

func init() {
	ety := types.LoadExecutorType(driverName)
	ety.InitFuncList(types.ListMethod(&Hellochain33{}))
}

//Init hellochain33 exec register
func Init(name string, sub []byte) {
	drivers.Register(GetName(), newHellochain33, types.GetDappFork(driverName, "Enable"))
}

//GetName return hellochain33 name
func GetName() string {
	return newHellochain33().GetName()
}

func newHellochain33() drivers.Driver {
	c := &Hellochain33{}
	c.SetChild(c)
	c.SetExecutorType(types.LoadExecutorType(driverName))
	return c
}

// GetDriverName return hellochain33 driver name
func (c *Hellochain33) GetDriverName() string {
	return pt.ParaHellochain33
}

func newAction(t *Hellochain33, tx *types.Transaction) *action {
	hash := tx.Hash()
	fromaddr := tx.From()
	return &action{t.GetCoinsAccount(), t.GetStateDB(), t.GetLocalDB(), hash, fromaddr,
		t.GetBlockTime(), t.GetHeight(), dapp.ExecAddress(string(tx.Execer)), t.GetAPI(), tx, t}
}

func (a *action) SetSelfIntro(info *pt.SetSelfIntroAction) (*types.Receipt, error) {
	//生成key，获取当前的状态，并生成该笔交易需要设置的新状态的KV
	key := calcSelfIntroKey(a.fromaddr)
	prevByte, _ := a.db.Get(key)
	var prev pt.SetSelfIntroAction
	if nil != prevByte {
		types.Decode(prevByte, &prev)
	}

	log := &pt.LogSetSelfIntro{
		Addr:a.fromaddr,
		Previous:&prev,
		Current: info,
	}
	receipt := &types.Receipt{
		Ty: types.ExecOk,
		KV: []*types.KeyValue{
			{Key: key, Value: types.Encode(info)},
		},
		Logs: []*types.ReceiptLog{
			{
				Ty:  pt.TyLogSetSelfIntroduction,
				Log: types.Encode(log),
			},
		},
	}
	//操作ping，pong状态
	receipt1, _ := a.PlayPingPong(a.fromaddr)
	receipt.KV = append(receipt.KV, receipt1.KV...)
	receipt.Logs = append(receipt.Logs, receipt1.Logs...)

	return receipt, nil
}

func (a *action) PingPong(info *pt.PingPongAction) (*types.Receipt, error) {
	return a.PlayPingPong(info.Account)
}

func (a *action) PlayPingPong(addr string) (*types.Receipt, error) {
	key1 := calcPingPongStatusKey(a.fromaddr)
	prevStatusByte, _ := a.db.Get(key1)
	//if nil == prevStatusByte {
	//	return nil, errors.New("Account is not setted")
	//}
	prevStatus := pt.PingPongStatus{false}
	if nil != prevStatusByte {
		types.Decode(prevStatusByte, &prevStatus)
	}

	val1 := types.Encode(&pt.PingPongStatus{
		Status:inversePingPongStatus(prevStatus.Status),
	})
	kv1 := &types.KeyValue{Key:key1, Value:val1}
	pinglog := &pt.LogPingPong{
		Addr:a.fromaddr,
		Previous:&prevStatus,
		Current:&pt.PingPongStatus{inversePingPongStatus(prevStatus.Status)},
	}
	//pinglog := &types.ReceiptLog{
	//	Ty:  pt.TyLogPlayPingPong,
	//	Log: types.Encode(log1),
	//}

	receipt := &types.Receipt{
		Ty: types.ExecOk,
		KV: []*types.KeyValue{kv1},
		Logs: []*types.ReceiptLog{
			{
				Ty:  pt.TyLogPlayPingPong,
				Log: types.Encode(pinglog),
			},
		},
	}

	return receipt, nil
}

func inversePingPongStatus(pingpong bool) bool {
	if pingpong {
		return false
	} else {
		return true
	}
}