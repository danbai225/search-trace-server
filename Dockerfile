FROM golang:1.18-alpine AS build-env
MAINTAINER DanBai

#修改镜像源为国内
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk update
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOPATH="/go"
RUN cd / && ls
#安装所需工具
RUN apk add gcc g++ make upx git
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
RUN npm config set registry https://registry.npmmirror.com

RUN npm install webpack -g
#拉取代码
RUN mkdir /build
ADD ./ /build
#构建后端
WORKDIR /build
RUN go build -ldflags '-w -s' -o search_trace
RUN upx search_trace
#构建前端
WORKDIR /build/web
RUN npm install
RUN npm run build

FROM alpine:latest
#运行环境
LABEL maintainer="danbai@88.com"
LABEL description="search-trace-server build image file"
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk update
RUN apk --no-cache add tzdata ca-certificates libc6-compat libgcc libstdc++ apache2-utils
#时区
ENV TZ=Asia/Shanghai

#配置时区为中国
RUN apk add tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

RUN mkdir /app
WORKDIR /app
COPY --from=build-env /build/search_trace /app/search_trace
COPY --from=build-env /build/web/dist/web /app/dist
COPY --from=build-env /go/pkg/mod/github.com/yanyiwu/ /go/pkg/mod/github.com/yanyiwu/
RUN chmod +x /app/search_trace
CMD ["/app/search_trace"]