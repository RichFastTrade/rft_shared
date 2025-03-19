package conf

import (
	"sync"
)

var cfg *Config
var once sync.Once

type Config struct {
	AppName string `yaml:"appName"`
	Version string `yaml:"version"`
	Url     string `yaml:"url"`
	Server  struct {
		Http struct {
			Addr string `yaml:"addr"`
		} `yaml:"http"`
		Grpc struct {
			Addr string `yaml:"addr"`
		} `yaml:"grpc"`
	} `yaml:"server"`
	Postgresql struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"postgresql"`
	Rabbitmq struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"rabbitmq"`
	Pyroscope struct {
		Addr string `yaml:"addr"`
	} `yaml:"pyroscope"`
	KlinePeriod   string `yaml:"klinePeriod"`
	AnalyzeAmount int    `yaml:"analyzeAmount"`
	StopLoss      int    `yaml:"stopLoss"`
	Exchange      struct {
		Binance struct {
			ApiKey    string `yaml:"apiKey"`
			ApiSecret string `yaml:"apiSecret"`
			Spot      struct {
				Url string `yaml:"url"`
			} `yaml:"spot"`
			Futures struct {
				Url string `yaml:"url"`
			} `yaml:"futures"`
			CopyFutures struct {
				Url string `yaml:"url"`
			} `yaml:"copyFutures"`
		} `yaml:"binance"`
	} `yaml:"datahub"`
	Feishu struct {
		AppId     string            `yaml:"appId"`
		AppSecret string            `yaml:"appSecret"`
		Groups    map[string]string `yaml:"groups"`
	} `yaml:"feishu"`
}

func Get() *Config {
	once.Do(func() {
		cfg = &Config{}
	})
	return cfg
}
