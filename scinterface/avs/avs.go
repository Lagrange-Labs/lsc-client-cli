// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package avs

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

// AvsMetaData contains all meta data concerning the Avs contract.
var AvsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_AVS_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsOperatorStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIAVSDirectory.OperatorAVSRegistrationStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorAVSRegistrationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelSalt\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorSaltIsSpent\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSRegistrationStatusUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIAVSDirectory.OperatorAVSRegistrationStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106101425760003560e01c80638da5cb5b116100b8578063d79aceab1161007c578063d79aceab146102f8578063df5cf7231461031f578063ec76f44214610346578063f2fde38b14610359578063f698da251461036c578063fabc1cbc1461037457600080fd5b80638da5cb5b1461029b5780639926ee7d146102ac578063a1060c88146102bf578063a364f4da146102d2578063a98fb355146102e557600080fd5b806349075da31161010a57806349075da3146101fa578063595c6a67146102355780635ac86ab71461023d5780635c975abb14610260578063715018a614610268578063886f11951461027057600080fd5b806310d67a2f14610147578063136439dd1461015c5780631794bb3c1461016f57806320606b7014610182578063374823b5146101bc575b600080fd5b61015a6101553660046118ab565b610387565b005b61015a61016a3660046118cf565b610443565b61015a61017d3660046118e8565b610582565b6101a97f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a86681565b6040519081526020015b60405180910390f35b6101ea6101ca366004611929565b609960209081526000928352604080842090915290825290205460ff1681565b60405190151581526020016101b3565b610228610208366004611955565b609860209081526000928352604080842090915290825290205460ff1681565b6040516101b391906119a4565b61015a6106ac565b6101ea61024b3660046119cc565b606654600160ff9092169190911b9081161490565b6066546101a9565b61015a610773565b606554610283906001600160a01b031681565b6040516001600160a01b0390911681526020016101b3565b6033546001600160a01b0316610283565b61015a6102ba366004611a5f565b610787565b6101a96102cd366004611b46565b610b1a565b61015a6102e03660046118ab565b610bd3565b61015a6102f3366004611b8c565b610d3c565b6101a97fda2c89bafdd34776a2b8bb9c83c82f419e20cc8c67207f70edd58249b92661bd81565b6102837f000000000000000000000000000000000000000000000000000000000000000081565b61015a6103543660046118cf565b610d83565b61015a6103673660046118ab565b610e2e565b6101a9610ea4565b61015a6103823660046118cf565b610ee2565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103fe9190611bfe565b6001600160a01b0316336001600160a01b0316146104375760405162461bcd60e51b815260040161042e90611c1b565b60405180910390fd5b6104408161103e565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561048b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104af9190611c65565b6104cb5760405162461bcd60e51b815260040161042e90611c87565b606654818116146105445760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c6974790000000000000000606482015260840161042e565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600054610100900460ff16158080156105a25750600054600160ff909116105b806105bc5750303b1580156105bc575060005460ff166001145b61061f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161042e565b6000805460ff191660011790558015610642576000805461ff0019166101001790555b61064c8383611135565b61065461121f565b609755610660846112b6565b80156106a6576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107189190611c65565b6107345760405162461bcd60e51b815260040161042e90611c87565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b61077b611308565b61078560006112b6565b565b606654600090600190811614156107dc5760405162461bcd60e51b815260206004820152601960248201527814185d5cd8589b194e881a5b99195e081a5cc81c185d5cd959603a1b604482015260640161042e565b42826040015110156108445760405162461bcd60e51b815260206004820152603e6024820152600080516020611dac83398151915260448201527f56533a206f70657261746f72207369676e617475726520657870697265640000606482015260840161042e565b60013360009081526098602090815260408083206001600160a01b038816845290915290205460ff16600181111561087e5761087e61198e565b14156108e05760405162461bcd60e51b815260206004820152603f6024820152600080516020611dac83398151915260448201527f56533a206f70657261746f7220616c7265616479207265676973746572656400606482015260840161042e565b6001600160a01b038316600090815260996020908152604080832085830151845290915290205460ff16156109645760405162461bcd60e51b81526020600482015260366024820152600080516020611dac8339815191526044820152751594ce881cd85b1d08185b1c9958591e481cdc195b9d60521b606482015260840161042e565b6040516336b87bd760e11b81526001600160a01b0384811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa1580156109ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ee9190611c65565b610a645760405162461bcd60e51b815260206004820152604d6024820152600080516020611dac83398151915260448201527f56533a206f70657261746f72206e6f74207265676973746572656420746f204560648201526c1a59d95b93185e595c881e595d609a1b608482015260a40161042e565b6000610a7a843385602001518660400151610b1a565b9050610a8b84828560000151611362565b3360008181526098602090815260408083206001600160a01b0389168085529083528184208054600160ff199182168117909255609985528386208a860151875290945293829020805490931684179092555190917ff0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b4191610b0c91906119a4565b60405180910390a350505050565b604080517fda2c89bafdd34776a2b8bb9c83c82f419e20cc8c67207f70edd58249b92661bd6020808301919091526001600160a01b0387811683850152861660608301526080820185905260a08083018590528351808403909101815260c0909201909252805191012060009081610b90610ea4565b60405161190160f01b602082015260228101919091526042810183905260620160408051808303601f190181529190528051602090910120979650505050505050565b60665460009060019081161415610c285760405162461bcd60e51b815260206004820152601960248201527814185d5cd8589b194e881a5b99195e081a5cc81c185d5cd959603a1b604482015260640161042e565b60013360009081526098602090815260408083206001600160a01b038716845290915290205460ff166001811115610c6257610c6261198e565b14610cd55760405162461bcd60e51b815260206004820152603f60248201527f4156534469726563746f72792e646572656769737465724f70657261746f724660448201527f726f6d4156533a206f70657261746f72206e6f74207265676973746572656400606482015260840161042e565b3360008181526098602090815260408083206001600160a01b0387168085529252808320805460ff191690555190917ff0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b4191610d3091906119a4565b60405180910390a35050565b336001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c9437138383604051610d77929190611ccf565b60405180910390a25050565b33600090815260996020908152604080832084845290915290205460ff1615610e085760405162461bcd60e51b815260206004820152603160248201527f4156534469726563746f72792e63616e63656c53616c743a2063616e6e6f742060448201527018d85b98d95b081cdc195b9d081cd85b1d607a1b606482015260840161042e565b33600090815260996020908152604080832093835292905220805460ff19166001179055565b610e36611308565b6001600160a01b038116610e9b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161042e565b610440816112b6565b60007f0000000000000000000000000000000000000000000000000000000000000000461415610ed5575060975490565b610edd61121f565b905090565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f35573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f599190611bfe565b6001600160a01b0316336001600160a01b031614610f895760405162461bcd60e51b815260040161042e90611c1b565b6066541981196066541916146110075760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c6974790000000000000000606482015260840161042e565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610577565b6001600160a01b0381166110cc5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a40161042e565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6065546001600160a01b031615801561115657506001600160a01b03821615155b6111d85760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a40161042e565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261121b8261103e565b5050565b604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6033546001600160a01b031633146107855760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161042e565b6001600160a01b0383163b1561148157604051630b135d3f60e11b808252906001600160a01b03851690631626ba7e906113a29086908690600401611cfe565b602060405180830381865afa1580156113bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113e39190611d5b565b6001600160e01b0319161461147c5760405162461bcd60e51b815260206004820152605360248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a2045524331323731207369676e6174757265206064820152721d995c9a599a58d85d1a5bdb8819985a5b1959606a1b608482015260a40161042e565b505050565b826001600160a01b03166114958383611521565b6001600160a01b03161461147c5760405162461bcd60e51b815260206004820152604760248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a207369676e6174757265206e6f742066726f6d6064820152661039b4b3b732b960c91b608482015260a40161042e565b60008060006115308585611545565b9150915061153d816115b5565b509392505050565b60008082516041141561157c5760208301516040840151606085015160001a61157087828585611770565b945094505050506115ae565b8251604014156115a6576020830151604084015161159b86838361185d565b9350935050506115ae565b506000905060025b9250929050565b60008160048111156115c9576115c961198e565b14156115d25750565b60018160048111156115e6576115e661198e565b14156116345760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161042e565b60028160048111156116485761164861198e565b14156116965760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161042e565b60038160048111156116aa576116aa61198e565b14156117035760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161042e565b60048160048111156117175761171761198e565b14156104405760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b606482015260840161042e565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156117a75750600090506003611854565b8460ff16601b141580156117bf57508460ff16601c14155b156117d05750600090506004611854565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611824573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661184d57600060019250925050611854565b9150600090505b94509492505050565b6000806001600160ff1b0383168161187a60ff86901c601b611d85565b905061188887828885611770565b935093505050935093915050565b6001600160a01b038116811461044057600080fd5b6000602082840312156118bd57600080fd5b81356118c881611896565b9392505050565b6000602082840312156118e157600080fd5b5035919050565b6000806000606084860312156118fd57600080fd5b833561190881611896565b9250602084013561191881611896565b929592945050506040919091013590565b6000806040838503121561193c57600080fd5b823561194781611896565b946020939093013593505050565b6000806040838503121561196857600080fd5b823561197381611896565b9150602083013561198381611896565b809150509250929050565b634e487b7160e01b600052602160045260246000fd5b60208101600283106119c657634e487b7160e01b600052602160045260246000fd5b91905290565b6000602082840312156119de57600080fd5b813560ff811681146118c857600080fd5b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715611a2857611a286119ef565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715611a5757611a576119ef565b604052919050565b60008060408385031215611a7257600080fd5b8235611a7d81611896565b915060208381013567ffffffffffffffff80821115611a9b57600080fd5b9085019060608288031215611aaf57600080fd5b611ab7611a05565b823582811115611ac657600080fd5b8301601f81018913611ad757600080fd5b803583811115611ae957611ae96119ef565b611afb601f8201601f19168701611a2e565b93508084528986828401011115611b1157600080fd5b808683018786013760008682860101525050818152838301358482015260408301356040820152809450505050509250929050565b60008060008060808587031215611b5c57600080fd5b8435611b6781611896565b93506020850135611b7781611896565b93969395505050506040820135916060013590565b60008060208385031215611b9f57600080fd5b823567ffffffffffffffff80821115611bb757600080fd5b818501915085601f830112611bcb57600080fd5b813581811115611bda57600080fd5b866020828501011115611bec57600080fd5b60209290920196919550909350505050565b600060208284031215611c1057600080fd5b81516118c881611896565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215611c7757600080fd5b815180151581146118c857600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b82815260006020604081840152835180604085015260005b81811015611d3257858101830151858201606001528201611d16565b81811115611d44576000606083870101525b50601f01601f191692909201606001949350505050565b600060208284031215611d6d57600080fd5b81516001600160e01b0319811681146118c857600080fd5b60008219821115611da657634e487b7160e01b600052601160045260246000fd5b50019056fe4156534469726563746f72792e72656769737465724f70657261746f72546f41a2646970667358221220cb5e8d0c40d6ac9a73119826942e085f364cead534ac646ead53145559007e4a64736f6c634300080c0033",
}

