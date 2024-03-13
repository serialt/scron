package main

var (
	// 版本信息
	appVersion bool // 控制是否显示版本
	APPVersion = "v0.0.2"
	BuildTime  = "2006-01-02 15:04:05"
	GitCommit  = "xxxxxxxxxxx"
	ConfigFile = "config.yaml"
	config     *Config
)

type Config struct {
	Cron []Cron `yaml:"cron"`
	Log  Log    `yaml:"log"`
}

type Cron struct {
	Name   string `yaml:"name"`
	Time   string `yaml:"time"`
	Action string `yaml:"action"`
}

type Log struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}
