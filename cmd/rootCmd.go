package cmd

import (
	"errors"
	"fmt"
	"github.com/wgpsec/EndpointSearch/define"
	"github.com/wgpsec/EndpointSearch/internal/config"
	"github.com/wgpsec/EndpointSearch/internal/log"
	"github.com/wgpsec/EndpointSearch/pkg"
	"github.com/wgpsec/EndpointSearch/utils/Compare"
	"github.com/wgpsec/EndpointSearch/utils/Error"
	"os"
	"strings"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "EndpointSearch",
	Short: "EndpointSearch is a scanner that probes the endpoint of a cloud service",
	Long: "  _____           _             _       _   ____                      _     \n" +
		" | ____|_ __   __| |_ __   ___ (_)_ __ | |_/ ___|  ___  __ _ _ __ ___| |__  \n" +
		" |  _| | '_ \\ / _` | '_ \\ / _ \\| | '_ \\| __\\___ \\ / _ \\/ _` | '__/ __| '_ \\ \n" +
		" | |___| | | | (_| | |_) | (_) | | | | | |_ ___) |  __/ (_| | | | (__| | | |\n" +
		" |_____|_| |_|\\__,_| .__/ \\___/|_|_| |_|\\__|____/ \\___|\\__,_|_|  \\___|_| |_|\n" +
		"                   |_|                                                      \n" +
		`
		github.com/wgpsec/EndpointSearch

EndpointSearch 是一个探测云服务 endpoint 的扫描器
EndpointSearch is a scanner that probes the endpoint of a cloud service 
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.Init(logLevel)
		if define.Url != "" && define.File != "" {
			Error.HandleFatal(errors.New("参数不可以同时存在"))
			return
		}
		if define.Url == "" && define.File == "" {
			Error.HandleFatal(errors.New("必选参数为空，请输入 -u 参数或 -f 参数"))
			return
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if define.Endpoint == "" {
			define.Endpoint = config.C.CloudEndpoint
		}
		if define.Port == "" {
			define.Port = config.C.PortList
		}
		portList := strings.Split(define.Port, ",")

		var hostList []string
		if define.File != "" {
			hostList = pkg.ParseFileParameter(define.File)
		} else {
			hostList = append(hostList, define.Url)
		}
		hostList = Compare.RemoveDuplicates(hostList)

		reqList := pkg.ConvertToReqList(define.Endpoint, hostList...)
		ipRecordList := pkg.SearchDomain(reqList...)
		recordList := pkg.SearchSRVRecord(ipRecordList...)

		client := pkg.GenerateHTTPClient(define.TimeOut)
		respList := pkg.SearchEndpoint(client, portList, recordList...)

		resultList := Compare.RemoveDuplicates(pkg.JudgeEndpoint(respList...))
		pkg.WriteToFile(resultList, define.OutPut)
		fmt.Printf("[+] The output is in %s\n", define.OutPut)
	},
}

var logLevel string

func init() {
	RootCmd.PersistentFlags().StringVar(&logLevel, "logLevel", "info", "设置日志等级 (Set log level) [trace|debug|info|warn|error|fatal|panic]")
	RootCmd.CompletionOptions.DisableDefaultCmd = true
	RootCmd.SetHelpFunc(customHelpFunc)
	RootCmd.Flags().StringVarP(&define.File, "file", "f", "", "从文件中读取目标地址 (Input filename)")
	RootCmd.Flags().StringVarP(&define.Url, "url", "u", "", "输入目标地址 (Input [domain|url])")
	RootCmd.Flags().StringVarP(&define.Endpoint, "endpoint", "e", "", "输入")
	RootCmd.Flags().IntVarP(&define.TimeOut, "timeout", "t", 2, "输入每个 http 请求的超时时间 (Enter the timeout period for every http request)")
	RootCmd.Flags().StringVarP(&define.OutPut, "output", "o", "./result.txt", "输入结果文件输出的位置 (Enter the location of the scan result output)")
	RootCmd.Flags().StringVarP(&define.Port, "port", "p", "", "输入需要被扫描的端口，逗号分割 (Enter the port to be scanned, separated by commas (,))")
}

func Execute() {
	cc.Init(&cc.Config{
		RootCmd:  RootCmd,
		Headings: cc.HiGreen + cc.Underline,
		Commands: cc.Cyan + cc.Bold,
		Example:  cc.Italic,
		ExecName: cc.Bold,
		Flags:    cc.Cyan + cc.Bold,
	})
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func customHelpFunc(cmd *cobra.Command, args []string) {
	cmd.Usage()
}
