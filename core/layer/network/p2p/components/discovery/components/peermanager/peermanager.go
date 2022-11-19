package peermanager

import (
	"context"
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/common/enums"
	sdk "github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/component/codec"
	"sync"
)

var (
	_ types.IPeerManager = (*DefaultPeerManager)(nil)
)

type DefaultPeerManager struct {
	*component.BaseComponent
	mutex   sync.RWMutex
	members map[types.PeerId]types.IPeerNode

	self types.IPeerNode

	forwardMessageChan chan interface{}

	seal bool
}

func NewDefaultPeerManager(ctx context.Context, ddd *component.DDDComponent, cdc *codec.CodecComponent) types.IPeerManager {
	ret := &DefaultPeerManager{
		mutex:   sync.RWMutex{},
		members: make(map[types.PeerId]types.IPeerNode),
	}
	ret.BaseComponent = component.NewBaseComponent(ctx, enums.PeerManagerModule, ret, ddd, cdc)
	// TODO, configurable
	ret.forwardMessageChan = make(chan interface{}, 100)
	return ret
}
func (d *DefaultPeerManager) Seal() {
	d.seal = true
}
func (d *DefaultPeerManager) Sealed() bool {
	return d.seal
}

func (d *DefaultPeerManager) SetupSelfNode(n types.IPeerNode) {
	if d.Sealed() {
		panic("asd")
	}
	d.self = n
}

func (d *DefaultPeerManager) Have(node types.PeerId) bool {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.members[node] != nil
}

func (d *DefaultPeerManager) GetByPeerId(id types.PeerId) types.IPeerNode {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.members[id]
}

func (d *DefaultPeerManager) Register(wrapper *types.PeerWrapper) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	node := &DefaultPeerNode{}
	node.From(wrapper)
	d.members[node.PeerId()] = node

	d.Send(sdk.EmptyCellContext(context.Background()), func() sdk.CellRequest {
		return wrapper
	})
}

func (d *DefaultPeerManager) ForwardMessage() chan<- interface{} {
	return d.forwardMessageChan
}

func (d *DefaultPeerManager) GetMembership() map[types.PeerId]types.IPeerNode {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	return d.members
}

func (d *DefaultPeerManager) OnStart(ctx *services.StartCTX) error {
	go d.routine()
	return nil
}
func (d *DefaultPeerManager) routine() {
	for {
		select {
		case msg := <-d.forwardMessageChan:
			d.handleMsg(msg)
		}
	}
}
func (d *DefaultPeerManager) handleMsg(msg interface{}) {
	d.Logger.Info("receive msg", "info", msg)
}

// TODO, copy?
func (d *DefaultPeerManager) GetSelfNode() types.IPeerNode {
	return d.self
}
