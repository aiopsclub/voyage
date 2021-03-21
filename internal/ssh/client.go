package ssh

import (
	"github.com/helloyi/go-sshclient"
	"golang.org/x/crypto/ssh"
)

func NewSshClinet(username, password, address string) (client *sshclient.Client, err error) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}
	client, err = sshclient.Dial("network", address, config)
	return
}
