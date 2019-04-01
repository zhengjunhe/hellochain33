// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (

	"github.com/33cn/chain33/types"
	pt "github.com/33cn/plugin/plugin/dapp/hellochain33/types"
)

func (t *Hellochain33) ExecLocal_SetSelfIntro(payload *pt.SetSelfIntroAction, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	set := &types.LocalDBSet{}
	if receiptData.GetTy() != types.ExecOk {
		return set, nil
	}

	for _, log := range receiptData.Logs {
		if log.Ty == pt.TyLogSetSelfIntroduction {
			log1 := pt.LogSetSelfIntro{}
			err := types.Decode(log.Log, &log1)
			if nil != err {
				return set, err
			}
			if log1.Current.Married {
				key := calcLoalSelfIntroKeyMarried(log1.Addr)
				value := types.Encode(log1.Current)
				set.KV = append(set.KV, &types.KeyValue{Key:key, Value:value})
			}
			//处理其他属性，如是否男性，年龄是否超过35周岁


		} else if log.Ty == pt.TyLogPlayPingPong {
			log2 := pt.LogPingPong{}
			err := types.Decode(log.Log, &log2)
			if nil != err {
				return set, err
			}

			key2 := calcLocalPingPongStatusKey(log2.Current.Status, log2.Addr)
			value2 := []byte{}
			set.KV = append(set.KV, &types.KeyValue{Key:key2, Value:value2})
		}
	}

	return set, nil
}

func (t *Hellochain33) ExecLocal_PingPong(payload *pt.PingPongAction, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	set := &types.LocalDBSet{}
	if receiptData.GetTy() != types.ExecOk {
		return set, nil
	}

	for _, log := range receiptData.Logs {
		if log.Ty == pt.TyLogPlayPingPong {
			log2 := pt.LogPingPong{}
			err := types.Decode(log.Log, &log2)
			if nil != err {
				return set, err
			}

			key2 := calcLocalPingPongStatusKey(log2.Current.Status, log2.Addr)
			value2 := []byte{}
			set.KV = append(set.KV, &types.KeyValue{Key:key2, Value:value2})

			key3 := calcLocalPingPongStatusKey(inversePingPongStatus(log2.Current.Status), log2.Addr)
			set.KV = append(set.KV, &types.KeyValue{Key:key3, Value:nil})

		}
	}

	return set, nil
}