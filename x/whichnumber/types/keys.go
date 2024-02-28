package types

const (
	// ModuleName defines the module name
	ModuleName = "whichnumber"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_whichnumber"

	SystemInfoKey = "SystemInfo-value-"
)

const (
	ParamKeyCommitTimeout     = "CommitTimeout"
	ParamKeyRevealTimeout     = "RevealTimeout"
	ParamKeyMaxPlayersPerGame = "MaxPlayersPerGame"
	ParamKeyMinDistanceToWin  = "MinDistanceToWin"
)

const (
	GameNewGas       = 10000
	GameCommitGas    = 5000
	GameRevealGas    = 5000
	LoopIterationGas = 100
	ParamsUpdateGas  = 1000
)

const (
	NoFifoId int64 = -1
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
