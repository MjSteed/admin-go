package config

type Sign struct {
	Auth map[string]string `mapstructure:"auth" json:"auth" yaml:"auth"`
	Exp  int64             `mapstructure:"exp" json:"exp" yaml:"exp"` //过期时间，秒
}
