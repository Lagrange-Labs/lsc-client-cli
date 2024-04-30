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
)

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// LagrangeMetaData contains all meta data concerning the Lagrange contract.
var LagrangeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_committee\",\"type\":\"address\",\"internalType\":\"contractILagrangeCommittee\"},{\"name\":\"_stakeManager\",\"type\":\"address\",\"internalType\":\"contractIStakeManager\"},{\"name\":\"_avsDirectoryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_voteWeigher\",\"type\":\"address\",\"internalType\":\"contractIVoteWeigher\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addBlsPubKeys\",\"inputs\":[{\"name\":\"additionalBlsPubKeys\",\"type\":\"uint256[2][]\",\"internalType\":\"uint256[2][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOperatorsToWhitelist\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"committee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILagrangeCommittee\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregister\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorWhitelist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"signAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blsPubKeys\",\"type\":\"uint256[2][]\",\"internalType\":\"uint256[2][]\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeBlsPubKeys\",\"inputs\":[{\"name\":\"indices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeOperatorsFromWhitelist\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"subscribe\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unsubscribe\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateBlsPubKey\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blsPubKey\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSignAddress\",\"inputs\":[{\"name\":\"newSignAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"voteWeigher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIVoteWeigher\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"serveUntilBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSubscribed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorUnsubscribed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106101425760003560e01c8063992b098e116100b8578063d822af3c1161007c578063d822af3c146102ae578063d864e740146102c1578063e03c8632146102e8578063e481af9d1461031b578063ef03067314610323578063f2fde38b1461034a57600080fd5b8063992b098e1461025a578063a98fb3551461026d578063aff5edb114610280578063c4d66de814610288578063c8d8e0131461029b57600080fd5b80635fa7306a1161010a5780635fa7306a146101be5780636b3aa72e146101d1578063715018a6146102105780637542ff95146102185780637a79ea331461023f5780638da5cb5b1461025257600080fd5b80630512d04c146101475780631be4b9f71461015c5780632e94d67b1461016f57806333cfb7b7146101825780635e332fe4146101ab575b600080fd5b61015a610155366004611275565b61035d565b005b61015a61016a3660046112e3565b610450565b61015a61017d366004611275565b6104cf565b61019561019036600461133a565b6105b9565b6040516101a29190611357565b60405180910390f35b61015a6101b93660046112e3565b610652565b61015a6101cc36600461133a565b61070a565b6101f87f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101a2565b61015a6107c5565b6101f87f000000000000000000000000000000000000000000000000000000000000000081565b61015a61024d366004611459565b6107d9565b6101f86109e2565b61015a610268366004611567565b6109fb565b61015a61027b36600461159d565b610b80565b61015a610bd6565b61015a61029636600461133a565b610ea5565b61015a6102a936600461160f565b610fb8565b61015a6102bc3660046112e3565b611039565b6101f87f000000000000000000000000000000000000000000000000000000000000000081565b61030b6102f636600461133a565b60656020526000908152604090205460ff1681565b60405190151581526020016101a2565b6101956110aa565b6101f87f000000000000000000000000000000000000000000000000000000000000000081565b61015a61035836600461133a565b611132565b3360009081526065602052604090205460ff166103955760405162461bcd60e51b815260040161038c90611649565b60405180910390fd5b604051630e9f564b60e01b815233600482015263ffffffff821660248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690630e9f564b90604401600060405180830381600087803b15801561040257600080fd5b505af1158015610416573d6000803e3d6000fd5b505060405163ffffffff841692503391507fc839611458a2fab2b8b94182f828cd8d886dc80a56b20f644bd651d3e128c6d890600090a350565b6104586111ab565b60005b818110156104ca5760016065600085858581811061047b5761047b611680565b9050602002016020810190610490919061133a565b6001600160a01b031681526020810191909152604001600020805460ff1916911515919091179055806104c281611696565b91505061045b565b505050565b3360009081526065602052604090205460ff166104fe5760405162461bcd60e51b815260040161038c90611649565b604051633588e1c760e11b815233600482015263ffffffff821660248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690636b11c38e90604401600060405180830381600087803b15801561056b57600080fd5b505af115801561057f573d6000803e3d6000fd5b505060405163ffffffff841692503391507fa7aebf234bd4af17c69e39555c1da417a31714b8efc90375f4019f3024bbf4bc90600090a350565b6040516356bf7c2560e01b81526001600160a01b0382811660048301526060917f0000000000000000000000000000000000000000000000000000000000000000909116906356bf7c2590602401600060405180830381865afa158015610624573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261064c91908101906116bf565b92915050565b3360009081526065602052604090205460ff166106815760405162461bcd60e51b815260040161038c90611649565b604051630315ee7560e21b815233906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690630c57b9d4906106d390849087908790600401611771565b600060405180830381600087803b1580156106ed57600080fd5b505af1158015610701573d6000803e3d6000fd5b50505050505050565b3360009081526065602052604090205460ff166107395760405162461bcd60e51b815260040161038c90611649565b604051630586d8e360e31b815233600482018190526001600160a01b03838116602484015290917f000000000000000000000000000000000000000000000000000000000000000090911690632c36c718906044015b600060405180830381600087803b1580156107a957600080fd5b505af11580156107bd573d6000803e3d6000fd5b505050505050565b6107cd6111ab565b6107d7600061120a565b565b3360009081526065602052604090205460ff166108085760405162461bcd60e51b815260040161038c90611649565b60405163c7e8a4f360e01b815233906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c7e8a4f39061085c90849089908990899060040161180a565b600060405180830381600087803b15801561087657600080fd5b505af115801561088a573d6000803e3d6000fd5b50506040516316fd18dd60e11b81526001600160a01b03848116600483015263ffffffff6024830181905293507f0000000000000000000000000000000000000000000000000000000000000000169150632dfa31ba90604401600060405180830381600087803b1580156108fe57600080fd5b505af1158015610912573d6000803e3d6000fd5b5050604051639926ee7d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169250639926ee7d91506109649085908790600401611841565b600060405180830381600087803b15801561097e57600080fd5b505af1158015610992573d6000803e3d6000fd5b505060405163ffffffff841681526001600160a01b03851692507f3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a915060200160405180910390a2505050505050565b60006109f66033546001600160a01b031690565b905090565b3360009081526065602052604090205460ff16610a2a5760405162461bcd60e51b815260040161038c90611649565b60405163bd7cd8ab60e01b815233906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063bd7cd8ab90610a7c908490879087906004016118c1565b600060405180830381600087803b158015610a9657600080fd5b505af1158015610aaa573d6000803e3d6000fd5b50506040516316fd18dd60e11b81526001600160a01b03848116600483015263ffffffff6024830181905293507f0000000000000000000000000000000000000000000000000000000000000000169150632dfa31ba90604401600060405180830381600087803b158015610b1e57600080fd5b505af1158015610b32573d6000803e3d6000fd5b505060405163ffffffff841681526001600160a01b03851692507f3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a915060200160405180910390a250505050565b610b886111ab565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb3559061078f90859085906004016118ef565b3360009081526065602052604090205460ff16610c055760405162461bcd60e51b815260040161038c90611649565b6040516319a74c5f60e01b815233600482018190529060009081907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906319a74c5f906024016040805180830381865afa158015610c70573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c94919061191e565b9150915081610cf45760405162461bcd60e51b815260206004820152602660248201527f546865206f70657261746f72206973206e6f742061626c6520746f206465726560448201526533b4b9ba32b960d11b606482015260840161038c565b6040516316fd18dd60e11b81526001600160a01b038481166004830152602482018390527f00000000000000000000000000000000000000000000000000000000000000001690632dfa31ba90604401600060405180830381600087803b158015610d5e57600080fd5b505af1158015610d72573d6000803e3d6000fd5b50506040516356452c2560e11b81526001600160a01b0386811660048301527f000000000000000000000000000000000000000000000000000000000000000016925063ac8a584a9150602401600060405180830381600087803b158015610dd957600080fd5b505af1158015610ded573d6000803e3d6000fd5b50506040516351b27a6d60e11b81526001600160a01b0386811660048301527f000000000000000000000000000000000000000000000000000000000000000016925063a364f4da9150602401600060405180830381600087803b158015610e5457600080fd5b505af1158015610e68573d6000803e3d6000fd5b50506040516001600160a01b03861692507f6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d9150600090a2505050565b600054610100900460ff1615808015610ec55750600054600160ff909116105b80610edf5750303b158015610edf575060005460ff166001145b610f425760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161038c565b6000805460ff191660011790558015610f65576000805461ff0019166101001790555b610f6e8261120a565b8015610fb4576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b3360009081526065602052604090205460ff16610fe75760405162461bcd60e51b815260040161038c90611649565b60405163e464b31f60e01b815233906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063e464b31f906106d390849087908790600401611951565b6110416111ab565b60005b818110156104ca576065600084848481811061106257611062611680565b9050602002016020810190611077919061133a565b6001600160a01b031681526020810191909152604001600020805460ff19169055806110a281611696565b915050611044565b60607f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663273cbaa06040518163ffffffff1660e01b8152600401600060405180830381865afa15801561110a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526109f691908101906116bf565b61113a6111ab565b6001600160a01b03811661119f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161038c565b6111a88161120a565b50565b336111b46109e2565b6001600160a01b0316146107d75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161038c565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b803563ffffffff8116811461127057600080fd5b919050565b60006020828403121561128757600080fd5b6112908261125c565b9392505050565b60008083601f8401126112a957600080fd5b50813567ffffffffffffffff8111156112c157600080fd5b6020830191508360208260051b85010111156112dc57600080fd5b9250929050565b600080602083850312156112f657600080fd5b823567ffffffffffffffff81111561130d57600080fd5b61131985828601611297565b90969095509350505050565b6001600160a01b03811681146111a857600080fd5b60006020828403121561134c57600080fd5b813561129081611325565b6020808252825182820181905260009190848201906040850190845b818110156113985783516001600160a01b031683529284019291840191600101611373565b50909695505050505050565b60008083601f8401126113b657600080fd5b50813567ffffffffffffffff8111156113ce57600080fd5b6020830191508360208260061b85010111156112dc57600080fd5b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715611422576114226113e9565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715611451576114516113e9565b604052919050565b6000806000806060858703121561146f57600080fd5b843561147a81611325565b935060208581013567ffffffffffffffff8082111561149857600080fd5b6114a489838a016113a4565b909650945060408801359150808211156114bd57600080fd5b908701906060828a0312156114d157600080fd5b6114d96113ff565b8235828111156114e857600080fd5b8301601f81018b136114f957600080fd5b80358381111561150b5761150b6113e9565b61151d601f8201601f19168701611428565b93508084528b8682840101111561153357600080fd5b8086830187860137600090840186015250908152818301359281019290925260409081013590820152939692955090935050565b6000806020838503121561157a57600080fd5b823567ffffffffffffffff81111561159157600080fd5b611319858286016113a4565b600080602083850312156115b057600080fd5b823567ffffffffffffffff808211156115c857600080fd5b818501915085601f8301126115dc57600080fd5b8135818111156115eb57600080fd5b8660208285010111156115fd57600080fd5b60209290920196919550909350505050565b6000806060838503121561162257600080fd5b61162b8361125c565b91508360608401111561163d57600080fd5b50926020919091019150565b6020808252601b908201527f4f70657261746f72206973206e6f742077686974656c69737465640000000000604082015260600190565b634e487b7160e01b600052603260045260246000fd5b60006000198214156116b857634e487b7160e01b600052601160045260246000fd5b5060010190565b600060208083850312156116d257600080fd5b825167ffffffffffffffff808211156116ea57600080fd5b818501915085601f8301126116fe57600080fd5b815181811115611710576117106113e9565b8060051b9150611721848301611428565b818152918301840191848101908884111561173b57600080fd5b938501935b83851015611765578451925061175583611325565b8282529385019390850190611740565b98975050505050505050565b6001600160a01b038416815260406020808301829052908201839052600090849060608401835b868110156117c15763ffffffff6117ae8561125c565b1682529282019290820190600101611798565b50979650505050505050565b8183526020830192506000816000805b858110156117ff576040808489379687018281529692909201916001016117dd565b509495945050505050565b6001600160a01b0385811682528416602082015260606040820181905260009061183790830184866117cd565b9695505050505050565b60018060a01b03831681526000602060408184015283516060604085015280518060a086015260005b818110156118865782810184015186820160c00152830161186a565b8181111561189857600060c083880101525b509185015160608501525060409093015160808301525060c0601f909201601f19160101919050565b6001600160a01b03841681526040602082018190526000906118e690830184866117cd565b95945050505050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b6000806040838503121561193157600080fd5b8251801515811461194157600080fd5b6020939093015192949293505050565b6001600160a01b038416815263ffffffff8316602082015260808101604083818401376000815294935050505056fea26469706673582212204510c251a2ce04af1a567bcfe189383a7b7a9bd4e1f4194a8f2771a0e8e1615964736f6c634300080c0033",
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
	parsed, err := abi.JSON(strings.NewReader(LagrangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// AddBlsPubKeys is a paid mutator transaction binding the contract method 0x992b098e.
//
// Solidity: function addBlsPubKeys(uint256[2][] additionalBlsPubKeys) returns()
func (_Lagrange *LagrangeTransactor) AddBlsPubKeys(opts *bind.TransactOpts, additionalBlsPubKeys [][2]*big.Int) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "addBlsPubKeys", additionalBlsPubKeys)
}

// AddBlsPubKeys is a paid mutator transaction binding the contract method 0x992b098e.
//
// Solidity: function addBlsPubKeys(uint256[2][] additionalBlsPubKeys) returns()
func (_Lagrange *LagrangeSession) AddBlsPubKeys(additionalBlsPubKeys [][2]*big.Int) (*types.Transaction, error) {
	return _Lagrange.Contract.AddBlsPubKeys(&_Lagrange.TransactOpts, additionalBlsPubKeys)
}

// AddBlsPubKeys is a paid mutator transaction binding the contract method 0x992b098e.
//
// Solidity: function addBlsPubKeys(uint256[2][] additionalBlsPubKeys) returns()
func (_Lagrange *LagrangeTransactorSession) AddBlsPubKeys(additionalBlsPubKeys [][2]*big.Int) (*types.Transaction, error) {
	return _Lagrange.Contract.AddBlsPubKeys(&_Lagrange.TransactOpts, additionalBlsPubKeys)
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

// Register is a paid mutator transaction binding the contract method 0x7a79ea33.
//
// Solidity: function register(address signAddress, uint256[2][] blsPubKeys, (bytes,bytes32,uint256) operatorSignature) returns()
func (_Lagrange *LagrangeTransactor) Register(opts *bind.TransactOpts, signAddress common.Address, blsPubKeys [][2]*big.Int, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "register", signAddress, blsPubKeys, operatorSignature)
}

