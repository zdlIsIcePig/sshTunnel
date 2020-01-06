package ssh

import (
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"time"
)

//Choose one of the two options to connect to the springboard by password or key
func ClientInit(sshConfig *Config) (conn *ssh.Client, err error) {

	sshClientConfig := makeConfig(sshConfig.JumpServerUserName, sshConfig.JumpServerUserPassword,
		sshConfig.JumpServerUserPrivateKey)

	if sshConfig.DialTimeoutSecond > 0 {
		connNet, err := net.DialTimeout("tcp", sshConfig.JumpServerHost+":"+sshConfig.JumpServerPort,
			time.Duration(sshConfig.DialTimeoutSecond)*time.Second)
		if err != nil {
			log.Println(err)
		}
		sc, chanS, reqs, err := ssh.NewClientConn(connNet, sshConfig.JumpServerHost+":"+sshConfig.JumpServerPort,
			sshClientConfig)
		if err != nil {
			log.Println(err)
		}
		conn = ssh.NewClient(sc, chanS, reqs)
	} else {
		conn, err = ssh.Dial("tcp", sshConfig.JumpServerHost+":"+sshConfig.JumpServerPort, sshClientConfig)
		if err != nil {
			return
		}
	}
	log.Println("dial ssh success")
	return
}

func makeConfig(user string, password string, privateKey string) (config *ssh.ClientConfig) {

	if password == "" && user == "" {
		log.Fatal("No password or private key available")
	}
	config = &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	if privateKey != "" {
		signer, err := ssh.ParsePrivateKey([]byte(privateKey))
		if err != nil {
			log.Fatalf("ssh.ParsePrivateKey error:%v", err)
		}
		clientKey := ssh.PublicKeys(signer)
		config = &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				clientKey,
			},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		}
	}
	return
}
