package system

var DisableSwap = `
swapoff -a
sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab 
`
var DisableFirewallod = `
systemctl stop firewalld
systemctl disable firewalld
iptables -F && iptables -X && iptables -F -t nat && iptables -X -t nat
iptables -P FORWARD ACCEPT
`
var DisableSelinux = `
setenforce 0
sed -i 's/^SELINUX=.*/SELINUX=disabled/' /etc/selinux/config
`
var SetDefaultTimeZone = `
timedatectl set-timezone Asia/Shanghai
`
