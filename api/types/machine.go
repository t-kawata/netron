package types

import "github.com/mudler/netron/pkg/types"

type Machine struct {
	types.Machine
	Connected bool
	OnChain   bool
	Online    bool
}
