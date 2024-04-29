package define

import "net"

type Configure struct {
	Mode          string `mapstructure:"Mode" json:"Mode" yaml:"Mode"`
	CloudEndpoint string `mapstructure:"CloudEndpoint" json:"CloudEndpoint" yaml:"CloudEndpoint"`
	PortList      string `mapstructure:"PortList" json:"PortList" yaml:"PortList"`
}

var (
	File     string
	Url      string
	Endpoint string
	OutPut   string
	TimeOut  int
	Port     string
)

type IPRecord struct {
	Domain string
	Ip     []net.IP
}

type Record struct {
	Ip         []net.IP
	SvcDomain  string
	SrvRecords []*net.SRV
}
