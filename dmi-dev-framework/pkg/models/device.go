package models

import (
	"encoding/json"
)

type Device struct {
	ID               string `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	ProtocolName     string `json:"protocol,omitempty"`
	PProtocol        Protocol
	Model            string            `json:"model,omitempty"`
	Twins            []Twin            `json:"twins,omitempty"`
	Datas            []Data            `json:"data,omitempty"`
	PropertyVisitors []PropertyVisitor `json:"propertyVisitors,omitempty"`
}

// Protocol is structure to store protocol in deviceProfile.json in configmap.
type Protocol struct {
	Name                 string          `json:"name,omitempty"`
	Protocol             string          `json:"protocol,omitempty"`
	ProtocolConfigs      json.RawMessage `json:"protocolConfig,omitempty"`
	ProtocolCommonConfig json.RawMessage `json:"protocolCommonConfig,omitempty"`
}

// PropertyVisitor is structure to store propertyVisitor in deviceProfile.json in configmap.
type PropertyVisitor struct {
	Name          string `json:"name,omitempty"`
	PropertyName  string `json:"propertyName,omitempty"`
	ModelName     string `json:"modelName,omitempty"`
	CollectCycle  int64  `json:"collectCycle"`
	ReportCycle   int64  `json:"reportCycle,omitempty"`
	PProperty     Property
	Protocol      string          `json:"protocol,omitempty"`
	VisitorConfig json.RawMessage `json:"visitorConfig"`
}

// Data is data structure for the message that only be subscribed in edge node internal.
type Data struct {
	Properties []DataProperty `json:"dataProperties,omitempty"`
	Topic      string         `json:"dataTopic,omitempty"`
}

// DataProperty is data property.
type DataProperty struct {
	Metadatas    DataMetadata `json:"metadata,omitempty"`
	PropertyName string       `json:"propertyName,omitempty"`
	PVisitor     *PropertyVisitor
}

// DataMetadata data metadata.
type DataMetadata struct {
	Timestamp int64  `json:"timestamp"`
	Type      string `json:"type"`
}

// Metadata is the metadata for data.
type Metadata struct {
	Timestamp string `json:"timestamp,omitempty"`
	Type      string `json:"type,omitempty"`
}

// Twin is the set/get pair to one register.
type Twin struct {
	PropertyName string `json:"propertyName,omitempty"`
	PVisitor     *PropertyVisitor
	Desired      DesiredData  `json:"desired,omitempty"`
	Reported     ReportedData `json:"reported,omitempty"`
}

// DesiredData is the desired data.
type DesiredData struct {
	Value     string   `json:"value,omitempty"`
	Metadatas Metadata `json:"metadata,omitempty"`
}

// ReportedData is the reported data.
type ReportedData struct {
	Value     string   `json:"value,omitempty"`
	Metadatas Metadata `json:"metadata,omitempty"`
}
