package tunnel

import (
	"io/ioutil"
	"sshTunnel/ssh"
	"testing"
)

func TestInitListener(t *testing.T) {

	buf, _ := ioutil.ReadFile("If you use the ssh key enter your key file path here")
	InitListener("56655", &ssh.Config{
		RemoteAddr:               "Enter the remote address you need to bind here",
		JumpServerPort:           "Enter the remote port you need to bind here",
		JumpServerHost:           "Enter your springboard address here",
		JumpServerUserName:       "Enter your username here",
		JumpServerUserPassword:   "Enter your password here",
		JumpServerUserPrivateKey: string(buf),
		DialTimeoutSecond:        0,
	})
}
