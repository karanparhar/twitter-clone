package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"

	"github.com/ethereum_project/models"
	ethereum "github.com/ethereum_project/usecase/ethereum_repository"
)

type transactionUsecase struct {
	blocksRepo     ethereum.Repository
	blocks         *types.Block
	contextTimeout time.Duration
	c              *Client
}

func NewtransactionUsecase(e ethereum.Repository, timeout time.Duration, client *Client) ethereum.Usecase {

	return &transactionUsecase{

		blocksRepo:     e,
		contextTimeout: timeout,
		c:              client,
	}

}

func (t *transactionUsecase) AddBlockToDb(input interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), t.contextTimeout)
	return t.blocksRepo.Store(ctx, input)

}

func (t *transactionUsecase) GetAllTransactions() []models.Transaction {
	tr := []models.Transaction{}

	b := t.blocks

	for _, tx := range b.Transactions() {

		temp := models.Transaction{
			Hash: tx.Hash().Hex(),
			To:   tx.To(),
		}
		tr = append(tr, temp)

	}

	return tr

}

func (t *transactionUsecase) GetBlocks(blocks uint64) error {
	result := make(chan *types.Block, blocks)
	ctx, _ := context.WithTimeout(context.Background(), t.contextTimeout)
	go func() {

		err := t.c.Subscribenewhead(ctx, blocks, result)
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	for {

		select {

		case block := <-result:

			t.blocks = block

			m := &models.Block{
				Hash:            block.Hash().Hex(),
				TransactionHash: block.TxHash().Hex(),
				BlockNumber:     block.Number().Uint64(),
				From:            block.ReceivedFrom,
				Transactions:    t.GetAllTransactions(),
			}

			err := t.AddBlockToDb(m)

			if err != nil {
				fmt.Println("Error", err.Error())

			}
		case <-time.After(10 * time.Second):
			return nil

		}

	}

	return nil
}

func (t *transactionUsecase) GetTransactions(hex string) (models.Block, error) {
	ctx, _ := context.WithTimeout(context.Background(), t.contextTimeout)

	q := bson.M{"transactionhash": hex}

	result, err := t.blocksRepo.FetchOne(ctx, q)

	if err != nil {

		return result, err
	}

	return result, nil
}
