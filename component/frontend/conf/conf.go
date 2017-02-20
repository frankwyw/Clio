package conf

import (
	"Clio/pkg/env"
	"Clio/pkg/utlis"
)

var (
	Backend_host string
	Backend_port string

	Registry_host string
	Registry_port string
)

func init() {
	Backend_host = utils.GetStringEnv(env.BACKEND_HOST, "127.0.0.1")
	Backend_port = utils.GetStringEnv(env.BACKEND_PORT, "8070")

	Registry_host = utils.GetStringEnv(env.REGISTRY_HOST, "127.0.0.1")
	Registry_port = utils.GetStringEnv(env.REGISTRY_PORT, "5000")

}
