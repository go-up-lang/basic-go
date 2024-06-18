package main

import (
	"fmt"

	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/types"
)

func main() {
	// Create a new Tendermint client
	c, err := client.New("tcp://localhost:26657", "/websocket")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Query the latest block height
	status, err := c.Status()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Latest block height: %d\n", status.SyncInfo.LatestBlockHeight)

	// Generate a new key pair for transactions
	privKey := crypto.GenPrivKeyEd25519()
	// pubKey := privKey.PubKey()

	// Example of sending a transaction
	tx := types.NewTransaction(
		status.ValidatorInfo.Address,
		status.ValidatorInfo.PubKey,
		1, // Amount of ATOM to transfer
		0, // Gas limit
		"RecipientAddress",
		"",
	)

	// Sign the transaction
	signature, err := privKey.Sign(tx.SignBytes(""))
	if err != nil {
		fmt.Println(err)
		return
	}
	tx.Signature = signature

	// Broadcast the transaction
	res, err := c.BroadcastTxSync(tx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Transaction broadcasted. Result: %v\n", res)
}
