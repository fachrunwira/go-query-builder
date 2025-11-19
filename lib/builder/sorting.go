package builder

import "fmt"

func (q *queryStruct) OrderByAsc(column string) manipulateData {
	if column != "" {
		q.grouping = append(q.grouping, fmt.Sprintf("%s ASC", column))
	}

	return q
}

func (q *queryStruct) OrderByDesc(column string) manipulateData {
	if column != "" {
		q.grouping = append(q.grouping, fmt.Sprintf("%s DESC", column))
	}

	return q
}
