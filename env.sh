mkdir -p bin
mkdir -p logs
CURDIR=`pwd`
export GO111MODULE=on
export GOPROXY="https://goproxy.cn"

workspace=$(cd $(dirname $0) && pwd -P)
echo "$workspace"