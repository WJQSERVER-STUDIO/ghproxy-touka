# GHProxy-Touka

现已作为 [WJQSERVER-STUDIO/ghproxy](https://github.com/WJQSERVER-STUDIO/ghproxy) v4

支持 Git clone、raw、releases的 Github 加速项目, 支持自托管的同时带来卓越的性能与极低的资源占用, 同时支持多种额外功能

## 项目说明

### 项目特点

- ⚡ **基于 Go 语言实现，跨平台的同时提供高并发性能**
- 🌐 **使用 [Touka](https://github.com/infinite-iroha/touka)作为 Web 框架**
- 📡 **使用 [Touka-HTTPC](https://github.com/satomitouka/touka-httpc) 作为 HTTP 客户端**
- 📥 **支持 Git clone、raw、releases 等文件拉取**
- 🐳 **支持反代Docker, GHCR等镜像仓库**
- 🎨 **支持多个前端主题**
- 🚫 **支持自定义黑名单/白名单**
- 🗄️ **支持 Git Clone 缓存（配合 [Smart-Git](https://github.com/WJQSERVER-STUDIO/smart-git)）**
- 🐳 **支持自托管与Docker容器化部署**
- ⚡ **支持速率限制**
- ⚡ **支持带宽速率限制**
- 🔒 **支持用户鉴权**
- 🐚 **支持 shell 脚本多层嵌套加速**

### 项目相关

[DEMO](https://ghproxy.1888866.xyz)

[TG讨论群组](https://t.me/ghproxy_go)

[相关文章](https://blog.wjqserver.com/categories/my-program/)

[GHProxy项目文档](https://wjqserver-docs.pages.dev/docs/ghproxy/) 感谢 [@redbunnys](https://github.com/redbunnys)的维护

### 使用示例

```bash 
# 下载文件
https://ghproxy.1888866.xyz/raw.githubusercontent.com/WJQSERVER-STUDIO/tools-stable/main/tools-stable-ghproxy.sh
https://ghproxy.1888866.xyz/https://raw.githubusercontent.com/WJQSERVER-STUDIO/tools-stable/main/tools-stable-ghproxy.sh

# 克隆仓库
git clone https://ghproxy.1888866.xyz/github.com/WJQSERVER-STUDIO/ghproxy.git
git clone https://ghproxy.1888866.xyz/https://github.com/WJQSERVER-STUDIO/ghproxy.git

# Docker(OCI) 代理
docker pull gh.example.com/wjqserver/ghproxy
docker pull gh.example.com/adguard/adguardhome

docker pull gh.example.com/docker.io/wjqserver/ghproxy
docker pull gh.example.com/docker.io/adguard/adguardhome

docker pull gh.example.com/ghcr.io/openfaas/queue-worker 
```

### Docker部署

- Docker-cli

```
docker run -p 7210:8080 -v ./ghproxy/log/run:/data/ghproxy/log -v ./ghproxy/log/caddy:/data/caddy/log -v ./ghproxy/config:/data/ghproxy/config  --restart always wjqserver/ghproxy-touka
```

- Docker-Compose (建议使用)

```yml
version: '3.9'
services:
    ghproxy:
        image: 'wjqserver/ghproxy-touka:latest'
        restart: always
        volumes:
            - './ghproxy/log:/data/ghproxy/log'
            - './ghproxy/config:/data/ghproxy/config'
        ports:
            - '7210:8080'
```

## 配置说明

可参考原项目文档

参看[项目文档](https://github.com/WJQSERVER-STUDIO/ghproxy/blob/main/docs/config.md)

### 前端页面

参看[GHProxy-Frontend](https://github.com/WJQSERVER-STUDIO/GHProxy-Frontend)

## 项目简史

**本项目是[WJQSERVER-STUDIO/ghproxy](https://github.com/WJQSERVER-STUDIO/ghproxy)的移植版本, 旨在作为GHProxy v2的替代版本

## LICENSE

继承于[WJQSERVER-STUDIO/ghproxy](https://github.com/WJQSERVER-STUDIO/ghproxy)的Mozila Mozilla Public License Version 2.0

## 赞助

如果您觉得本项目对您有帮助,欢迎赞助支持,您的赞助将用于Demo服务器开支及开发者时间成本支出,感谢您的支持!

USDT(TRC20): `TNfSYG6F2vkiibd6J6mhhHNWDgWgNdF5hN`

### 捐赠列表

| 赞助人    |金额|
|--------|------|
| starry | 8 USDT (TRC20)   |
