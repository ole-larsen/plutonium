package blockchain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ole-larsen/plutonium/contracts/auction"
	"github.com/ole-larsen/plutonium/contracts/collection"
	"github.com/ole-larsen/plutonium/contracts/marketplace"
	"github.com/ole-larsen/plutonium/contracts/nft"
	"github.com/ole-larsen/plutonium/internal/log"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/ole-larsen/plutonium/models"
)

const GATEWAY = "https://nftstorage.link/ipfs/"

type Marketplace struct {
	Fee              *big.Int
	CollectionsCount *big.Int
	Collections      map[string]models.MarketplaceCollection
	Contract         *marketplace.Marketplace
	Name             string
	ABI              string
	Owner            common.Address
	Address          common.Address
}

type NFT struct {
	Contract *nft.NFT
	Address  string
	ABI      string
}

type Collection struct {
	Contract *collection.NFTCollection
	Name     string
	ABI      string
	Address  common.Address
}

type Auction struct {
	Contract *auction.NFTAuction
	Name     string
	ABI      string
	Address  common.Address
}

type Collections map[string]Collection
type Auctions []Auction

type Market struct {
	Collections
	NFT
	Auctions
	Marketplace
}

type Web3Dialer struct {
	Web3Client *ethclient.Client
	Contracts  *repository.ContractsRepository

	Logger *log.Logger
	Market
}

func NewWeb3Dialer(
	logger *log.Logger,
	address string,
	contracts *repository.ContractsRepository,
) (*Web3Dialer, error) {
	client, err := ethclient.Dial(address)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}

	return &Web3Dialer{
		Web3Client: client,
		Contracts:  contracts,
		Logger:     logger,
		Market:     Market{},
	}, nil
}

func (w *Web3Dialer) Load(ctx context.Context) error {
	if err := w.LoadMarketContractFromDB(ctx); err != nil {
		return err
	}

	if err := w.LoadMarketCollectionContractsFromDB(ctx); err != nil {
		return err
	}

	if err := w.LoadMarketAuctionContractsFromDB(ctx); err != nil {
		return err
	}
	// After loading market contract, start load market information from blockchain
	if err := w.LoadMarketplaceInfoFromWeb3(); err != nil {
		return err
	}

	return nil
}

func (w *Web3Dialer) LoadMarketContractFromDB(ctx context.Context) error {
	contract, err := w.Contracts.GetOne(ctx, "marketplace")
	if err != nil {
		return err
	}

	address := common.HexToAddress(contract.Address)

	instance, err := marketplace.NewMarketplace(address, w.Web3Client)

	if err != nil {
		return err
	}

	w.Market.Marketplace = Marketplace{
		Contract: instance,
		Address:  address,
		ABI:      contract.Abi,
	}

	return nil
}

func (w *Web3Dialer) LoadMarketCollectionContractsFromDB(ctx context.Context) error {
	w.Market.Collections = make(map[string]Collection)

	collections, err := w.Contracts.GetCollectionsContracts(ctx)
	if err != nil {
		return err
	}

	for _, _collection := range collections {
		collectionID := new(big.Int).SetInt64(_collection.ID)

		instance, err := collection.NewNFTCollection(common.HexToAddress(_collection.Address), w.Web3Client)

		if err != nil {
			return err
		}

		w.Market.Collections[collectionID.String()] = Collection{
			Name:     _collection.Name,
			Address:  common.HexToAddress(_collection.Address),
			ABI:      _collection.Abi,
			Contract: instance,
		}
	}

	return nil
}

func (w *Web3Dialer) LoadMarketAuctionContractsFromDB(ctx context.Context) error {
	w.Market.Auctions = make([]Auction, 0)

	auctions, err := w.Contracts.GetAuctions(ctx)
	if err != nil {
		return err
	}

	for _, _auction := range auctions {
		instance, err := auction.NewNFTAuction(common.HexToAddress(_auction.Address), w.Web3Client)
		if err != nil {
			return err
		}

		w.Market.Auctions = append(w.Market.Auctions, Auction{
			Name:     _auction.Name,
			Address:  common.HexToAddress(_auction.Address),
			ABI:      _auction.Abi,
			Contract: instance,
		})
	}

	return nil
}

func (w *Web3Dialer) LoadMarketplaceInfoFromWeb3() error {
	// 1. First load market name
	if err := w.LoadMarketName(); err != nil {
		return err
	}

	// 2. load fee account
	if err := w.LoadMarketOwner(); err != nil {
		return err
	}

	// 3. load market fee percent
	if err := w.LoadMarketFee(); err != nil {
		return err
	}

	return nil
}

func (w *Web3Dialer) LoadMarketName() error {
	marketName, err := w.Market.Marketplace.Contract.GetName(nil)

	if err != nil {
		return err
	}

	w.Market.Name = marketName

	return nil
}

func (w *Web3Dialer) LoadMarketOwner() error {
	account, err := w.Market.Marketplace.Contract.GetOwner(nil)

	if err != nil {
		return err
	}

	w.Market.Owner = account

	return nil
}

func (w *Web3Dialer) LoadMarketFee() error {
	percent, err := w.Market.Marketplace.Contract.GetFee(nil)

	if err != nil {
		return err
	}

	w.Market.Fee = percent

	return nil
}
