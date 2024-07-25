// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lagrange

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

// IBLSKeyCheckerBLSKeyWithProof is an auto generated low-level Go binding around an user-defined struct.
type IBLSKeyCheckerBLSKeyWithProof struct {
	BlsG1PublicKeys [][2]*big.Int
	AggG2PublicKey  [2][2]*big.Int
	Signature       [2]*big.Int
	Salt            [32]byte
	Expiry          *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// LagrangeMetaData contains all meta data concerning the Lagrange contract.
var LagrangeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_committee\",\"type\":\"address\",\"internalType\":\"contractILagrangeCommittee\"},{\"name\":\"_stakeManager\",\"type\":\"address\",\"internalType\":\"contractIStakeManager\"},{\"name\":\"_avsDirectoryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_voteWeigher\",\"type\":\"address\",\"internalType\":\"contractIVoteWeigher\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addBlsPubKeys\",\"inputs\":[{\"name\":\"blsKeyWithProof\",\"type\":\"tuple\",\"internalType\":\"structIBLSKeyChecker.BLSKeyWithProof\",\"components\":[{\"name\":\"blsG1PublicKeys\",\"type\":\"uint256[2][]\",\"internalType\":\"uint256[2][]\"},{\"name\":\"aggG2PublicKey\",\"type\":\"uint256[2][2]\",\"internalType\":\"uint256[2][2]\"},{\"name\":\"signature\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOperatorsToWhitelist\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"committee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILagrangeCommittee\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregister\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorWhitelist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"signAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blsKeyWithProof\",\"type\":\"tuple\",\"internalType\":\"structIBLSKeyChecker.BLSKeyWithProof\",\"components\":[{\"name\":\"blsG1PublicKeys\",\"type\":\"uint256[2][]\",\"internalType\":\"uint256[2][]\"},{\"name\":\"aggG2PublicKey\",\"type\":\"uint256[2][2]\",\"internalType\":\"uint256[2][2]\"},{\"name\":\"signature\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeBlsPubKeys\",\"inputs\":[{\"name\":\"indices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeOperatorsFromWhitelist\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"subscribe\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unsubscribe\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unsubscribeByAdmin\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"chainID\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateBlsPubKey\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blsKeyWithProof\",\"type\":\"tuple\",\"internalType\":\"structIBLSKeyChecker.BLSKeyWithProof\",\"components\":[{\"name\":\"blsG1PublicKeys\",\"type\":\"uint256[2][]\",\"internalType\":\"uint256[2][]\"},{\"name\":\"aggG2PublicKey\",\"type\":\"uint256[2][2]\",\"internalType\":\"uint256[2][2]\"},{\"name\":\"signature\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSignAddress\",\"inputs\":[{\"name\":\"newSignAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"voteWeigher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIVoteWeigher\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"serveUntilBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSubscribed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorUnsubscribed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnsubscribedByAdmin\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b506004361061014d5760003560e01c80638da5cb5b116100c3578063d822af3c1161007c578063d822af3c146102cc578063d864e740146102df578063e03c863214610306578063e481af9d14610339578063ef03067314610341578063f2fde38b1461036857600080fd5b80638da5cb5b14610270578063935a9b6a14610278578063a98fb3551461028b578063aff5edb11461029e578063c4d66de8146102a6578063c5bdc30a146102b957600080fd5b80635e332fe4116101155780635e332fe4146101c95780635fa7306a146101dc57806363d944ea146101ef5780636b3aa72e14610202578063715018a6146102415780637542ff951461024957600080fd5b80630512d04c146101525780631be4b9f7146101675780632e94d67b1461017a57806333cfb7b71461018d5780634e32e904146101b6575b600080fd5b6101656101603660046113f2565b61037b565b005b61016561017536600461145f565b61046e565b6101656101883660046113f2565b6104ed565b6101a061019b3660046114b5565b6105d7565b6040516101ad91906114d2565b60405180910390f35b6101656101c43660046115ea565b610670565b6101656101d736600461145f565b610876565b6101656101ea3660046114b5565b61092e565b6101656101fd3660046116f4565b6109e9565b6102297f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101ad565b610165610b6b565b6102297f000000000000000000000000000000000000000000000000000000000000000081565b610229610b7f565b610165610286366004611730565b610b98565b610165610299366004611783565b610cfd565b610165610d53565b6101656102b43660046114b5565b611022565b6101656102c73660046118b6565b611135565b6101656102da36600461145f565b6111b6565b6102297f000000000000000000000000000000000000000000000000000000000000000081565b6103296103143660046114b5565b60656020526000908152604090205460ff1681565b60405190151581526020016101ad565b6101a0611227565b6102297f000000000000000000000000000000000000000000000000000000000000000081565b6101656103763660046114b5565b6112af565b3360009081526065602052604090205460ff166103b35760405162461bcd60e51b81526004016103aa906119d1565b60405180910390fd5b604051630e9f564b60e01b815233600482015263ffffffff821660248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690630e9f564b90604401600060405180830381600087803b15801561042057600080fd5b505af1158015610434573d6000803e3d6000fd5b505060405163ffffffff841692503391507fc839611458a2fab2b8b94182f828cd8d886dc80a56b20f644bd651d3e128c6d890600090a350565b610476611328565b60005b818110156104e85760016065600085858581811061049957610499611a08565b90506020020160208101906104ae91906114b5565b6001600160a01b031681526020810191909152604001600020805460ff1916911515919091179055806104e081611a1e565b915050610479565b505050565b3360009081526065602052604090205460ff1661051c5760405162461bcd60e51b81526004016103aa906119d1565b604051633588e1c760e11b815233600482015263ffffffff821660248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690636b11c38e90604401600060405180830381600087803b15801561058957600080fd5b505af115801561059d573d6000803e3d6000fd5b505060405163ffffffff841692503391507fa7aebf234bd4af17c69e39555c1da417a31714b8efc90375f4019f3024bbf4bc90600090a350565b6040516356bf7c2560e01b81526001600160a01b0382811660048301526060917f0000000000000000000000000000000000000000000000000000000000000000909116906356bf7c2590602401600060405180830381865afa158015610642573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261066a9190810190611a47565b92915050565b3360009081526065602052604090205460ff1661069f5760405162461bcd60e51b81526004016103aa906119d1565b6040516316c72b2f60e01b815233906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906316c72b2f906106f190849088908890600401611bec565b600060405180830381600087803b15801561070b57600080fd5b505af115801561071f573d6000803e3d6000fd5b50506040516316fd18dd60e11b81526001600160a01b03848116600483015263ffffffff6024830181905293507f0000000000000000000000000000000000000000000000000000000000000000169150632dfa31ba90604401600060405180830381600087803b15801561079357600080fd5b505af11580156107a7573d6000803e3d6000fd5b5050604051639926ee7d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169250639926ee7d91506107f99085908790600401611c21565b600060405180830381600087803b15801561081357600080fd5b505af1158015610827573d6000803e3d6000fd5b505060405163ffffffff841681526001600160a01b03851692507f3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a915060200160405180910390a25050505050565b3360009081526065602052604090205460ff166108a55760405162461bcd60e51b81526004016103aa906119d1565b604051630315ee7560e21b815233906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690630c57b9d4906108f790849087908790600401611ca1565b600060405180830381600087803b15801561091157600080fd5b505af1158015610925573d6000803e3d6000fd5b50505050505050565b3360009081526065602052604090205460ff1661095d5760405162461bcd60e51b81526004016103aa906119d1565b604051630586d8e360e31b815233600482018190526001600160a01b03838116602484015290917f000000000000000000000000000000000000000000000000000000000000000090911690632c36c718906044015b600060405180830381600087803b1580156109cd57600080fd5b505af11580156109e1573d6000803e3d6000fd5b505050505050565b3360009081526065602052604090205460ff16610a185760405162461bcd60e51b81526004016103aa906119d1565b60405163cfedf42760e01b815233906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063cfedf42790610a689084908690600401611cfd565b600060405180830381600087803b158015610a8257600080fd5b505af1158015610a96573d6000803e3d6000fd5b50506040516316fd18dd60e11b81526001600160a01b03848116600483015263ffffffff6024830181905293507f0000000000000000000000000000000000000000000000000000000000000000169150632dfa31ba90604401600060405180830381600087803b158015610b0a57600080fd5b505af1158015610b1e573d6000803e3d6000fd5b505060405163ffffffff841681526001600160a01b03851692507f3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a915060200160405180910390a2505050565b610b73611328565b610b7d6000611387565b565b6000610b936033546001600160a01b031690565b905090565b610ba0611328565b4661426814610bf15760405162461bcd60e51b815260206004820152601f60248201527f4f6e6c7920486f6c65736b7920746573746e657420697320616c6c6f7765640060448201526064016103aa565b6040516349ad4db560e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063935a9b6a90610c4190869086908690600401611d21565b600060405180830381600087803b158015610c5b57600080fd5b505af1158015610c6f573d6000803e3d6000fd5b5084925060009150505b81811015610cf6578263ffffffff16858583818110610c9a57610c9a611a08565b9050602002016020810190610caf91906114b5565b6001600160a01b03167f2bb1cce1818b83dc9a614be5ae7471f32f000b48764252d856e61dd2c439d97f60405160405180910390a380610cee81611a1e565b915050610c79565b5050505050565b610d05611328565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb355906109b39085908590600401611d7e565b3360009081526065602052604090205460ff16610d825760405162461bcd60e51b81526004016103aa906119d1565b6040516319a74c5f60e01b815233600482018190529060009081907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906319a74c5f906024016040805180830381865afa158015610ded573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e119190611dad565b9150915081610e715760405162461bcd60e51b815260206004820152602660248201527f546865206f70657261746f72206973206e6f742061626c6520746f206465726560448201526533b4b9ba32b960d11b60648201526084016103aa565b6040516316fd18dd60e11b81526001600160a01b038481166004830152602482018390527f00000000000000000000000000000000000000000000000000000000000000001690632dfa31ba90604401600060405180830381600087803b158015610edb57600080fd5b505af1158015610eef573d6000803e3d6000fd5b50506040516356452c2560e11b81526001600160a01b0386811660048301527f000000000000000000000000000000000000000000000000000000000000000016925063ac8a584a9150602401600060405180830381600087803b158015610f5657600080fd5b505af1158015610f6a573d6000803e3d6000fd5b50506040516351b27a6d60e11b81526001600160a01b0386811660048301527f000000000000000000000000000000000000000000000000000000000000000016925063a364f4da9150602401600060405180830381600087803b158015610fd157600080fd5b505af1158015610fe5573d6000803e3d6000fd5b50506040516001600160a01b03861692507f6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d9150600090a2505050565b600054610100900460ff16158080156110425750600054600160ff909116105b8061105c5750303b15801561105c575060005460ff166001145b6110bf5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103aa565b6000805460ff1916600117905580156110e2576000805461ff0019166101001790555b6110eb82611387565b8015611131576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b3360009081526065602052604090205460ff166111645760405162461bcd60e51b81526004016103aa906119d1565b60405163829c7dd760e01b815233906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063829c7dd7906108f790849087908790600401611e37565b6111be611328565b60005b818110156104e857606560008484848181106111df576111df611a08565b90506020020160208101906111f491906114b5565b6001600160a01b031681526020810191909152604001600020805460ff191690558061121f81611a1e565b9150506111c1565b60607f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663273cbaa06040518163ffffffff1660e01b8152600401600060405180830381865afa158015611287573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b939190810190611a47565b6112b7611328565b6001600160a01b03811661131c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103aa565b61132581611387565b50565b33611331610b7f565b6001600160a01b031614610b7d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103aa565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b803563ffffffff811681146113ed57600080fd5b919050565b60006020828403121561140457600080fd5b61140d826113d9565b9392505050565b60008083601f84011261142657600080fd5b5081356001600160401b0381111561143d57600080fd5b6020830191508360208260051b850101111561145857600080fd5b9250929050565b6000806020838503121561147257600080fd5b82356001600160401b0381111561148857600080fd5b61149485828601611414565b90969095509350505050565b6001600160a01b038116811461132557600080fd5b6000602082840312156114c757600080fd5b813561140d816114a0565b6020808252825182820181905260009190848201906040850190845b818110156115135783516001600160a01b0316835292840192918401916001016114ee565b50909695505050505050565b6000610120828403121561153257600080fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b038111828210171561157057611570611538565b60405290565b60405160a081016001600160401b038111828210171561157057611570611538565b604080519081016001600160401b038111828210171561157057611570611538565b604051601f8201601f191681016001600160401b03811182821017156115e2576115e2611538565b604052919050565b6000806000606084860312156115ff57600080fd5b833561160a816114a0565b92506020848101356001600160401b038082111561162757600080fd5b6116338883890161151f565b9450604087013591508082111561164957600080fd5b908601906060828903121561165d57600080fd5b61166561154e565b82358281111561167457600080fd5b8301601f81018a1361168557600080fd5b80358381111561169757611697611538565b6116a9601f8201601f191687016115ba565b93508084528a868284010111156116bf57600080fd5b808683018786013760008682860101525050818152838301358482015260408301356040820152809450505050509250925092565b60006020828403121561170657600080fd5b81356001600160401b0381111561171c57600080fd5b6117288482850161151f565b949350505050565b60008060006040848603121561174557600080fd5b83356001600160401b0381111561175b57600080fd5b61176786828701611414565b909450925061177a9050602085016113d9565b90509250925092565b6000806020838503121561179657600080fd5b82356001600160401b03808211156117ad57600080fd5b818501915085601f8301126117c157600080fd5b8135818111156117d057600080fd5b8660208285010111156117e257600080fd5b60209290920196919550909350505050565b60006001600160401b0382111561180d5761180d611538565b5060051b60200190565b600082601f83011261182857600080fd5b611830611598565b80604084018581111561184257600080fd5b845b8181101561185c578035845260209384019301611844565b509095945050505050565b600082601f83011261187857600080fd5b611880611598565b80608084018581111561189257600080fd5b845b8181101561185c576118a68782611817565b8452602090930192604001611894565b60008060408084860312156118ca57600080fd5b6118d3846113d9565b92506020808501356001600160401b03808211156118f057600080fd5b90860190610120828903121561190557600080fd5b61190d611576565b82358281111561191c57600080fd5b83019150601f8201891361192f57600080fd5b813561194261193d826117f4565b6115ba565b81815260069190911b8301850190858101908b83111561196157600080fd5b938601935b82851015611987576119788c86611817565b82529387019390860190611966565b835250611998905089848601611867565b848201526119a98960a08501611817565b8582015260e08301356060820152610100830135608082015280955050505050509250929050565b6020808252601b908201527f4f70657261746f72206973206e6f742077686974656c69737465640000000000604082015260600190565b634e487b7160e01b600052603260045260246000fd5b6000600019821415611a4057634e487b7160e01b600052601160045260246000fd5b5060010190565b60006020808385031215611a5a57600080fd5b82516001600160401b03811115611a7057600080fd5b8301601f81018513611a8157600080fd5b8051611a8f61193d826117f4565b81815260059190911b82018301908381019087831115611aae57600080fd5b928401925b82841015611ad5578351611ac6816114a0565b82529284019290840190611ab3565b979650505050505050565b6040818337506000604082015250565b8183526020830192506000816000805b85811015611b2257604080848937968701828152969290920191600101611b00565b509495945050505050565b806000805b6002811015610cf657604080848737948501828152949290920191600101611b32565b60006101208235601e19843603018112611b6e57600080fd5b830180356001600160401b03811115611b8657600080fd5b8060061b3603851315611b9857600080fd5b828652611bab8387018260208501611af0565b92505050611bbf6020850160208501611b2d565b611bcf60a0850160a08501611ae0565b60e083810135908501526101009283013592909301919091525090565b6001600160a01b03848116825283166020820152606060408201819052600090611c1890830184611b55565b95945050505050565b60018060a01b03831681526000602060408184015283516060604085015280518060a086015260005b81811015611c665782810184015186820160c001528301611c4a565b81811115611c7857600060c083880101525b509185015160608501525060409093015160808301525060c0601f909201601f19160101919050565b6001600160a01b038416815260406020808301829052908201839052600090849060608401835b86811015611cf15763ffffffff611cde856113d9565b1682529282019290820190600101611cc8565b50979650505050505050565b6001600160a01b038316815260406020820181905260009061172890830184611b55565b6040808252810183905260008460608301825b86811015611d64578235611d47816114a0565b6001600160a01b0316825260209283019290910190600101611d34565b50809250505063ffffffff83166020830152949350505050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60008060408385031215611dc057600080fd5b82518015158114611dd057600080fd5b6020939093015192949293505050565b8060005b6002811015611e03578151845260209384019390910190600101611de4565b50505050565b8060005b6002811015611e0357611e21848351611de0565b6040939093019260209190910190600101611e0d565b60018060a01b03841681526000602063ffffffff85168184015260406060818501526101808401855161012060608701528181518084526101a0880191508583019350600092505b80831015611ea657611e92828551611de0565b928501926001929092019190840190611e7f565b509387015193611eb96080880186611e09565b838801519450611ecd610100880186611de0565b60608801516101408801526080880151610160880152809550505050505094935050505056fea2646970667358221220a486c798bb7efa82554b35c21bd40bb0d69142b569e2b26388e67873a3a4a09e64736f6c634300080c0033",
}

