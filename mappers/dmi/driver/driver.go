package driver

import "sync"

type DmiDevice struct {
	mutex                sync.RWMutex
	protocolConfig       DmiProtocolConfig
	protocolCommonConfig DmiProtocolCommonConfig
	visitorConfig        DmiVisitorConfig
}

type DmiProtocolConfig struct {
	ProtocolName       string `json:"protocolName"`
	ProtocolConfigData `json:"configData"`
}

type ProtocolConfigData struct {
	// TODO: add your config data according to configmap
}

type DmiProtocolCommonConfig struct {
	CommonCustomizedValues `json:"customizedValues"`
}

type CommonCustomizedValues struct {
	// TODO: add your CommonCustomizedValues according to configmap
}
type DmiVisitorConfig struct {
	ProtocolName      string `json:"protocolName"`
	VisitorConfigData `json:"configData"`
}

type VisitorConfigData struct {
	// TODO: add your Visitor ConfigData according to configmap
}
