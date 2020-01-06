package ssh

import (
	"io/ioutil"
	"testing"
)

func TestClientInit(t *testing.T) {

	buf, _ := ioutil.ReadFile("If you use the ssh key enter your key file path here")
	sshConfig := Config{
		RemoteAddr:               "Enter the remote address you need to bind here",
		JumpServerPort:           "Enter the remote port you need to bind here",
		JumpServerHost:           "Enter your springboard address here",
		JumpServerUserName:       "Enter your username here",
		JumpServerUserPassword:   "Enter your password here",
		JumpServerUserPrivateKey: string(buf),
		DialTimeoutSecond:        0,
	}
	_, err := ClientInit(&sshConfig)

	if err != nil {

		t.Errorf("ClientInit: %v, err: %s", sshConfig, err)
	}
}
