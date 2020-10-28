package std

//---------- Env ---------

// Env defines the state of the blockchain environment this contract is
// running in. This must contain only trusted data - nothing from the Tx itself
// that has not been verfied (like Signer).
//
// Env are json encoded to a byte slice before passing to the wasm contract.
type Env struct {
	Block    BlockInfo    `json:"block"`
	Contract ContractInfo `json:"contract"`
}

type BlockInfo struct {
	// block height this transaction is executed
	Height uint64 `json:"height"`
	// time in seconds since unix epoch - since cosmwasm 0.3
	Time    uint64 `json:"time"`
	ChainID string `json:"chain_id"`
}

type MessageInfo struct {
	// bech32 encoding of sdk.AccAddress executing the contract
	Sender string `json:"sender"`
	// amount of funds send to the contract along with this message
	SentFunds []Coin `json:"sent_funds"`
}

type ContractInfo struct {
	// bech32 encoding of sdk.AccAddress of the contract, to be used when sending messages
	Address string `json:"address"`
}
