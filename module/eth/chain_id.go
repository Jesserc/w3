package eth

import (
	"github.com/Jesserc/w3/internal/module"
	"github.com/Jesserc/w3/w3types"
)

// ChainID requests the chains ID.
func ChainID() w3types.RPCCallerFactory[uint64] {
	return module.NewFactory(
		"eth_chainId",
		nil,
		module.WithRetWrapper(module.HexUint64RetWrapper),
	)
}