// AvsABI is the input ABI used to generate the binding from.
// Deprecated: Use AvsMetaData.ABI instead.
var AvsABI = AvsMetaData.ABI

// AvsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AvsMetaData.Bin instead.
var AvsBin = AvsMetaData.Bin

// DeployAvs deploys a new Ethereum contract, binding an instance of Avs to it.
func DeployAvs(auth *bind.TransactOpts, backend bind.ContractBackend, _delegation common.Address) (common.Address, *types.Transaction, *Avs, error) {
	parsed, err := AvsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AvsBin), backend, _delegation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Avs{AvsCaller: AvsCaller{contract: contract}, AvsTransactor: AvsTransactor{contract: contract}, AvsFilterer: AvsFilterer{contract: contract}}, nil
}

// Avs is an auto generated Go binding around an Ethereum contract.
type Avs struct {
	AvsCaller     // Read-only binding to the contract
	AvsTransactor // Write-only binding to the contract
	AvsFilterer   // Log filterer for contract events
}

// AvsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AvsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AvsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AvsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AvsSession struct {
	Contract     *Avs              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AvsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AvsCallerSession struct {
	Contract *AvsCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AvsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AvsTransactorSession struct {
	Contract     *AvsTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AvsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AvsRaw struct {
	Contract *Avs // Generic contract binding to access the raw methods on
}

// AvsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AvsCallerRaw struct {
	Contract *AvsCaller // Generic read-only contract binding to access the raw methods on
}

// AvsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AvsTransactorRaw struct {
	Contract *AvsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAvs creates a new instance of Avs, bound to a specific deployed contract.
func NewAvs(address common.Address, backend bind.ContractBackend) (*Avs, error) {
	contract, err := bindAvs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Avs{AvsCaller: AvsCaller{contract: contract}, AvsTransactor: AvsTransactor{contract: contract}, AvsFilterer: AvsFilterer{contract: contract}}, nil
}

// NewAvsCaller creates a new read-only instance of Avs, bound to a specific deployed contract.
func NewAvsCaller(address common.Address, caller bind.ContractCaller) (*AvsCaller, error) {
	contract, err := bindAvs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AvsCaller{contract: contract}, nil
}

// NewAvsTransactor creates a new write-only instance of Avs, bound to a specific deployed contract.
func NewAvsTransactor(address common.Address, transactor bind.ContractTransactor) (*AvsTransactor, error) {
	contract, err := bindAvs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AvsTransactor{contract: contract}, nil
}

// NewAvsFilterer creates a new log filterer instance of Avs, bound to a specific deployed contract.
func NewAvsFilterer(address common.Address, filterer bind.ContractFilterer) (*AvsFilterer, error) {
	contract, err := bindAvs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AvsFilterer{contract: contract}, nil
}

// bindAvs binds a generic wrapper to an already deployed contract.
func bindAvs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AvsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Avs *AvsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Avs.Contract.AvsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Avs *AvsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Avs.Contract.AvsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Avs *AvsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Avs.Contract.AvsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Avs *AvsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Avs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Avs *AvsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Avs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Avs *AvsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Avs.Contract.contract.Transact(opts, method, params...)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Avs *AvsCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Avs.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Avs *AvsSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Avs.Contract.DOMAINTYPEHASH(&_Avs.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Avs *AvsCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Avs.Contract.DOMAINTYPEHASH(&_Avs.CallOpts)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_Avs *AvsCaller) OPERATORAVSREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Avs.contract.Call(opts, &out, "OPERATOR_AVS_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_Avs *AvsSession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _Avs.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_Avs.CallOpts)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_Avs *AvsCallerSession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _Avs.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_Avs.CallOpts)
}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address , address ) view returns(uint8)
func (_Avs *AvsCaller) AvsOperatorStatus(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (uint8, error) {
	var out []interface{}
	err := _Avs.contract.Call(opts, &out, "avsOperatorStatus", arg0, arg1)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address , address ) view returns(uint8)
func (_Avs *AvsSession) AvsOperatorStatus(arg0 common.Address, arg1 common.Address) (uint8, error) {
	return _Avs.Contract.AvsOperatorStatus(&_Avs.CallOpts, arg0, arg1)
}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address , address ) view returns(uint8)
func (_Avs *AvsCallerSession) AvsOperatorStatus(arg0 common.Address, arg1 common.Address) (uint8, error) {
	return _Avs.Contract.AvsOperatorStatus(&_Avs.CallOpts, arg0, arg1)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_Avs *AvsCaller) CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Avs.contract.Call(opts, &out, "calculateOperatorAVSRegistrationDigestHash", operator, avs, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_Avs *AvsSession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _Avs.Contract.CalculateOperatorAVSRegistrationDigestHash(&_Avs.CallOpts, operator, avs, salt, expiry)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_Avs *AvsCallerSession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _Avs.Contract.CalculateOperatorAVSRegistrationDigestHash(&_Avs.CallOpts, operator, avs, salt, expiry)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_Avs *AvsCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Avs.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_Avs *AvsSession) Delegation() (common.Address, error) {
	return _Avs.Contract.Delegation(&_Avs.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_Avs *AvsCallerSession) Delegation() (common.Address, error) {
	return _Avs.Contract.Delegation(&_Avs.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Avs *AvsCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Avs.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Avs *AvsSession) DomainSeparator() ([32]byte, error) {
	return _Avs.Contract.DomainSeparator(&_Avs.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Avs *AvsCallerSession) DomainSeparator() ([32]byte, error) {
	return _Avs.Contract.DomainSeparator(&_Avs.CallOpts)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_Avs *AvsCaller) OperatorSaltIsSpent(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _Avs.contract.Call(opts, &out, "operatorSaltIsSpent", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_Avs *AvsSession) OperatorSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _Avs.Contract.OperatorSaltIsSpent(&_Avs.CallOpts, arg0, arg1)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_Avs *AvsCallerSession) OperatorSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _Avs.Contract.OperatorSaltIsSpent(&_Avs.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Avs *AvsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Avs.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Avs *AvsSession) Owner() (common.Address, error) {
	return _Avs.Contract.Owner(&_Avs.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Avs *AvsCallerSession) Owner() (common.Address, error) {
	return _Avs.Contract.Owner(&_Avs.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Avs *AvsCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _Avs.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Avs *AvsSession) Paused(index uint8) (bool, error) {
	return _Avs.Contract.Paused(&_Avs.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Avs *AvsCallerSession) Paused(index uint8) (bool, error) {
	return _Avs.Contract.Paused(&_Avs.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Avs *AvsCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Avs.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Avs *AvsSession) Paused0() (*big.Int, error) {
	return _Avs.Contract.Paused0(&_Avs.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Avs *AvsCallerSession) Paused0() (*big.Int, error) {
	return _Avs.Contract.Paused0(&_Avs.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Avs *AvsCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Avs.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Avs *AvsSession) PauserRegistry() (common.Address, error) {
	return _Avs.Contract.PauserRegistry(&_Avs.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Avs *AvsCallerSession) PauserRegistry() (common.Address, error) {
	return _Avs.Contract.PauserRegistry(&_Avs.CallOpts)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_Avs *AvsTransactor) CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error) {
	return _Avs.contract.Transact(opts, "cancelSalt", salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_Avs *AvsSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _Avs.Contract.CancelSalt(&_Avs.TransactOpts, salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_Avs *AvsTransactorSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _Avs.Contract.CancelSalt(&_Avs.TransactOpts, salt)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_Avs *AvsTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Avs.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_Avs *AvsSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _Avs.Contract.DeregisterOperatorFromAVS(&_Avs.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_Avs *AvsTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _Avs.Contract.DeregisterOperatorFromAVS(&_Avs.TransactOpts, operator)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_Avs *AvsTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _Avs.contract.Transact(opts, "initialize", initialOwner, _pauserRegistry, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_Avs *AvsSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _Avs.Contract.Initialize(&_Avs.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_Avs *AvsTransactorSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _Avs.Contract.Initialize(&_Avs.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Avs *AvsTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Avs.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Avs *AvsSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Avs.Contract.Pause(&_Avs.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Avs *AvsTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Avs.Contract.Pause(&_Avs.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Avs *AvsTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Avs.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Avs *AvsSession) PauseAll() (*types.Transaction, error) {
	return _Avs.Contract.PauseAll(&_Avs.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Avs *AvsTransactorSession) PauseAll() (*types.Transaction, error) {
	return _Avs.Contract.PauseAll(&_Avs.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_Avs *AvsTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Avs.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_Avs *AvsSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Avs.Contract.RegisterOperatorToAVS(&_Avs.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_Avs *AvsTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Avs.Contract.RegisterOperatorToAVS(&_Avs.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Avs *AvsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Avs.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Avs *AvsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Avs.Contract.RenounceOwnership(&_Avs.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Avs *AvsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Avs.Contract.RenounceOwnership(&_Avs.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_Avs *AvsTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _Avs.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_Avs *AvsSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _Avs.Contract.SetPauserRegistry(&_Avs.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_Avs *AvsTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _Avs.Contract.SetPauserRegistry(&_Avs.TransactOpts, newPauserRegistry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Avs *AvsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Avs.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Avs *AvsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Avs.Contract.TransferOwnership(&_Avs.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Avs *AvsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Avs.Contract.TransferOwnership(&_Avs.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Avs *AvsTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Avs.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Avs *AvsSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Avs.Contract.Unpause(&_Avs.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Avs *AvsTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Avs.Contract.Unpause(&_Avs.TransactOpts, newPausedStatus)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_Avs *AvsTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _Avs.contract.Transact(opts, "updateAVSMetadataURI", metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_Avs *AvsSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _Avs.Contract.UpdateAVSMetadataURI(&_Avs.TransactOpts, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_Avs *AvsTransactorSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _Avs.Contract.UpdateAVSMetadataURI(&_Avs.TransactOpts, metadataURI)
}

// AvsAVSMetadataURIUpdatedIterator is returned from FilterAVSMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for AVSMetadataURIUpdated events raised by the Avs contract.
type AvsAVSMetadataURIUpdatedIterator struct {
	Event *AvsAVSMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *AvsAVSMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsAVSMetadataURIUpdated)
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
		it.Event = new(AvsAVSMetadataURIUpdated)
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
func (it *AvsAVSMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsAVSMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsAVSMetadataURIUpdated represents a AVSMetadataURIUpdated event raised by the Avs contract.
type AvsAVSMetadataURIUpdated struct {
	Avs         common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAVSMetadataURIUpdated is a free log retrieval operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_Avs *AvsFilterer) FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*AvsAVSMetadataURIUpdatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _Avs.contract.FilterLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return &AvsAVSMetadataURIUpdatedIterator{contract: _Avs.contract, event: "AVSMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchAVSMetadataURIUpdated is a free log subscription operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_Avs *AvsFilterer) WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *AvsAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _Avs.contract.WatchLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsAVSMetadataURIUpdated)
				if err := _Avs.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
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

// ParseAVSMetadataURIUpdated is a log parse operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_Avs *AvsFilterer) ParseAVSMetadataURIUpdated(log types.Log) (*AvsAVSMetadataURIUpdated, error) {
	event := new(AvsAVSMetadataURIUpdated)
	if err := _Avs.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Avs contract.
type AvsInitializedIterator struct {
	Event *AvsInitialized // Event containing the contract specifics and raw log

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
func (it *AvsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsInitialized)
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
		it.Event = new(AvsInitialized)
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
func (it *AvsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsInitialized represents a Initialized event raised by the Avs contract.
type AvsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Avs *AvsFilterer) FilterInitialized(opts *bind.FilterOpts) (*AvsInitializedIterator, error) {

	logs, sub, err := _Avs.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AvsInitializedIterator{contract: _Avs.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Avs *AvsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AvsInitialized) (event.Subscription, error) {

	logs, sub, err := _Avs.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsInitialized)
				if err := _Avs.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Avs *AvsFilterer) ParseInitialized(log types.Log) (*AvsInitialized, error) {
	event := new(AvsInitialized)
	if err := _Avs.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOperatorAVSRegistrationStatusUpdatedIterator is returned from FilterOperatorAVSRegistrationStatusUpdated and is used to iterate over the raw logs and unpacked data for OperatorAVSRegistrationStatusUpdated events raised by the Avs contract.
type AvsOperatorAVSRegistrationStatusUpdatedIterator struct {
	Event *AvsOperatorAVSRegistrationStatusUpdated // Event containing the contract specifics and raw log

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
func (it *AvsOperatorAVSRegistrationStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorAVSRegistrationStatusUpdated)
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
		it.Event = new(AvsOperatorAVSRegistrationStatusUpdated)
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
func (it *AvsOperatorAVSRegistrationStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorAVSRegistrationStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorAVSRegistrationStatusUpdated represents a OperatorAVSRegistrationStatusUpdated event raised by the Avs contract.
type AvsOperatorAVSRegistrationStatusUpdated struct {
	Operator common.Address
	Avs      common.Address
	Status   uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAVSRegistrationStatusUpdated is a free log retrieval operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_Avs *AvsFilterer) FilterOperatorAVSRegistrationStatusUpdated(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*AvsOperatorAVSRegistrationStatusUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _Avs.contract.FilterLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorAVSRegistrationStatusUpdatedIterator{contract: _Avs.contract, event: "OperatorAVSRegistrationStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorAVSRegistrationStatusUpdated is a free log subscription operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_Avs *AvsFilterer) WatchOperatorAVSRegistrationStatusUpdated(opts *bind.WatchOpts, sink chan<- *AvsOperatorAVSRegistrationStatusUpdated, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _Avs.contract.WatchLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorAVSRegistrationStatusUpdated)
				if err := _Avs.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
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

// ParseOperatorAVSRegistrationStatusUpdated is a log parse operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_Avs *AvsFilterer) ParseOperatorAVSRegistrationStatusUpdated(log types.Log) (*AvsOperatorAVSRegistrationStatusUpdated, error) {
	event := new(AvsOperatorAVSRegistrationStatusUpdated)
	if err := _Avs.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Avs contract.
type AvsOwnershipTransferredIterator struct {
	Event *AvsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AvsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOwnershipTransferred)
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
		it.Event = new(AvsOwnershipTransferred)
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
func (it *AvsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOwnershipTransferred represents a OwnershipTransferred event raised by the Avs contract.
type AvsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Avs *AvsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AvsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Avs.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AvsOwnershipTransferredIterator{contract: _Avs.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Avs *AvsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AvsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Avs.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOwnershipTransferred)
				if err := _Avs.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Avs *AvsFilterer) ParseOwnershipTransferred(log types.Log) (*AvsOwnershipTransferred, error) {
	event := new(AvsOwnershipTransferred)
	if err := _Avs.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Avs contract.
type AvsPausedIterator struct {
	Event *AvsPaused // Event containing the contract specifics and raw log

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
func (it *AvsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsPaused)
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
		it.Event = new(AvsPaused)
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
func (it *AvsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsPaused represents a Paused event raised by the Avs contract.
type AvsPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_Avs *AvsFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*AvsPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Avs.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &AvsPausedIterator{contract: _Avs.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_Avs *AvsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AvsPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Avs.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsPaused)
				if err := _Avs.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_Avs *AvsFilterer) ParsePaused(log types.Log) (*AvsPaused, error) {
	event := new(AvsPaused)
	if err := _Avs.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the Avs contract.
type AvsPauserRegistrySetIterator struct {
	Event *AvsPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *AvsPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsPauserRegistrySet)
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
		it.Event = new(AvsPauserRegistrySet)
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
func (it *AvsPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsPauserRegistrySet represents a PauserRegistrySet event raised by the Avs contract.
type AvsPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_Avs *AvsFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*AvsPauserRegistrySetIterator, error) {

	logs, sub, err := _Avs.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &AvsPauserRegistrySetIterator{contract: _Avs.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_Avs *AvsFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *AvsPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _Avs.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsPauserRegistrySet)
				if err := _Avs.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_Avs *AvsFilterer) ParsePauserRegistrySet(log types.Log) (*AvsPauserRegistrySet, error) {
	event := new(AvsPauserRegistrySet)
	if err := _Avs.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Avs contract.
type AvsUnpausedIterator struct {
	Event *AvsUnpaused // Event containing the contract specifics and raw log

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
func (it *AvsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsUnpaused)
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
		it.Event = new(AvsUnpaused)
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
func (it *AvsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsUnpaused represents a Unpaused event raised by the Avs contract.
type AvsUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_Avs *AvsFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*AvsUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Avs.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &AvsUnpausedIterator{contract: _Avs.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_Avs *AvsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AvsUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Avs.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsUnpaused)
				if err := _Avs.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_Avs *AvsFilterer) ParseUnpaused(log types.Log) (*AvsUnpaused, error) {
	event := new(AvsUnpaused)
	if err := _Avs.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
