// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package auction

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// NFTAuctionMetaData contains all meta data concerning the NFTAuction contract.
var NFTAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"marketPlaceAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftCollectionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftItemId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AuctionEndAlreadyCalled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionNotYetEnded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"}],\"name\":\"BidNotHighEnough\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionStartTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"started\",\"type\":\"bool\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"auctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEnded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isStarted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"itemId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161176c38038061176c833981810160405281019061003291906101c1565b600160008190555086600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550856002819055508460038190555033600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550836005819055508260068190555081600781905550806008819055506000600c60016101000a81548160ff0219169083151502179055506000600c60006101000a81548160ff02191690831515021790555050505050505050610263565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101588261012d565b9050919050565b6101688161014d565b811461017357600080fd5b50565b6000815190506101858161015f565b92915050565b6000819050919050565b61019e8161018b565b81146101a957600080fd5b50565b6000815190506101bb81610195565b92915050565b600080600080600080600060e0888a0312156101e0576101df610128565b5b60006101ee8a828b01610176565b97505060206101ff8a828b016101ac565b96505060406102108a828b016101ac565b95505060606102218a828b016101ac565b94505060806102328a828b016101ac565b93505060a06102438a828b016101ac565b92505060c06102548a828b016101ac565b91505092959891949750929550565b6114fa806102726000396000f3fe6080604052600436106100e85760003560e01c8063a4fd6f561161008a578063eb54f9ec11610059578063eb54f9ec146102c4578063ecfc7ecc146102ef578063f1a9af89146102f9578063fe67a54b14610324576100e8565b8063a4fd6f5614610218578063ca6158cb14610243578063d57bde791461026e578063db2e1eed14610299576100e8565b80634b449cba116100c65780634b449cba14610180578063544736e6146101ab5780636b64c769146101d657806391f90157146101ed576100e8565b8063150b7a02146100ed57806338af3eed1461012a5780633d26bb6714610155575b600080fd5b3480156100f957600080fd5b50610114600480360381019061010f9190610f23565b61033b565b6040516101219190610fe1565b60405180910390f35b34801561013657600080fd5b5061013f61034f565b60405161014c919061101d565b60405180910390f35b34801561016157600080fd5b5061016a610379565b6040516101779190611047565b60405180910390f35b34801561018c57600080fd5b50610195610383565b6040516101a29190611047565b60405180910390f35b3480156101b757600080fd5b506101c061038d565b6040516101cd919061107d565b60405180910390f35b3480156101e257600080fd5b506101eb6103a4565b005b3480156101f957600080fd5b50610202610619565b60405161020f919061101d565b60405180910390f35b34801561022457600080fd5b5061022d610643565b60405161023a919061107d565b60405180910390f35b34801561024f57600080fd5b5061025861065a565b6040516102659190611047565b60405180910390f35b34801561027a57600080fd5b50610283610664565b6040516102909190611047565b60405180910390f35b3480156102a557600080fd5b506102ae61066e565b6040516102bb9190611047565b60405180910390f35b3480156102d057600080fd5b506102d9610678565b6040516102e69190611047565b60405180910390f35b6102f7610682565b005b34801561030557600080fd5b5061030e6109fd565b60405161031b9190611047565b60405180910390f35b34801561033057600080fd5b50610339610a07565b005b600063150b7a0260e01b9050949350505050565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600254905090565b6000600854905090565b6000600c60019054906101000a900460ff16905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061044d5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b61048c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104839061111b565b60405180910390fd5b6007544210156104d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104c890611187565b60405180910390fd5b6008544210610515576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050c906111f3565b60405180910390fd5b600c60019054906101000a900460ff1615610565576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055c9061125f565b60405180910390fd5b6001600c60016101000a81548160ff021916908315150217905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f0dbe763818e0c26abc06649c68eab10b0f73d7491fea8b4b0062c2b77546c066600254600354600654600554600754600854600c60019054906101000a900460ff1660405161060f979695949392919061127f565b60405180910390a2565b6000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600c60009054906101000a900460ff16905090565b6000600354905090565b6000600954905090565b6000600654905090565b6000600754905090565b61068a610cdc565b600c60019054906101000a900460ff166106d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d09061133a565b60405180910390fd5b600c60009054906101000a900460ff1615610729576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610720906113a6565b60405180910390fd5b60075442101561076e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076590611187565b60405180910390fd5b60085442106107b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107a9906111f3565b60405180910390fd5b60065434116107fa576009546040517f4e12c1bb0000000000000000000000000000000000000000000000000000000081526004016107f19190611047565b60405180910390fd5b6005543411610842576009546040517f4e12c1bb0000000000000000000000000000000000000000000000000000000081526004016108399190611047565b60405180910390fd5b600954341161088a576009546040517f4e12c1bb0000000000000000000000000000000000000000000000000000000081526004016108819190611047565b60405180910390fd5b6000600954141580156108ec5750600073ffffffffffffffffffffffffffffffffffffffff16600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b1561095d57600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc6009549081150290604051600060405180830381858888f1935050505015801561095b573d6000803e3d6000fd5b505b3460098190555033600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167f3fabff0a9c3ecd6814702e247fa9733e5d0aa69e3a38590f92cb18f623a2254d346040516109eb9190611047565b60405180910390a26109fb610d2b565b565b6000600554905090565b600c60019054906101000a900460ff16610a56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4d9061133a565b60405180910390fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610aff5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b610b3e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3590611438565b60405180910390fd5b600754421015610b83576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7a90611187565b60405180910390fd5b6008544210610bc7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bbe906111f3565b60405180910390fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc6009549081150290604051600060405180830381858888f19350505050158015610c31573d6000803e3d6000fd5b506000600c60016101000a81548160ff0219169083151502179055506001600c60006101000a81548160ff021916908315150217905550600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda600954604051610cd29190611047565b60405180910390a2565b600260005403610d21576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d18906114a4565b60405180910390fd5b6002600081905550565b6001600081905550565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610d7482610d49565b9050919050565b610d8481610d69565b8114610d8f57600080fd5b50565b600081359050610da181610d7b565b92915050565b6000819050919050565b610dba81610da7565b8114610dc557600080fd5b50565b600081359050610dd781610db1565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610e3082610de7565b810181811067ffffffffffffffff82111715610e4f57610e4e610df8565b5b80604052505050565b6000610e62610d35565b9050610e6e8282610e27565b919050565b600067ffffffffffffffff821115610e8e57610e8d610df8565b5b610e9782610de7565b9050602081019050919050565b82818337600083830152505050565b6000610ec6610ec184610e73565b610e58565b905082815260208101848484011115610ee257610ee1610de2565b5b610eed848285610ea4565b509392505050565b600082601f830112610f0a57610f09610ddd565b5b8135610f1a848260208601610eb3565b91505092915050565b60008060008060808587031215610f3d57610f3c610d3f565b5b6000610f4b87828801610d92565b9450506020610f5c87828801610d92565b9350506040610f6d87828801610dc8565b925050606085013567ffffffffffffffff811115610f8e57610f8d610d44565b5b610f9a87828801610ef5565b91505092959194509250565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b610fdb81610fa6565b82525050565b6000602082019050610ff66000830184610fd2565b92915050565b600061100782610d49565b9050919050565b61101781610ffc565b82525050565b6000602082019050611032600083018461100e565b92915050565b61104181610da7565b82525050565b600060208201905061105c6000830184611038565b92915050565b60008115159050919050565b61107781611062565b82525050565b6000602082019050611092600083018461106e565b92915050565b600082825260208201905092915050565b7f6f6e6c79207468652073656c6c65722063616e2073746172742074686520617560008201527f6374696f6e000000000000000000000000000000000000000000000000000000602082015250565b6000611105602583611098565b9150611110826110a9565b604082019050919050565b60006020820190508181036000830152611134816110f8565b9050919050565b7f61756374696f6e20686173206e6f742073746172746564207965740000000000600082015250565b6000611171601b83611098565b915061117c8261113b565b602082019050919050565b600060208201905081810360008301526111a081611164565b9050919050565b7f61756374696f6e2068617320616c726561647920656e64656400000000000000600082015250565b60006111dd601983611098565b91506111e8826111a7565b602082019050919050565b6000602082019050818103600083015261120c816111d0565b9050919050565b7f61756374696f6e206d757374206e6f7420626520737461727465640000000000600082015250565b6000611249601b83611098565b915061125482611213565b602082019050919050565b600060208201905081810360008301526112788161123c565b9050919050565b600060e082019050611294600083018a611038565b6112a16020830189611038565b6112ae6040830188611038565b6112bb6060830187611038565b6112c86080830186611038565b6112d560a0830185611038565b6112e260c083018461106e565b98975050505050505050565b7f61756374696f6e2073686f756c64206265207374617274656400000000000000600082015250565b6000611324601983611098565b915061132f826112ee565b602082019050919050565b6000602082019050818103600083015261135381611317565b9050919050565b7f61756374696f6e2073686f756c64206265206e6f7420656e6465640000000000600082015250565b6000611390601b83611098565b915061139b8261135a565b602082019050919050565b600060208201905081810360008301526113bf81611383565b9050919050565b7f6f6e6c79207468652073656c6c65722063616e20656e6420746865206175637460008201527f696f6e0000000000000000000000000000000000000000000000000000000000602082015250565b6000611422602383611098565b915061142d826113c6565b604082019050919050565b6000602082019050818103600083015261145181611415565b9050919050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b600061148e601f83611098565b915061149982611458565b602082019050919050565b600060208201905081810360008301526114bd81611481565b905091905056fea264697066735822122054971454fe3551ab96cf6c14fd7f930c7e568cc60444bd1a53b33c40a47e12f664736f6c634300081c0033",
}

// NFTAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTAuctionMetaData.ABI instead.
var NFTAuctionABI = NFTAuctionMetaData.ABI

// NFTAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NFTAuctionMetaData.Bin instead.
var NFTAuctionBin = NFTAuctionMetaData.Bin

// DeployNFTAuction deploys a new Ethereum contract, binding an instance of NFTAuction to it.
func DeployNFTAuction(auth *bind.TransactOpts, backend bind.ContractBackend, marketPlaceAddress common.Address, nftCollectionId *big.Int, nftItemId *big.Int, sPrice *big.Int, rPrice *big.Int, startTime *big.Int, endTime *big.Int) (common.Address, *types.Transaction, *NFTAuction, error) {
	parsed, err := NFTAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NFTAuctionBin), backend, marketPlaceAddress, nftCollectionId, nftItemId, sPrice, rPrice, startTime, endTime)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NFTAuction{NFTAuctionCaller: NFTAuctionCaller{contract: contract}, NFTAuctionTransactor: NFTAuctionTransactor{contract: contract}, NFTAuctionFilterer: NFTAuctionFilterer{contract: contract}}, nil
}

// NFTAuction is an auto generated Go binding around an Ethereum contract.
type NFTAuction struct {
	NFTAuctionCaller     // Read-only binding to the contract
	NFTAuctionTransactor // Write-only binding to the contract
	NFTAuctionFilterer   // Log filterer for contract events
}

// NFTAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTAuctionSession struct {
	Contract     *NFTAuction       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTAuctionCallerSession struct {
	Contract *NFTAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NFTAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTAuctionTransactorSession struct {
	Contract     *NFTAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NFTAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTAuctionRaw struct {
	Contract *NFTAuction // Generic contract binding to access the raw methods on
}

// NFTAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTAuctionCallerRaw struct {
	Contract *NFTAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// NFTAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTAuctionTransactorRaw struct {
	Contract *NFTAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFTAuction creates a new instance of NFTAuction, bound to a specific deployed contract.
func NewNFTAuction(address common.Address, backend bind.ContractBackend) (*NFTAuction, error) {
	contract, err := bindNFTAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFTAuction{NFTAuctionCaller: NFTAuctionCaller{contract: contract}, NFTAuctionTransactor: NFTAuctionTransactor{contract: contract}, NFTAuctionFilterer: NFTAuctionFilterer{contract: contract}}, nil
}

// NewNFTAuctionCaller creates a new read-only instance of NFTAuction, bound to a specific deployed contract.
func NewNFTAuctionCaller(address common.Address, caller bind.ContractCaller) (*NFTAuctionCaller, error) {
	contract, err := bindNFTAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionCaller{contract: contract}, nil
}

// NewNFTAuctionTransactor creates a new write-only instance of NFTAuction, bound to a specific deployed contract.
func NewNFTAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTAuctionTransactor, error) {
	contract, err := bindNFTAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionTransactor{contract: contract}, nil
}

// NewNFTAuctionFilterer creates a new log filterer instance of NFTAuction, bound to a specific deployed contract.
func NewNFTAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTAuctionFilterer, error) {
	contract, err := bindNFTAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionFilterer{contract: contract}, nil
}

