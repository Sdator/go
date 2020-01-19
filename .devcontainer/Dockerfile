
FROM golang:alpine

# 设置了这个导致vs无法正常调试 build无法执行
# ENV GO111MODULE=on

# 设置镜像站
ENV GOPROXY=https://goproxy.cn


# 替换alpine源为阿里云镜像
# 安装常用程序
RUN sed -i -e 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories \
    && apk --update add --no-cache \
    git \
    openssh

RUN rm -rf /var/cache/apk/*