// LagrangeABI is the input ABI used to generate the binding from.
// Deprecated: Use LagrangeMetaData.ABI instead.
var LagrangeABI = LagrangeMetaData.ABI

// LagrangeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LagrangeMetaData.Bin instead.
var LagrangeBin = LagrangeMetaData.Bin

// DeployLagrange deploys a new Ethereum contract, binding an instance of Lagrange to it.
func DeployLagrange(auth *bind.TransactOpts, backend bind.ContractBackend, _committee common.Address, _stakeManager common.Address, _avsDirectoryAddress common.Address, _voteWeigher common.Address) (common.Address, *types.Transaction, *Lagrange, error) {
	parsed, err := LagrangeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LagrangeBin), backend, _committee, _stakeManager, _avsDirectoryAddress, _voteWeigher)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lagrange{LagrangeCaller: LagrangeCaller{contract: contract}, LagrangeTransactor: LagrangeTransactor{contract: contract}, LagrangeFilterer: LagrangeFilterer{contract: contract}}, nil
}

// Lagrange is an auto generated Go binding around an Ethereum contract.
type Lagrange struct {
	LagrangeCaller     // Read-only binding to the contract
	LagrangeTransactor // Write-only binding to the contract
	LagrangeFilterer   // Log filterer for contract events
}

// LagrangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type LagrangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LagrangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LagrangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LagrangeSession struct {
	Contract     *Lagrange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LagrangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LagrangeCallerSession struct {
	Contract *LagrangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LagrangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LagrangeTransactorSession struct {
	Contract     *LagrangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LagrangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type LagrangeRaw struct {
	Contract *Lagrange // Generic contract binding to access the raw methods on
}

// LagrangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LagrangeCallerRaw struct {
	Contract *LagrangeCaller // Generic read-only contract binding to access the raw methods on
}

// LagrangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LagrangeTransactorRaw struct {
	Contract *LagrangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLagrange creates a new instance of Lagrange, bound to a specific deployed contract.
func NewLagrange(address common.Address, backend bind.ContractBackend) (*Lagrange, error) {
	contract, err := bindLagrange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lagrange{LagrangeCaller: LagrangeCaller{contract: contract}, LagrangeTransactor: LagrangeTransactor{contract: contract}, LagrangeFilterer: LagrangeFilterer{contract: contract}}, nil
}

// NewLagrangeCaller creates a new read-only instance of Lagrange, bound to a specific deployed contract.
func NewLagrangeCaller(address common.Address, caller bind.ContractCaller) (*LagrangeCaller, error) {
	contract, err := bindLagrange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LagrangeCaller{contract: contract}, nil
}

// NewLagrangeTransactor creates a new write-only instance of Lagrange, bound to a specific deployed contract.
func NewLagrangeTransactor(address common.Address, transactor bind.ContractTransactor) (*LagrangeTransactor, error) {
	contract, err := bindLagrange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LagrangeTransactor{contract: contract}, nil
}

// NewLagrangeFilterer creates a new log filterer instance of Lagrange, bound to a specific deployed contract.
func NewLagrangeFilterer(address common.Address, filterer bind.ContractFilterer) (*LagrangeFilterer, error) {
	contract, err := bindLagrange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LagrangeFilterer{contract: contract}, nil
}

// bindLagrange binds a generic wrapper to an already deployed contract.
func bindLagrange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LagrangeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lagrange *LagrangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lagrange.Contract.LagrangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lagrange *LagrangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lagrange.Contract.LagrangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lagrange *LagrangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lagrange.Contract.LagrangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lagrange *LagrangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lagrange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lagrange *LagrangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lagrange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lagrange *LagrangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lagrange.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_Lagrange *LagrangeCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_Lagrange *LagrangeSession) AvsDirectory() (common.Address, error) {
	return _Lagrange.Contract.AvsDirectory(&_Lagrange.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_Lagrange *LagrangeCallerSession) AvsDirectory() (common.Address, error) {
	return _Lagrange.Contract.AvsDirectory(&_Lagrange.CallOpts)
}

// Committee is a free data retrieval call binding the contract method 0xd864e740.
//
// Solidity: function committee() view returns(address)
func (_Lagrange *LagrangeCaller) Committee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "committee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Committee is a free data retrieval call binding the contract method 0xd864e740.
//
// Solidity: function committee() view returns(address)
func (_Lagrange *LagrangeSession) Committee() (common.Address, error) {
	return _Lagrange.Contract.Committee(&_Lagrange.CallOpts)
}

// Committee is a free data retrieval call binding the contract method 0xd864e740.
//
// Solidity: function committee() view returns(address)
func (_Lagrange *LagrangeCallerSession) Committee() (common.Address, error) {
	return _Lagrange.Contract.Committee(&_Lagrange.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_Lagrange *LagrangeCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_Lagrange *LagrangeSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _Lagrange.Contract.GetOperatorRestakedStrategies(&_Lagrange.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_Lagrange *LagrangeCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _Lagrange.Contract.GetOperatorRestakedStrategies(&_Lagrange.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_Lagrange *LagrangeCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_Lagrange *LagrangeSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _Lagrange.Contract.GetRestakeableStrategies(&_Lagrange.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_Lagrange *LagrangeCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _Lagrange.Contract.GetRestakeableStrategies(&_Lagrange.CallOpts)
}

// OperatorWhitelist is a free data retrieval call binding the contract method 0xe03c8632.
//
// Solidity: function operatorWhitelist(address ) view returns(bool)
func (_Lagrange *LagrangeCaller) OperatorWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "operatorWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorWhitelist is a free data retrieval call binding the contract method 0xe03c8632.
//
// Solidity: function operatorWhitelist(address ) view returns(bool)
func (_Lagrange *LagrangeSession) OperatorWhitelist(arg0 common.Address) (bool, error) {
	return _Lagrange.Contract.OperatorWhitelist(&_Lagrange.CallOpts, arg0)
}

// OperatorWhitelist is a free data retrieval call binding the contract method 0xe03c8632.
//
// Solidity: function operatorWhitelist(address ) view returns(bool)
func (_Lagrange *LagrangeCallerSession) OperatorWhitelist(arg0 common.Address) (bool, error) {
	return _Lagrange.Contract.OperatorWhitelist(&_Lagrange.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lagrange *LagrangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lagrange *LagrangeSession) Owner() (common.Address, error) {
	return _Lagrange.Contract.Owner(&_Lagrange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lagrange *LagrangeCallerSession) Owner() (common.Address, error) {
	return _Lagrange.Contract.Owner(&_Lagrange.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_Lagrange *LagrangeCaller) StakeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "stakeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_Lagrange *LagrangeSession) StakeManager() (common.Address, error) {
	return _Lagrange.Contract.StakeManager(&_Lagrange.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_Lagrange *LagrangeCallerSession) StakeManager() (common.Address, error) {
	return _Lagrange.Contract.StakeManager(&_Lagrange.CallOpts)
}

// VoteWeigher is a free data retrieval call binding the contract method 0xef030673.
//
// Solidity: function voteWeigher() view returns(address)
func (_Lagrange *LagrangeCaller) VoteWeigher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "voteWeigher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteWeigher is a free data retrieval call binding the contract method 0xef030673.
//
// Solidity: function voteWeigher() view returns(address)
func (_Lagrange *LagrangeSession) VoteWeigher() (common.Address, error) {
	return _Lagrange.Contract.VoteWeigher(&_Lagrange.CallOpts)
}

// VoteWeigher is a free data retrieval call binding the contract method 0xef030673.
//
// Solidity: function voteWeigher() view returns(address)
func (_Lagrange *LagrangeCallerSession) VoteWeigher() (common.Address, error) {
	return _Lagrange.Contract.VoteWeigher(&_Lagrange.CallOpts)
}

// AddBlsPubKeys is a paid mutator transaction binding the contract method 0x63d944ea.
//
// Solidity: function addBlsPubKeys((uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof) returns()
func (_Lagrange *LagrangeTransactor) AddBlsPubKeys(opts *bind.TransactOpts, blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "addBlsPubKeys", blsKeyWithProof)
}

// AddBlsPubKeys is a paid mutator transaction binding the contract method 0x63d944ea.
//
// Solidity: function addBlsPubKeys((uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof) returns()
func (_Lagrange *LagrangeSession) AddBlsPubKeys(blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof) (*types.Transaction, error) {
	return _Lagrange.Contract.AddBlsPubKeys(&_Lagrange.TransactOpts, blsKeyWithProof)
}

// AddBlsPubKeys is a paid mutator transaction binding the contract method 0x63d944ea.
//
// Solidity: function addBlsPubKeys((uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof) returns()
func (_Lagrange *LagrangeTransactorSession) AddBlsPubKeys(blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof) (*types.Transaction, error) {
	return _Lagrange.Contract.AddBlsPubKeys(&_Lagrange.TransactOpts, blsKeyWithProof)
}

// AddOperatorsToWhitelist is a paid mutator transaction binding the contract method 0x1be4b9f7.
//
// Solidity: function addOperatorsToWhitelist(address[] operators) returns()
func (_Lagrange *LagrangeTransactor) AddOperatorsToWhitelist(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "addOperatorsToWhitelist", operators)
}

// AddOperatorsToWhitelist is a paid mutator transaction binding the contract method 0x1be4b9f7.
//
// Solidity: function addOperatorsToWhitelist(address[] operators) returns()
func (_Lagrange *LagrangeSession) AddOperatorsToWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.AddOperatorsToWhitelist(&_Lagrange.TransactOpts, operators)
}

// AddOperatorsToWhitelist is a paid mutator transaction binding the contract method 0x1be4b9f7.
//
// Solidity: function addOperatorsToWhitelist(address[] operators) returns()
func (_Lagrange *LagrangeTransactorSession) AddOperatorsToWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.AddOperatorsToWhitelist(&_Lagrange.TransactOpts, operators)
}

// Deregister is a paid mutator transaction binding the contract method 0xaff5edb1.
//
// Solidity: function deregister() returns()
func (_Lagrange *LagrangeTransactor) Deregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "deregister")
}

// Deregister is a paid mutator transaction binding the contract method 0xaff5edb1.
//
// Solidity: function deregister() returns()
func (_Lagrange *LagrangeSession) Deregister() (*types.Transaction, error) {
	return _Lagrange.Contract.Deregister(&_Lagrange.TransactOpts)
}

// Deregister is a paid mutator transaction binding the contract method 0xaff5edb1.
//
// Solidity: function deregister() returns()
func (_Lagrange *LagrangeTransactorSession) Deregister() (*types.Transaction, error) {
	return _Lagrange.Contract.Deregister(&_Lagrange.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Lagrange *LagrangeTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Lagrange *LagrangeSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.Initialize(&_Lagrange.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Lagrange *LagrangeTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.Initialize(&_Lagrange.TransactOpts, initialOwner)
}

// Register is a paid mutator transaction binding the contract method 0x4e32e904.
//
// Solidity: function register(address signAddress, (uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof, (bytes,bytes32,uint256) operatorSignature) returns()
func (_Lagrange *LagrangeTransactor) Register(opts *bind.TransactOpts, signAddress common.Address, blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "register", signAddress, blsKeyWithProof, operatorSignature)
}

// Register is a paid mutator transaction binding the contract method 0x4e32e904.
//
// Solidity: function register(address signAddress, (uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof, (bytes,bytes32,uint256) operatorSignature) returns()
func (_Lagrange *LagrangeSession) Register(signAddress common.Address, blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Lagrange.Contract.Register(&_Lagrange.TransactOpts, signAddress, blsKeyWithProof, operatorSignature)
}

// Register is a paid mutator transaction binding the contract method 0x4e32e904.
//
// Solidity: function register(address signAddress, (uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof, (bytes,bytes32,uint256) operatorSignature) returns()
func (_Lagrange *LagrangeTransactorSession) Register(signAddress common.Address, blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Lagrange.Contract.Register(&_Lagrange.TransactOpts, signAddress, blsKeyWithProof, operatorSignature)
}

// RemoveBlsPubKeys is a paid mutator transaction binding the contract method 0x5e332fe4.
//
// Solidity: function removeBlsPubKeys(uint32[] indices) returns()
func (_Lagrange *LagrangeTransactor) RemoveBlsPubKeys(opts *bind.TransactOpts, indices []uint32) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "removeBlsPubKeys", indices)
}

// RemoveBlsPubKeys is a paid mutator transaction binding the contract method 0x5e332fe4.
//
// Solidity: function removeBlsPubKeys(uint32[] indices) returns()
func (_Lagrange *LagrangeSession) RemoveBlsPubKeys(indices []uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.RemoveBlsPubKeys(&_Lagrange.TransactOpts, indices)
}

// RemoveBlsPubKeys is a paid mutator transaction binding the contract method 0x5e332fe4.
//
// Solidity: function removeBlsPubKeys(uint32[] indices) returns()
func (_Lagrange *LagrangeTransactorSession) RemoveBlsPubKeys(indices []uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.RemoveBlsPubKeys(&_Lagrange.TransactOpts, indices)
}

// RemoveOperatorsFromWhitelist is a paid mutator transaction binding the contract method 0xd822af3c.
//
// Solidity: function removeOperatorsFromWhitelist(address[] operators) returns()
func (_Lagrange *LagrangeTransactor) RemoveOperatorsFromWhitelist(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "removeOperatorsFromWhitelist", operators)
}

