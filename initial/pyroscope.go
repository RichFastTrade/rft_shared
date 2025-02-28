package initial

import (
	"github.com/RichFastTrade/rft_shared/conf"
	"github.com/grafana/pyroscope-go"
)

func initPyroscope() error {
	if _, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: conf.Get().AppName,
		ServerAddress:   conf.Get().Pyroscope.Addr,
	}); err != nil {
		return err
	}
	return nil
}
