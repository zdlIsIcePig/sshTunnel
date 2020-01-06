package main

import (
	"fmt"
	"net"
	"sshTunnel/ssh"
)

func InitListener(localListenPort string, forwardConfig *ssh.Config) {

	bindListenPort, bindListenPortResult := net.ResolveTCPAddr("tcp", "127.0.0.1"+":"+localListenPort)

	if bindListenPortResult != nil {

		fmt.Printf("tcp listener init failed.Err: %s", bindListenPortResult)
	}

	initListener, initListenerResult := net.ListenTCP("tcp", bindListenPort)

	if initListenerResult != nil {

		fmt.Printf("tcp listener init failed.Err: %s", initListenerResult)
	}

	defer func(listener *net.TCPListener) {

		closeListenResult := initListener.Close()

		if closeListenResult != nil {

			fmt.Printf("Close tcp listen failed.Err: %s", closeListenResult)
		}
	}(initListener)

	for {

		acceptConn, acceptResult := initListener.Accept()

		if acceptResult != nil {

			fmt.Printf("Accept conn is failed.Err: %s", acceptResult)
			continue
		}

		go ssh.Forward(forwardConfig, &acceptConn)
	}
}
