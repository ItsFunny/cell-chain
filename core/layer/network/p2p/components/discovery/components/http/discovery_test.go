package http

import (
	"context"
	"fmt"
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/core/layer/common/dispatcher"
	"github.com/itsfunny/go-cell/application"
	"github.com/itsfunny/go-cell/base/node/core/extension"
	"github.com/itsfunny/go-cell/di"
	"github.com/stretchr/testify/suite"
	"strconv"
	"sync"
	"testing"
)

type HttpDiscoveryTestSuit struct {
	suite.Suite
}

func modules() []di.OptionBuilder {
	return []di.OptionBuilder{
		HttpDiscoveryModule,
		component.DIDDDModule,
		component.DIDefaultRoutineModule,
		dispatcher.DIMsgDispatcherModule,
	}
}

func (suite *HttpDiscoveryTestSuit) SetupTest() {
	count := 1
	wg := sync.WaitGroup{}
	wg.Add(count)
	// TODO, 需要修改go-cell#sdk#configuration
	cellHome := "/Users/lvcong/go/src/github.com/itsfunny/cell-chain/testdata/config  "
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
		go func() {
			for {
				select {
				case msg := <-notify.Out():
					data := msg.Data()
					if _, ok := data.(extension.ExtensionLoadedEvent); ok {
						wg.Done()
						return
					}
				}
			}
		}()
	}
	wg.Wait()
	fmt.Println("setup successfully")
}

func TestIBCHttpDiscoverySuite(t *testing.T) {
	suite.Run(t, new(HttpDiscoveryTestSuit))
}
