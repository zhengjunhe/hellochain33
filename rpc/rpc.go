// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"context"

	"github.com/33cn/chain33/types"
	pt "github.com/33cn/plugin/plugin/dapp/hellochain33/types"
	log "github.com/33cn/chain33/common/log/log15"
)

var clog = log.New("module", "execs.hellochain33")

func (c *channelClient) GetSelfIntro(ctx context.Context, req *pt.ReqAccount) (*pt.SelfIntroInfo, error) {
	data, err := c.Query(pt.GetExecName(), "GetSelfIntro", req)
	if err != nil {
		return nil, err
	}
	if resp, ok := data.(*pt.SelfIntroInfo); ok {
		return resp, nil
	}
	return nil, types.ErrDecode

}
func (c *channelClient) GetPingPongStatus(ctx context.Context, req *pt.ReqAccount) (*pt.PingPongStatus, error) {
	data, err := c.Query(pt.GetExecName(), "GetPingPongStatus", req)
	if err != nil {
		return nil, err
	}
	if resp, ok := data.(*pt.PingPongStatus); ok {
		return resp, nil
	}
	return nil, types.ErrDecode
}

func (c *channelClient) GetAccountsSelfIntroSetted(ctx context.Context, req *pt.ReqAccountSpecified) (*pt.ResAccountsSelfIntro, error) {
	data, err := c.Query(pt.GetExecName(), "GetAccountsSelfIntroSetted", req)
	if err != nil {
		return nil, err
	}
	if resp, ok := data.(*pt.ResAccountsSelfIntro); ok {
		return resp, nil
	}
	return nil, types.ErrDecode

}
func (c *channelClient) GetPingOrPongs(ctx context.Context, req *pt.ReqPingPongStatus) (*pt.PingPongCount, error) {
	data, err := c.Query(pt.GetExecName(), "GetPingOrPongs", req)
	if err != nil {
		return nil, err
	}
	if resp, ok := data.(*pt.PingPongCount); ok {
		return resp, nil
	}
	return nil, types.ErrDecode
}

func (c *Jrpc) GetSelfIntro(req *pt.ReqAccount, result *interface{}) error {
	clog.Info("GetSelfIntro", "req para is", req)
	if req == nil {
		return types.ErrInvalidParam
	}
	data, err := c.cli.GetSelfIntro(context.Background(), req)
	*result = data
	return err
}

func (c *Jrpc) GetPingPongStatus(req *pt.ReqAccount, result *interface{}) error {
	if req == nil {
		return types.ErrInvalidParam
	}
	data, err := c.cli.GetPingPongStatus(context.Background(), req)
	*result = data
	return err
}

func (c *Jrpc) GetAccountsSelfIntroSetted(req *pt.ReqAccountSpecified, result *interface{}) error {
	if req == nil {
		return types.ErrInvalidParam
	}
	data, err := c.cli.GetAccountsSelfIntroSetted(context.Background(), req)
	*result = data
	return err
}

func (c *Jrpc) GetPingOrPongs(req *pt.ReqPingPongStatus, result *interface{}) error {
	if req == nil {
		return types.ErrInvalidParam
	}
	data, err := c.cli.GetPingOrPongs(context.Background(), req)
	*result = data
	return err
}
