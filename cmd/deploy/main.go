package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ole-larsen/plutonium/contracts/marketplace"
	"github.com/ole-larsen/plutonium/internal/log"
	"github.com/ole-larsen/plutonium/internal/plutonium/settings"
	"github.com/ole-larsen/plutonium/internal/storage"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
)

const (
	MaxItemsInCollection = 10000
	DefaultTimeout       = 250
	DefaultGasLimit      = 30000000
)

func main() {
	logger := log.NewLogger("info", log.DefaultBuildLogger)
	cfg := settings.LoadConfig(".env")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize storage and set it on the service
	store, err := storage.SetupStorage(ctx, cfg.DSN)
	if err != nil {
		panic(err)
	}

	err = deployMarketplace(ctx, store.GetContractsRepository(), cfg)
	if err != nil {
		logger.Fatalln(err)
	}
}

func deployMarketplace(ctx context.Context, repo *repository.ContractsRepository, cfg *settings.Settings) error {
	client, err := ethclient.Dial(cfg.Network)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return err
	}

	auth.Nonce = new(big.Int).SetUint64(nonce)
	auth.Value = big.NewInt(0) // in wei

	auth.GasLimit = uint64(DefaultGasLimit) // in units
	auth.GasPrice = gasPrice

	var feePercent big.Int

	feePercent.SetUint64(1)

	address, tx, instance, err := marketplace.DeployMarketplace(auth, client, cfg.MarketName, &feePercent)
	if err != nil {
		return err
	}

	_ = instance

	contractMap := make(map[string]interface{})
	contractMap["name"] = "marketplace"
	contractMap["address"] = address.Hex()
	contractMap["tx"] = tx.Hash().Hex()
	contractMap["abi"] = marketplace.MarketplaceMetaData.ABI

	err = repo.Create(ctx, contractMap)
	if err != nil {
		return err
	}

	fmt.Printf("Contract Marketplace pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	return nil
}
