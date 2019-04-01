// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (

	"github.com/33cn/chain33/types"
	pt "github.com/33cn/plugin/plugin/dapp/hellochain33/types"
)

// Query_GetSelfIntro 查询 账户对应的个人信息
func (e *Hellochain33) Query_GetSelfIntro(in *pt.ReqAccount) (types.Message, error) {
	key := calcSelfIntroKey(in.Account)
	value, err := e.GetStateDB().Get(key)
	if err != nil {
		return nil, err
	}
	info := pt.SetSelfIntroAction{}
	types.Decode(value, &info)
	res := pt.SelfIntroInfo{
		info.Introduction,
		info.Age,
		info.Male,
		info.Married,
	}
	return &res, nil
}

//获取指定账户的ping，pong状态
//rpc GetPingPongStatus(ReqAccount) returns (PingPongStatus) {}
func (e *Hellochain33) Query_GetPingPongStatus(in *pt.ReqAccount) (types.Message, error) {
	key := calcPingPongStatusKey(in.Account)
	value, err := e.GetStateDB().Get(key)
	if err != nil {
		return nil, err
	}
	info := pt.PingPongStatus{}
	types.Decode(value, &info)
	return &info, nil
}

//获取所有满足条件的用户
//rpc GetAccountsSelfIntroSetted(ReqAccountSpecified) returns (ResAccountsSelfIntro) {}
func (e *Hellochain33) Query_GetAccountsSelfIntroSetted(in *pt.ReqAccountSpecified) (types.Message, error) {
    //todo:有待实现
	return nil, nil
}

//获取所有处在ping或pong的状态的用户
//rpc GetPingPongPlayCount(ReqPingPongCount) returns (PingPongCount) {}
func (e *Hellochain33) Query_GetPingOrPongs(in *pt.ReqPingPongStatus) (types.Message, error) {
	keyPrefix := calcLocalPingPongStatusPrefix(in.Status)
	value, _ := e.GetLocalDB().List(keyPrefix, nil, 0, 0)
	res := pt.PingPongCount{
		0,
	}
	res.Count = int32(len(value))

	return &res, nil
}



