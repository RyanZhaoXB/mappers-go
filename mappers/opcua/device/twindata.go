/*
Copyright 2021 The KubeEdge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package device

import (
	"strings"

	"github.com/kubeedge/mappers-go/pkg/common"
	"github.com/kubeedge/mappers-go/pkg/driver/opcua"
	"github.com/kubeedge/mappers-go/pkg/global"

	"k8s.io/klog/v2"
)

// TwinData is the timer structure for getting twin/data.
type TwinData struct {
	Client *opcua.OPCUAClient
	Name   string
	Type   string
	NodeID string
	Result string
	Topic  string
}

// Run timer function.
func (td *TwinData) Run() {
	var err error
	td.Result, err = td.Client.Get(td.NodeID)
	if err != nil {
		klog.Errorf("Get register failed: %v", err)
		return
	}
	// construct payload
	var payload []byte
	if strings.Contains(td.Topic, "$hw") {
		if payload, err = common.CreateMessageTwinUpdate(td.Name, td.Type, td.Result); err != nil {
			klog.Errorf("Create message twin update failed: %v", err)
			return
		}
	} else {
		if payload, err = common.CreateMessageData(td.Name, td.Type, td.Result); err != nil {
			klog.Errorf("Create message data failed: %v", err)
			return
		}
	}
	if err = global.MqttClient.Publish(td.Topic, payload); err != nil {
		klog.Errorf("Publish topic %v failed, err: %v", td.Topic, err)
	}

	klog.V(2).Infof("Update value: %s, topic: %s", td.Result, td.Topic)
}
