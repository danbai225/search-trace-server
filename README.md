# search-trace-server

记录浏览历史内容，在需要时进行检索。

寻迹服务端,插件仓库[search-trace](https://github.com/danbai225/search-trace)

# 安装

## docker-compose

使用docker-compose快速搭建

### 拉取代码

`git clone https://github.com/danbai225/search-trace-server.git`

### 修改配置

编辑配置`conf.json`

- lite_mode 低配模式
- db mysql数据库配置
- production 是否为正式环境

### 启动

`docker-compose up -d`

## 使用

通过安装配套的[浏览器插件](https://github.com/danbai225/search-trace)来使用