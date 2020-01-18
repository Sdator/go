
FROM golang:alpine


ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Configure apt, install packages and tools
# RUN apt-get update \
# && apt-get -y install --no-install-recommends apt-utils dialog 2>&1 \
#
# Verify git, process tools, lsb-release (common in install instructions for CLIs) installed
# && apt-get -y install git iproute2 procps lsb-release \
#

# 开启模块 使用阿里云镜像站
ENV GO111MODULE=on
ENV GOPROXY="https://mirrors.aliyun.com/goproxy/,direct"

# # 替换alpine源为阿里云镜像
RUN sed -i -e 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories \
    && apk --update add --no-cache \
    # 已内置 openssl go gcc 
    git \
    sudo \
    bash

    # # Install gocode-gomod
    # && go get -x -d github.com/stamblerre/gocode 2>&1 \
    # && go build -o gocode-gomod github.com/stamblerre/gocode \
    # && mv gocode-gomod $GOPATH/bin/ \

    # # Install Go tools
    # && go get -u -v \
    # github.com/mdempsky/gocode \
    # github.com/uudashr/gopkgs/cmd/gopkgs \
    # github.com/ramya-rao-a/go-outline \
    # github.com/acroca/go-symbols \
    # github.com/godoctor/godoctor \
    # golang.org/x/tools/cmd/guru \
    # golang.org/x/tools/cmd/gorename \
    # github.com/rogpeppe/godef \
    # github.com/zmb3/gogetdoc \
    # github.com/haya14busa/goplay/cmd/goplay \
    # github.com/sqs/goreturns \
    # github.com/josharian/impl \
    # github.com/davidrjenni/reftools/cmd/fillstruct \
    # github.com/fatih/gomodifytags \
    # github.com/cweill/gotests/... \
    # golang.org/x/tools/cmd/goimports \
    # golang.org/x/lint/golint \
    # golang.org/x/tools/cmd/gopls \
    # github.com/alecthomas/gometalinter \
    # honnef.co/go/tools/... \
    # github.com/golangci/golangci-lint/cmd/golangci-lint \
    # github.com/mgechev/revive \
    # github.com/derekparker/delve/cmd/dlv 2>&1

# 添加用户组
RUN addgroup -g $USER_GID $USERNAME \
    # 添加用户
    && adduser -s /bin/bash -u $USER_UID -G $USERNAME -D $USERNAME \
    # 添加到sudo用户组 使用户可以使用sudo指令
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME