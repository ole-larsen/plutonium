# Plutonium NFT Marketplace

Welcome to Plutonium NFT Marketplace, a powerful and feature-rich platform for creating, buying, and selling NFTs (Non-Fungible Tokens).

## Develop Installation

Plutonium requires [golang](https://go.dev/) v1.23+ to run.

Install the dependencies and devDependencies and start the server.

```sh
cd plutonium
docker pull quay.io/goswagger/swagger
alias swagger='docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger'
swagger generate server -A service -P models.Principal -f ./schema/swagger.yml   
```

## Tests
```bash
go test -v -coverpkg=./... -coverprofile=coverage.out -covermode=count ./...
go test -race -covermode=atomic -coverprofile=coverage.out ./... &&
go tool cover -func coverage.out | grep total | awk '{print $3}' &&
go tool cover -html="coverage.out"
```

## Make sure your gopath is correct

```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

### Install abigen

```
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```

# Compile contracts

```
solc --abi contracts/nft/NFT.sol -o contracts/nft/abi --overwrite && \
solc --abi contracts/collection/NFTCollection.sol -o contracts/collection/abi --overwrite && \
solc --abi contracts/marketplace/Marketplace.sol -o contracts/marketplace/abi --overwrite && \
solc --abi contracts/ballot/Ballot.sol -o contracts/ballot/abi --overwrite && \
solc --abi contracts/auction/NFTAuction.sol -o contracts/auction/abi --overwrite
```

# Generate bindings

```
abigen --abi contracts/nft/abi/NFT.abi --pkg nft --type NFT --out contracts/nft/NFT.go && \
abigen --abi contracts/collection/abi/NFTCollection.abi --pkg collection --type NFTCollection --out contracts/collection/NFTCollection.go && \
abigen --abi contracts/marketplace/abi/Marketplace.abi --pkg marketplace --type Marketplace --out contracts/marketplace/Marketplace.go && \
abigen --abi contracts/ballot/abi/Ballot.abi --pkg ballot --type Ballot --out contracts/ballot/Ballot.go && \
abigen --abi contracts/auction/abi/NFTAuction.abi --pkg auction --type Auction --out contracts/auction/NFTAuction.go
```

# Deploying contracts

```
solc --bin contracts/nft/NFT.sol -o contracts/nft/bin --overwrite && \
solc --bin contracts/collection/NFTCollection.sol -o contracts/collection/bin --overwrite && \
solc --bin contracts/marketplace/Marketplace.sol -o contracts/marketplace/bin --overwrite && \
solc --bin contracts/ballot/Ballot.sol -o contracts/ballot/bin --overwrite && \
solc --bin contracts/auction/NFTAuction.sol -o contracts/auction/bin --overwrite

# One Command

```
solc --evm-version paris --abi contracts/nft/NFT.sol -o contracts/nft/abi --overwrite && \
solc --evm-version paris --abi contracts/collection/NFTCollection.sol -o contracts/collection/abi --overwrite && \
solc --evm-version paris --abi contracts/marketplace/Marketplace.sol -o contracts/marketplace/abi --overwrite && \
solc --evm-version paris --abi contracts/ballot/Ballot.sol -o contracts/ballot/abi --overwrite && \
solc --evm-version paris --abi contracts/auction/NFTAuction.sol -o contracts/auction/abi --overwrite && \
abigen --abi contracts/nft/abi/NFT.abi --pkg nft --type NFT --out contracts/nft/NFT.go && \
abigen --abi contracts/collection/abi/NFTCollection.abi --pkg collection --type NFTCollection --out contracts/collection/NFTCollection.go && \
abigen --abi contracts/marketplace/abi/Marketplace.abi --pkg marketplace --type Marketplace --out contracts/marketplace/Marketplace.go && \
abigen --abi contracts/ballot/abi/Ballot.abi --pkg ballot --type Ballot --out contracts/ballot/Ballot.go && \
abigen --abi contracts/auction/abi/NFTAuction.abi --pkg auction --type NFTAuction --out contracts/auction/NFTAuction.go && \
solc --evm-version paris --bin contracts/nft/NFT.sol -o contracts/nft/bin --overwrite && \
solc --evm-version paris --bin contracts/collection/NFTCollection.sol -o contracts/collection/bin --overwrite && \
solc --evm-version paris --bin contracts/marketplace/Marketplace.sol -o contracts/marketplace/bin --overwrite && \
solc --evm-version paris --bin contracts/ballot/Ballot.sol -o contracts/ballot/bin --overwrite && \
solc --evm-version paris --bin contracts/auction/NFTAuction.sol -o contracts/auction/bin --overwrite && \
abigen --abi contracts/nft/abi/NFT.abi --pkg nft --type NFT --out contracts/nft/NFT.go --bin contracts/nft/bin/NFT.bin && \
abigen --abi contracts/collection/abi/NFTCollection.abi --pkg collection   --type NFTCollection --out contracts/collection/NFTCollection.go --bin contracts/collection/bin/NFTCollection.bin && \
abigen --abi contracts/marketplace/abi/Marketplace.abi   --pkg marketplace --type Marketplace   --out contracts/marketplace/Marketplace.go  --bin contracts/marketplace/bin/Marketplace.bin && \
abigen --abi contracts/ballot/abi/Ballot.abi --pkg ballot --type Ballot --out contracts/ballot/Ballot.go --bin contracts/ballot/bin/Ballot.bin && \
abigen --abi contracts/auction/abi/NFTAuction.abi --pkg auction --type NFTAuction --out contracts/auction/NFTAuction.go --bin contracts/auction/bin/NFTAuction.bin
```

### Mocks
```bash
mockgen -source=internal/storage/db/repository/users.go -destination=internal/storage/mocks/users_mock.go -package=mocks
mockgen -source=internal/storage/db/repository/contracts.go -destination=internal/storage/mocks/contracts_mock.go -package=mocks
mockgen -source=internal/storage/dbstorage.go -destination=internal/storage/mocks/dbstorage_mock.go -package=mocks
```