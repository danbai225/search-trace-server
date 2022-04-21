FROM golang:1.18-alpine AS build-env
MAINTAINER DanBai

#需要如下链接操作,否则运行程序会提示not found
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

#修改镜像源为国内
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk update
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
#安装git
RUN apk add git
#安装gcc
RUN apk add gcc g++ make
#配置时区为中国
RUN apk add tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

#安装node
RUN apk add nodejs
RUN node -v
RUN apk add npm
RUN npm -v
#npm 镜像
RUN alias cnpm="npm --registry=https://registry.npmmirror.com \
--cache=$HOME/.npm/.cache/cnpm \
--disturl=https://npmmirror.com/mirrors/node \
--userconfig=$HOME/.cnpmrc"
RUN npm install -g cnpm --registry=https://registry.npm.taobao.org

RUN cnpm install webpack -g
#拉取代码
RUN mkdir /build
WORKDIR /build
RUN git clone https://ghproxy.com/https://github.com/danbai225/search-trace-server
#构建后端
WORKDIR /build/search-trace-server
RUN go build
#构建前端
WORKDIR /build/search-trace-server/web
RUN cnpm install
RUN cnpm run build

FROM ubuntu:20.04
#运行环境

LABEL maintainer="danbai@88.com"
LABEL version="0.1"
LABEL description="search-trace-server build image file"
#时区
ENV TZ=Asia/Shanghai
RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list \
          && apt-get update \
          && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone \
          && apt-get install tzdata \
          && apt-get clean \
          && apt-get autoclean \
          && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN mkdir /app
WORKDIR /app
COPY --from=build-env /build/search-trace-server /app/search-trace-server
COPY --from=build-env /build/search-trace-server/web/dist /app/dist
RUN chmod +x /build/search-trace-server
CMD ["/build/search-trace-server"]