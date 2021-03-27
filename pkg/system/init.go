package system

var CheckSshConnect = `:`

var disableSwap = `
#!/bin/bash
set -e
swapoff -a
sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab 
`
var disableFirewallod = `
#!/bin/bash
set -e
systemctl stop firewalld
systemctl disable firewalld
iptables -F && iptables -X && iptables -F -t nat && iptables -X -t nat
iptables -P FORWARD ACCEPT
`
var disableSelinux = `
#!/bin/bash
setenforce 0
sed -i 's/^SELINUX=.*/SELINUX=disabled/' /etc/selinux/config
`
var setDefaultTimeZone = `
#!/bin/bash
set -e
timedatectl set-timezone Asia/Shanghai
`
var installDependSoft = `
#!/bin/bash
set -e
yum install -y epel-release
yum install -y chrony conntrack ipvsadm ipset jq iptables curl sysstat libseccomp wget socat git
`

var SystemOperation = []map[string]string{
	{"name": "installDependSoft", "script": installDependSoft},
	{"name": "disableFirewallod", "script": disableFirewallod},
	{"name": "disableSelinux", "script": disableSelinux},
	{"name": "disableSwap", "script": disableSwap},
	{"name": "setDefaultTimeZone", "script": setDefaultTimeZone},
}
