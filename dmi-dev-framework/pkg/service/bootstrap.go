package service

import (
	"errors"
	"os"

	"k8s.io/klog/v2"

	"github.com/kubeedge/mappers-go/dmi-dev-framework/pkg/config"
	"github.com/kubeedge/mappers-go/dmi-dev-framework/pkg/device"
	"github.com/kubeedge/mappers-go/dmi-dev-framework/pkg/grpcclient"
	"github.com/kubeedge/mappers-go/dmi-dev-framework/pkg/grpcserver"
)

func Bootstrap(protocolName string, deviceInterface interface{}) {
	var err error
	var c config.Config

	klog.InitFlags(nil)
	defer klog.Flush()

	if err = c.Parse(); err != nil {
		klog.Fatal(err)
		os.Exit(1)
	}
	klog.Infof("config: %+v", c)

	grpcclient.Init(&c)

	// start grpc server
	grpcServer := grpcserver.NewServer(
		grpcserver.Config{
			SockPath: c.GrpcServer.SocketPath,
			Protocol: protocolName,
		},
	)

	panel := device.NewDevPanel(deviceInterface)
	err = panel.DevInit(&c)
	if err != nil && !errors.Is(err, parse.ErrEmptyData) {
		klog.Fatal(err)
	}
	klog.Infoln("devInit finished")

	// register to edgecore
	klog.Infoln("======dev init mode is not register, will register to edgecore")
	// TODO health check
	if _, _, err = grpcclient.RegisterMapper(&c, false); err != nil {
		klog.Fatal(err)
	}
	klog.Infoln("registerMapper finished")

	panel.DevStart()

	if err = grpcServer.Start(); err != nil {
		klog.Fatal(err)
	}
	klog.Infoln("grpc server start finished")
}
