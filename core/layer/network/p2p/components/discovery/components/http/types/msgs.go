package types

import (
	"bytes"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/base/serialize"
	"github.com/itsfunny/go-cell/component/codec"
	"github.com/itsfunny/go-cell/component/codec/types"
	types2 "github.com/itsfunny/go-cell/framework/rpc/grpc/common/types"
	"net/http"
)

var (
	_ reactor.ICommandSerialize = (*Peer2PeerRequest)(nil)
)

type Peer2PeerRequest struct {
	Envelope types2.Envelope
}

func NewPeer2PeerRequest(envelope types2.Envelope) *Peer2PeerRequest {
	return &Peer2PeerRequest{Envelope: envelope}
}

func (p *Peer2PeerRequest) Send(cdc *codec.CodecComponent, address string) error {
	data := cdc.MustMarshal(p)
	// TODO ,handle response
	_, err := http.Post(address, "application/json;charset=utf-8", bytes.NewBuffer(data))
	return err
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
