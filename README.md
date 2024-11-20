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