// bindNFTAuction binds a generic wrapper to an already deployed contract.
func bindNFTAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NFTAuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTAuction *NFTAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTAuction.Contract.NFTAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTAuction *NFTAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTAuction.Contract.NFTAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTAuction *NFTAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTAuction.Contract.NFTAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTAuction *NFTAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTAuction *NFTAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTAuction *NFTAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTAuction.Contract.contract.Transact(opts, method, params...)
}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_NFTAuction *NFTAuctionCaller) AuctionEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "auctionEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_NFTAuction *NFTAuctionSession) AuctionEndTime() (*big.Int, error) {
	return _NFTAuction.Contract.AuctionEndTime(&_NFTAuction.CallOpts)
}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_NFTAuction *NFTAuctionCallerSession) AuctionEndTime() (*big.Int, error) {
	return _NFTAuction.Contract.AuctionEndTime(&_NFTAuction.CallOpts)
}

// AuctionStartTime is a free data retrieval call binding the contract method 0xeb54f9ec.
//
// Solidity: function auctionStartTime() view returns(uint256)
func (_NFTAuction *NFTAuctionCaller) AuctionStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "auctionStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionStartTime is a free data retrieval call binding the contract method 0xeb54f9ec.
//
// Solidity: function auctionStartTime() view returns(uint256)
func (_NFTAuction *NFTAuctionSession) AuctionStartTime() (*big.Int, error) {
	return _NFTAuction.Contract.AuctionStartTime(&_NFTAuction.CallOpts)
}

