package main

import (
	"fmt"

	"github.com/babylonchain/babylon-da-sdk/sdk"
)

func checkBlockFinalized(height uint64, hash string) {
	client, err := sdk.NewClient(sdk.Config{
		ChainType:    0,
		ContractAddr: "bbn1eyfccmjm6732k7wp4p6gdjwhxjwsvje44j0hfx8nkgrm8fs7vqfsa3n3gc",
	})

	if err != nil {
		fmt.Printf("error creating client: %v\n", err)
		return
	}

	isFinalized, err := client.QueryIsBlockBabylonFinalized(&sdk.L2Block{
		BlockHeight:    height,
		BlockHash:      hash,
		BlockTimestamp: uint64(1718332131),
	})
	if err != nil {
		fmt.Printf("error checking block %d: %v\n", height, err)
	} else {
		fmt.Printf("is block %d finalized?: %t\n", height, isFinalized)
	}
}

func main() {
	blockHash := "0x3aa074144a25d3ed71c7353a20c579650e0c56a993444c6156d44bb90b932f0d"
	// TODO: now you will see `error checking block 2: not enough voting power`
	// in the future, we will deploy a better stub contract
	checkBlockFinalized(uint64(2), blockHash)
}
