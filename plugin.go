// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hellochain33

import (
	"github.com/33cn/chain33/pluginmgr"
	"github.com/33cn/plugin/plugin/dapp/hellochain33/commands"
	"github.com/33cn/plugin/plugin/dapp/hellochain33/executor"
	"github.com/33cn/plugin/plugin/dapp/hellochain33/rpc"
	"github.com/33cn/plugin/plugin/dapp/hellochain33/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.ParaHellochain33,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.Hellochain33Cmd,
		RPC:      rpc.Init,
	})
}
