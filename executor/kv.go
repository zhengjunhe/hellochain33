// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"fmt"
	pt "github.com/33cn/plugin/plugin/dapp/hellochain33/types"
)

var (
	SelfIntroPrefix = "setSelfIntro-"
	//setSelfIntroPrefix = "mavl-hellochain33-setSelfIntro-"
	pingPongStatusPrefix = "pingpongStatus-"
	//localSelfIntroPrefix = "LODB-hellochain33-setSelfIntro-"
	//localpingPongStatusPrefix = "LODB-hellochain33-pingpongStatus-"
)

func calcStateStatePrefix() string {
	return "mavl-" + pt.ParaHellochain33 + "-" + SelfIntroPrefix +"-"
}

func calcLocalDBPrefix() string {
	return "LODB-" + pt.ParaHellochain33 + "-"
}
//mavl-hellochain33-setSelfIntro-addr
func calcSelfIntroKey(addr string) []byte {
	return []byte(fmt.Sprintf(calcStateStatePrefix() + SelfIntroPrefix+"%s", addr))
}
//mavl-hellochain33-pingpongStatus-addr
func calcPingPongStatusKey(addr string) []byte {
	return []byte(fmt.Sprintf(calcStateStatePrefix() + pingPongStatusPrefix+"%s", addr))
}
//LODB-hellochain33-setSelfIntro-addr
func calcLoalSelfIntroKeyMarried(addr string) []byte {
	return []byte(fmt.Sprintf(calcLocalDBPrefix() + SelfIntroPrefix+"married-"+"%s", addr))
}
//LODB-hellochain33-pingpongStatus-addr
func calcLocalPingPongStatusKey(status bool, addr string) []byte {
	return []byte(fmt.Sprintf(calcLocalDBPrefix() + pingPongStatusPrefix+"%v-"+"%s", status, addr))
}

func calcLocalPingPongStatusPrefix(status bool) []byte {
	return []byte(fmt.Sprintf(calcLocalDBPrefix() + pingPongStatusPrefix+"%v-", status))
}

