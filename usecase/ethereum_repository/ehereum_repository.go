package ehereum_repository

import (
	"context"

	"github.com/ethereum_project/models"
	"gopkg.in/mgo.v2/bson"
)

// Repository represent the blocks's repository contract
type Repository interface {
	Fetch(ctx context.Context) ([]models.Block, error)
	FetchOne(ctx context.Context, query bson.M) (models.Block, error)
	Store(ctx context.Context, docs interface{}) error
}

type Usecase interface {
	GetBlocks(blocks uint64) error
	GetTransactions(hex string) (models.Block, error)
}
