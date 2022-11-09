package command

import (
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/common/types"
	types2 "github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/http/types"
	types3 "github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/base/serialize"
	"github.com/itsfunny/go-cell/framework/http/context"
)

func newPeer2PeerCommand(ddd *component.DDDComponent) reactor.ICommand {
	return &reactor.Command{
		ProtocolID: types3.Peer2Peer,
		PreRun:     nil,
		Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
			cellCtx := types.CellContext{}
			cellCtx = cellCtx.FromHttpCtx(ctx)
			ddd.Send(cellCtx, reqData.(*types2.Peer2PeerRequest))
			return nil
		},
		Property: reactor.CommandProperty{
			Async: false,
			RequestDataCreateF: func() reactor.ICommandSerialize {
				return &types2.Peer2PeerRequest{}
			},
			GetInputArchiveFromCtxFunc: func(ctx reactor.IBuzzContext) (serialize.IInputArchive, error) {
				return context.GetByteJSONInputArchiveFromCtx(ctx)
			},
		},
		RunType:     reactor.RunTypeHttpPost,
		Description: "端到端发送",
		MetaData: reactor.MetaData{
			Description: "端到端发送",
			Produces:    nil,
			Tags:        nil,
			Summary:     "",
			Response:    nil,
		},
	}
}
