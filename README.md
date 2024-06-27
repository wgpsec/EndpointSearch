<p align="center">
  <a href="https://github.com/wgpsec/ENScan_GO">
    <img src="https://github.com/wgpsec/EndpointSearch/assets/16091665/9a26fcef-26fe-4f6b-8c8f-905cdd066296" alt="Logo" width="200" height="200">
  </a>
  <h3 align="center">Endpoint Search</h3>
  <p align="center">

Endpoint Search is a reconnaissance tool tailored for identifying and enumerating cloud service endpoints. Inspired by the Black Hat talk "Evading Logging in the Cloud: Bypassing AWS CloudTrail," it facilitates stealthy detection of potentially exposed services.
  </p>

<a href="https://github.com/wgpsec/EndpointSearch/stargazers"><img alt="GitHub stars" src="https://img.shields.io/github/stars/wgpsec/EndpointSearch"/></a>
<a href="https://github.com/wgpsec/EndpointSearch/releases"><img alt="GitHub releases" src="https://img.shields.io/github/release/wgpsec/EndpointSearch"/></a>
<a href="https://github.com/wgpsec/EndpointSearch/blob/main/LICENSE"><img alt="License" src="https://img.shields.io/badge/License-Apache%202.0-blue.svg"/></a>
<a href="https://github.com/wgpsec/EndpointSearch/releases"><img alt="Downloads" src="https://img.shields.io/github/downloads/wgpsec/EndpointSearch/total?color=brightgreen"/></a>
<a href="https://goreportcard.com/report/github.com/wgpsec/EndpointSearch"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/wgpsec/EndpointSearch"/></a>
<a href="https://twitter.com/wgpsec"><img alt="Twitter" src="https://img.shields.io/twitter/follow/wgpsec?label=Followers&style=social" /></a>
<br>
<br>
<a href="https://github.com/wgpsec/EndpointSearch/discussions"><strong>More Tricks »</strong></a>
      <br/>
    <br />
        <a href="https://github.com/wgpsec/EndpointSearch/blob/main/README_CN.md">中文文档</a>
    .
    <a href="https://github.com/wgpsec/EndpointSearch/releases">Download</a>
    ·
    <a href="https://github.com/wgpsec/EndpointSearch/issues">Issues</a>
    ·
    <a href="https://github.com/wgpsec/EndpointSearch/discussions">Discussions</a>
  </p>

## Features
* **DNS Enumeration**: Constructs and queries DNS for probable endpoint URLs based on predefined patterns.
* **SRV Record Inspection**: Automatically detects SRV records to uncover associated ports.
* **HTTP/HTTPS Probing**: Tests targets with both HTTP and HTTPS protocols when SRV records are not present.
* **PROXY SUPPORT**: Supports HTTP/HTTPS and SOCK5 protocol proxy traffic during the HTTP request phase.
* **Endpoint Judgment**: Determine the likelihood of the endpoint based on whether the response content hits the rule.
* **Automatic Domain Extraction**: Extracts domains from input URLs automatically.
* **Output Redundancy Removal**: Ensures unique results by deduplication.
* **Configurable Behavior**: Offers a flexible configuration file for customizing service names, connection modes, and more.

## Installation

* Download precompiled binaries from the [releases page](https://github.com/wgpsec/EndpointSearch/releases).

* Alternatively, use the included Makefile to compile from source.

## Configuration
When EndpointSearch is run for the first time, config.yaml and rule-yaml are detected and the default config.yaml and rule-yaml are generated if they are not present

config.yaml fill in as follows:
```
CloudService: oss,ecs
Mode: .
Mode2: -,.
PortList: 80,443
Prefix: sonic,legacy,preprod,gamma,beta,staging
Suffix: sonic,legacy,preprod,gamma,beta,staging
```
CloudService is an enumerated cloud service name. Mode is the mode used to connect CloudService to target. Mode2 is the mode used to connect prefixes and suffixes to CloudService

The content of rule.yaml is as follows:
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
Multiple groups of rules can be defined. Header and Body in the Rule of each group must match exactly to be identified as endpoints. If there is only one feature, the other part can be left blank.

Note that if both Header and Body in a rule are empty, all HTTP requests will pass the rule

## Workflow
![EndpointSearch](https://github.com/wgpsec/EndpointSearch/assets/50265741/bbe62843-ff98-46d8-85ad-5cb0ce1fcc51)

1. Enter the domain name example.com. DNS is used to enumerate example.com.
* Prefix + Mode2 + CloudService + Mode + Host
* CloudService + Mode2 + Suffix + Mode + Host
* CloudService + Mode +Host

For example, if Prefix is sonic, Suffix is legacy, CloudService is oss, Mode is., and Mode2 is -, the system will enumerate:
```
sonic-oss.examlpe.com
oss-legacy.example.com
oss.example.com
```

2. If the domain name exists, the system queries the srv records in the dns to discover the port

3. If srv records already exist, HTTP/HTTPS is used to request the URL instead of enumerating the port

4. Otherwise, HTTP and HTTPS are used to access ports in the target domain name + PortList. If the proxy option is used, traffic can be forwarded to the proxy server

5. Finally, the HTTP request result determines whether the entire URL is an Endpoint, and the access is to determine whether the request traffic matches the rule in rule.yaml

If there are other characteristics, feel free to raise them in the Issues, or launch a PR directly.

## Usage
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

EndpointSearch can also override configuration parameters manually. For example, -e is set to CloudEndpoint by default, and -p is set to PortList by default

When parameters are actively specified, the default values in the configuration file are no longer used


## TODO
1. Added more ways to determine endpoints

