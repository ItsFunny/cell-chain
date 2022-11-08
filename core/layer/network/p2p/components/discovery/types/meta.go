package types

import (
	"errors"
	"fmt"
	"github.com/itsfunny/cell-chain/common/enums"
	"github.com/itsfunny/go-cell/sdk/common"
)

type PeerMetaData struct {
	Domain string // optional
	Ip     string // optional
	Port   uint
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
	return fmt.Sprintf("%s:%d", m.Ip, m.Port)
}
