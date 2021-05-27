package systemcontract

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/congress/vmcaller"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"math"
	"math/big"
)

const (
	punishV1Code = "0x608060405234801561001057600080fd5b506004361061012c5760003560e01c8063741579b1116100ad578063d93d2cb911610071578063d93d2cb91461027b578063e0d8ea5314610298578063ea7221a1146102a0578063f5323feb146102c6578063f62af26c146102ce5761012c565b8063741579b1146101315780638129fc1c146102445780639001eed81461024c578063c967f90f14610254578063cb1ea725146102735761012c565b806332f3c17f116100f457806332f3c17f1461019c5780633a5381b5146101c25780633b58524d146101e657806344c1aa991461021657806363e1d4511461021e5761012c565b806303fab4f614610131578063158ef93e1461014b57806315de360e146101675780632897183d1461016f578063303a4a4614610177575b600080fd5b6101396102eb565b60408051918252519081900360200190f35b6101536102f7565b604080519115158252519081900360200190f35b610139610300565b610139610305565b61017f610300565b6040805167ffffffffffffffff9092168252519081900360200190f35b610139600480360360208110156101b257600080fd5b50356001600160a01b031661030b565b6101ca610326565b604080516001600160a01b039092168252519081900360200190f35b610214600480360360408110156101fc57600080fd5b506001600160a01b038135811691602001351661033a565b005b610139610372565b6101536004803603602081101561023457600080fd5b50356001600160a01b0316610378565b6102146105e3565b610139610601565b61025c61060d565b6040805161ffff9092168252519081900360200190f35b610139610612565b6102146004803603602081101561029157600080fd5b5035610618565b610139610875565b610214600480360360208110156102b657600080fd5b50356001600160a01b031661087b565b6101ca610b61565b6101ca600480360360208110156102e457600080fd5b5035610b70565b670de0b6b3a764000081565b60005460ff1681565b600081565b60045481565b6001600160a01b031660009081526005602052604090205490565b60005461010090046001600160a01b031681565b60008054610100600160a81b0319166101006001600160a01b0394851602179055600180546001600160a01b03191691909216179055565b60035481565b60008054604080516308ab66a960e41b81526001600160a01b03858116600483015291513393610100900490921691638ab66a9091602480820192602092909190829003018186803b1580156103cd57600080fd5b505afa1580156103e1573d6000803e3d6000fd5b505050506040513d60208110156103f757600080fd5b50516001600160a01b031614610454576040805162461bcd60e51b815260206004820152601860248201527f43616e646964617465206e6f7420726567697374657265640000000000000000604482015290519081900360640190fd5b6001600160a01b0382166000908152600560205260409020541561048c576001600160a01b0382166000908152600560205260408120555b6001600160a01b03821660009081526005602052604090206002015460ff1680156104b8575060065415155b156105db576006546001600160a01b0383166000908152600560205260409020600101546000199091011461058257600680546000919060001981019081106104fd57fe5b60009182526020808320909101546001600160a01b038681168452600590925260409092206001015460068054929093169350839291811061053b57fe5b600091825260208083209190910180546001600160a01b0319166001600160a01b039485161790558583168252600590526040808220600190810154949093168252902001555b600680548061058d57fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038416825260059052604081206001810191909155600201805460ff191690555b506001919050565b6018600281905560306003556004556000805460ff19166001179055565b674563918244f4000081565b601581565b60025481565b334114610659576040805162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b604482015290519081900360640190fd5b4360009081526008602052604090205460ff16156106b2576040805162461bcd60e51b8152602060048201526011602482015270105b1c9958591e48191958dc99585cd959607a1b604482015290519081900360640190fd5b808043816106bc57fe5b0615610702576040805162461bcd60e51b815260206004820152601060248201526f426c6f636b2065706f6368206f6e6c7960801b604482015290519081900360640190fd5b436000908152600860205260409020805460ff1916600117905560065461072857610871565b60005b600654811015610846576004546003548161074257fe5b04600560006006848154811061075457fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020541115610805576004546003548161078c57fe5b04600560006006848154811061079e57fe5b60009182526020808320909101546001600160a01b031683528201929092526040018120546006805493909103926005929190859081106107db57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205561083e565b6000600560006006848154811061081857fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020555b60010161072b565b506040517f181d51be54e8e8eaca6eae0eab32d4162099236bd519e7238d015d0870db464190600090a15b5050565b60065490565b3341146108bc576040805162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b604482015290519081900360640190fd5b4360009081526007602052604090205460ff1615610914576040805162461bcd60e51b815260206004820152601060248201526f105b1c9958591e481c1d5b9a5cda195960821b604482015290519081900360640190fd5b436000908152600760209081526040808320805460ff191660011790556001600160a01b0384168352600590915290206002015460ff166109bd57600680546001600160a01b038316600081815260056020526040812060018082018590558085019095557ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f90930180546001600160a01b0319168317905552600201805460ff191690911790555b6001600160a01b038116600090815260056020526040902080546001019081905560035490816109e957fe5b06610af95760008060019054906101000a90046001600160a01b03166001600160a01b0316638ab66a90836040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610a5157600080fd5b505afa158015610a65573d6000803e3d6000fd5b505050506040513d6020811015610a7b57600080fd5b50516040805163209b4f7b60e21b815290519192506001600160a01b0383169163826d3dec9160048082019260009290919082900301818387803b158015610ac257600080fd5b505af1158015610ad6573d6000803e3d6000fd5b505050506001600160a01b03821660009081526005602052604081205550610b1f565b6002546001600160a01b03821660009081526005602052604090205481610b1c57fe5b50505b6040805142815290516001600160a01b038316917f770e0cca42c35d00240986ce8d3ed438be04663c91dac6576b79537d7c180f1e919081900360200190a250565b6001546001600160a01b031681565b60068181548110610b7d57fe5b6000918252602090912001546001600160a01b031690508156fea2646970667358221220f60e92140d99ec76fc519b005247d3b68555cfeb78a6125abd0c7d2e0991d74964736f6c634300060c0033"
)

type hardForkPunishV1 struct {
}

func (s *hardForkPunishV1) GetName() string {
	return PunishV1ContractName
}

func (s *hardForkPunishV1) Update(config *params.ChainConfig, height *big.Int, state *state.StateDB) (err error) {
	contractCode := common.FromHex(punishV1Code)

	//write code to sys contract
	state.SetCode(PunishV1ContractAddr, contractCode)
	log.Debug("Write code to system contract account", "addr", PunishV1ContractAddr.String(), "code", punishV1Code)

	return
}

func (s *hardForkPunishV1) Execute(state *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (err error) {
	// initialize v1 contract
	method := "initialize"
	data, err := GetInteractiveABI()[s.GetName()].Pack(method)
	if err != nil {
		log.Error("Can't pack data for initialize", "error", err)
		return err
	}

	msg := types.NewMessage(header.Coinbase, &PunishV1ContractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)
	_, err = vmcaller.ExecuteMsg(msg, state, header, chainContext, config)

	return
}