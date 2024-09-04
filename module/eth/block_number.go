package eth

import (
	"math/big"

	"github.com/Jesserc/w3/internal/module"
	"github.com/Jesserc/w3/w3types"
)

// BlockNumber requests the number of the most recent block.
func BlockNumber() w3types.RPCCallerFactory[*big.Int] {
	return module.NewFactory(
		"eth_blockNumber",
		nil,
		module.WithRetWrapper(module.HexBigRetWrapper),
	)
}