// AuctionStartTime is a free data retrieval call binding the contract method 0xeb54f9ec.
//
// Solidity: function auctionStartTime() view returns(uint256)
func (_NFTAuction *NFTAuctionCallerSession) AuctionStartTime() (*big.Int, error) {
	return _NFTAuction.Contract.AuctionStartTime(&_NFTAuction.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_NFTAuction *NFTAuctionCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_NFTAuction *NFTAuctionSession) Beneficiary() (common.Address, error) {
	return _NFTAuction.Contract.Beneficiary(&_NFTAuction.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_NFTAuction *NFTAuctionCallerSession) Beneficiary() (common.Address, error) {
	return _NFTAuction.Contract.Beneficiary(&_NFTAuction.CallOpts)
}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_NFTAuction *NFTAuctionCaller) CollectionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "collectionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_NFTAuction *NFTAuctionSession) CollectionId() (*big.Int, error) {
	return _NFTAuction.Contract.CollectionId(&_NFTAuction.CallOpts)
}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_NFTAuction *NFTAuctionCallerSession) CollectionId() (*big.Int, error) {
	return _NFTAuction.Contract.CollectionId(&_NFTAuction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_NFTAuction *NFTAuctionCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "highestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_NFTAuction *NFTAuctionSession) HighestBid() (*big.Int, error) {
	return _NFTAuction.Contract.HighestBid(&_NFTAuction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_NFTAuction *NFTAuctionCallerSession) HighestBid() (*big.Int, error) {
	return _NFTAuction.Contract.HighestBid(&_NFTAuction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_NFTAuction *NFTAuctionCaller) HighestBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "highestBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_NFTAuction *NFTAuctionSession) HighestBidder() (common.Address, error) {
	return _NFTAuction.Contract.HighestBidder(&_NFTAuction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_NFTAuction *NFTAuctionCallerSession) HighestBidder() (common.Address, error) {
	return _NFTAuction.Contract.HighestBidder(&_NFTAuction.CallOpts)
}

// IsEnded is a free data retrieval call binding the contract method 0xa4fd6f56.
//
// Solidity: function isEnded() view returns(bool)
func (_NFTAuction *NFTAuctionCaller) IsEnded(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "isEnded")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEnded is a free data retrieval call binding the contract method 0xa4fd6f56.
//
// Solidity: function isEnded() view returns(bool)
func (_NFTAuction *NFTAuctionSession) IsEnded() (bool, error) {
	return _NFTAuction.Contract.IsEnded(&_NFTAuction.CallOpts)
}

// IsEnded is a free data retrieval call binding the contract method 0xa4fd6f56.
//
// Solidity: function isEnded() view returns(bool)
func (_NFTAuction *NFTAuctionCallerSession) IsEnded() (bool, error) {
	return _NFTAuction.Contract.IsEnded(&_NFTAuction.CallOpts)
}

// IsStarted is a free data retrieval call binding the contract method 0x544736e6.
//
// Solidity: function isStarted() view returns(bool)
func (_NFTAuction *NFTAuctionCaller) IsStarted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "isStarted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStarted is a free data retrieval call binding the contract method 0x544736e6.
//
// Solidity: function isStarted() view returns(bool)
func (_NFTAuction *NFTAuctionSession) IsStarted() (bool, error) {
	return _NFTAuction.Contract.IsStarted(&_NFTAuction.CallOpts)
}

// IsStarted is a free data retrieval call binding the contract method 0x544736e6.
//
// Solidity: function isStarted() view returns(bool)
func (_NFTAuction *NFTAuctionCallerSession) IsStarted() (bool, error) {
	return _NFTAuction.Contract.IsStarted(&_NFTAuction.CallOpts)
}

// ItemId is a free data retrieval call binding the contract method 0xca6158cb.
//
// Solidity: function itemId() view returns(uint256)
func (_NFTAuction *NFTAuctionCaller) ItemId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "itemId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ItemId is a free data retrieval call binding the contract method 0xca6158cb.
//
// Solidity: function itemId() view returns(uint256)
func (_NFTAuction *NFTAuctionSession) ItemId() (*big.Int, error) {
	return _NFTAuction.Contract.ItemId(&_NFTAuction.CallOpts)
}

// ItemId is a free data retrieval call binding the contract method 0xca6158cb.
//
// Solidity: function itemId() view returns(uint256)
func (_NFTAuction *NFTAuctionCallerSession) ItemId() (*big.Int, error) {
	return _NFTAuction.Contract.ItemId(&_NFTAuction.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_NFTAuction *NFTAuctionCaller) ReservePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "reservePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_NFTAuction *NFTAuctionSession) ReservePrice() (*big.Int, error) {
	return _NFTAuction.Contract.ReservePrice(&_NFTAuction.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_NFTAuction *NFTAuctionCallerSession) ReservePrice() (*big.Int, error) {
	return _NFTAuction.Contract.ReservePrice(&_NFTAuction.CallOpts)
}

// StartPrice is a free data retrieval call binding the contract method 0xf1a9af89.
//
// Solidity: function startPrice() view returns(uint256)
func (_NFTAuction *NFTAuctionCaller) StartPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "startPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartPrice is a free data retrieval call binding the contract method 0xf1a9af89.
//
// Solidity: function startPrice() view returns(uint256)
func (_NFTAuction *NFTAuctionSession) StartPrice() (*big.Int, error) {
	return _NFTAuction.Contract.StartPrice(&_NFTAuction.CallOpts)
}

// StartPrice is a free data retrieval call binding the contract method 0xf1a9af89.
//
// Solidity: function startPrice() view returns(uint256)
func (_NFTAuction *NFTAuctionCallerSession) StartPrice() (*big.Int, error) {
	return _NFTAuction.Contract.StartPrice(&_NFTAuction.CallOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_NFTAuction *NFTAuctionTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTAuction.contract.Transact(opts, "endAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_NFTAuction *NFTAuctionSession) EndAuction() (*types.Transaction, error) {
	return _NFTAuction.Contract.EndAuction(&_NFTAuction.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_NFTAuction *NFTAuctionTransactorSession) EndAuction() (*types.Transaction, error) {
	return _NFTAuction.Contract.EndAuction(&_NFTAuction.TransactOpts)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_NFTAuction *NFTAuctionTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _NFTAuction.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_NFTAuction *NFTAuctionSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _NFTAuction.Contract.OnERC721Received(&_NFTAuction.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_NFTAuction *NFTAuctionTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _NFTAuction.Contract.OnERC721Received(&_NFTAuction.TransactOpts, arg0, arg1, arg2, arg3)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xecfc7ecc.
//
// Solidity: function placeBid() payable returns()
func (_NFTAuction *NFTAuctionTransactor) PlaceBid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTAuction.contract.Transact(opts, "placeBid")
}

// PlaceBid is a paid mutator transaction binding the contract method 0xecfc7ecc.
//
// Solidity: function placeBid() payable returns()
func (_NFTAuction *NFTAuctionSession) PlaceBid() (*types.Transaction, error) {
	return _NFTAuction.Contract.PlaceBid(&_NFTAuction.TransactOpts)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xecfc7ecc.
//
// Solidity: function placeBid() payable returns()
func (_NFTAuction *NFTAuctionTransactorSession) PlaceBid() (*types.Transaction, error) {
	return _NFTAuction.Contract.PlaceBid(&_NFTAuction.TransactOpts)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns()
func (_NFTAuction *NFTAuctionTransactor) StartAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTAuction.contract.Transact(opts, "startAuction")
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns()
func (_NFTAuction *NFTAuctionSession) StartAuction() (*types.Transaction, error) {
	return _NFTAuction.Contract.StartAuction(&_NFTAuction.TransactOpts)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns()
func (_NFTAuction *NFTAuctionTransactorSession) StartAuction() (*types.Transaction, error) {
	return _NFTAuction.Contract.StartAuction(&_NFTAuction.TransactOpts)
}

// NFTAuctionAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the NFTAuction contract.
type NFTAuctionAuctionEndedIterator struct {
	sub      ethereum.Subscription
	fail     error
	Event    *NFTAuctionAuctionEnded
	contract *bind.BoundContract
	logs     chan types.Log
	event    string
	done     bool
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionAuctionEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionAuctionEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionAuctionEnded represents a AuctionEnded event raised by the NFTAuction contract.
type NFTAuctionAuctionEnded struct {
	Amount *big.Int
	Raw    types.Log
	Winner common.Address
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address indexed winner, uint256 amount)
func (_NFTAuction *NFTAuctionFilterer) FilterAuctionEnded(opts *bind.FilterOpts, winner []common.Address) (*NFTAuctionAuctionEndedIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _NFTAuction.contract.FilterLogs(opts, "AuctionEnded", winnerRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionAuctionEndedIterator{contract: _NFTAuction.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address indexed winner, uint256 amount)
func (_NFTAuction *NFTAuctionFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *NFTAuctionAuctionEnded, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _NFTAuction.contract.WatchLogs(opts, "AuctionEnded", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionAuctionEnded)
				if err := _NFTAuction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionEnded is a log parse operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address indexed winner, uint256 amount)
func (_NFTAuction *NFTAuctionFilterer) ParseAuctionEnded(log types.Log) (*NFTAuctionAuctionEnded, error) {
	event := new(NFTAuctionAuctionEnded)
	if err := _NFTAuction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the NFTAuction contract.
type NFTAuctionAuctionStartedIterator struct {
	sub      ethereum.Subscription
	fail     error
	Event    *NFTAuctionAuctionStarted
	contract *bind.BoundContract
	logs     chan types.Log
	event    string
	done     bool
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionAuctionStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionAuctionStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionAuctionStarted represents a AuctionStarted event raised by the NFTAuction contract.
type NFTAuctionAuctionStarted struct {
	CollectionId     *big.Int
	ItemId           *big.Int
	ReservePrice     *big.Int
	StartPrice       *big.Int
	AuctionStartTime *big.Int
	AuctionEndTime   *big.Int
	Raw              types.Log
	Beneficiary      common.Address
	Started          bool
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0x0dbe763818e0c26abc06649c68eab10b0f73d7491fea8b4b0062c2b77546c066.
//
// Solidity: event AuctionStarted(address indexed beneficiary, uint256 collectionId, uint256 itemId, uint256 reservePrice, uint256 startPrice, uint256 auctionStartTime, uint256 auctionEndTime, bool started)
func (_NFTAuction *NFTAuctionFilterer) FilterAuctionStarted(opts *bind.FilterOpts, beneficiary []common.Address) (*NFTAuctionAuctionStartedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _NFTAuction.contract.FilterLogs(opts, "AuctionStarted", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionAuctionStartedIterator{contract: _NFTAuction.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0x0dbe763818e0c26abc06649c68eab10b0f73d7491fea8b4b0062c2b77546c066.
//
// Solidity: event AuctionStarted(address indexed beneficiary, uint256 collectionId, uint256 itemId, uint256 reservePrice, uint256 startPrice, uint256 auctionStartTime, uint256 auctionEndTime, bool started)
func (_NFTAuction *NFTAuctionFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *NFTAuctionAuctionStarted, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _NFTAuction.contract.WatchLogs(opts, "AuctionStarted", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionAuctionStarted)
				if err := _NFTAuction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionStarted is a log parse operation binding the contract event 0x0dbe763818e0c26abc06649c68eab10b0f73d7491fea8b4b0062c2b77546c066.
//
// Solidity: event AuctionStarted(address indexed beneficiary, uint256 collectionId, uint256 itemId, uint256 reservePrice, uint256 startPrice, uint256 auctionStartTime, uint256 auctionEndTime, bool started)
func (_NFTAuction *NFTAuctionFilterer) ParseAuctionStarted(log types.Log) (*NFTAuctionAuctionStarted, error) {
	event := new(NFTAuctionAuctionStarted)
	if err := _NFTAuction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the NFTAuction contract.
type NFTAuctionBidPlacedIterator struct {
	sub      ethereum.Subscription
	fail     error
	Event    *NFTAuctionBidPlaced
	contract *bind.BoundContract
	logs     chan types.Log
	event    string
	done     bool
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionBidPlaced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionBidPlaced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionBidPlaced represents a BidPlaced event raised by the NFTAuction contract.
type NFTAuctionBidPlaced struct {
	Amount *big.Int
	Raw    types.Log
	Bidder common.Address
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x3fabff0a9c3ecd6814702e247fa9733e5d0aa69e3a38590f92cb18f623a2254d.
//
// Solidity: event BidPlaced(address indexed bidder, uint256 amount)
func (_NFTAuction *NFTAuctionFilterer) FilterBidPlaced(opts *bind.FilterOpts, bidder []common.Address) (*NFTAuctionBidPlacedIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _NFTAuction.contract.FilterLogs(opts, "BidPlaced", bidderRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionBidPlacedIterator{contract: _NFTAuction.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x3fabff0a9c3ecd6814702e247fa9733e5d0aa69e3a38590f92cb18f623a2254d.
//
// Solidity: event BidPlaced(address indexed bidder, uint256 amount)
func (_NFTAuction *NFTAuctionFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *NFTAuctionBidPlaced, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _NFTAuction.contract.WatchLogs(opts, "BidPlaced", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionBidPlaced)
				if err := _NFTAuction.contract.UnpackLog(event, "BidPlaced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBidPlaced is a log parse operation binding the contract event 0x3fabff0a9c3ecd6814702e247fa9733e5d0aa69e3a38590f92cb18f623a2254d.
//
// Solidity: event BidPlaced(address indexed bidder, uint256 amount)
func (_NFTAuction *NFTAuctionFilterer) ParseBidPlaced(log types.Log) (*NFTAuctionBidPlaced, error) {
	event := new(NFTAuctionBidPlaced)
	if err := _NFTAuction.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
