package kernel

import (
	"github.com/Masterminds/semver"
	"strings"
)

func VaildKernelVersion(kernelVersion string) (bool, error) {
	// check kernel min version, example: 5.10.12-1.el7.elrepo.x86_64
	versionInfo := strings.Split(kernelVersion, ".")
	c2, _ := semver.NewConstraint(">=4.4.0-0")
	v, err := semver.NewVersion(strings.Join(versionInfo[0:3], "."))
	if err != nil {
		return false, err
	}
	return c2.Check(v), nil
}
