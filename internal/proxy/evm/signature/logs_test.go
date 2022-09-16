package signature

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	expectedResult := "0x8d731a2ffef3ab1b31560b622ea11072b21835c2cb5c4df77b23f30711541eff"

	am, _ := big.NewInt(0).SetString("100000000000000000000", 10)
	log := Erc20Log{
		TokenAddress: "0x26B862f640357268Bd2d9E95bc81553a2Aa81D7E",
		Amount:       am,
		Receiver:     "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		TxHash:       common.HexToHash("0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d"),
		EventIndex:   1794147,
		ChainID:      big.NewInt(31378),
		IsWrapped:    true,
	}

	hash := hexutil.Encode(log.Hash())

	if hash != expectedResult {
		t.Log(hash)
		t.Errorf("Wrong hash")
	}
}

func TestErc721Log(t *testing.T) {
	expectedResult := "0x66d0f2f121e882581235c708b3bf5ce9f8e45724d9c78ab7c7d1fa284b32970e"

	log := Erc721Log{
		TokenAddress: "0xc7cDb7A2E5dDa1B7A0E792Fe1ef08ED20A6F56D4",
		TokenID:      big.NewInt(5000),
		Receiver:     "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		TxHash:       common.HexToHash("0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d"),
		EventIndex:   1794147,
		ChainID:      big.NewInt(31378),
		TokenUri:     "https://some.link",
		IsWrapped:    true,
	}

	hash := hexutil.Encode(log.Hash())

	if hash != expectedResult {
		t.Log(hash)
		t.Errorf("Wrong hash")
	}
}

func TestErc1155Log(t *testing.T) {
	expectedResult := "0xd0e3a90c86f4b3cfbe244401f01264fd5f3ce096c0caa1529328897f1fc5c122"

	log := Erc1155Log{
		TokenAddress: "0xCace1b78160AE76398F486c8a18044da0d66d86D",
		TokenID:      big.NewInt(5000),
		Amount:       big.NewInt(10),
		Receiver:     "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		TxHash:       common.HexToHash("0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d"),
		EventIndex:   1794147,
		ChainID:      big.NewInt(31378),
		TokenUri:     "https://some.link",
		IsWrapped:    true,
	}

	hash := hexutil.Encode(log.Hash())

	if hash != expectedResult {
		t.Log(hash)
		t.Errorf("Wrong hash")
	}
}
