<p align="center">
  <a href="https://github.com/wgpsec/ENScan_GO">
    <img src="https://github.com/wgpsec/EndpointSearch/assets/16091665/9a26fcef-26fe-4f6b-8c8f-905cdd066296" alt="Logo" width="200" height="200">
  </a>
  <h3 align="center">Endpoint Search</h3>
  <p align="center">
    一个用于侦查云服务端点的工具，此工具参考了 Black Hat议题 Evading Logging in the Cloud: Bypassing AWS CloudTrail
  </p>

<a href="https://github.com/wgpsec/EndpointSearch/stargazers"><img alt="GitHub stars" src="https://img.shields.io/github/stars/wgpsec/EndpointSearch"/></a>
<a href="https://github.com/wgpsec/EndpointSearch/releases"><img alt="GitHub releases" src="https://img.shields.io/github/release/wgpsec/EndpointSearch"/></a>
<a href="https://github.com/wgpsec/EndpointSearch/blob/main/LICENSE"><img alt="License" src="https://img.shields.io/badge/License-Apache%202.0-blue.svg"/></a>
<a href="https://github.com/wgpsec/EndpointSearch/releases"><img alt="Downloads" src="https://img.shields.io/github/downloads/wgpsec/EndpointSearch/total?color=brightgreen"/></a>
<a href="https://goreportcard.com/report/github.com/wgpsec/EndpointSearch"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/wgpsec/EndpointSearch"/></a>
<a href="https://twitter.com/wgpsec"><img alt="Twitter" src="https://img.shields.io/twitter/follow/wgpsec?label=Followers&style=social" /></a>
<br>
<br>
<a href="https://github.com/wgpsec/EndpointSearch/discussions"><strong>探索更多Tricks »</strong></a>
      <br/>
    <br />
    <a href="https://github.com/wgpsec/EndpointSearch/blob/main/README.md">英文文档</a>
    .
    <a href="https://github.com/wgpsec/EndpointSearch/releases">下载程序</a>
    ·
    <a href="https://github.com/wgpsec/EndpointSearch/issues">反馈Bug</a>
    ·
    <a href="https://github.com/wgpsec/EndpointSearch/discussions">提交需求</a>
  </p>

## 安装

下载release中的二进制文件使用

或使用Makefile进行编译二进制文件后使用

## 配置
当首次运行 EndpointSearch 时，会检测 config.yaml 和 rule.yaml 是否存在，不存在则会生成默认的 config.yaml 和 rule.yaml

config.yaml 的填写内容如下：
```
CloudService: oss,ecs
Mode: .
Mode2: -,.
PortList: 80,443
Prefix: sonic,legacy,preprod,gamma,beta,staging
Suffix: sonic,legacy,preprod,gamma,beta,staging
```
CloudService 为枚举的云服务名称，Mode 是连接 CloudService 与 target 的方式，Mode2 是连接前后缀与 CloudService 的方式, PortList 为扫描的端口，具体例子见工作流程

rule.yaml 的填写内容如下：
```
rules:
    - Header:
        - text/xml
        - application/xml
      Body:
        - InvalidVersion
    - Header:
        - "123"
      Body:
        - ""
```
可定义多组 Rule，每组的 Rule 中的 Header 和 Body 均要完全匹配才会被判断为 Endpoint，如果仅有一种特征，可以将另一部分置空。

注意，如果一个规则中的 Header 和 Body 均为空则会导致所有 HTTP 请求均会通过规则

## 工作流程
1. 输入域名 example.com，首先会使用 DNS 去枚举 example.com，枚举方式遵循以下特点：
* Prefix + Mode2 + CloudService + Mode + Host
* CloudService + Mode2 + Suffix + Mode + Host
* CloudService + Mode +Host

例如 Prefix 为 sonic，Suffix 为 legacy，CloudService 为 oss，Mode 为 .，Mode2 为 -，则会枚举：
```
sonic-oss.examlpe.com
oss-legacy.example.com
oss.example.com
```

2. 当域名存在时，会查询 dns 中的 srv 记录发现端口

3. 若已经存在 srv 记录，则不会去枚举端口，而是直接用 HTTP / HTTPS 协议去请求这个URL

4. 否则将通过 HTTP 和 HTTPS 协议去尝试访问目标域名 + PortList 中的端口，如果使用了代理选项，可代理流量至代理服务器

5. 最后通过 HTTP 的请求结果判断整个 URL 是否为 Endpoint，判断访问为判断请求流量是否命中 rule.yaml 中的规则

如果有其他特征，欢迎在 Issues 中提出，或者直接发起 PR。

## 用法
```
Usage:

  EndpointSearch [flags]


Flags:

  -f, --file string       从文件中读取目标地址 (Input filename)
  -h, --help              help for EndpointSearch
      --logLevel string   设置日志等级 (Set log level) [trace|debug|info|warn|error|fatal|panic] (default "info")
  -o, --output string     输入结果文件输出的位置 (Enter the location of the scan result output) (default "./result.txt")
  -p, --port string       输入需要被扫描的端口，逗号分割 (Enter the port to be scanned, separated by commas (,))
      --prefix string     输入枚举云服务的前缀 (Enter the prefix for enumerating the cloud service)
      --proxy string      使用 HTTP/SOCKS5代理，仅限web探测时 (List of http/socks5 proxy to use,Only for web detection
  -s, --service string    输入需要被枚举的服务名称 (Input Service Name)
      --suffix string     输入枚举云服务的后缀 (Enter a suffix for enumerating cloud services)
  -t, --timeout int       输入每个 http 请求的超时时间 (Enter the timeout period for every http request) (default 2)
  -u, --url string        输入目标地址 (Input [domain|url])
```

EndpointSearch 同样支持手动覆盖配置参数，例如 -e 参数默认为配置中的 CloudEndpoint，-p 参数为配置中的 PortList

当主动指定参数后，将不再使用配置文件中的默认值

## 功能列表

1. 利用 dns 服务枚举端点，隐蔽侦查
2. 当域名存在时，自动探测 srv 服务发现端口
3. 支持 HTTP / SOCSK5 代理
4. 自动去重
5. 输入的 url 将自动提取为域名

## TODO
1. 增加更多判断 Endpoint 的方法

