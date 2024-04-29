<p align="center">
  <a href="https://github.com/wgpsec/ENScan_GO">
    <img src="https://github.com/wgpsec/EndpointSearch/assets/16091665/9a26fcef-26fe-4f6b-8c8f-905cdd066296" alt="Logo" width="200" height="200">
  </a>
  <h3 align="center">Endpoint Search</h3>
  <p align="center">
    一个探测云服务 endpoint 的扫描器，主要用于嗅探私有云的 endpoint 地址
    <br />
          <br />
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
当首次运行 EndpointSearch 时，会检测 config.json 文件是否存在，不存在则会自动创建

config.json的填写内容应该如下：
```
{
	"CloudEndpoint":"oss,ecs",
	"Mode":",-,_,."
}
```
CloudEndpoint 为枚举的端点，Mode是端点与域名的连接方式

例如输入域名 example.com，则会枚举 ossexample.com, oss-example.com...

CloudEopint 可参考我的另一个字典项目: https://github.com/shadowabi/S-BlastingDictionary/blob/main/CloudEndpoint.txt

## 用法
```
Usage:

  EndpointSearch [flags]


Flags:

  -e, --endpoint string   输入
  -f, --file string       从文件中读取目标地址 (Input filename)
  -h, --help              help for EndpointSearch
      --logLevel string   设置日志等级 (Set log level) [trace|debug|info|warn|error|fatal|panic] (default "info")
  -o, --output string     输入结果文件输出的位置 (Enter the location of the scan result output) (default "./result.txt")
  -u, --url string        输入目标地址 (Input [domain|url])

```

## 功能列表

1. 利用 dns 服务枚举端点，隐蔽侦查
2. 当域名存在时，自动探测 srv 服务发现端口
3. 自动去重

## TODO
1. 添加 socket5 代理的支持

