package mem

import (
	"gitlab.com/tokend/bridge/core/resources"
	"golang.org/x/exp/constraints"
)

func contains[T constraints.Ordered](src []T, value T) bool {
	for _, v := range src {
		if v == value {
			return true
		}
	}

	return false
}

// TODO: Remove this function when we will be able to use generics
func containsStr(src []string, value string) bool {
	for _, v := range src {
		if v == value {
			return true
		}
	}

	return false
}

func containsTokenType(src []resources.TokenType, value resources.TokenType) bool {
	for _, v := range src {
		if v == value {
			return true
		}
	}

	return false
}

func containsChainType(src []resources.ChainType, value resources.ChainType) bool {
	for _, v := range src {
		if v == value {
			return true
		}
	}

	return false
}
