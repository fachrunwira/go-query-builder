package builder

import "fmt"

func (q *queryStruct) Join(tableJoin, table_1, table_2 string) manipulateData {
	switch {
	case tableJoin == "":
		q.errors = fmt.Errorf("join table is empty")
	case table_1 == "":
		q.errors = fmt.Errorf("join column is empty")
	case table_2 == "":
		q.errors = fmt.Errorf("join column is empty")
	default:
		q.joins = append(q.joins, fmt.Sprintf(" JOIN %s ON %s = %s", tableJoin, table_1, table_2))
	}
	return q
}
