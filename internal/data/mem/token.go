package mem

import (
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/resources"
)

func NewTokenQ(tokens []data.Token) data.TokensQ {
	return &tokensQ{
		tokens:  tokens,
		filters: make([]tokenFilter, 0),
	}
}

type tokensQ struct {
	tokens  []data.Token
	filters []tokenFilter
}

type tokenFilter func(token data.Token) bool

func (q *tokensQ) New() data.TokensQ {
	return NewTokenQ(q.tokens)
}

func (q *tokensQ) Get() (*data.Token, error) {
	for _, token := range q.tokens {
		if q.filter(token) {
			return &token, nil
		}
	}

	return nil, nil
}

func (q *tokensQ) Select() ([]data.Token, error) {
	result := make([]data.Token, 0, len(q.tokens))
	for _, token := range q.tokens {
		if q.filter(token) {
			result = append(result, token)
		}
	}

	return result, nil
}

func (q *tokensQ) FilterByID(ids ...string) data.TokensQ {
	q.filters = append(q.filters, func(token data.Token) bool {
		return containsStr(ids, token.ID)
	})
	return q
}

func (q *tokensQ) FilterByType(types ...resources.TokenType) data.TokensQ {
	q.filters = append(q.filters, func(token data.Token) bool {
		return containsTokenType(types, token.Type)
	})
	return q
}

func (q *tokensQ) filter(token data.Token) bool {
	for _, filter := range q.filters {
		if !filter(token) {
			return false
		}
	}

	return true
}

func (q *tokensQ) PageTokens(token data.Token) data.TokensQ {
	result := make([]data.Token, 0, len(q.tokens))
	//todo add pagination
	for _, token := range q.tokens {
		if q.filter(token) {
			result = append(result, token)
		}
	}
	return q
}

//
//func  ApplyTo(Tokens data.Token, cols ...string) squirrel.SelectBuilder {
//	if p.Limit == 0 {
//		p.Limit = 15
//	}
//	if p.Order == "" {
//		p.Order = OrderTypeDesc
//	}
//
//	offset := p.Limit * p.PageNumber
//
//	sql = sql.Limit(p.Limit).Offset(offset)
//
//	switch p.Order {
//	case OrderTypeAsc:
//		for _, col := range cols {
//			sql = sql.OrderBy(fmt.Sprintf("%s %s", col, "asc"))
//		}
//	case OrderTypeDesc:
//		for _, col := range cols {
//			sql = sql.OrderBy(fmt.Sprintf("%s %s", col, "desc"))
//		}
//	default:
//		panic(fmt.Errorf("unexpected order type: %v", p.Order))
//	}
//
//	return sql
//}
