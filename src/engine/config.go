package engine

type ConfigMysql struct {
	Enable      bool   `toml:"Enable"`
	Dsn         string `toml:"Dsn"`
	MaxLifetime int    `toml:"MaxLifetime"`
	MaxIdle     int    `toml:"MaxIdle"`
	MaxOpen     int    `toml:"MaxOpen"`
}
type ConfigWx struct {
	Token     string `toml:"Token"`
	AppId     string `toml:"AppId"`
	AppSecret string `toml:"AppSecret"`
}
type ConfigOss struct {
	Ak         string `toml:"Ak"`
	Sk         string `toml:"Sk"`
	EndPoint   string `toml:"End_point"`
	BucketName string `toml:"BucketName"`
}
type ConfigSystem struct {
	LogPath string `toml:"LogPath"`
	IsDebug bool   `toml:"IsDebug"`
	Addr    string `toml:"Addr"`
}
type ConfigDomains struct {
	ShortUrl string `toml:"ShortUrl"`
	ImgHttps string `toml:"ImgHttps"`
	ImgHttp  string `toml:"ImgHttp"`
}
type Config struct {
	Mysql   ConfigMysql
	Wx      ConfigWx
	Oss     ConfigOss
	Domains ConfigDomains
	System  ConfigSystem
}
