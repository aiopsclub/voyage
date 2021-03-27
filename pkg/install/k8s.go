package install

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/helloyi/go-sshclient"
	"strconv"
	"voyage/internal/config"
	"voyage/internal/logger"
	"voyage/internal/ssh"
	"voyage/internal/util"
	"voyage/pkg/kernel"
	"voyage/pkg/system"
)

type Connects struct {
	HostName string
	Conn     *sshclient.Client
}

func Run(initConfig *config.InitConfig) error {
	masterHostList, err := newMasterHostList(initConfig)
	if err != nil {
		return err
	}
	hostScript, err := masterHostList.MasterHostsInfo()

	masterSshConnect, err := newmasterSshConnect(initConfig)
	if err != nil {
		return err
	}

	defer func() {
		for _, connInfo := range masterSshConnect {
			connInfo.Conn.Close()
		}
	}()

	for _, sshConnInfo := range masterSshConnect {
		// check kernel version
		logger.Logger.Infof("Check node %s kernel version.", sshConnInfo.HostName)
		vaild, err := checkKernelVersion(sshConnInfo.Conn)
		if err != nil {
			return err
		}
		if !vaild {
			return errors.New(fmt.Sprintf("invaild kernel version for %s", sshConnInfo.HostName))
		}

		logger.Logger.Infof("Starting system init for node %s.", sshConnInfo.HostName)

		for _, opScript := range system.SystemOperation {
			logger.Logger.Infof("Starting %s node %s.", opScript["name"], sshConnInfo.HostName)
			stdout, stderr, err := scriptRun(sshConnInfo.Conn, opScript["script"])
			if err != nil {
				logger.Logger.Errorf("stdout: %s, stderr: %s", stdout, stderr)
				return err
			}
			logger.Logger.Infof("%s node %s successfully.", opScript["name"], sshConnInfo.HostName)
		}
		logger.Logger.Infof("Starting config hosts node %s.", sshConnInfo.HostName)
		stdout, stderr, err := scriptRun(sshConnInfo.Conn, hostScript)
		if err != nil {
			logger.Logger.Errorf("stdout: %s, stderr: %s", stdout, stderr)
			return err
		}
	}
	return nil
}

func newMasterHostList(initConfig *config.InitConfig) (util.MasterHostList, error) {
	var masterHostList util.MasterHostList
	if len(initConfig.MasterIPs) == 0 {
		return masterHostList, errors.New("master ips can't be null.")
	}
	for masterIndex, masterNode := range initConfig.MasterIPs {
		masterHostList.Hosts = append(masterHostList.Hosts, &util.MasterHost{
			HostName: "node" + strconv.Itoa(masterIndex) + "." + "voyage.com",
			Ip:       masterNode,
		})

	}
	return masterHostList, nil
}

func newmasterSshConnect(initConfig *config.InitConfig) ([]*Connects, error) {
	var masterSshConnect []*Connects
	for _, masterNode := range initConfig.MasterIPs {
		masterSsh, err := ssh.NewSshClinet(initConfig.UserName, initConfig.Passwd, masterNode+":"+initConfig.SshPort)
		if err != nil {
			return nil, err
		}
		masterSshConnect = append(masterSshConnect, &Connects{
			HostName: masterNode,
			Conn:     masterSsh,
		})
	}
	return masterSshConnect, nil
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
