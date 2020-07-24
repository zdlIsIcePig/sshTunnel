package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sshTunnel/ssh"
	"sshTunnel/tunnel"
	"sshTunnel/utils"
)

func main() {

	flag.Parse()

	var commandValidateResult = true

	if utils.Port != "" {

		used, _ := utils.CheckPortUsedForTCP(utils.Port)

		if used {

			fmt.Println("port is used.")
			commandValidateResult = false
		}
	}

	if utils.RemoteAddr == "" {

		fmt.Println("remoteAddr is required.")
		commandValidateResult = false
	}

	if utils.JumpServerHost == "" {

		fmt.Println("jumpServerHost is required.")
		commandValidateResult = false
	}

	if utils.JumpServerPort == "" {

		fmt.Println("jumpServerPort is required.")
		commandValidateResult = false
	}

	if !commandValidateResult {

		os.Exit(0)
	}

	var buf []byte

	if utils.PrivateRsaPath != "" {

		buf, _ = ioutil.ReadFile(utils.PrivateRsaPath)
	}

	tunnel.InitListener(utils.Port, &ssh.Config{
		RemoteAddr:               utils.RemoteAddr,
		JumpServerHost:           utils.JumpServerHost,
		JumpServerPort:           utils.JumpServerPort,
		JumpServerUserName:       utils.JumpServerUserName,
		JumpServerUserPassword:   utils.JumpServerUserPassword,
		JumpServerUserPrivateKey: string(buf),
		DialTimeoutSecond:        utils.DialTimeoutSecond,
	})
}
