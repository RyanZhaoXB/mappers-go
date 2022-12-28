package main

import (
	"github.com/kubeedge/mappers-go/dmi-dev-framework/pkg/service"
	"github.com/kubeedge/mappers-go/mappers/dmi/driver"
)

func main() {
	vd := &driver.DmiDevice{}
	service.Bootstrap("dmi", vd)
}
