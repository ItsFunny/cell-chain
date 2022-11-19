package config

import (
	"encoding/json"
	"github.com/itsfunny/go-cell/component/codec/types"
)

var (
	_ types.Unmarshaler = (*HttpDiscoveryConfiguration)(nil)
	_ types.Marshaller  = (*HttpDiscoveryConfiguration)(nil)
)

type HttpDiscoveryConfiguration struct {
	OutputAddress string `json:"outputAddress"`
	PeerId        string `json:"peerId"`
}

func (h *HttpDiscoveryConfiguration) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &h)
}

func (h *HttpDiscoveryConfiguration) Marshal() ([]byte, error) {
	return json.Marshal(h)
}
