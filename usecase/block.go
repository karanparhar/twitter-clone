package usecase

import (
	"context"
	"crypto/rand"
	"math/big"

	log "github.com/sirupsen/logrus"

	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

//url's to connect to ethereum node
const (
	//websocketurl web socket url
	Websocketurl = "wss://ropsten.infura.io/ws"
	//url to fetch blocks
	Trurl = "https://mainnet.infura.io"
)

//client struct  to connect to node
type Client struct {
	*ethclient.Client
}

func client(url string) (*ethclient.Client, error) {

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil

}

func NewClient(url string) (*Client, error) {

	c, err := client(url)

	if err != nil {

		return nil, err
	}

	return &Client{c}, nil

}

func (c *Client) Subscribenewhead(ctx context.Context, blockscount uint64, result chan<- *types.Block) error {
	headers := make(chan *types.Block, blockscount)
	go func() {
		for i := 0; i <= int(blockscount); i++ {
			blockNumber, _ := rand.Int(rand.Reader, big.NewInt(5000000))
			block, err := c.BlockByNumber(ctx, blockNumber)

			if err != nil {
				log.Error(err)
			}

			headers <- block
		}
	}()

	for {
		select {

		case header := <-headers:

			result <- header
		case <-time.After(10 * time.Second):
			return nil

		}
	}

	return nil

}
