package mem

import (
	"fmt"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/resources"
	"strconv"
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

func (q *tokensQ) Page(limitStr, currentPageStr, path string, tokens []data.Token) (data.TokensQList, error) {
	list := data.TokensQList{}
	limit, err := strconv.Atoi(limitStr)
	currentPage, err := strconv.Atoi(currentPageStr)
	if err != nil {
		return list, err
	}
	if limit == 0 {
		limit = 15
	}
	if currentPage < 1 {
		currentPage = 1
	}

	firstEntry := (currentPage - 1) * limit
	lastEntry := firstEntry + limit

	if lastEntry > len(tokens) {
		lastEntry = len(tokens)
	}
	list.Items = tokens[firstEntry:lastEntry]
	//index := strings.Index(path, "?") //
	//rootPath := path[:index]

	list.NextPageID = fmt.Sprint(path, "?page[page_number]=", currentPage+1) //todo add link

	return list, nil
}
