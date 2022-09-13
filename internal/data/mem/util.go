package mem

import "gitlab.com/tokend/bridge/core/resources"

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
