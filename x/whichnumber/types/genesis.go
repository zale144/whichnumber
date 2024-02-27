package types

import (
	"fmt"
	"strconv"
)

// DefaultId is the default capability global index
const DefaultId int64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		SystemInfo:     DefaultSystemInfo(),
		StoredGameList: []Game{},
		Params:         DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in storedGame
	storedGameIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredGameList {
		index := string(StoredGameKey(strconv.FormatInt(elem.Id, 10)))
		if _, ok := storedGameIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedGame")
		}
		storedGameIndexMap[index] = struct{}{}
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
