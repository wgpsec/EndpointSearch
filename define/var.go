package define

type Configure struct {
	Mode          string `mapstructure:"Mode" json:"Mode" yaml:"Mode"`
	CloudEndpoint string `mapstructure:"CloudEndpoint" json:"CloudEndpoint" yaml:"CloudEndpoint"`
}

var (
	File     string
	Url      string
	Endpoint string
	OutPut   string
)