// RemoveOperatorsFromWhitelist is a paid mutator transaction binding the contract method 0xd822af3c.
//
// Solidity: function removeOperatorsFromWhitelist(address[] operators) returns()
func (_Lagrange *LagrangeSession) RemoveOperatorsFromWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.RemoveOperatorsFromWhitelist(&_Lagrange.TransactOpts, operators)
}

// RemoveOperatorsFromWhitelist is a paid mutator transaction binding the contract method 0xd822af3c.
//
// Solidity: function removeOperatorsFromWhitelist(address[] operators) returns()
func (_Lagrange *LagrangeTransactorSession) RemoveOperatorsFromWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.RemoveOperatorsFromWhitelist(&_Lagrange.TransactOpts, operators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lagrange *LagrangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lagrange *LagrangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lagrange.Contract.RenounceOwnership(&_Lagrange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lagrange *LagrangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lagrange.Contract.RenounceOwnership(&_Lagrange.TransactOpts)
}

// Subscribe is a paid mutator transaction binding the contract method 0x2e94d67b.
//
// Solidity: function subscribe(uint32 chainID) returns()
func (_Lagrange *LagrangeTransactor) Subscribe(opts *bind.TransactOpts, chainID uint32) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "subscribe", chainID)
}

// Subscribe is a paid mutator transaction binding the contract method 0x2e94d67b.
//
// Solidity: function subscribe(uint32 chainID) returns()
func (_Lagrange *LagrangeSession) Subscribe(chainID uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.Subscribe(&_Lagrange.TransactOpts, chainID)
}

