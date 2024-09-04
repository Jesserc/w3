package eth_test

import (
	"testing"

	"github.com/Jesserc/w3"
	"github.com/Jesserc/w3/module/eth"
	"github.com/Jesserc/w3/rpctest"
)

func TestCode(t *testing.T) {
	rpctest.RunTestCases(t, []rpctest.TestCase[[]byte]{
		{
			Golden:  "get_code",
			Call:    eth.Code(w3.A("0x000000000000000000000000000000000000c0DE"), nil),
			WantRet: w3.B("0xdeadbeef"),
		},
	})
}
