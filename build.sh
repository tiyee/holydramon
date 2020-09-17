#!/usr/bin/env bash
workspace=$(cd $(dirname $0) && pwd -P)
cd ${workspace}
gitversion=.gitversion
app=holydramon

## function
function build() {
	# 设置golang环境变量
    echo -e "default `go version`"
    cd /usr/local && ls -al | grep "go"
    cd ${workspace}


#	local go="/usr/local/go"
#	if [[ -d "$go" ]]; then
#	    export GOROOT="$go"
#	    export PATH=${GOROOT}/bin:$PATH
#	    echo -e "use `go version`"
#	fi

    # 进行编译

#go env -w GOPROXY=https://goproxy.cn,direct
    go build -o bin/${app} src/main.go src/router.go src/hooks.go
    local sc=$?
    if [[ ${sc} -ne 0 ]];then
    	## 编译失败, 退出码为 非0
        echo "$app build error"
        exit ${sc}
    else
        echo -n "$app build ok, vsn="
        gitversion
    fi
}



## internals
function gitversion() {
    git log -1 --pretty=%h > ${gitversion}
    local gv=`cat ${gitversion}`
    echo "$gv"
}


##########################################
## main
## 其中,
## 		1.进行编译
##		2.生成部署包output
##########################################

# 1.进行编译
build


# 编译成功
echo -e "build done"
exit 0