// Subscribe is a paid mutator transaction binding the contract method 0x2e94d67b.
//
// Solidity: function subscribe(uint32 chainID) returns()
func (_Lagrange *LagrangeTransactorSession) Subscribe(chainID uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.Subscribe(&_Lagrange.TransactOpts, chainID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lagrange *LagrangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lagrange *LagrangeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.TransferOwnership(&_Lagrange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lagrange *LagrangeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.TransferOwnership(&_Lagrange.TransactOpts, newOwner)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0x0512d04c.
//
// Solidity: function unsubscribe(uint32 chainID) returns()
func (_Lagrange *LagrangeTransactor) Unsubscribe(opts *bind.TransactOpts, chainID uint32) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "unsubscribe", chainID)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0x0512d04c.
//
// Solidity: function unsubscribe(uint32 chainID) returns()
func (_Lagrange *LagrangeSession) Unsubscribe(chainID uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.Unsubscribe(&_Lagrange.TransactOpts, chainID)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0x0512d04c.
//
// Solidity: function unsubscribe(uint32 chainID) returns()
func (_Lagrange *LagrangeTransactorSession) Unsubscribe(chainID uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.Unsubscribe(&_Lagrange.TransactOpts, chainID)
}

// UnsubscribeByAdmin is a paid mutator transaction binding the contract method 0x935a9b6a.
//
// Solidity: function unsubscribeByAdmin(address[] operators, uint32 chainID) returns()
func (_Lagrange *LagrangeTransactor) UnsubscribeByAdmin(opts *bind.TransactOpts, operators []common.Address, chainID uint32) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "unsubscribeByAdmin", operators, chainID)
}

// UnsubscribeByAdmin is a paid mutator transaction binding the contract method 0x935a9b6a.
//
// Solidity: function unsubscribeByAdmin(address[] operators, uint32 chainID) returns()
func (_Lagrange *LagrangeSession) UnsubscribeByAdmin(operators []common.Address, chainID uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.UnsubscribeByAdmin(&_Lagrange.TransactOpts, operators, chainID)
}

// UnsubscribeByAdmin is a paid mutator transaction binding the contract method 0x935a9b6a.
//
// Solidity: function unsubscribeByAdmin(address[] operators, uint32 chainID) returns()
func (_Lagrange *LagrangeTransactorSession) UnsubscribeByAdmin(operators []common.Address, chainID uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.UnsubscribeByAdmin(&_Lagrange.TransactOpts, operators, chainID)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_Lagrange *LagrangeTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_Lagrange *LagrangeSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _Lagrange.Contract.UpdateAVSMetadataURI(&_Lagrange.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_Lagrange *LagrangeTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _Lagrange.Contract.UpdateAVSMetadataURI(&_Lagrange.TransactOpts, _metadataURI)
}

// UpdateBlsPubKey is a paid mutator transaction binding the contract method 0xc5bdc30a.
//
// Solidity: function updateBlsPubKey(uint32 index, (uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof) returns()
func (_Lagrange *LagrangeTransactor) UpdateBlsPubKey(opts *bind.TransactOpts, index uint32, blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "updateBlsPubKey", index, blsKeyWithProof)
}

// UpdateBlsPubKey is a paid mutator transaction binding the contract method 0xc5bdc30a.
//
// Solidity: function updateBlsPubKey(uint32 index, (uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof) returns()
func (_Lagrange *LagrangeSession) UpdateBlsPubKey(index uint32, blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof) (*types.Transaction, error) {
	return _Lagrange.Contract.UpdateBlsPubKey(&_Lagrange.TransactOpts, index, blsKeyWithProof)
}

// UpdateBlsPubKey is a paid mutator transaction binding the contract method 0xc5bdc30a.
//
// Solidity: function updateBlsPubKey(uint32 index, (uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof) returns()
func (_Lagrange *LagrangeTransactorSession) UpdateBlsPubKey(index uint32, blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof) (*types.Transaction, error) {
	return _Lagrange.Contract.UpdateBlsPubKey(&_Lagrange.TransactOpts, index, blsKeyWithProof)
}

// UpdateSignAddress is a paid mutator transaction binding the contract method 0x5fa7306a.
//
// Solidity: function updateSignAddress(address newSignAddress) returns()
func (_Lagrange *LagrangeTransactor) UpdateSignAddress(opts *bind.TransactOpts, newSignAddress common.Address) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "updateSignAddress", newSignAddress)
}

