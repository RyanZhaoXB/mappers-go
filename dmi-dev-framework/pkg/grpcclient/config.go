package grpcclient

import (
	"github.com/kubeedge/mappers-go/dmi-dev-framework/pkg/config"
)

var cfg *config.Config

func Init(c *config.Config) {
	cfg = &config.Config{}
	cfg = c
}
