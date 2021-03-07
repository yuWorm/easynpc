package main

import (
	"flag"
	"fmt"
	"os"

	"ehang.io/nps/client"
	"ehang.io/nps/lib/common"
)

var cl *client.TRPClient

func main() {
	var serverAddr string
	var vkey string
	var connType string

	flag.StringVar(&serverAddr, "server", "null", "nps服务服务器地址，包括端口号")
	flag.StringVar(&vkey, "vkey", "null", "客户端唯一密钥")
	flag.StringVar(&connType, "type", "tcp", "客户端连接类型")

	flag.Parse()

	if serverAddr == "null" {
		fmt.Println("请输入nps服务器地址")
		os.Exit(0)
	}

	if vkey == "null" {
		fmt.Println("请输入客户端密钥")
		os.Exit(0)
	}

	fmt.Println("npc版本：" + logs())
	fmt.Println("nps github:https://github.com/ehang-io/nps")
	fmt.Println("本项目Github：")
	fmt.Println("看到这条消息可以使用Ctrl + C关闭程序")
	cl = client.NewRPClient(serverAddr, vkey, connType, "", nil, 60)
	cl.Start()

}

func logs() string {
	return common.GetLogMsg()
}

