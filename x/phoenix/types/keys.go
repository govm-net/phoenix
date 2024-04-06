package types

const (
	// ModuleName defines the module name
	ModuleName = "phoenix"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_phoenix"
)

var (
	ParamsKey = []byte("p_phoenix")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
