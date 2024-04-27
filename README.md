# EndpointSearch
EndpointSearch 是一个探测云服务 endpoint 的扫描器，主要用于嗅探私有云的 endpoint 地址

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

