package ssh

import (
	"github.com/helloyi/go-sshclient"
	"golang.org/x/crypto/ssh"
	"net"
)

func NewSshClinet(username, password, address string) (client *sshclient.Client, err error) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	client, err = sshclient.Dial("tcp", address, config)
	return
}
