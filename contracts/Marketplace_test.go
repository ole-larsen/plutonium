package contracts_test

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ole-larsen/plutonium/contracts/marketplace"
)

func etherToWei(eth *big.Int) *big.Int {
	return new(big.Int).Mul(eth, big.NewInt(params.Ether))
}

func weiToEther(val *big.Int) *big.Int {
	return new(big.Int).Div(val, big.NewInt(params.Ether))
}

func etherFloatToWei(eth *big.Float) *big.Float {
	return new(big.Float).Mul(eth, big.NewFloat(params.Ether))
}

func weiToEtherFloat(wei *big.Float) *big.Float {
	return new(big.Float).Quo(wei, big.NewFloat(params.Ether))
}

func deployMarketPlace(marketName string, marketFee *big.Int, auth *bind.TransactOpts, client *backends.SimulatedBackend) (common.Address, *marketplace.Marketplace, error) {
	address, _, contract, err := marketplace.DeployMarketplace(
		auth,
		client,
		marketName,
		marketFee,
	)

	client.Commit()

	return address, contract, err
}
func getAuthAndClient() (ownerAuth, buyerAuth, sellerAuth *bind.TransactOpts, client *simulated.Backend) {
	privateOwnerKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	chainID := new(big.Int)
	chainID.SetString("1337", 10)

	ownerAuth, err = bind.NewKeyedTransactorWithChainID(privateOwnerKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	ownerBalance := new(big.Int)
	ownerBalance.SetString("1000000000000000000000", 10) // 1000 eth in wei

	address := ownerAuth.From

	genesisAlloc := map[common.Address]types.Account{
		address: {
			Balance: ownerBalance,
		},
	}

	privateBuyerKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	buyerAuth, err = bind.NewKeyedTransactorWithChainID(privateBuyerKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	buyerBalance := new(big.Int)
	buyerBalance.SetString("100000000000000000000000", 10) // 1000 eth in wei

	genesisAlloc[buyerAuth.From] = types.Account{
		Balance: buyerBalance,
	}

	privateSellerKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	sellerAuth, err = bind.NewKeyedTransactorWithChainID(privateSellerKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	sellerBalance := new(big.Int)
	sellerBalance.SetString("100000000000000000000000", 10) // 1000 eth in wei

	genesisAlloc[sellerAuth.From] = types.Account{
		Balance: sellerBalance,
	}

	client = simulated.NewBackend(genesisAlloc)

	return ownerAuth, buyerAuth, sellerAuth, client
}

const MarketFee = 2.5
const Collection1Fee = 5
const Collection1UpdatedFee = 4.5
const Collection1Price = 100
const Collection1UpdatedPrice = 99.9
const Collectible1Price = 0.5
const Collectible1SellPrice = 0.6
const Collectible2SellPrice = 100

// func TestMarketplaceDeploy(t *testing.T) {
// 	auth, _, _, client := getAuthAndClient()

// 	// fee must be in wei to store float numbers
// 	var fee float64 = MARKET_FEE
// 	var feeInEther *big.Float = big.NewFloat(fee)
// 	var feeInWei *big.Float = etherFloatToWei(feeInEther)
// 	var feeInWeiString string = fmt.Sprintf("%.0f", feeInWei)

// 	marketFee, _ := new(big.Int).SetString(feeInWeiString, 10)

// 	if marketFee.String() != feeInWeiString {
// 		t.Errorf("wrong market fee conversion %f", fee)
// 	}

// 	marketName := "Ploutonion"

// 	if gasPrice, err := client.SuggestGasPrice(context.Background()); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		auth.GasPrice = gasPrice
// 	}

// 	marketAddress, marketPlace, err := deployMarketPlace(marketName, marketFee, auth, client)

// 	if err != nil {
// 		t.Fatalf("Failed to deploy the marketPlace contract: %v", err)
// 	}

// 	if len(marketAddress.Bytes()) == 0 {
// 		t.Error("Expected a valid deployment address. Received empty address byte array instead")
// 	}

// 	log.Println("Marketplace contract address", marketAddress)

// 	if got, _ := marketPlace.GetName(nil); got != marketName {
// 		t.Errorf("Expected marketPlace name %s: %s", marketName, got)
// 	}

// 	if got, _ := marketPlace.GetFee(nil); got.Int64() != marketFee.Int64() {
// 		t.Errorf("Expected marketFee %s: %s", marketFee.String(), got.String())
// 	}

// 	if got, _ := marketPlace.GetOwner(nil); got != auth.From {
// 		t.Errorf("Expected marketPlace owner %s: %s", auth.From, got)
// 	}

// 	// change name
// 	newMarketName := "Ploutonion 2.0"

// 	if gasPrice, err := client.SuggestGasPrice(context.Background()); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		auth.GasPrice = gasPrice
// 	}
// 	if tx, err := marketPlace.SetName(auth, newMarketName); err != nil {
// 		log.Fatalf("Failed to update value: %v", err)
// 	} else {
// 		log.Printf("Update pending setName: 0x%x\n", tx.Hash())
// 		client.Commit()
// 	}

// 	if got, _ := marketPlace.GetName(nil); got != newMarketName {
// 		t.Errorf("Expected new marketPlace name %s: %s", newMarketName, got)
// 	}

// 	// change fee
// 	var newFee float64 = 0.55
// 	var newFeeInEther *big.Float = big.NewFloat(newFee)
// 	var newFeeInWei *big.Float = etherFloatToWei(newFeeInEther)
// 	var newFeeInWeiString string = fmt.Sprintf("%.0f", newFeeInWei)

// 	newMarketFee, _ := new(big.Int).SetString(newFeeInWeiString, 10)

// 	if gasPrice, err := client.SuggestGasPrice(context.Background()); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		auth.GasPrice = gasPrice
// 	}
// 	if tx, err := marketPlace.SetFee(auth, newMarketFee); err != nil {
// 		log.Fatalf("Failed to update value: %v", err)
// 	} else {
// 		log.Printf("Update pending setFee: 0x%x\n", tx.Hash())
// 		client.Commit()
// 	}

// 	if got, _ := marketPlace.GetFee(nil); got.Int64() != newMarketFee.Int64() {
// 		t.Errorf("Expected new marketFee %s: %s", newMarketFee.String(), got.String())
// 	}
// }

// func TestMarketplaceCollection(t *testing.T) {
// 	auth, _, _, client := getAuthAndClient()

// 	// fee must be in wei to store float numbers
// 	var fee float64 = MARKET_FEE
// 	var feeInEther *big.Float = big.NewFloat(fee)
// 	var feeInWei *big.Float = etherFloatToWei(feeInEther)
// 	var feeInWeiString string = fmt.Sprintf("%.0f", feeInWei)

// 	marketFee, _ := new(big.Int).SetString(feeInWeiString, 10)

// 	var marketName = "Ploutonion"

// 	if gasPrice, err := client.SuggestGasPrice(context.Background()); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		auth.GasPrice = gasPrice
// 	}

// 	_, marketPlace, err := deployMarketPlace(marketName, marketFee, auth, client)

// 	collectionName := "Art Collecion"
// 	collectionSymbol := "TKN"

// 	if gasPrice, err := client.SuggestGasPrice(context.Background()); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		auth.GasPrice = gasPrice
// 	}

// 	// deploy collection
// 	collectionAddress, _, collectionContract, err := collection.DeployNFTCollection(auth, client, collectionName, collectionSymbol)
// 	if err != nil {
// 		t.Fatalf("Failed to deploy the NFT Collection %s contract: %v", collectionName, err)
// 	}

// 	if len(collectionAddress.Bytes()) == 0 {
// 		t.Error("Expected a valid deployment address. Received empty address byte array instead")
// 	}

// 	log.Printf("%s contract address: %s", collectionName, collectionAddress)

// 	if collectionContract != nil {
// 	}

// 	collectionDescription := "Modern Art Collection"

// 	var _collectionFee float64 = COLLECTION_1_FEE
// 	var collectionFeeInEther *big.Float = big.NewFloat(_collectionFee)
// 	var collectionFeeInWei *big.Float = etherFloatToWei(collectionFeeInEther)
// 	var collectionFeeInWeiString string = fmt.Sprintf("%.0f", collectionFeeInWei)

// 	collectionFee, _ := new(big.Int).SetString(collectionFeeInWeiString, 10)

// 	// change fee
// 	var _collectionPrice float64 = COLLECTION_1_PRICE
// 	var collectionPriceInEther *big.Float = big.NewFloat(_collectionPrice)
// 	var collectionPriceInWei *big.Float = etherFloatToWei(collectionPriceInEther)
// 	var collectionPriceInWeiString string = fmt.Sprintf("%.0f", collectionPriceInWei)

// 	collectionPrice, _ := new(big.Int).SetString(collectionPriceInWeiString, 10)

// 	if gasPrice, err := client.SuggestGasPrice(context.Background()); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		auth.GasPrice = new(big.Int).Mul(gasPrice, big.NewInt(1000))
// 	}

// 	if collection1Tx, err := marketPlace.CreateCollection(
// 		auth,
// 		collectionName,
// 		collectionSymbol,
// 		collectionDescription,
// 		collectionFee,
// 		collectionPrice,
// 		collectionAddress,
// 		auth.From,
// 	); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		log.Printf("Update pending createCollection %s: 0x%x\n", collectionName, collection1Tx.Hash())
// 		client.Commit()
// 	}

// 	// test created collection
// 	collectionCount, _ := marketPlace.GetCollectionCounter(nil)
// 	if collectionCount.Int64() != 1 {
// 		t.Errorf("Expected marketPlace collections %d: %s", 1, collectionCount.String())
// 	}

// 	// get created collection1
// 	collection1, err := marketPlace.GetCollection(nil, collectionCount)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if collection1.Name != collectionName {
// 		t.Errorf("Collection Name %s does not match %s", collection1.Name, collectionName)
// 	}

// 	if collection1.Symbol != collectionSymbol {
// 		t.Errorf("Collection Symbol %s does not match %s", collection1.Symbol, collectionSymbol)
// 	}

// 	if collection1.Description != collectionDescription {
// 		t.Errorf("Collection Symbol %s does not match %s", collection1.Description, collectionDescription)
// 	}

// 	if collection1.NftCollection != collectionAddress {
// 		t.Errorf("Collection address %s does not match %s", collection1.NftCollection, collectionAddress)
// 	}

// 	if collection1.Fee.Int64() != collectionFee.Int64() {
// 		t.Errorf("Collection fee %s does not match %s", collection1.Fee.String(), collection1.Fee.String())
// 	}

// 	if collection1.Owner != auth.From {
// 		t.Errorf("Collection owner %s does not match %s", collection1.Owner, auth.From)
// 	}

// 	if collection1.Creator != auth.From {
// 		t.Errorf("Collection creator %s does not match %s", collection1.Creator, auth.From)
// 	}

// 	if collection1.IsApproved != false {
// 		t.Errorf("Collection isApproved %v does not match false", collection1.IsApproved)
// 	}

// 	if collection1.IsLocked != false {
// 		t.Errorf("Collection isLocked %v does not match false", collection1.IsLocked)
// 	}

// 	collectionUpdatedDescription := "Updated Modern Art Collection"

// 	var _collectionUpdatedFee float64 = COLLECTION_1_UPDATED_FEE
// 	var collectionUpdatedFeeInEther *big.Float = big.NewFloat(_collectionUpdatedFee)
// 	var collectionUpdatedFeeInWei *big.Float = etherFloatToWei(collectionUpdatedFeeInEther)
// 	var collectionUpdatedFeeInWeiString string = fmt.Sprintf("%.0f", collectionUpdatedFeeInWei)

// 	collectionUpdatedFee, _ := new(big.Int).SetString(collectionUpdatedFeeInWeiString, 10)

// 	// change fee
// 	var _collectionUpdatedPrice float64 = COLLECTION_1_UPDATED_PRICE
// 	var collectionUpdatedPriceInEther *big.Float = big.NewFloat(_collectionUpdatedPrice)
// 	var collectionUpdatedPriceInWei *big.Float = etherFloatToWei(collectionUpdatedPriceInEther)
// 	var collectionUpdatedPriceInWeiString string = fmt.Sprintf("%.0f", collectionUpdatedPriceInWei)

// 	collectionUpdatedPrice, _ := new(big.Int).SetString(collectionUpdatedPriceInWeiString, 10)

// 	if gasPrice, err := client.SuggestGasPrice(context.Background()); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		auth.GasPrice = new(big.Int).Mul(gasPrice, big.NewInt(1000))
// 	}

// 	if _, err := marketPlace.EditCollection(
// 		auth,
// 		collectionCount,
// 		collection1.Name,
// 		collection1.Symbol,
// 		collectionUpdatedDescription,
// 		collectionUpdatedFee,
// 		collectionUpdatedPrice,
// 		collection1.NftCollection,
// 		collection1.Owner,
// 		collection1.IsApproved,
// 		collection1.IsLocked,
// 	); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		client.Commit()
// 		collection1, err := marketPlace.GetCollection(nil, collectionCount)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		if collection1.Name != collectionName {
// 			t.Errorf("Collection Name %s does not match %s", collection1.Name, collectionName)
// 		}

// 		if collection1.Symbol != collectionSymbol {
// 			t.Errorf("Collection Symbol %s does not match %s", collection1.Symbol, collectionSymbol)
// 		}

// 		if collection1.Description != collectionUpdatedDescription {
// 			t.Errorf("Collection Symbol %s does not match %s", collection1.Description, collectionUpdatedDescription)
// 		}

// 		if collection1.NftCollection != collectionAddress {
// 			t.Errorf("Collection address %s does not match %s", collection1.NftCollection, collectionAddress)
// 		}

// 		if collection1.Fee.Int64() != collectionUpdatedFee.Int64() {
// 			t.Errorf("Collection fee %s does not match %s", collection1.Fee.String(), collectionUpdatedFee.String())
// 		}

// 		if collection1.Owner != auth.From {
// 			t.Errorf("Collection owner %s does not match %s", collection1.Owner, auth.From)
// 		}

// 		if collection1.Creator != auth.From {
// 			t.Errorf("Collection creator %s does not match %s", collection1.Creator, auth.From)
// 		}

// 		if collection1.IsApproved != false {
// 			t.Errorf("Collection isApproved %v does not match false", collection1.IsApproved)
// 		}

// 		if collection1.IsLocked != false {
// 			t.Errorf("Collection isLocked %v does not match false", collection1.IsLocked)
// 		}
// 	}
// }

// func TestMarketplaceCollectibleSafeMint(t *testing.T) {
// 	auth, buyerAuth, sellerAuth, client := getAuthAndClient()

// 	// fee must be in wei to store float numbers
// 	var fee float64 = MARKET_FEE
// 	var feeInEther *big.Float = big.NewFloat(fee)
// 	var feeInWei *big.Float = etherFloatToWei(feeInEther)
// 	var feeInWeiString string = fmt.Sprintf("%.0f", feeInWei)

// 	marketFee, _ := new(big.Int).SetString(feeInWeiString, 10)

// 	var marketName = "Ploutonion"

// 	marketAddress, marketPlace, err := deployMarketPlace(marketName, marketFee, auth, client)

// 	collectionName := "Art Collecion"
// 	collectionSymbol := "TKN"

// 	if gasPrice, err := client.SuggestGasPrice(context.Background()); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		auth.GasPrice = gasPrice
// 	}

// 	// deploy collection
// 	collectionAddress, _, collectionContract, err := collection.DeployNFTCollection(auth, client, collectionName, collectionSymbol)
// 	if err != nil {
// 		t.Fatalf("Failed to deploy the NFT Collection %s contract: %v", collectionName, err)
// 	}

// 	if len(collectionAddress.Bytes()) == 0 {
// 		t.Error("Expected a valid deployment address. Received empty address byte array instead")
// 	}

// 	log.Printf("%s contract address: %s", collectionName, collectionAddress)

// 	if collectionContract != nil {
// 	}

// 	collectionDescription := "Modern Art Collection"

// 	var _collectionFee float64 = COLLECTION_1_FEE
// 	var collectionFeeInEther *big.Float = big.NewFloat(_collectionFee)
// 	var collectionFeeInWei *big.Float = etherFloatToWei(collectionFeeInEther)
// 	var collectionFeeInWeiString string = fmt.Sprintf("%.0f", collectionFeeInWei)

// 	collectionFee, _ := new(big.Int).SetString(collectionFeeInWeiString, 10)

// 	// change fee
// 	var _collectionPrice float64 = COLLECTION_1_PRICE
// 	var collectionPriceInEther *big.Float = big.NewFloat(_collectionPrice)
// 	var collectionPriceInWei *big.Float = etherFloatToWei(collectionPriceInEther)
// 	var collectionPriceInWeiString string = fmt.Sprintf("%.0f", collectionPriceInWei)

// 	collectionPrice, _ := new(big.Int).SetString(collectionPriceInWeiString, 10)

// 	if gasPrice, err := client.SuggestGasPrice(context.Background()); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		auth.GasPrice = new(big.Int).Mul(gasPrice, big.NewInt(1000))
// 	}

// 	if collection1Tx, err := marketPlace.CreateCollection(
// 		auth,
// 		collectionName,
// 		collectionSymbol,
// 		collectionDescription,
// 		collectionFee,
// 		collectionPrice,
// 		collectionAddress,
// 		auth.From,
// 	); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		log.Printf("Update pending createCollection %s: 0x%x\n", collectionName, collection1Tx.Hash())
// 		client.Commit()
// 	}

// 	// test created collection
// 	collectionCount, _ := marketPlace.GetCollectionCounter(nil)
// 	if collectionCount.Int64() != 1 {
// 		t.Errorf("Expected marketPlace collections %d: %s", 1, collectionCount.String())
// 	}

// 	// get created collection1
// 	collection1, err := marketPlace.GetCollection(nil, collectionCount)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if collection1.Name != collectionName {
// 		t.Errorf("Collection Name %s does not match %s", collection1.Name, collectionName)
// 	}

// 	// create first token in and add it to collection
// 	// 1. Check collection has no tokens
// 	if got, err := collectionContract.TotalSupply(nil); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		if got.Int64() != 0 {
// 			t.Errorf("Collection totalSupply %d does not match 0", got.Int64())
// 		}
// 	}

// 	if gasPrice, err := client.SuggestGasPrice(context.Background()); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		auth.GasPrice = gasPrice
// 	}

// 	if safeMintTx, err := collectionContract.SafeMint(auth, "testURI"); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		log.Printf("Update pending safeMint: 0x%x\n", safeMintTx.Hash())
// 		client.Commit()
// 	}

// 	// check token created
// 	token1ID, err := collectionContract.TotalSupply(nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if token1ID.Int64() != 1 {
// 		t.Errorf("Collection totalSupply %d does not match 2", token1ID.Int64())
// 	}

// 	// check balances
// 	ownerBalance, err := collectionContract.BalanceOf(nil, auth.From)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// because firwst token was listed owner balance become 0
// 	if ownerBalance.Int64() != 1 {
// 		t.Errorf("NFT Collection owner balance, who created token does not match %d", ownerBalance.Int64())
// 	}

// 	buyerBalance, err := collectionContract.BalanceOf(nil, buyerAuth.From)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if buyerBalance.Int64() != 0 {
// 		t.Errorf("NFT Collection buyer balance, who created token does not match %d", buyerBalance.Int64())
// 	}

// 	sellerBalance, err := collectionContract.BalanceOf(nil, sellerAuth.From)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if sellerBalance.Int64() != 0 {
// 		t.Errorf("NFT Collection seller balance, who created token does not match %d", sellerBalance.Int64())
// 	}

// 	// check the owner of created token
// 	owner, err := collectionContract.OwnerOf(nil, token1ID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if owner != auth.From {
// 		t.Errorf("Token owner %s does not match %s", owner, auth.From)
// 	}

// 	// createCollectible(uint256[] memory _tokenIds, uint256 _collectionId, bool _isAuction, address _nftAuction)
// 	if got, err := marketPlace.GetCollectibleCount(nil, collectionCount); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		if got.Int64() != 0 {
// 			t.Errorf("Collection 1 collectible %d does not match 0", got.Int64())
// 		}
// 	}

// 	var price float64 = COLLECTIBLE_1_PRICE
// 	var priceInEther *big.Float = big.NewFloat(price)
// 	var priceInWei *big.Float = etherFloatToWei(priceInEther)
// 	var priceInWeiString string = fmt.Sprintf("%.0f", priceInWei)

// 	collectiblePrice, _ := new(big.Int).SetString(priceInWeiString, 10)

// 	log.Println("current collectible price in eth:", priceInEther)
// 	log.Println("current collectible price in wei:", collectiblePrice)

// 	// add collectible
// 	// add collectible
// 	if _, err = collectionContract.SetApprovalForAll(auth, marketAddress, true); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		client.Commit()
// 	}

// 	isAuthApproved, err := collectionContract.IsApprovedForAll(nil, auth.From, marketAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if isAuthApproved == false {
// 		t.Error("auth is not approved")
// 	}

// 	// try to set apprival for buyer
// 	if _, err = collectionContract.SetApprovalForAll(buyerAuth, marketAddress, true); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		client.Commit()
// 	}

// 	isBuyerApproved, err := collectionContract.IsApprovedForAll(nil, buyerAuth.From, marketAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if isBuyerApproved == false {
// 		t.Error("buyer auth is not approved")
// 	}

// 	tokenIds := make([]*big.Int, 0)
// 	tokenIds = append(tokenIds, token1ID)

// 	if gasPrice, err := client.SuggestGasPrice(context.Background()); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		auth.GasPrice = gasPrice
// 	}

// 	if collectibleTx, err := marketPlace.CreateCollectible(auth, tokenIds, collectionCount, false, collectiblePrice); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		log.Printf("Update pending createCollectible: 0x%x\n", collectibleTx.Hash())
// 		client.Commit()
// 	}

// 	if got, err := marketPlace.GetCollectibleCount(nil, collectionCount); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		if got.Int64() != 1 {
// 			t.Errorf("wrong collectible count %d", got.Int64())
// 		}
// 	}

// 	collectibleID, _ := marketPlace.GetCollectibleCount(nil, collectionCount)

// 	if got, err := marketPlace.GetCollectible(nil, collectionCount, collectibleID); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		if got.Id.Int64() != collectibleID.Int64() {
// 			t.Errorf("wrong id %d", got.Id.Int64())
// 		}
// 		if got.CollectionId.Int64() != collectionCount.Int64() {
// 			t.Errorf("wrong id %d", got.Id.Int64())
// 		}
// 		if len(got.TokenIds) != len(tokenIds) {
// 			t.Error("wrong tokens number", tokenIds)
// 		}
// 		if got.Owners[0] != auth.From {
// 			t.Errorf("wrong owner %s", got.Owners[0])
// 		}
// 		if got.Creator != auth.From {
// 			t.Errorf("wrong creator %s", got.Creator)
// 		}
// 		if got.IsAuction != false {
// 			t.Error("wrong auction", got.IsAuction)
// 		}
// 	}

// 	// check token owner
// 	// check balances
// 	ownerBalance, err = collectionContract.BalanceOf(nil, auth.From)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// token moved from owner balance to contract
// 	if ownerBalance.Int64() != 0 {
// 		t.Errorf("NFT Collection owner balance, who created token does not match %d", ownerBalance.Int64())
// 	}

// 	marketBalance, err := collectionContract.BalanceOf(nil, marketAddress)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// token moved from owner balance to contract
// 	if marketBalance.Int64() != 1 {
// 		t.Errorf("NFT Collection owner balance, who created token does not match %d", ownerBalance.Int64())
// 	}
// 	buyerBalance, err = collectionContract.BalanceOf(nil, buyerAuth.From)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if buyerBalance.Int64() != 0 {
// 		t.Errorf("NFT Collection buyer balance, who created token does not match %d", buyerBalance.Int64())
// 	}

// 	sellerBalance, err = collectionContract.BalanceOf(nil, sellerAuth.From)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if sellerBalance.Int64() != 0 {
// 		t.Errorf("NFT Collection seller balance, who created token does not match %d", sellerBalance.Int64())
// 	}

// 	// check the owner of created token
// 	owner, err = collectionContract.OwnerOf(nil, token1ID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if owner != marketAddress {
// 		t.Errorf("Token owner %s does not match %s", owner, marketAddress)
// 	}

// 	if got, err := marketPlace.GetCollectible(nil, collectionCount, collectibleID); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		if got.IsAuction != false {
// 			t.Error("auction must have false value")
// 		}
// 		if got.Fulfilled[0] != false {
// 			t.Error("fulfilled must be false")
// 		}
// 		if got.IsLocked != false {
// 			t.Error("locked must be false")
// 		}

// 		if new(big.Int).Div(got.Price, big.NewInt(params.Ether)).Int64() != int64(price) {
// 			t.Errorf("wrong price conversion %d", int64(price))
// 		}

// 		// collection1, _ := marketPlace.GetCollection(nil, collection1Count)

// 		//var _collectionFee float64 = COLLECTION_1_FEE

// 		var price float64 = COLLECTIBLE_1_PRICE

// 		total := price + price*_collectionFee/100
// 		total = total * params.Ether
// 		totalString := fmt.Sprintf("%.0f", total)

// 		collectible1TotalPrice, _ := new(big.Int).SetString(totalString, 10)
// 		buyerAuth.Value = collectible1TotalPrice
// 		quantity := big.NewInt(1)

// 		if buyTx, err := marketPlace.Buy(buyerAuth, collectionCount, got.Id, quantity); err != nil {
// 			log.Fatal(err)
// 		} else {
// 			log.Printf("Update pending buyCollectible: 0x%x\n", buyTx.Hash())
// 			client.Commit()
// 		}

// 		collectible, err := marketPlace.GetCollectible(nil, collectionCount, collectibleID)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if collectible.Owners[0] != buyerAuth.From {
// 			t.Errorf("wrong owner")
// 		}
// 		// check token owner
// 		// check balances
// 		ownerBalance, err = collectionContract.BalanceOf(nil, auth.From)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// token moved from owner balance to contract
// 		if ownerBalance.Int64() != 0 {
// 			t.Errorf("NFT Collection owner balance, who created token does not match %d", ownerBalance.Int64())
// 		}

// 		marketBalance, err := collectionContract.BalanceOf(nil, marketAddress)

// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// token moved from owner balance to contract
// 		if marketBalance.Int64() != 0 {
// 			t.Errorf("NFT Collection owner balance, who created token does not match %d", ownerBalance.Int64())
// 		}
// 		buyerBalance, err = collectionContract.BalanceOf(nil, buyerAuth.From)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if buyerBalance.Int64() != 1 {
// 			t.Errorf("NFT Collection buyer balance, who created token does not match %d", buyerBalance.Int64())
// 		}

// 		sellerBalance, err = collectionContract.BalanceOf(nil, sellerAuth.From)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if sellerBalance.Int64() != 0 {
// 			t.Errorf("NFT Collection seller balance, who created token does not match %d", sellerBalance.Int64())
// 		}

// 		// check the owner of created token
// 		owner, err = collectionContract.OwnerOf(nil, token1ID)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		if owner != buyerAuth.From {
// 			t.Errorf("Token owner %s does not match %s", owner, buyerAuth.From)
// 		}
// 	}

// 	isBuyerApproved, err = collectionContract.IsApprovedForAll(nil, buyerAuth.From, marketAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var newPrice float64 = COLLECTIBLE_1_SELL_PRICE
// 	var newPriceInEther *big.Float = big.NewFloat(newPrice)
// 	var newPriceInWei *big.Float = etherFloatToWei(newPriceInEther)
// 	var newPriceInWeiString string = fmt.Sprintf("%.0f", newPriceInWei)

// 	collectibleNewPrice, _ := new(big.Int).SetString(newPriceInWeiString, 10)

// 	if sellTx, err := marketPlace.Sell(buyerAuth, collectionCount, big.NewInt(1), collectibleNewPrice); err != nil {
// 		log.Fatal("sell error", err)
// 	} else {
// 		log.Printf("Update pending sell collectible: 0x%x\n", sellTx.Hash())
// 		client.Commit()
// 	}
// }
