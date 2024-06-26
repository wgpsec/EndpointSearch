package define

type Configure struct {
	Mode         string `mapstructure:"Mode" json:"Mode" yaml:"Mode"`
	Mode2        string `mapstructure:"Mode2" json:"Mode2" yaml:"Mode2"`
	CloudService string `mapstructure:"CloudService" json:"CloudService" yaml:"CloudService"`
	PortList     string `mapstructure:"PortList" json:"PortList" yaml:"PortList"`
	Prefix       string `mapstructure:"Prefix" json:"Prefix" yaml:"Prefix"`
	Suffix       string `mapstructure:"Suffix" json:"Suffix" yaml:"Suffix"`
}

type Rules struct {
	RuleText []RuleText `mapstructure:"rules" json:"rules" yaml:"rules"`
}

type RuleText struct {
	Header []string `mapstructure:"Header" json:"Header" yaml:"Header"`
	Body   []string `mapstructure:"Body" json:"Body" yaml:"Body"`
}

var (
	File     string
	Url      string
	Service  string
	OutPut   string
	TimeOut  int
	Port     string
	Prefix   string
	Suffix   string
	ProxyURL string
)
