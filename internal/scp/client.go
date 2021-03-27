package scp

import (
	"errors"
	"fmt"
	scpclient "github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
	"os"
)

func CopyFile(sshClient *ssh.Client, srcFile string, destFile string, address string) error {
	client, err := scpclient.NewClientBySSH(sshClient)
	if err != nil {
		return errors.New(fmt.Sprint("Error creating new SSH session from existing connection", err))
	}

	// Open a file
	f, _ := os.Open("/path/to/local/file")

	// Close client connection after the file has been copied
	defer client.Close()

	// Close the file after it has been copied
	defer f.Close()

	// Finaly, copy the file over
	// Usage: CopyFile(fileReader, remotePath, permission)

	err = client.CopyFile(f, "/home/server/test.txt", "0655")

	if err != nil {
		return errors.New(fmt.Sprintln("Error while copying file ", err))
	}
	return nil
}
