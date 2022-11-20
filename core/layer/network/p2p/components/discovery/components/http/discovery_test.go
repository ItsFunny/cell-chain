package http

import (
	"context"
	"fmt"
	"github.com/itsfunny/cell-chain/common/component"
	types2 "github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/cell-chain/core/layer/common/dispatcher"
	types3 "github.com/itsfunny/cell-chain/core/layer/common/types"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/application"
	"github.com/itsfunny/go-cell/base/common/utils"
	"github.com/itsfunny/go-cell/base/core/eventbus"
	"github.com/itsfunny/go-cell/base/node/core/extension"
	"github.com/itsfunny/go-cell/component/codec"
	"github.com/itsfunny/go-cell/di"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
	"reflect"
	"strconv"
	"sync"
	"testing"
)

type HttpDiscoveryTestSuit struct {
	suite.Suite
	Count          int
	Contexts       map[int]extension.INodeContext
	TestExtensions map[int]*TestExtension
}

func modules() []di.OptionBuilder {
	return []di.OptionBuilder{
		HttpDiscoveryModule,
		component.DIDDDModule,
		component.DIDefaultRoutineModule,
		dispatcher.DIMsgDispatcherModule,
		testExtension,
	}
}

var testExtension di.OptionBuilder = func() fx.Option {
	return di.RegisterExtension(NewTestExtension)
}

type TestExtension struct {
	*extension.BaseExtension

	discovery types.DiscoveryComponent
	Cdc       *codec.CodecComponent
	Ddd       *component.DDDComponent
	Bus       eventbus.ICommonEventBus
}

func NewTestExtension(discovery types.DiscoveryComponent,
	cdc *codec.CodecComponent,
	ddd *component.DDDComponent,
	bus eventbus.ICommonEventBus) extension.INodeExtension {
	ret := &TestExtension{}
	ret.BaseExtension = extension.NewBaseExtension(ret)
	ret.discovery = discovery
	ret.Cdc = cdc
	ret.Ddd = ddd
	ret.Bus = bus
	return ret
}

func (suite *HttpDiscoveryTestSuit) SetupTest() {
	count := suite.Count
	wg := sync.WaitGroup{}
	wg.Add(count)
	// TODO, 需要修改go-cell#sdk#configuration
	cellHome := "/Users/lvcong/go/src/github.com/itsfunny/cell-chain/testdata/config"
	configType := "test%d"
	for i := 0; i < count; i++ {
		app := application.New(context.Background(), modules()...)
		go func(index int) {
			testType := fmt.Sprintf(configType, index)
			args := []string{fmt.Sprintf("-cellHome=%s ", cellHome), fmt.Sprintf("-configType=%s", testType)}
			app.Run(args)
		}(i)
		bus := app.GetApplicationBus()
		notify, err := bus.SubscribeApplicationEvent(context.Background(), strconv.Itoa(i))
		if nil != err {
			panic(err)
		}
		go func(index int) {
			for {
				select {
				case msg := <-notify.Out():
					data := msg.Data()
					if v, ok := data.(extension.ExtensionLoadedEvent); ok {
						//v.Context.SwitchTo(&reflect.TypeOf())
						suite.Contexts[index] = v.Context
						suite.TestExtensions[index] = v.Context.SwitchTo(reflect.TypeOf(&TestExtension{})).(*TestExtension)
						wg.Done()
						return
					}
				}
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("setup successfully")
}

func TestHttpDiscoverySuite(t *testing.T) {
	h := HttpDiscoveryTestSuit{
		Count:          2,
		Contexts:       make(map[int]extension.INodeContext),
		TestExtensions: make(map[int]*TestExtension),
	}
	suite.Run(t, &h)
}

func (suite *HttpDiscoveryTestSuit) TestProbe() {
	testE := suite.TestExtensions[0]
	cdc := testE.Cdc

	test2 := suite.TestExtensions[1]
	peer2 := test2.discovery.GetPeerManager()
	out := peer2.GetSelfNode().MetaData().GetOutPutAddress()

	seq := utils.GenerateSequenceId()

	sub := types.SubscribePeerManagerCommonEvent(context.Background(), "test", testE.Bus, 10)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case msg := <-sub.Out():
				switch v := msg.Data().(type) {
				case *types.PeerWrapper:
					fmt.Println(v)
					close(done)
				}
			}
		}
	}()
	newMemEnv := types.CreateNewMemberEnvelopeRequest(cdc.GetCodec(), seq, types.NewNewMemberRequest(out))
	testE.Ddd.Send(types2.EmptyCellContext(context.Background()).WithSeq(seq), types3.NewPeer2PeerRequest(newMemEnv))
	<-done
}
