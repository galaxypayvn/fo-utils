package mqtt

const (
	_defaultReconnectPeriod = 1000
	_defaultConnectTimeout  = 30000
)

type Config struct {
	Broker          string `json:"broker" mapstructure:"broker"`
	ConnectTimeout  int    `json:"connect_timeout" mapstructure:"connect_timeout"`
	ReconnectPeriod int    `json:"reconnect_period" mapstructure:"reconnect_period"`
	UserName        string `json:"username" mapstructure:"username"`
	Password        string `json:"password" mapstructure:"password"`
}

func (o *Config) Prepare() {
	if o == nil {
		return
	}

	if o.ConnectTimeout < 1 {
		o.ConnectTimeout = _defaultConnectTimeout
	}

	if o.ReconnectPeriod < 1 {
		o.ReconnectPeriod = _defaultReconnectPeriod
	}
}
