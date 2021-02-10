package main

import (
	"fmt"
	"os"

	"ehang.io/nps/client"
	"ehang.io/nps/lib/version"
)

var cls *client.TRPClient

func main() {
	if len(os.Args) != 0 {
		if len(os.Args) > 3 {
			var serverAddr string = os.Args[1]
			var verifyKey string = os.Args[2]
			var connType string = os.Args[3]
			cls = client.NewRPClient(serverAddr, verifyKey, connType, "", nil, 60)
			fmt.Println("客户端版本：" + getVersion() + "\n连接地址：" + serverAddr + "\n客户端密钥：" + verifyKey + "\n连接方式：" + connType)
			fmt.Println("开始连接")
			cls.Start()
		} else {
			fmt.Println("请确认输入了服务地址，客户端密钥和连接类型")
		}
	} else {
		fmt.Println("请输入参数")
	}
}

func getVersion() string {
	return version.VERSION
}
