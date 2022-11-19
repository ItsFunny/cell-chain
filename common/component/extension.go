package component

import "github.com/itsfunny/go-cell/base/node/core/extension"

type DDDExtension struct {
	*extension.BaseExtension
	Handlers []DDDHandler
	ddd      *DDDComponent
}

func NewDDDExtension(ddd *DDDComponent, holder DDDHandlerHolder) extension.INodeExtension {
	ret := &DDDExtension{}
	ret.BaseExtension = extension.NewBaseExtension(ret)
	ret.Handlers = holder.Handlers
	ret.ddd = ddd
	return ret
}
func (dd *DDDExtension) OnExtensionInit(ctx extension.INodeContext) error {
	for _, v := range dd.Handlers {
		dd.ddd.RegisterDDDHandler(v)
	}
	return nil
}
