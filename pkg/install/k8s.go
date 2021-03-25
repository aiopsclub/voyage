package install

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/helloyi/go-sshclient"
	"voyage/internal/config"
	"voyage/internal/logger"
	"voyage/internal/ssh"
	"voyage/pkg/kernel"
	"voyage/pkg/system"
)

func Run(initConfig *config.InitConfig) error {
	// check kernel version
	for _, masterNode := range initConfig.MasterIPs {
		masterSsh, err := ssh.NewSshClinet(initConfig.UserName, initConfig.Passwd, masterNode+":"+initConfig.SshPort)
		if err != nil {
			return err
		}
		logger.Logger.Infof("Check node %s kernel version.", masterNode)
		vaild, err := checkKernelVersion(masterSsh)
		if err != nil {
			return err
		}
		if !vaild {
			return errors.New(fmt.Sprintf("invaild kernel version for %s", masterNode))
		}

		logger.Logger.Infof("Starting system init for node %s.", masterNode)

		for opName, opScripts := range system.SystemOperation {
			logger.Logger.Infof("Starting %s node %s.", opName, masterNode)
			stdout, stderr, err := scriptRun(masterSsh, opScripts)
			if err != nil {
				logger.Logger.Errorf("stdout: %s, stderr: %s", stdout, stderr)
				return err
			}
			logger.Logger.Infof("%s node %s successfully.", opName, masterNode)
		}

	}
	return nil
}

func scriptRun(client *sshclient.Client, script string) (string, string, error) {
	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}
	err := client.Script(script).SetStdio(&stdout, &stderr).Run()
	if err != nil {
		return "", "", err
	}
	return stdout.String(), stderr.String(), nil
}

func checkKernelVersion(client *sshclient.Client) (bool, error) {
	out, err := client.Cmd("uname -r").SmartOutput()
	if err != nil {
		return false, err
	}
	vaild, err := kernel.VaildKernelVersion(string(out))
	return vaild, err
}
