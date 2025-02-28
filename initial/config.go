package initial

import (
	"github.com/RichFastTrade/rft_shared/conf"
	"github.com/RichFastTrade/rft_shared/consul"
	kratosconsul "github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"gopkg.in/yaml.v3"
)

func initConfig() error {
	cs, err := kratosconsul.New(consul.Get(), kratosconsul.WithPath("rft/config"))
	if err != nil {
		return err
	}
	rs, err := kratosconsul.New(consul.Get(), kratosconsul.WithPath("rft/rft"))
	if err != nil {
		return err
	}
	c := config.New(
		config.WithSource(cs, rs),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	defer c.Close()
	if err = c.Load(); err != nil {
		return err
	}
	if err = c.Scan(conf.Get()); err != nil {
		return err
	}
	return nil
}
