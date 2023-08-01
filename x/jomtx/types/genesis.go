package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TxnList: []Txn{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in txn
	txnIdMap := make(map[uint64]bool)
	txnCount := gs.GetTxnCount()
	for _, elem := range gs.TxnList {
		if _, ok := txnIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for txn")
		}
		if elem.Id >= txnCount {
			return fmt.Errorf("txn id should be lower or equal than the last id")
		}
		txnIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
