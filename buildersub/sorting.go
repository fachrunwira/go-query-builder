package buildersub

import "fmt"

func (q *queryStruct) OrderByAsc(column string) SubQuery {
	if column != "" {
		q.ordering = append(q.ordering, fmt.Sprintf("%s ASC", column))
	}

	return q
}

func (q *queryStruct) OrderByDesc(column string) SubQuery {
	if column != "" {
		q.ordering = append(q.ordering, fmt.Sprintf("%s DESC", column))
	}

	return q
}
