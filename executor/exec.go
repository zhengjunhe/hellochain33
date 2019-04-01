// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"encoding/hex"

	"github.com/33cn/chain33/types"
	pt "github.com/33cn/plugin/plugin/dapp/hellochain33/types"
	"github.com/pkg/errors"
)

func (e *Hellochain33) Exec_SetSelfIntro(payload *pt.SetSelfIntroAction, tx *types.Transaction, index int) (*types.Receipt, error) {
	a := newAction(e, tx)
	receipt, err := a.SetSelfIntro(payload)
	if err != nil {
		clog.Error("Paracross commit failed", "error", err, "hash", hex.EncodeToString(tx.Hash()))
		return nil, errors.Cause(err)
	}
	return receipt, nil
}

func (e *Hellochain33) Exec_PingPong(payload *pt.PingPongAction, tx *types.Transaction, index int) (*types.Receipt, error) {
	a := newAction(e, tx)
	receipt, err := a.PingPong(payload)
	if err != nil {
		clog.Error("Paracross commit failed", "error", err, "hash", hex.EncodeToString(tx.Hash()))
		return nil, errors.Cause(err)
	}
	return receipt, nil
}