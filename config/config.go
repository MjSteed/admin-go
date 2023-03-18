package config

type ApplicationConfig struct {
	Logger   Logger   `mapstructure:"log" json:"log" yaml:"log"`
	Jwt      Jwt      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	AwsS3    AwsS3    `mapstructure:"aws-s3" json:"aws-s3" yaml:"aws-s3"`
	Sign     Sign     `mapstructure:"sign" json:"sign" yaml:"sign"`
}
