#!/bin/bash -
#===============================================================================
#
#          FILE: cfssl.sh
#
#         USAGE: ./cfssl.sh
#
#   DESCRIPTION: init cfssl env. 
#
#       OPTIONS: ---
#  REQUIREMENTS: ---
#          BUGS: ---
#         NOTES: ---
#        AUTHOR: apiopsclub, 
#  ORGANIZATION: 
#       CREATED: 03/27/2021 09:53:49 PM
#      REVISION:  ---
#===============================================================================

set -e

readonly K8S_CERT_DIR="/opt/k8s/cert"
readonly CFSSL_BINARYS="cfssl cfssljson cfssl-certinfo"
readonly DEFAULT_CFSSL_VERSION="1.4.1"

check_cfssl_binary(){
    version=${1}
    mkdir -p $K8S_CERT_DIR
    for item in ${CFSSL_BINARYS}
    do
        which $item || wget https://github.com/cloudflare/cfssl/releases/download/v${version}/${item}_${version}_linux_amd64 -O /usr/bin/${item} && chmod +x /usr/bin/${item} || exit 1
    done
}

main(){
    version=${1:-${DEFAULT_CFSSL_VERSION}}
    check_cfssl_binary $version
}

main $@
