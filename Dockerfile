FROM golang:1.18-alpine AS build-env
MAINTAINER DanBai

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
ADD ./ /build
#构建后端
WORKDIR /build
RUN go build -o search_trace
#构建前端
WORKDIR /build/web
RUN cnpm install
RUN cnpm run build

FROM alpine:latest
#运行环境

LABEL maintainer="danbai@88.com"
LABEL version="0.1"
LABEL description="search-trace-server build image file"
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk update
#需要如下链接操作,否则运行程序会提示not found
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
#时区
ENV TZ=Asia/Shanghai
RUN apk add --no-cache ca-certificates apache2-utils
#配置时区为中国
RUN apk add tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

RUN mkdir /app
WORKDIR /app
COPY --from=build-env /build/search_trace /app/search_trace
COPY --from=build-env /build/web/dist/web /app/dist
RUN chmod +x /app/search_trace
CMD ["bash"]