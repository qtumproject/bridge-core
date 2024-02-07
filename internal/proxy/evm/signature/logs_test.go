package signature

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/enums"
	"math/big"
	"testing"
)

func TestNativeLog(t *testing.T) {
	expectedResult := "0x6514d8c53feac8b0e102b3e25b8778fb849645e19e1badeadc2453b8232368a5"

	am, _ := big.NewInt(0).SetString("10000000000000000000", 10)
	log := NativeLog{
		Amount:     am,
		Receiver:   "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		TxHash:     common.HexToHash("0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d"),
		EventIndex: 1794147,
		ChainID:    big.NewInt(31378),
	}

	hash := hexutil.Encode(log.Hash())

	if hash != expectedResult {
		t.Log(hash)
		t.Errorf("Wrong hash")
	}
}

func TestErc20Log(t *testing.T) {
	expectedResult := "0x384f9cbeb9f719bd7a883505ae21aef2ab7608b1ff7c54658dcba5bca69b860d"

	am, _ := big.NewInt(0).SetString("100000000000000000000", 10)
	log := Erc20Log{
		TokenAddress: "0x322813Fd9A801c5507c9de605d63CEA4f2CE6c44",
		Amount:       am,
		Receiver:     "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		TxHash:       common.HexToHash("0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d"),
		EventIndex:   1794147,
		ChainID:      big.NewInt(31378),
		BridgingType: enums.ToErc20BridgingType(data.BridgingTypeWrapped),
	}

	hash := hexutil.Encode(log.Hash())
	println(hash)

	if hash != expectedResult {
		t.Log(hash)
		t.Errorf("Wrong hash")
	}
}

func TestErc721Log(t *testing.T) {
	expectedResult := "0x1bf4a3053a90252cfa79d72727d1cd31f8c04de8ba416cf7153ce0632f542bdc"

	log := Erc721Log{
		TokenAddress: "0xc3e53F4d16Ae77Db1c982e75a937B9f60FE63690",
		TokenID:      big.NewInt(5000),
		Receiver:     "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		TxHash:       common.HexToHash("0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d"),
		EventIndex:   1794147,
		ChainID:      big.NewInt(31378),
		TokenUri:     "https://some.link",
		BridgingType: enums.ToErc721BridgingType(data.BridgingTypeWrapped),
	}

	hash := hexutil.Encode(log.Hash())

	if hash != expectedResult {
		t.Log(hash)
		t.Errorf("Wrong hash")
	}
}

func TestErc1155Log(t *testing.T) {
	expectedResult := "0x64905dbe4e25d5915269edc98c27cab61b147d4ccf4e0e0b8dba20830a81f6a3"

	log := Erc1155Log{
		TokenAddress: "0x68B1D87F95878fE05B998F19b66F4baba5De1aed",
		TokenID:      big.NewInt(5000),
		Amount:       big.NewInt(10),
		Receiver:     "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		TxHash:       common.HexToHash("0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d"),
		EventIndex:   1794147,
		ChainID:      big.NewInt(31378),
		TokenUri:     "https://some.link",
		BridgingType: enums.ToErc1155BridgingType(data.BridgingTypeWrapped),
	}

	hash := hexutil.Encode(log.Hash())

	if hash != expectedResult {
		t.Log(hash)
		t.Errorf("Wrong hash")
	}
}
