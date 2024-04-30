package define

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
