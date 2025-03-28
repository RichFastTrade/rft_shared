package consul

import (
	"github.com/hashicorp/consul/api"
	"os"
	"sync"
)

var client *api.Client
var once sync.Once

func Get() *api.Client {
	once.Do(func() {
		addr := os.Getenv("CONSUL_ADDRESS")
		if addr == "" {
			addr = "consul:8500"
		}
		consulClient, err := api.NewClient(&api.Config{
			Address: addr,
		})
		if err != nil {
			panic(err)
		}
		client = consulClient
	})
	return client
}
