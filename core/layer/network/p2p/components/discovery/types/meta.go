package types

import (
	"errors"
	"fmt"
	"github.com/itsfunny/cell-chain/common/enums"
	"github.com/itsfunny/go-cell/sdk/common"
)

// TODO, add validbasic
type PeerMetaData struct {
	Domain   string // optional
	Ip       string // optional
	Protocol string // http,https
	Port     uint
}

func NewPeerMetaData(domain string, ip string, protocol string, port uint) *PeerMetaData {
	return &PeerMetaData{Domain: domain, Ip: ip, Protocol: protocol, Port: port}
}

func (m PeerMetaData) ValidBasic() error {
	if len(m.Domain) == 0 {
		if len(m.Ip) == 0 && m.Port == 0 {
			return common.NewWrappedError(enums.ErrInvalidMetaData, errors.New("invalid meta data"))
		}
	}
	return nil
}

func (m PeerMetaData) GetOutPutAddress() string {
	if len(m.Domain) > 0 {
		return m.Domain
	}
	return fmt.Sprintf("%s://%s:%d", m.Protocol, m.Ip, m.Port)
}
