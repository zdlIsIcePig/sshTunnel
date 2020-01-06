package ssh

import (
	"io"
	"log"
	"net"
)

type Config struct {
	RemoteAddr               string
	JumpServerPort           string
	JumpServerHost           string
	JumpServerUserName       string
	JumpServerUserPassword   string
	JumpServerUserPrivateKey string
	DialTimeoutSecond        uint64
}

func Forward(config *Config, conn *net.Conn) {

	sshConn, err := ClientInit(config)

	if err != nil {

		log.Fatalf("init ssh client error:%v", err)
	}

	sshForwardConn, err := sshConn.Dial("tcp", config.RemoteAddr)
	if err != nil {
		log.Fatalf("ssh client dial error:%v", err)
	}
	log.Println("create ssh connection ok")
	go localReaderToRemoteWriter(*conn, sshForwardConn)
	go remoteReaderToLocalWriter(sshForwardConn, *conn)
}

func localReaderToRemoteWriter(localConn net.Conn, sshConn net.Conn) {

	_, err := io.Copy(sshConn, localConn)
	if err != nil {
		log.Fatalf("io copy error:%v", err)
	}
}

func remoteReaderToLocalWriter(sshConn net.Conn, localConn net.Conn) {

	_, err := io.Copy(localConn, sshConn)
	if err != nil {
		log.Fatalf("io copy error:%v", err)
	}
}
