package models

// Type Block represents a transactions
type Block struct {
	Hash            string
	TransactionHash string
	To              interface{}
	From            interface{}
	BlockNumber     uint64
	Transactions    []Transaction
}

type Transaction struct {
	Hash string
	To   interface{}
}