// Register is a paid mutator transaction binding the contract method 0x7a79ea33.
//
// Solidity: function register(address signAddress, uint256[2][] blsPubKeys, (bytes,bytes32,uint256) operatorSignature) returns()
func (_Lagrange *LagrangeSession) Register(signAddress common.Address, blsPubKeys [][2]*big.Int, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Lagrange.Contract.Register(&_Lagrange.TransactOpts, signAddress, blsPubKeys, operatorSignature)
}

// Register is a paid mutator transaction binding the contract method 0x7a79ea33.
//
// Solidity: function register(address signAddress, uint256[2][] blsPubKeys, (bytes,bytes32,uint256) operatorSignature) returns()
func (_Lagrange *LagrangeTransactorSession) Register(signAddress common.Address, blsPubKeys [][2]*big.Int, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Lagrange.Contract.Register(&_Lagrange.TransactOpts, signAddress, blsPubKeys, operatorSignature)
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

// UpdateBlsPubKey is a paid mutator transaction binding the contract method 0xc8d8e013.
//
// Solidity: function updateBlsPubKey(uint32 index, uint256[2] blsPubKey) returns()
func (_Lagrange *LagrangeTransactor) UpdateBlsPubKey(opts *bind.TransactOpts, index uint32, blsPubKey [2]*big.Int) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "updateBlsPubKey", index, blsPubKey)
}

// UpdateBlsPubKey is a paid mutator transaction binding the contract method 0xc8d8e013.
//
// Solidity: function updateBlsPubKey(uint32 index, uint256[2] blsPubKey) returns()
func (_Lagrange *LagrangeSession) UpdateBlsPubKey(index uint32, blsPubKey [2]*big.Int) (*types.Transaction, error) {
	return _Lagrange.Contract.UpdateBlsPubKey(&_Lagrange.TransactOpts, index, blsPubKey)
}

// UpdateBlsPubKey is a paid mutator transaction binding the contract method 0xc8d8e013.
//
// Solidity: function updateBlsPubKey(uint32 index, uint256[2] blsPubKey) returns()
func (_Lagrange *LagrangeTransactorSession) UpdateBlsPubKey(index uint32, blsPubKey [2]*big.Int) (*types.Transaction, error) {
	return _Lagrange.Contract.UpdateBlsPubKey(&_Lagrange.TransactOpts, index, blsPubKey)
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
