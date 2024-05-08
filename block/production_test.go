package block_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dymensionxyz/dymint/mempool"
	mempoolv1 "github.com/dymensionxyz/dymint/mempool/v1"
	"github.com/dymensionxyz/dymint/node/events"
	"github.com/dymensionxyz/dymint/types"
	uevent "github.com/dymensionxyz/dymint/utils/event"
	tmcfg "github.com/tendermint/tendermint/config"

	"github.com/dymensionxyz/dymint/testutil"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/libs/pubsub"
	"github.com/tendermint/tendermint/proxy"
)

func TestCreateEmptyBlocksEnableDisable(t *testing.T) {
	const blockTime = 200 * time.Millisecond
	const EmptyBlocksMaxTime = blockTime * 10
	const runTime = 10 * time.Second

	assert := assert.New(t)
	require := require.New(t)
	app := testutil.GetAppMock()
	// Create proxy app
	clientCreator := proxy.NewLocalClientCreator(app)
	proxyApp := proxy.NewAppConns(clientCreator)
	err := proxyApp.Start()
	require.NoError(err)

	// Init manager with empty blocks feature disabled
	managerConfigCreatesEmptyBlocks := testutil.GetManagerConfig()
	managerConfigCreatesEmptyBlocks.BlockTime = blockTime
	managerConfigCreatesEmptyBlocks.EmptyBlocksMaxTime = 0 * time.Second
	managerWithEmptyBlocks, err := testutil.GetManager(managerConfigCreatesEmptyBlocks, nil, nil, 1, 1, 0, proxyApp, nil)
	require.NoError(err)

	// Init manager with empty blocks feature enabled
	managerConfig := testutil.GetManagerConfig()
	managerConfig.BlockTime = blockTime
	managerConfig.EmptyBlocksMaxTime = EmptyBlocksMaxTime
	managerConfig.PriorityMaxIdleTime = EmptyBlocksMaxTime
	manager, err := testutil.GetManager(managerConfig, nil, nil, 1, 1, 0, proxyApp, nil)
	require.NoError(err)

	// Check initial height
	initialHeight := uint64(0)
	require.Equal(initialHeight, manager.Store.Height())

	mCtx, cancel := context.WithTimeout(context.Background(), runTime)
	defer cancel()
	go manager.ProduceBlockLoop(mCtx)
	go managerWithEmptyBlocks.ProduceBlockLoop(mCtx)

	buf1 := make(chan bool, 100) //dummy to avoid unhealthy event
	buf2 := make(chan bool, 100) //dummy to avoid unhealthy event
	go manager.AccumulatedDataLoop(mCtx, buf1)
	go managerWithEmptyBlocks.AccumulatedDataLoop(mCtx, buf2)
	<-mCtx.Done()

	require.Greater(manager.Store.Height(), initialHeight)
	require.Greater(managerWithEmptyBlocks.Store.Height(), initialHeight)
	assert.Greater(managerWithEmptyBlocks.Store.Height(), manager.Store.Height())

	// Check that blocks are created with empty blocks feature disabled
	assert.LessOrEqual(manager.Store.Height(), uint64(runTime/EmptyBlocksMaxTime))
	assert.LessOrEqual(managerWithEmptyBlocks.Store.Height(), uint64(runTime/blockTime))

	for i := uint64(2); i < managerWithEmptyBlocks.Store.Height(); i++ {
		prevBlock, err := managerWithEmptyBlocks.Store.LoadBlock(i - 1)
		assert.NoError(err)

		block, err := managerWithEmptyBlocks.Store.LoadBlock(i)
		assert.NoError(err)
		assert.NotZero(block.Header.Time)

		diff := time.Unix(0, int64(block.Header.Time)).Sub(time.Unix(0, int64(prevBlock.Header.Time)))
		assert.Greater(diff, blockTime-blockTime/10)
		assert.Less(diff, blockTime+blockTime/10)
	}

	for i := uint64(2); i < manager.Store.Height(); i++ {
		prevBlock, err := manager.Store.LoadBlock(i - 1)
		assert.NoError(err)

		block, err := manager.Store.LoadBlock(i)
		assert.NoError(err)
		assert.NotZero(block.Header.Time)

		diff := time.Unix(0, int64(block.Header.Time)).Sub(time.Unix(0, int64(prevBlock.Header.Time)))
		assert.Greater(diff, manager.Conf.EmptyBlocksMaxTime)
	}
}

