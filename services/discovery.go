package services

import (
	consulapi "github.com/hashicorp/consul/api"
)

func Discovery() map[string]*consulapi.AgentService {
	config := consulapi.DefaultConfig()
	consul, error := consulapi.NewClient(config)

	if error != nil {
		panic(error)
	}

	services, error := consul.Agent().Services()

	if error != nil {
		panic(error)
	}

	return services
}
