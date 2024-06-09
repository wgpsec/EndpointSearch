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
        <a href="https://github.com/wgpsec/EndpointSearch/blob/main/README.md">中文文档</a>
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
* **HTTP/HTTPS Probing**: Tests endpoints with both HTTP and HTTPS protocols when SRV records are not present.
* **Endpoint Judgment**: Determines endpoints likelihood based on response content, currently focusing on XML format.
* **Automatic Domain Extraction**: Extracts domains from input URLs automatically.
* **Output Redundancy Removal**: Ensures unique results by deduplication.
* **Configurable Behavior**: Offers a flexible configuration file for customizing service names, connection modes, and more.

## Installation

* Download precompiled binaries from the [releases page](https://github.com/wgpsec/EndpointSearch/releases).

* Alternatively, use the included Makefile to compile from source.

## Configuration
Upon first run, the tool checks for config.json. If missing, it generates one with default settings including:
* CloudService: Enumerated cloud services.
* Mode & Mode2: Patterns connecting services, prefixes/suffixes, and targets.
* PortList: Ports to scan if no SRV record is found.
* Prefix & Suffix: Common naming conventions for prefixing or suffixing service names.
```
{
	"CloudService":"oss,ecs",
	"Mode":".",
	"Mode2" :"-,.",
	"PortList":"80,443",
	"Prefix":"sonic,legacy,preprod,gamma,beta,staging",
	"Suffix":"sonic,legacy,preprod,gamma,beta,staging",
}
```

CloudService can refer to another of my dictionary projects : https://github.com/shadowabi/S-BlastingDictionary/blob/main/CloudService.txt

## Workflow
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

4. Otherwise, the system attempts to access the port in the target domain name + PortList through HTTP or HTTPS

5. Finally, determine whether the entire URL is an Endpoint based on the HTTP request result. At present, determine whether the destination data is returned in xml format

If there are other characteristics, feel free to raise them in the Issues, or launch a PR directly.

The method to determine the Endpoint is implemented in the JudgeEndpoint function in the pkg directory data.go

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
      --prefix string     输入需要被枚举的服务名称 (Input Service Name)
  -s, --service string    输入需要被枚举的服务名称 (Input Service Name)
      --suffix string     输入需要被枚举的服务名称 (Input Service Name)
  -t, --timeout int       输入每个 http 请求的超时时间 (Enter the timeout period for every http request) (default 2)
  -u, --url string        输入目标地址 (Input [domain|url])
```
EndpointSearch also supports manually overwriting configuration parameters. By default, the -e parameter is CloudEndpoint in the configuration, and the -p parameter is PortList in the configuration

When parameters are actively specified, the default values in the configuration file are no longer used


## TODO
1. Proxy Support: Implementation of SOCKS5 proxy support for enhanced anonymity.
2. Enhanced Endpoint Detection: Expanding endpoint validation criteria beyond XML responses.

