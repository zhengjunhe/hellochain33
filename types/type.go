// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"encoding/json"
	"reflect"

	log "github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/types"
)

var (
	// ParaX paracross exec name
	ParaHellochain33 = "hellochain33"
	glog             = log.New("module", ParaHellochain33)
)

func init() {
	// init executor type
	types.AllowUserExec = append(types.AllowUserExec, []byte(ParaHellochain33))
	types.RegistorExecutor(ParaHellochain33, NewType())
	types.RegisterDappFork(ParaHellochain33, "Enable", 0)
}

// GetExecName get para exec name
func GetExecName() string {
	return types.ExecName(ParaHellochain33)
}

type Hellochain33Type struct {
	types.ExecTypeBase
}

// NewType get Hellochain33 type
func NewType() *Hellochain33Type {
	c := &Hellochain33Type{}
	c.SetChild(c)
	return c
}

// GetName 获取执行器名称
func (p *Hellochain33Type) GetName() string {
	return ParaHellochain33
}

// GetLogMap get receipt log map
func (p *Hellochain33Type) GetLogMap() map[int64]*types.LogInfo {
	return map[int64]*types.LogInfo{
		TyLogSetSelfIntroduction: {Ty: reflect.TypeOf(LogSetSelfIntro{}), Name: "LogSetSelfIntro"},
		TyLogPlayPingPong: {Ty: reflect.TypeOf(LogPingPong{}), Name: "LogPingPong"},
	}
}

// GetTypeMap get action type
func (p *Hellochain33Type) GetTypeMap() map[string]int32 {
	return map[string]int32{
		"SetSelfIntro": TypeSetSelfIntroduction,
		"PingPong":     TypePingPong,
	}
}

// GetPayload Hellochain33 get action payload
func (p *Hellochain33Type) GetPayload() types.Message {
	return &Hellochain33Action{}
}

// CreateTx Hellochain33 create tx by different action
func (p Hellochain33Type) CreateTx(action string, message json.RawMessage) (*types.Transaction, error) {
	glog.Info("CreateTx", "action", action)

	switch action {
	case "SetSelfIntroAction":
		var param SetSelfIntroAction
		err := types.JSONToPB(message, &param)
		if err != nil {
			glog.Error("CreateTx to SetSelfIntroAction", "Error", err)
			return nil, types.ErrInvalidParam
		}
		return createSetSelfIntroActionTx(&param)
	case "PingPongAction":
		var param PingPongAction
		err := json.Unmarshal(message, &param)
		if err != nil {
			glog.Error("CreateTx to PingPongAction", "Error", err)
			return nil, types.ErrInvalidParam
		}
		return createPingPongActionTx(&param)
	default:
		glog.Error("CreateTx", "action type is not supported with name", action)
		return nil, types.ErrNotSupport
	}
}