func TestCreateEmptyBlocksNew(t *testing.T) {
	// TODO(https://github.com/dymensionxyz/dymint/issues/352)
	t.Skip("FIXME: fails to submit tx to test the empty blocks feature")
	assert := assert.New(t)
	require := require.New(t)
	app := testutil.GetAppMock()
	// Create proxy app
	clientCreator := proxy.NewLocalClientCreator(app)
	proxyApp := proxy.NewAppConns(clientCreator)
	err := proxyApp.Start()
	require.NoError(err)
	// Init manager
	managerConfig := testutil.GetManagerConfig()
	managerConfig.BlockTime = 200 * time.Millisecond
	managerConfig.EmptyBlocksMaxTime = 1 * time.Second
	manager, err := testutil.GetManager(managerConfig, nil, nil, 1, 1, 0, proxyApp, nil)
	require.NoError(err)

	abciClient, err := clientCreator.NewABCIClient()
	require.NoError(err)
	require.NotNil(abciClient)
	require.NoError(abciClient.Start())

	defer func() {
		// Capture the error returned by abciClient.Stop and handle it.
		err := abciClient.Stop()
		if err != nil {
			t.Logf("Error stopping ABCI client: %v", err)
		}
	}()

	mempoolCfg := tmcfg.DefaultMempoolConfig()
	mempoolCfg.KeepInvalidTxsInCache = false
	mempoolCfg.Recheck = false

	mpool := mempoolv1.NewTxMempool(log.TestingLogger(), tmcfg.DefaultMempoolConfig(), proxy.NewAppConnMempool(abciClient), 0)

	// Check initial height
	expectedHeight := uint64(0)
	assert.Equal(expectedHeight, manager.Store.Height())

	mCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go manager.ProduceBlockLoop(mCtx)

	<-time.Tick(1 * time.Second)
	err = mpool.CheckTx([]byte{1, 2, 3, 4}, nil, mempool.TxInfo{})
	require.NoError(err)

	<-mCtx.Done()
	foundTx := false
	assert.LessOrEqual(manager.Store.Height(), uint64(10))
	for i := uint64(2); i < manager.Store.Height(); i++ {
		prevBlock, err := manager.Store.LoadBlock(i - 1)
		assert.NoError(err)

		block, err := manager.Store.LoadBlock(i)
		assert.NoError(err)
		assert.NotZero(block.Header.Time)

		diff := time.Unix(0, int64(block.Header.Time)).Sub(time.Unix(0, int64(prevBlock.Header.Time)))
		txsCount := len(block.Data.Txs)
		if txsCount == 0 {
			assert.Greater(diff, manager.Conf.EmptyBlocksMaxTime)
			assert.Less(diff, manager.Conf.EmptyBlocksMaxTime+1*time.Second)
		} else {
			foundTx = true
			assert.Less(diff, manager.Conf.BlockTime+100*time.Millisecond)
		}

		fmt.Println("time diff:", diff, "tx len", 0)
	}
	assert.True(foundTx)
}

func TestInvalidBatch(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	manager, err := testutil.GetManager(testutil.GetManagerConfig(), nil, nil, 1, 1, 0, nil, nil)
	require.NoError(err)

	batchSize := uint64(5)
	syncTarget := uint64(10)

	// Create cases
	cases := []struct {
		startHeight uint64
		endHeight   uint64
		shouldError bool
	}{
		{startHeight: syncTarget + 1, endHeight: syncTarget + batchSize, shouldError: false},
		// batch with endHight < startHeight
		{startHeight: syncTarget + 1, endHeight: syncTarget, shouldError: true},
		// batch with startHeight != previousEndHeight + 1
		{startHeight: syncTarget, endHeight: syncTarget + batchSize + batchSize, shouldError: true},
	}
	for _, c := range cases {
		batch := &types.Batch{
			StartHeight: c.startHeight,
			EndHeight:   c.endHeight,
		}

		manager.UpdateSyncParams(syncTarget)
		err := manager.ValidateBatch(batch)
		if c.shouldError {
			assert.Error(err)
		} else {
			assert.NoError(err)
		}
	}
}

// TestStopBlockProduction tests the block production stops when submitter is full
// and resumes when submitter is ready to accept more batches
func TestStopBlockProduction(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	managerConfig := testutil.GetManagerConfig()
	managerConfig.BlockBatchMaxSizeBytes = 1000 // small batch size to fill up quickly
	manager, err := testutil.GetManager(managerConfig, nil, nil, 1, 1, 0, nil, nil)
	require.NoError(err)

	// validate initial accumulated is zero
	require.Equal(manager.AccumulatedBatchSize.Load(), uint64(0))
	assert.Equal(manager.Store.Height(), uint64(0))

	// subscribe to health status event
	eventReceivedCh := make(chan error)
	cb := func(event pubsub.Message) {
		eventReceivedCh <- event.Data().(*events.DataHealthStatus).Error
	}
	go uevent.MustSubscribe(context.Background(), manager.Pubsub, "HealthStatusHandler", events.QueryHealthStatus, cb, log.TestingLogger())

	var wg sync.WaitGroup
	wg.Add(2) // Add 2 because we have 2 goroutines

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		manager.ProduceBlockLoop(ctx)
		wg.Done() // Decrease counter when this goroutine finishes
	}()

	toSubmit := make(chan bool)
	go func() {
		manager.AccumulatedDataLoop(ctx, toSubmit)
		wg.Done() // Decrease counter when this goroutine finishes
	}()

	// validate block production works
	time.Sleep(400 * time.Millisecond)
	assert.Greater(manager.Store.Height(), uint64(0))
	assert.Greater(manager.AccumulatedBatchSize.Load(), uint64(0))

	// we don't read from the submit channel, so we assume it get full
	// we expect the block production to stop and unhealthy event to be emitted
	select {
	case <-ctx.Done():
		t.Error("expected unhealthy event")
	case err := <-eventReceivedCh:
		assert.Error(err)
	}

	stoppedHeight := manager.Store.Height()

	// make sure block production is stopped
	time.Sleep(400 * time.Millisecond)
	assert.Equal(stoppedHeight, manager.Store.Height())

	// consume the signal
	<-toSubmit

	// check for health status event and block production to continue
	select {
	case <-ctx.Done():
		t.Error("expected health event")
	case err := <-eventReceivedCh:
		assert.NoError(err)
	}

	// make sure block production is resumed
	time.Sleep(400 * time.Millisecond)
	assert.Greater(manager.Store.Height(), stoppedHeight)
}
