package service

import (
	"github.com/kubeedge/mappers-go/dmi-dev-framework/pkg/config"
)

type MapperService struct {
}

func (ms *MapperService) InitMapperService(protocolName string, c config.Config, deviceInterface interface{}) {

}

// // waitExit create a goroutine to monitor exit signal
// func (ms *MapperService) waitExit() {
// 	go func() {
// 		<-ms.quit
// 		err := ms.driver.StopDevice()
// 		if err != nil {
// 			klog.Errorf("Service has stopped but failed to stop device:%v", err)
// 			os.Exit(1)
// 		}
// 		klog.V(1).Info("Exit mapper safely")
// 		os.Exit(1)
// 	}()
// }

// // initDeviceMutex init the mutex of device
// func (ms *MapperService) initDeviceMutex() {
// 	for i := range ms.deviceInstances {
// 		ms.deviceMutex[i] = new(common.Lock)
// 		ms.deviceMutex[i].DeviceLock = new(sync.Mutex)
// 	}
// }
