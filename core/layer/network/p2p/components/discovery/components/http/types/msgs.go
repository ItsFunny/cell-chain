package types

import (
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/base/serialize"
	"github.com/itsfunny/go-cell/component/codec/types"
)

var (
	_ reactor.ICommandSerialize = (*Peer2PeerRequest)(nil)
)

type Peer2PeerRequest struct {
}

func (p *Peer2PeerRequest) Read(archive serialize.IInputArchive, cdc types.Codec) error {
	data, err := archive.ReadByte()
	if nil != err {
		return err
	}

	return cdc.Unmarshal(data, &p)
}

func (p *Peer2PeerRequest) ToBytes(cdc types.Codec) ([]byte, error) {
	return cdc.Marshal(p)
}

func (p *Peer2PeerRequest) ValidateBasic(ctx reactor.IBuzzContext) error {
	return nil
}
