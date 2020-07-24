package utils

import "flag"

//PrivateRsaPath ssh私钥绝对路径
var PrivateRsaPath string

//Port 隧道监听端口
var Port string

//RemoteAddr 隧道目标地址
var RemoteAddr string

//JumpServerHost 跳板机服务地址
var JumpServerHost string

//JumpServerPort 跳板机服务端口
var JumpServerPort string

//JumpServerUserName 跳板机登录名
var JumpServerUserName string

//JumpServerUserPassword 跳板机登录密码
var JumpServerUserPassword string

//DialTimeoutSecond 跳板机连接初始化超时时间
var DialTimeoutSecond uint64

func init() {

	flag.StringVar(&PrivateRsaPath, "privateRsaPath", "", "path of private rsa for ssh.")
	flag.StringVar(&Port, "port", "", "service listen port.If nil, it's round")
	flag.StringVar(&RemoteAddr, "remoteAddr", "", "ssh target addr")
	flag.StringVar(&JumpServerHost, "jumpServerHost", "", "jump server host")
	flag.StringVar(&JumpServerPort, "jumpServerPort", "", "jump server port.")
	flag.StringVar(&JumpServerUserName, "jumpServerUserName", "", "jump server username")
	flag.StringVar(&JumpServerUserPassword, "jumpServerUserPassword", "", "jump server password")
	flag.Uint64Var(&DialTimeoutSecond, "dialTimeoutSecond", 0, "connect to jump server timeout")
}
