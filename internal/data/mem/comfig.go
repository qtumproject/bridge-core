package mem

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/resources"

	"reflect"
)

func NewChainer(getter kv.Getter) Chainer {
	return &chainer{
		getter: getter,
	}
}

type Chainer interface {
	Chains() []data.Chain
	Tokens() []data.Token
	TokenChains() []data.TokenChain
}

type chainer struct {
	getter kv.Getter
	once   comfig.Once

	chains      []data.Chain
	tokens      []data.Token
	tokenChains []data.TokenChain
}

func (c *chainer) Chains() []data.Chain {
	c.readConfig()
	return c.chains
}

func (c *chainer) Tokens() []data.Token {
	c.readConfig()
	return c.tokens
}

func (c *chainer) TokenChains() []data.TokenChain {
	c.readConfig()
	return c.tokenChains
}

func (c *chainer) readConfig() {
	c.once.Do(func() interface{} {
		cfg := struct {
			Tokens []data.Token `fig:"tokens,required"`
			Chains []data.Chain `fig:"chains,required"`
		}{}
		err := figure.
			Out(&cfg).
			With(figure.BaseHooks, topHooks).
			From(kv.MustGetStringMap(c.getter, "data")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out signer"))
		}

		for i, chain := range cfg.Chains {
			chain.Tokens = make([]data.TokenChain, 0)
			cfg.Chains[i] = chain
		}

		tokenChains := make([]data.TokenChain, 0)
		for _, token := range cfg.Tokens {
			for i, tokenChain := range token.Chains {
				tokenChain.TokenID = token.ID
				token.Chains[i] = tokenChain
				tokenChains = append(tokenChains, tokenChain)
				for k, chain := range cfg.Chains {
					if chain.ID == tokenChain.ChainID {
						chain.Tokens = append(chain.Tokens, tokenChain)
						cfg.Chains[k] = chain
					}
				}
			}
		}

		c.tokens = cfg.Tokens
		c.chains = cfg.Chains
		c.tokenChains = tokenChains

		return nil
	})
}

var topHooks = figure.Hooks{
	"[]data.Token": func(value interface{}) (reflect.Value, error) {
		switch s := value.(type) {
		case []interface{}:
			chains := make([]data.Token, 0, len(s))
			var err error
			for _, rawElem := range s {
				mapElem, ok := rawElem.(map[interface{}]interface{})
				if !ok {
					return reflect.Value{}, errors.Wrap(err,
						"failed to cast mapElem to interface")
				}

				normMap := make(map[string]interface{}, len(mapElem))
				for key, value := range mapElem {
					strKey := fmt.Sprintf("%v", key)
					normMap[strKey] = value
				}

				var data data.Token

				err := figure.
					Out(&data).
					With(figure.BaseHooks, hooks).
					From(normMap).
					Please()
				if err != nil {
					return reflect.Value{}, errors.Wrap(err, "failed to figure out")
				}

				chains = append(chains, data)
			}

			return reflect.ValueOf(chains), nil
		default:
			return reflect.Value{}, errors.New("unexpected type while figure []data.TokenChain")
		}
	},
	"[]data.Chain": func(value interface{}) (reflect.Value, error) {
		switch s := value.(type) {
		case []interface{}:
			chains := make([]data.Chain, 0, len(s))
			var err error
			for _, rawElem := range s {
				mapElem, ok := rawElem.(map[interface{}]interface{})
				if !ok {
					return reflect.Value{}, errors.Wrap(err,
						"failed to cast mapElem to interface")
				}

				normMap := make(map[string]interface{}, len(mapElem))
				for key, value := range mapElem {
					strKey := fmt.Sprintf("%v", key)
					normMap[strKey] = value
				}

				var data data.Chain

				err := figure.
					Out(&data).
					With(figure.BaseHooks, hooks).
					From(normMap).
					Please()
				if err != nil {
					return reflect.Value{}, errors.Wrap(err, "failed to figure out")
				}

				chains = append(chains, data)
			}

			return reflect.ValueOf(chains), nil
		default:
			return reflect.Value{}, errors.New("unexpected type while figure []data.TokenChain")
		}
	},
}

var hooks = figure.Hooks{
	"resources.TokenType": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToStringE(value)
		if err != nil {
			return reflect.Value{}, errors.Wrap(err, "failed to parse string")
		}
		return reflect.ValueOf(resources.TokenType(result)), nil
	},
	"resources.ChainType": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToStringE(value)
		if err != nil {
			return reflect.Value{}, errors.Wrap(err, "failed to parse string")
		}
		return reflect.ValueOf(resources.ChainType(result)), nil
	},
	"json.RawMessage": func(value interface{}) (reflect.Value, error) {
		if value == nil {
			return reflect.Value{}, nil
		}

		var params map[string]interface{}
		switch s := value.(type) {
		case map[interface{}]interface{}:
			params = make(map[string]interface{})
			for key, value := range s {
				params[key.(string)] = value
			}
		default:
			return reflect.Value{}, errors.New("unexpected type while figure []json.RawMessage")
		}

		result, err := json.Marshal(params)
		if err != nil {
			return reflect.Value{}, errors.Wrap(err, "failed to parse json.RawMessage")
		}
		return reflect.ValueOf(json.RawMessage(result)), nil
	},
	"[]data.TokenChain": func(value interface{}) (reflect.Value, error) {
		switch s := value.(type) {
		case []interface{}:
			chains := make([]data.TokenChain, 0, len(s))
			var err error
			for _, rawElem := range s {
				mapElem, ok := rawElem.(map[interface{}]interface{})
				if !ok {
					return reflect.Value{}, errors.Wrap(err,
						"failed to cast mapElem to interface")
				}

				normMap := make(map[string]interface{}, len(mapElem))
				for key, value := range mapElem {
					strKey := fmt.Sprintf("%v", key)
					normMap[strKey] = value
				}

				var data data.TokenChain

				err := figure.
					Out(&data).
					With(figure.BaseHooks).
					From(normMap).
					Please()
				if err != nil {
					return reflect.Value{}, errors.Wrap(err, "failed to figure out")
				}

				chains = append(chains, data)
			}

			return reflect.ValueOf(chains), nil
		default:
			return reflect.Value{}, errors.New("unexpected type while figure []data.TokenChain")
		}
	},
}
