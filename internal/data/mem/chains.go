package mem

import (
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/resources"
	"strconv"
)

const (
	// OrderTypeAsc means result should be sorted in ascending order.
	OrderTypeAsc = "asc"
	// OrderTypeDesc means result should be sorted in descending order.
	OrderTypeDesc = "desc"
)

func NewChainsQ(chains []data.Chain) data.ChainsQ {
	return &chainsQ{
		chains:  chains,
		filters: make([]chainFilter, 0),
	}
}

type chainsQ struct {
	chains  []data.Chain
	filters []chainFilter
}

type chainFilter func(value data.Chain) bool

func (q *chainsQ) New() data.ChainsQ {
	return NewChainsQ(q.chains)
}

func (q *chainsQ) Get() (*data.Chain, error) {
	for _, value := range q.chains {
		if q.filter(value) {
			return &value, nil
		}
	}

	return nil, nil
}

func (q *chainsQ) Select() ([]data.Chain, error) {
	result := make([]data.Chain, 0, len(q.chains))
	for _, value := range q.chains {
		if q.filter(value) {
			result = append(result, value)
		}
	}

	return result, nil
}

func (q *chainsQ) FilterByID(ids ...string) data.ChainsQ {
	q.filters = append(q.filters, func(value data.Chain) bool {
		return containsStr(ids, value.ID)
	})
	return q
}

func (q *chainsQ) FilterByType(types ...resources.ChainType) data.ChainsQ {
	q.filters = append(q.filters, func(value data.Chain) bool {
		return containsChainType(types, value.Type)
	})
	return q
}

func (q *chainsQ) filter(value data.Chain) bool {
	for _, filter := range q.filters {
		if !filter(value) {
			return false
		}
	}

	return true
}

// ArticleList contains a list of articles

// Return each page content
func (q *chainsQ) Paginate(limitStr, currentPageStr string, chains []data.Chain) (data.ChainsQList, error) {
	list := data.ChainsQList{}
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

	if lastEntry > len(chains) {
		lastEntry = len(chains)
	}
	list.Items = chains[firstEntry:lastEntry]
	list.NextPageID = currentPage + 1 //todo add link
	return list, nil
}

//func Pagination(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		PageID := r.URL.Query().Get("id")
//		intPageID := 0
//		var err error
//		if PageID != "" {
//			intPageID, err = strconv.Atoi(PageID)
//			if err != nil {
//				ape.Render(w, err)
//				return
//			}
//		}
//		ctx := context.WithValue(r.Context(), "chains", intPageID)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}