// UpdateSignAddress is a paid mutator transaction binding the contract method 0x5fa7306a.
//
// Solidity: function updateSignAddress(address newSignAddress) returns()
func (_Lagrange *LagrangeSession) UpdateSignAddress(newSignAddress common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.UpdateSignAddress(&_Lagrange.TransactOpts, newSignAddress)
}

// UpdateSignAddress is a paid mutator transaction binding the contract method 0x5fa7306a.
//
// Solidity: function updateSignAddress(address newSignAddress) returns()
func (_Lagrange *LagrangeTransactorSession) UpdateSignAddress(newSignAddress common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.UpdateSignAddress(&_Lagrange.TransactOpts, newSignAddress)
}

// LagrangeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Lagrange contract.
type LagrangeInitializedIterator struct {
	Event *LagrangeInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LagrangeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeInitialized)
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
		it.Event = new(LagrangeInitialized)
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
func (it *LagrangeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeInitialized represents a Initialized event raised by the Lagrange contract.
type LagrangeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Lagrange *LagrangeFilterer) FilterInitialized(opts *bind.FilterOpts) (*LagrangeInitializedIterator, error) {

	logs, sub, err := _Lagrange.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LagrangeInitializedIterator{contract: _Lagrange.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Lagrange *LagrangeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LagrangeInitialized) (event.Subscription, error) {

	logs, sub, err := _Lagrange.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeInitialized)
				if err := _Lagrange.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Lagrange *LagrangeFilterer) ParseInitialized(log types.Log) (*LagrangeInitialized, error) {
	event := new(LagrangeInitialized)
	if err := _Lagrange.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the Lagrange contract.
type LagrangeOperatorDeregisteredIterator struct {
	Event *LagrangeOperatorDeregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LagrangeOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeOperatorDeregistered)
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
		it.Event = new(LagrangeOperatorDeregistered)
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
func (it *LagrangeOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeOperatorDeregistered represents a OperatorDeregistered event raised by the Lagrange contract.
type LagrangeOperatorDeregistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d.
//
// Solidity: event OperatorDeregistered(address indexed operator)
func (_Lagrange *LagrangeFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address) (*LagrangeOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Lagrange.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeOperatorDeregisteredIterator{contract: _Lagrange.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d.
//
// Solidity: event OperatorDeregistered(address indexed operator)
func (_Lagrange *LagrangeFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *LagrangeOperatorDeregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Lagrange.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeOperatorDeregistered)
				if err := _Lagrange.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d.
//
// Solidity: event OperatorDeregistered(address indexed operator)
func (_Lagrange *LagrangeFilterer) ParseOperatorDeregistered(log types.Log) (*LagrangeOperatorDeregistered, error) {
	event := new(LagrangeOperatorDeregistered)
	if err := _Lagrange.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the Lagrange contract.
type LagrangeOperatorRegisteredIterator struct {
	Event *LagrangeOperatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LagrangeOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeOperatorRegistered)
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
		it.Event = new(LagrangeOperatorRegistered)
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
func (it *LagrangeOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeOperatorRegistered represents a OperatorRegistered event raised by the Lagrange contract.
type LagrangeOperatorRegistered struct {
	Operator        common.Address
	ServeUntilBlock uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32 serveUntilBlock)
func (_Lagrange *LagrangeFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*LagrangeOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Lagrange.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeOperatorRegisteredIterator{contract: _Lagrange.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32 serveUntilBlock)
func (_Lagrange *LagrangeFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *LagrangeOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Lagrange.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeOperatorRegistered)
				if err := _Lagrange.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32 serveUntilBlock)
func (_Lagrange *LagrangeFilterer) ParseOperatorRegistered(log types.Log) (*LagrangeOperatorRegistered, error) {
	event := new(LagrangeOperatorRegistered)
	if err := _Lagrange.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeOperatorSubscribedIterator is returned from FilterOperatorSubscribed and is used to iterate over the raw logs and unpacked data for OperatorSubscribed events raised by the Lagrange contract.
type LagrangeOperatorSubscribedIterator struct {
	Event *LagrangeOperatorSubscribed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LagrangeOperatorSubscribedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeOperatorSubscribed)
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
		it.Event = new(LagrangeOperatorSubscribed)
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
func (it *LagrangeOperatorSubscribedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeOperatorSubscribedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeOperatorSubscribed represents a OperatorSubscribed event raised by the Lagrange contract.
type LagrangeOperatorSubscribed struct {
	Operator common.Address
	ChainID  uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSubscribed is a free log retrieval operation binding the contract event 0xa7aebf234bd4af17c69e39555c1da417a31714b8efc90375f4019f3024bbf4bc.
//
// Solidity: event OperatorSubscribed(address indexed operator, uint32 indexed chainID)
func (_Lagrange *LagrangeFilterer) FilterOperatorSubscribed(opts *bind.FilterOpts, operator []common.Address, chainID []uint32) (*LagrangeOperatorSubscribedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _Lagrange.contract.FilterLogs(opts, "OperatorSubscribed", operatorRule, chainIDRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeOperatorSubscribedIterator{contract: _Lagrange.contract, event: "OperatorSubscribed", logs: logs, sub: sub}, nil
}

// WatchOperatorSubscribed is a free log subscription operation binding the contract event 0xa7aebf234bd4af17c69e39555c1da417a31714b8efc90375f4019f3024bbf4bc.
//
// Solidity: event OperatorSubscribed(address indexed operator, uint32 indexed chainID)
func (_Lagrange *LagrangeFilterer) WatchOperatorSubscribed(opts *bind.WatchOpts, sink chan<- *LagrangeOperatorSubscribed, operator []common.Address, chainID []uint32) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _Lagrange.contract.WatchLogs(opts, "OperatorSubscribed", operatorRule, chainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeOperatorSubscribed)
				if err := _Lagrange.contract.UnpackLog(event, "OperatorSubscribed", log); err != nil {
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

// ParseOperatorSubscribed is a log parse operation binding the contract event 0xa7aebf234bd4af17c69e39555c1da417a31714b8efc90375f4019f3024bbf4bc.
//
// Solidity: event OperatorSubscribed(address indexed operator, uint32 indexed chainID)
func (_Lagrange *LagrangeFilterer) ParseOperatorSubscribed(log types.Log) (*LagrangeOperatorSubscribed, error) {
	event := new(LagrangeOperatorSubscribed)
	if err := _Lagrange.contract.UnpackLog(event, "OperatorSubscribed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeOperatorUnsubscribedIterator is returned from FilterOperatorUnsubscribed and is used to iterate over the raw logs and unpacked data for OperatorUnsubscribed events raised by the Lagrange contract.
type LagrangeOperatorUnsubscribedIterator struct {
	Event *LagrangeOperatorUnsubscribed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LagrangeOperatorUnsubscribedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeOperatorUnsubscribed)
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
		it.Event = new(LagrangeOperatorUnsubscribed)
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
func (it *LagrangeOperatorUnsubscribedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeOperatorUnsubscribedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeOperatorUnsubscribed represents a OperatorUnsubscribed event raised by the Lagrange contract.
type LagrangeOperatorUnsubscribed struct {
	Operator common.Address
	ChainID  uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorUnsubscribed is a free log retrieval operation binding the contract event 0xc839611458a2fab2b8b94182f828cd8d886dc80a56b20f644bd651d3e128c6d8.
//
// Solidity: event OperatorUnsubscribed(address indexed operator, uint32 indexed chainID)
func (_Lagrange *LagrangeFilterer) FilterOperatorUnsubscribed(opts *bind.FilterOpts, operator []common.Address, chainID []uint32) (*LagrangeOperatorUnsubscribedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _Lagrange.contract.FilterLogs(opts, "OperatorUnsubscribed", operatorRule, chainIDRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeOperatorUnsubscribedIterator{contract: _Lagrange.contract, event: "OperatorUnsubscribed", logs: logs, sub: sub}, nil
}

// WatchOperatorUnsubscribed is a free log subscription operation binding the contract event 0xc839611458a2fab2b8b94182f828cd8d886dc80a56b20f644bd651d3e128c6d8.
//
// Solidity: event OperatorUnsubscribed(address indexed operator, uint32 indexed chainID)
func (_Lagrange *LagrangeFilterer) WatchOperatorUnsubscribed(opts *bind.WatchOpts, sink chan<- *LagrangeOperatorUnsubscribed, operator []common.Address, chainID []uint32) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _Lagrange.contract.WatchLogs(opts, "OperatorUnsubscribed", operatorRule, chainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeOperatorUnsubscribed)
				if err := _Lagrange.contract.UnpackLog(event, "OperatorUnsubscribed", log); err != nil {
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

// ParseOperatorUnsubscribed is a log parse operation binding the contract event 0xc839611458a2fab2b8b94182f828cd8d886dc80a56b20f644bd651d3e128c6d8.
//
// Solidity: event OperatorUnsubscribed(address indexed operator, uint32 indexed chainID)
func (_Lagrange *LagrangeFilterer) ParseOperatorUnsubscribed(log types.Log) (*LagrangeOperatorUnsubscribed, error) {
	event := new(LagrangeOperatorUnsubscribed)
	if err := _Lagrange.contract.UnpackLog(event, "OperatorUnsubscribed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Lagrange contract.
type LagrangeOwnershipTransferredIterator struct {
	Event *LagrangeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LagrangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeOwnershipTransferred)
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
		it.Event = new(LagrangeOwnershipTransferred)
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
func (it *LagrangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeOwnershipTransferred represents a OwnershipTransferred event raised by the Lagrange contract.
type LagrangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lagrange *LagrangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LagrangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lagrange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeOwnershipTransferredIterator{contract: _Lagrange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lagrange *LagrangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LagrangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lagrange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeOwnershipTransferred)
				if err := _Lagrange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lagrange *LagrangeFilterer) ParseOwnershipTransferred(log types.Log) (*LagrangeOwnershipTransferred, error) {
	event := new(LagrangeOwnershipTransferred)
	if err := _Lagrange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeUnsubscribedByAdminIterator is returned from FilterUnsubscribedByAdmin and is used to iterate over the raw logs and unpacked data for UnsubscribedByAdmin events raised by the Lagrange contract.
type LagrangeUnsubscribedByAdminIterator struct {
	Event *LagrangeUnsubscribedByAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LagrangeUnsubscribedByAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeUnsubscribedByAdmin)
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
		it.Event = new(LagrangeUnsubscribedByAdmin)
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
func (it *LagrangeUnsubscribedByAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeUnsubscribedByAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeUnsubscribedByAdmin represents a UnsubscribedByAdmin event raised by the Lagrange contract.
type LagrangeUnsubscribedByAdmin struct {
	Operator common.Address
	ChainID  uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnsubscribedByAdmin is a free log retrieval operation binding the contract event 0x2bb1cce1818b83dc9a614be5ae7471f32f000b48764252d856e61dd2c439d97f.
//
// Solidity: event UnsubscribedByAdmin(address indexed operator, uint32 indexed chainID)
func (_Lagrange *LagrangeFilterer) FilterUnsubscribedByAdmin(opts *bind.FilterOpts, operator []common.Address, chainID []uint32) (*LagrangeUnsubscribedByAdminIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _Lagrange.contract.FilterLogs(opts, "UnsubscribedByAdmin", operatorRule, chainIDRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeUnsubscribedByAdminIterator{contract: _Lagrange.contract, event: "UnsubscribedByAdmin", logs: logs, sub: sub}, nil
}

// WatchUnsubscribedByAdmin is a free log subscription operation binding the contract event 0x2bb1cce1818b83dc9a614be5ae7471f32f000b48764252d856e61dd2c439d97f.
//
// Solidity: event UnsubscribedByAdmin(address indexed operator, uint32 indexed chainID)
func (_Lagrange *LagrangeFilterer) WatchUnsubscribedByAdmin(opts *bind.WatchOpts, sink chan<- *LagrangeUnsubscribedByAdmin, operator []common.Address, chainID []uint32) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _Lagrange.contract.WatchLogs(opts, "UnsubscribedByAdmin", operatorRule, chainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeUnsubscribedByAdmin)
				if err := _Lagrange.contract.UnpackLog(event, "UnsubscribedByAdmin", log); err != nil {
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

// ParseUnsubscribedByAdmin is a log parse operation binding the contract event 0x2bb1cce1818b83dc9a614be5ae7471f32f000b48764252d856e61dd2c439d97f.
//
// Solidity: event UnsubscribedByAdmin(address indexed operator, uint32 indexed chainID)
func (_Lagrange *LagrangeFilterer) ParseUnsubscribedByAdmin(log types.Log) (*LagrangeUnsubscribedByAdmin, error) {
	event := new(LagrangeUnsubscribedByAdmin)
	if err := _Lagrange.contract.UnpackLog(event, "UnsubscribedByAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
