package consul

import (
	"github.com/hashicorp/consul/api"
	"sync"
)

var client *api.Client
var once sync.Once

func Get() *api.Client {
	once.Do(func() {
		consulClient, err := api.NewClient(&api.Config{
			Address: "consul.richfast.trade:8500",
		})
		if err != nil {
			panic(err)
		}
		client = consulClient
	})
	return client
}
