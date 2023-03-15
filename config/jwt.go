package config

type Jwt struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Issuer string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
	Ttl    int    `mapstructure:"ttl" json:"ttl" yaml:"ttl"` //秒
}
