package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/33cn/chain33/rpc/jsonclient"
	rpctypes "github.com/33cn/chain33/rpc/types"
	"github.com/33cn/chain33/types"
	pt "github.com/33cn/plugin/plugin/dapp/hellochain33/types"
	"github.com/spf13/cobra"
)

// Hellochain33Cmd 本执行器的命令行初始化总入口
func Hellochain33Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hellochain33",
		Short: "hellochain33 commandline interface",
		Args:  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(
		CreateRawSetSelfIntroTxCmd(),
		CreatePlayPingPongTxCmd(),
		QueryOnesInfoCmd(),      // 查询某个用户的具体信息
		QueryPingPongCountCmd(), //查寻并统计当前处在ping或者pong状态的用户个数
		// 如果有其它命令，在这里加入
	)
	return cmd
}

// QueryCmd query 命令
func QueryOnesInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query_i",
		Short: "query one's self-introduction",
		Run:   queryOnesSelfIntroduction,
	}
	addQueryInfoFlags(cmd)
	return cmd
}

func addQueryInfoFlags(cmd *cobra.Command) {
	// type参数，指定查询的消息类型，为uint32类型，默认值为1，通过-t参数指定
	cmd.Flags().StringP("account", "a", "", "account address")
	cmd.MarkFlagRequired("account")
}

func queryOnesSelfIntroduction(cmd *cobra.Command, args []string) {
	// 这个是命令行的默认参数，可以制定调用哪一个服务地址
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	account, _ := cmd.Flags().GetString("account")
	// 创建RPC客户端，调用我们实现的QueryPing服务接口
	client, err := jsonclient.NewJSONClient(rpcLaddr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// 初始化查询参数结构
	var action = &pt.ReqAccount{Account: account}

	var result pt.SelfIntroInfo
	err = client.Call("hellochain33.GetSelfIntro", action, &result)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(string(data))
}

// QueryCmd query 命令
func QueryPingPongCountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query_pingpong",
		Short: "query total count with status:ping or pong",
		Run:   queryPingPongCount,
	}
	addPingPongCountFlags(cmd)
	return cmd
}

func addPingPongCountFlags(cmd *cobra.Command) {
	//cmd.Flags().StringP("account", "a", "", "account address")
	//cmd.MarkFlagRequired("account")
	// type参数，指定查询的消息类型，为uint32类型，默认值为1，通过-t参数指定
	cmd.Flags().BoolP("status", "s", true, "account status")
	cmd.MarkFlagRequired("status")
}

func queryPingPongCount(cmd *cobra.Command, args []string) {
	// 这个是命令行的默认参数，可以制定调用哪一个服务地址
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	//account, _ := cmd.Flags().GetString("account")
	status, _ := cmd.Flags().GetBool("status")
	// 创建RPC客户端，调用我们实现的QueryPing服务接口
	client, err := jsonclient.NewJSONClient(rpcLaddr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	// 初始化查询参数结构
	var action = &pt.ReqPingPongStatus{
		Status: status}

	var result pt.PingPongCount
	err = client.Call("hellochain33.GetPingOrPongs", action, &result)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(string(data))
}

func CreateRawSetSelfIntroTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "seti",
		Short: "Create a Set Self Introducion transaction",
		Run:   setSelfIntro,
	}
	addSetSelfIntroFlags(cmd)
	return cmd
}

func addSetSelfIntroFlags(cmd *cobra.Command) {
	// type参数，指定查询的消息类型，为uint32类型，默认值为1，通过-t参数指定
	cmd.Flags().StringP("intro", "i", "", "self introduction")
	cmd.MarkFlagRequired("intro")

	cmd.Flags().Int32P("age", "a", 0, "age")
	cmd.MarkFlagRequired("age")

	cmd.Flags().BoolP("male", "x", true, "male")
	//cmd.MarkFlagRequired("male")

	cmd.Flags().BoolP("married", "m", true, "married")
	//cmd.MarkFlagRequired("married")

	cmd.Flags().StringP("notes", "n", "", "notes")
}

func setSelfIntro(cmd *cobra.Command, args []string) {
	intro, _ := cmd.Flags().GetString("intro")
	age, _ := cmd.Flags().GetInt32("age")
	male, _ := cmd.Flags().GetBool("male")
	married, _ := cmd.Flags().GetBool("married")
	notes, _ := cmd.Flags().GetString("notes")

	para := &pt.SetSelfIntroAction{
		intro,
		age,
		male,
		married,
		notes,
	}

	params := &rpctypes.CreateTxIn{
		Execer:     types.ExecName(pt.ParaHellochain33),
		ActionName: "SetSelfIntroAction",
		Payload:    types.MustPBToJSON(para),
	}

	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	ctx := jsonclient.NewRPCCtx(rpcLaddr, "Chain33.CreateTransaction", params, nil)
	ctx.RunWithoutMarshal()
}

func CreatePlayPingPongTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "play ping pong",
		Short: "Inverse one's status",
		Run:   createPlayPingPong,
	}
	addPingPongFlags(cmd)
	return cmd
}

func addPingPongFlags(cmd *cobra.Command) {
	// type参数，指定查询的消息类型，为uint32类型，默认值为1，通过-t参数指定
	cmd.Flags().StringP("account", "a", "", "account address")
	cmd.MarkFlagRequired("account")

	cmd.Flags().StringP("notes", "n", "", "notes")
}

func createPlayPingPong(cmd *cobra.Command, args []string) {
	account, _ := cmd.Flags().GetString("account")
	notes, _ := cmd.Flags().GetString("notes")

	para := &pt.PingPongAction{
		account,
		notes,
	}

	params := &rpctypes.CreateTxIn{
		Execer:     types.ExecName(pt.ParaHellochain33),
		ActionName: "PingPongAction",
		Payload:    types.MustPBToJSON(para),
	}

	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	ctx := jsonclient.NewRPCCtx(rpcLaddr, "Chain33.CreateTransaction", params, nil)
	ctx.RunWithoutMarshal()
}
