package types

const (
	// ModuleName defines the module name
	ModuleName = "jomtx"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_jomtx"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	TxnKey      = "Txn/value/"
	TxnCountKey = "Txn/count/"
)

const (
	TxnCreatedEventType    = "new-txn-created" // Indicates what event type to listen to
	TxnCreatedEventCreator = "creator"         // Subsidiary information
	TxnCreatedEventTxnId   = "txn-id"          // What txn is relevant
)
