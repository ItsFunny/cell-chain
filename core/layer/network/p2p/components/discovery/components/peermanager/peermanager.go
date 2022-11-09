package peermanager

import (
	"context"
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/common/enums"
	sdk "github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"sync"
)

var (
	_ types.IPeerManager = (*DefaultPeerManager)(nil)
)

type DefaultPeerManager struct {
	*component.BaseComponent
	mutex   sync.RWMutex
	members map[types.PeerId]types.IPeerNode
}

func NewDefaultPeerManager(ddd *component.DDDComponent) *DefaultPeerManager {
	ret := &DefaultPeerManager{
		mutex:   sync.RWMutex{},
		members: make(map[types.PeerId]types.IPeerNode),
	}
	ret.BaseComponent = component.NewBaseComponent(enums.PeerManagerModule, ret, ddd)
	return ret
}

func (d *DefaultPeerManager) Register(wrapper types.PeerWrapper) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	node := &DefaultPeerNode{}
	node.From(wrapper)
	d.members[node.PeerId()] = node

	d.Send(sdk.EmptyCellContext(context.Background()), func() sdk.CellRequest {
		return wrapper
	})
}

func (d *DefaultPeerManager) GetMembership() map[types.PeerId]types.IPeerNode {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	return d.members
}
