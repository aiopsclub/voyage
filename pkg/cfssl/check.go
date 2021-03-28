package cfssl

import (
	"errors"
	"fmt"
	"os"
	"voyage/internal/logger"
	"voyage/internal/util"
)

var cfsslBinarys = []string{"cfssl", "cfssljson", "cfssl-certinfo"}
var cfsslDestDir = "/usr/bin"
var cfsslVersion = "1.4.1"

func CheckAndInstall() error {
	for _, cfsslCmd := range cfsslBinarys {
		logger.Logger.Info(fmt.Sprintf("Check and install %s", cfsslCmd))
		if !util.IsCommandExist(cfsslCmd) {
			_, stderr, err := util.CommandRun(fmt.Sprintf("wget https://github.com/cloudflare/cfssl/releases/download/v%s/%s_%s_linux_amd64 -O %s/%s", cfsslVersion, cfsslCmd, cfsslVersion, cfsslDestDir, cfsslCmd))
			if err != nil {
				_, _, _ = util.CommandRun(fmt.Sprintf("rm -f %s/%s", cfsslDestDir, cfsslCmd))
				return errors.New(fmt.Sprintf("Install %s error. reason: %s", cfsslCmd, stderr))
			}
			os.Chmod(fmt.Sprintf("%s/%s", cfsslDestDir, cfsslCmd), 0755)
		}

	}
	return nil
}
