package peermanager

import (
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/base/node/core/extension"
	"github.com/itsfunny/go-cell/di"
	"go.uber.org/fx"
)

var DIDefaultPeerManagerModule = fx.Options(
	fx.Provide(NewDefaultPeerManager),
	di.RegisterExtension(NewPeerManagerExtension),
)

type PeerManagerExtension struct {
	*extension.BaseExtension
	p types.IPeerManager
}

func NewPeerManagerExtension(p types.IPeerManager) *PeerManagerExtension {
	ret := &PeerManagerExtension{}
	ret.BaseExtension = extension.NewBaseExtension(ret)
	ret.p = p
	return ret
}
func (p *PeerManagerExtension) OnExtensionStart(ctx extension.INodeContext) error {
	return p.p.BStart(services.SyncStartOpt)
}
