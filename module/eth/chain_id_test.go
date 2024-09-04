package eth_test

import (
	"testing"

	"github.com/Jesserc/w3/module/eth"
	"github.com/Jesserc/w3/rpctest"
)

func TestChainID(t *testing.T) {
	rpctest.RunTestCases(t, []rpctest.TestCase[uint64]{
		{
			Golden:  "chain_id",
			Call:    eth.ChainID(),
			WantRet: 1,
		},
	})
}
