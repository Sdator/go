
FROM golang:alpine

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# 设置了这个导致vs无法正常调试 build无法执行
# ENV GO111MODULE=on

# 设置镜像站
ENV GOPROXY=https://goproxy.cn


# 替换alpine源为阿里云镜像
# 安装常用程序
RUN sed -i -e 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories \
    && apk --update add --no-cache \
    git \
    openssh \
    sudo

RUN rm -rf /var/cache/apk/*


# 添加用户组
RUN addgroup -g $USER_GID $USERNAME \
    # 添加用户
    && adduser -s /bin/bash -u $USER_UID -G $USERNAME -D $USERNAME \
    # 添加到sudo用户组 使用户可以使用sudo指令
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME