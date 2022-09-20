package search

type Criteria int

const (
	UserByUsernameCriteria Criteria = 0
)

func IsASearchCriteria(criteria int) bool {
	if criteria == 0 {
		return true
	}
	return false
}
