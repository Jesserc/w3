package eth_test

import (
	"math/big"
	"testing"

	"github.com/Jesserc/w3"
	"github.com/Jesserc/w3/module/eth"
	"github.com/Jesserc/w3/rpctest"
)

func TestBlockNumber(t *testing.T) {
	rpctest.RunTestCases(t, []rpctest.TestCase[*big.Int]{
		{
			Golden:  "block_number",
			Call:    eth.BlockNumber(),
			WantRet: w3.I("0xc0fe"),
		},
	})
}
